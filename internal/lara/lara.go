package lara

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	dpos_contract "github.com/Taraxa-project/taraxa-indexer/abi/dpos"
	lara_contract "github.com/Taraxa-project/taraxa-indexer/abi/lara"
	multicall_contract "github.com/Taraxa-project/taraxa-indexer/abi/multicall"
	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/transact"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type State struct {
	epochDuration                 *big.Int
	lastSnapshotBlock             *big.Int
	lastSnapshotID                *big.Int
	lastRebalance                 *big.Int
	lastEpochTotalDelegatedAmount *big.Int
	validatorStakes               map[common.Address]*big.Int
	validators                    []oracle.NodeData
	isMakingSnapshot              bool
	isRebalancing                 bool
}
type Lara struct {
	deploymentAddress string
	Eth               *ethclient.Client
	signer            *bind.TransactOpts
	chainID           *int
	contract          *lara_contract.LaraContract
	oracle            *apy_oracle.ApyOracle
	dpos              *dpos_contract.DposContract
	state             State
	graphQLEndpoint   string
	multicall         *MulticallContract
	lastSnapshotId    uint64
}

type SnapshotTakenEvent struct {
	SnapshotId           *big.Int
	TotalDelegation      *big.Int
	DistributableRewards *big.Int
}

func MakeLara(rpc *ethclient.Client, signing_key, deployment_address, oracle_address, graphQLEndpoint string, chainID int, lastSnapshotId uint64) *Lara {
	l := new(Lara)
	l.Eth = rpc
	l.signer = transact.MakeSigner(signing_key, chainID)
	l.deploymentAddress = deployment_address
	l.chainID = &chainID
	l.graphQLEndpoint = graphQLEndpoint
	l.lastSnapshotId = lastSnapshotId
	contract, err := lara_contract.NewLaraContract(common.HexToAddress(l.deploymentAddress), l.Eth)
	if err != nil {
		log.Fatalf("Failed to create contract: %v", err)
	}
	l.oracle, err = apy_oracle.NewApyOracle(common.HexToAddress(oracle_address), l.Eth)
	if err != nil {
		log.Fatalf("Failed to create oracle: %v", err)
	}
	l.dpos, err = dpos_contract.NewDposContract(common.HexToAddress("0x00000000000000000000000000000000000000fe"), l.Eth)
	if err != nil {
		log.Fatalf("Failed to create dpos: %v", err)
	}
	l.contract = contract
	l.SyncState()
	l.multicall = MakeMulticallContract(l.Eth, signing_key, "0xfce7a3121b42664aad145712e1c2bf2e38f60aa1", *l.chainID, 31500000)
	return l
}

func (l *Lara) Run(interval int, generalBlockTime int) {
	if l.Eth == nil {
		log.Fatalf("Eth client is nil")
	}
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	log.WithFields(log.Fields{"interval": interval, "generalBlockTime": generalBlockTime}).Info("LARA: Registering rewards distribution ticker")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start listening for new snapshots
	go l.ListenForSnapshots(ctx)

	for {
		select {
		case <-ticker.C:
			ctx := context.Background()
			currentBlock, err := l.Eth.BlockNumber(ctx)
			if err != nil {
				log.Fatalf("Lara: Failed to get current block: %v", err)
			}
			// if we pass the time to end epoch
			expectedSnapshotTime := l.state.lastSnapshotBlock.Int64() + l.state.epochDuration.Int64()
			expectedRebalanceTime := l.state.lastRebalance.Int64() + l.state.epochDuration.Int64()
			l.SyncState()

			if int64(currentBlock) > expectedSnapshotTime {
				l.Compound()

				l.SyncState()
			}
			if int64(currentBlock) > expectedRebalanceTime {
				log.Warnf("Triggering rebalance at block: %d, expected rebalance time: %d", currentBlock, expectedRebalanceTime)
				l.Rebalance()
			}
		}
	}
}

func (l *Lara) retryTransaction(txFunc func() (*types.Transaction, error), description string) error {
	maxRetries := 5
	initialDelay := 1 * time.Second
	maxDelay := 16 * time.Second

	for attempt := 0; attempt < maxRetries; attempt++ {
		tx, err := txFunc()
		if err != nil {
			if strings.Contains(err.Error(), "Transaction already in transactions pool") {
				log.Warnf("%s tx already in pool", description)
				return nil
			} else if strings.Contains(err.Error(), "No nodes available for delegation") {
				log.Warnf("No nodes available for delegation")
				return nil
			} else {
				log.Errorf("Failed to %s: %v", description, err)
			}
		} else {
			log.WithFields(log.Fields{"txhash": tx.Hash().Hex()}).Infof("LARA %s: ", strings.ToUpper(description))
			receipt, err := bind.WaitMined(context.Background(), l.Eth, tx)
			if err != nil {
				log.Errorf("Failed to wait for transaction to be mined: %v", err)
				continue // Retry the transaction
			}

			// Check if the transaction was successful
			if receipt.Status == types.ReceiptStatusFailed {
				log.Errorf("Transaction %s failed", description)
				continue // Retry the transaction
			}
			return nil
		}

		delay := initialDelay * time.Duration(math.Pow(2, float64(attempt)))
		jitter := time.Duration(rand.Int63n(int64(delay) / 2))
		delay = delay + jitter

		if delay > maxDelay {
			delay = maxDelay
		}

		log.Infof("Retrying %s in %v...", description, delay)
		time.Sleep(delay)
	}

	return fmt.Errorf("failed to %s after maximum retries", description)
}

func (l *Lara) IsSnapshotDistributedToUser(snapshotId uint64, userAddress common.Address) bool {
	opts := &bind.CallOpts{
		Pending:     false,
		From:        l.signer.From,
		BlockNumber: nil,
		Context:     nil,
	}
	laraDistributedAlready, err := l.contract.StakerSnapshotClaimed(opts, userAddress, big.NewInt(int64(snapshotId)))
	if err != nil {
		log.Fatalf("Failed to get staker snapshot claimed: %v", err)
	}
	return laraDistributedAlready
}

func (l *Lara) GetRewardsPerSnapshot(snapshotId uint64) *big.Int {
	opts := &bind.CallOpts{
		Pending:     false,
		From:        l.signer.From,
		BlockNumber: nil,
		Context:     nil,
	}
	rewards, err := l.contract.RewardsPerSnapshot(opts, big.NewInt(int64(snapshotId)))
	if err != nil {
		log.Fatalf("Failed to get rewards per snapshot: %v", err)
	}
	return rewards
}

func (l *Lara) retryMulticallDistributeRewards(holderAddresses []common.Address, snapshotId uint64) error {
	log.Infof("LARA: Distributing rewards for snapshot %d to %d holders", snapshotId, len(holderAddresses))

	const laraABI = `[{"type":"function","name":"distributeRewardsForSnapshot","inputs":[{"name":"staker","type":"address","internalType": "address"},{"name":"snapshotId","type":"uint256","internalType": "uint256"}],"outputs":[],"stateMutability":"nonpayable"}]`

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(laraABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}
	if len(holderAddresses) > 150 {
		log.Warnf("Multicall is too large, splitting into chunks")
		// execute in chunks of 150
		for i := 0; i < len(holderAddresses); i += 150 {
			chunk := holderAddresses[i:min(i+150, len(holderAddresses))]
			err := l.retryMulticallDistributeRewards(chunk, snapshotId)
			if err != nil {
				return err
			}
		}
	} else {
		multicall := make([]multicall_contract.MulticallCall, 0)
		for _, holderAddress := range holderAddresses {
			data, err := parsedABI.Pack("distributeRewardsForSnapshot", holderAddress, big.NewInt(int64(snapshotId)))
			if err != nil {
				log.Fatalf("Failed to pack data: %v", err)
			}
			multicall = append(multicall, multicall_contract.MulticallCall{
				Target:   common.HexToAddress(l.deploymentAddress),
				CallData: data,
			})
		}

		if len(multicall) == 0 {
			log.Infof("Empty multicall")
			return nil
		}

		tx, err := l.multicall.Multicall(multicall)
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{"txhash": tx.TxHash.Hex()}).Infof("LARA: Distribute rewards for snapshot %d to holders", snapshotId)
	}

	return nil
}

func (l *Lara) retryDistributeRewards(holderAddress common.Address, snapshotId *big.Int) error {

	const laraABI = `[{"type":"function","name":"distributeRewardsForSnapshot","inputs":[{"name":"staker","type":"address","internalType": "address"},{"name":"snapshotId","type":"uint256","internalType": "uint256"}],"outputs":[],"stateMutability":"nonpayable"}]`

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(laraABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	err = l.retryTransaction(func() (*types.Transaction, error) {
		deployAddress := common.HexToAddress(l.deploymentAddress)
		data, err := parsedABI.Pack("distributeRewardsForSnapshot", holderAddress, snapshotId)
		if err != nil {
			log.Fatalf("Failed to pack data: %v", err)
		}
		gasLimit, err := l.Eth.EstimateGas(context.Background(), ethereum.CallMsg{
			From: l.signer.From,
			To:   &deployAddress,
			Data: data,
		})
		if err != nil {
			log.Errorf("Failed to estimate gas: %v", err)
			return nil, err
		}

		// Add a buffer to the estimated gas limit
		gasLimit = gasLimit + (gasLimit / 20) // Add 20% buffer

		opts := &bind.TransactOpts{
			From:     l.signer.From,
			Signer:   l.signer.Signer,
			GasLimit: gasLimit,
			Context:  context.Background(),
		}
		return l.contract.DistributeRewardsForSnapshot(opts, holderAddress, snapshotId)
	}, fmt.Sprintf("distribute rewards for snapshot %s to holder %s", snapshotId.String(), holderAddress.Hex()))

	return err
}

func (l *Lara) DisburseRewardsBetweenHolders(snapshot SnapshotWithBlock) uint64 {
	if snapshot.TotalRewards.Cmp(big.NewInt(0)) == 0 {
		l.lastSnapshotId = snapshot.ID
		log.WithFields(log.Fields{"snapshotID": snapshot.ID}).Info("LARA: No rewards to distribute")
		return 0
	}
	log.WithFields(log.Fields{"blockNumber": snapshot.Block, "snapshotID": snapshot.ID}).Info("LARA: Getting staked tara holders")
	holders := GetStakedTaraHolders(l.graphQLEndpoint, snapshot.Block)

	log.WithFields(log.Fields{"# of holders": len(holders), "snapshotID": snapshot.ID}).Info("LARA: Disbursing rewards to holders for snapshot")

	holdersToDistribute := make([]common.Address, 0)
	for _, holder := range holders {
		holderAddress := common.HexToAddress(holder)

		isDistributed := l.IsSnapshotDistributedToUser(snapshot.ID, holderAddress)
		if isDistributed {
			log.WithFields(log.Fields{"holder": holder, "snapshotID": snapshot.ID}).Info("LARA: Snapshot already distributed to holder")
			continue
		} else {
			holdersToDistribute = append(holdersToDistribute, holderAddress)
		}
	}
	if len(holdersToDistribute) == 0 {
		log.WithFields(log.Fields{"snapshotID": snapshot.ID}).Warn("LARA: No holders to distribute rewards to")
		return 0
	}

	err := l.retryMulticallDistributeRewards(holdersToDistribute, snapshot.ID)
	if err != nil {
		log.Fatalf("Failed to disburse rewards for snapshot: %v", err)
	}
	return snapshot.ID
}

func (l *Lara) Compound() {
	laraEthBalance, err := l.Eth.BalanceAt(context.Background(), common.HexToAddress(l.deploymentAddress), nil)
	if err == nil {
		laraEthBalance = laraEthBalance.Sub(laraEthBalance, big.NewInt(1e18)) // Subtract one ETH (1 ETH = 10^18 wei)
	}
	if err != nil {
		log.Errorf("Failed to get lara eth balance: %v", err)
		return
	}

	taraThreshold := big.NewInt(1000)
	taraThreshold.Mul(taraThreshold, big.NewInt(1e18)) // Convert to wei

	if laraEthBalance.Cmp(taraThreshold) < 0 {
		log.Info("LARA: Balance is less than 1000 TARA, skipping compounding")
		return
	}

	err = l.retryTransaction(func() (*types.Transaction, error) {
		opts := &bind.TransactOpts{
			From:     l.signer.From,
			Signer:   l.signer.Signer,
			GasLimit: 0,
			Context:  nil,
		}
		return l.contract.Compound(opts, laraEthBalance)
	}, "compound")

	if err != nil {
		log.Error(err)
	}
	log.WithFields(log.Fields{"laraEthBalance": laraEthBalance}).Info("LARA: Compounded")
}

func (l *Lara) SyncState() {
	opts := &bind.CallOpts{
		Pending:     false,
		From:        l.signer.From,
		BlockNumber: nil,
		Context:     nil,
	}
	epochDuration, err := l.contract.EpochDuration(opts)
	if err != nil {
		log.Fatalf("Failed to get epoch duration: %v", err)
	}

	lastSnapshotBlock, err := l.contract.LastSnapshotBlock(opts)
	if err != nil {
		log.Fatalf("Failed to get last snapshot: %v", err)
	}

	lastSnapshotID, err := l.contract.LastSnapshotId(opts)
	if err != nil {
		log.Fatalf("Failed to get last snapshot ID: %v", err)
	}

	lastRebalanceBlock, err := l.contract.LastRebalance(opts)
	if err != nil {
		log.Fatalf("Failed to get last rebalance block: %v", err)
	}

	lastEpochTotalDelegatedAmount, err := l.dpos.GetTotalDelegation(opts, common.HexToAddress(l.deploymentAddress))
	if err != nil {
		log.Fatalf("Failed to get last epoch total delegated amount: %v", err)
	}

	// fetch validators from 0 until reverted and put in map
	pos := big.NewInt(0)
	validatorStakes := make(map[common.Address]*big.Int)
	validators := make([]oracle.NodeData, 0)
	for {
		validatorsFromDpos, err := l.dpos.GetDelegations(opts, common.HexToAddress(l.deploymentAddress), uint32(pos.Uint64()))
		if err != nil {
			if strings.Contains(err.Error(), "reverted") {
				break
			} else {
				log.Fatalf("Failed to get validator: %v", err)
				break
			}
		}
		for _, delegation := range validatorsFromDpos.Delegations {
			validator := delegation.Account
			totalStakeAtValidator, err := l.contract.ProtocolTotalStakeAtValidator(opts, validator)
			if err != nil {
				log.Fatalf("Failed to get total stake at validator: %v", err)
			}
			validatorStakes[validator] = totalStakeAtValidator

			// fetch node data from oracle
			nodeData, err := l.oracle.Nodes(opts, validator)
			if err != nil {
				log.Fatalf("Failed to get node data: %v", err)
			}
			validators = append(validators, nodeData)
		}
		if !validatorsFromDpos.End {
			pos.Add(pos, big.NewInt(1))
		} else {
			break
		}
	}

	//set commission to 8%
	commission, err := l.contract.Commission(opts)
	if err != nil {
		log.Fatalf("Failed to get commission: %v", err)
	}
	if commission.Cmp(big.NewInt(8)) != 0 {
		_, err = l.contract.SetCommission(l.signer, big.NewInt(8))

		if err != nil && !strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Fatalf("Failed to set commission: %v", err)
		}
		// wait 1 sec
		time.Sleep(1 * time.Second)
	}
	l.state = State{
		epochDuration:                 epochDuration,
		lastSnapshotBlock:             lastSnapshotBlock,
		lastSnapshotID:                lastSnapshotID,
		lastRebalance:                 lastRebalanceBlock,
		lastEpochTotalDelegatedAmount: lastEpochTotalDelegatedAmount,
		validatorStakes:               validatorStakes,
		validators:                    validators,
	}
	nextSnapshot := big.NewInt(0).Add(lastSnapshotBlock, epochDuration)
	currentBlock, err := l.Eth.BlockNumber(context.Background())
	if err != nil {
		log.Fatalf("SyncState: Failed to get current block: %v", err)
	}
	log.WithFields(log.Fields{"currentBlock": currentBlock, "lastRebalance": l.state.lastRebalance, "lastSnapshotBlock": l.state.lastSnapshotBlock, "nextSnapshotBlock": nextSnapshot, "nodesDelegatedTo": len(l.state.validators), "totalDelegated": l.state.lastEpochTotalDelegatedAmount}).Info("LARA STATE: ")
}

func (l *Lara) GetState() State {
	return l.state
}

func (l *Lara) Rebalance() {
	if l.state.isRebalancing {
		log.Error("WARN: PENDING REBALANCE")
		return
	}
	if l.state.isMakingSnapshot {
		log.Error("WARN: SNAPSHOT IN PROGRESS")
		return
	}

	l.state.isRebalancing = true
	defer func() {
		l.state.isRebalancing = false
	}()

	err := l.retryTransaction(func() (*types.Transaction, error) {
		opts := &bind.TransactOpts{
			From:     l.signer.From,
			Signer:   l.signer.Signer,
			GasLimit: 0,
			Context:  context.Background(),
		}
		return l.contract.Rebalance(opts)
	}, "rebalance")

	if err != nil {
		if strings.Contains(err.Error(), "Transaction already in transactions pool") {
			log.Warn("Rebalance tx already in pool")
		} else {
			log.Errorf("Failed to rebalance: %v", err)
		}
	}
}

func (l *Lara) ProcessSnapshot(snapshot SnapshotWithBlock) {
	if snapshot.TotalRewards.Cmp(big.NewInt(0)) == 0 {
		log.WithFields(log.Fields{"snapshotID": snapshot.ID}).Info("Skipping snapshot with zero rewards")
		l.lastSnapshotId = snapshot.ID
		return
	}

	log.WithFields(log.Fields{
		"blockNumber":  snapshot.Block,
		"snapshotID":   snapshot.ID,
		"totalRewards": snapshot.TotalRewards,
	}).Info("LARA: Processing snapshot")

	stakers := l.DisburseRewardsBetweenHolders(snapshot)
	if stakers == 0 {
		l.lastSnapshotId = snapshot.ID
		log.WithFields(log.Fields{"snapshotID": snapshot.ID}).Info("No stakers to distribute rewards to")
	}
}

func (l *Lara) ListenForSnapshots(ctx context.Context) {
	snapshotTakenSig := []byte("SnapshotTaken(uint256,uint256,uint256)")
	snapshotTakenHash := crypto.Keccak256Hash(snapshotTakenSig)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(l.deploymentAddress)},
		Topics:    [][]common.Hash{{snapshotTakenHash}},
	}

	eventCh := make(chan types.Log)
	log.Info("Listening for SnapshotTaken events")
	sub, err := l.Eth.SubscribeFilterLogs(ctx, query, eventCh)
	log.Info("Subscribed to SnapshotTaken events")
	if err != nil {
		log.Fatalf("Failed to subscribe to SnapshotTaken events: %v", err)
	}

	healthCheckTicker := time.NewTicker(1 * time.Minute)
	defer healthCheckTicker.Stop()

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("Error in event subscription: %v", err)
			return
		case event := <-eventCh:
			snapshot := SnapshotWithBlock{
				ID:           event.Topics[1].Big().Uint64(),
				Block:        event.BlockNumber,
				TotalRewards: event.Topics[3].Big(),
			}
			l.ProcessSnapshot(snapshot)
		case <-healthCheckTicker.C:
			log.WithFields(log.Fields{
				"lastSnapshotId":  l.lastSnapshotId,
				"contractAddress": l.deploymentAddress,
			}).Info("SnapshotTaken event listener is active")
		case <-ctx.Done():
			return
		}
	}
}

func (l *Lara) FetchAndDistributePastRewards() {
	log.Info("Starting to fetch and distribute past rewards")

	snapshots := GetSnapshotsWithBlocks(l.graphQLEndpoint)
	if len(snapshots) == 0 {
		log.Info("No snapshots to process")
		return
	}

	sort.Slice(snapshots, func(i, j int) bool {
		return snapshots[i].ID > snapshots[j].ID
	})
	finalSnapshotId := l.lastSnapshotId
	log.WithFields(log.Fields{"maxSnapshotID": snapshots[0].ID}).Info("Max snapshot ID")

	var wg sync.WaitGroup
	var mu sync.Mutex

	concurrencyLimit := 5
	semaphore := make(chan struct{}, concurrencyLimit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, snapshot := range snapshots {
		if snapshot.ID >= l.lastSnapshotId {
			// Check if we should process this snapshot
			if snapshot.TotalRewards.Cmp(big.NewInt(0)) == 0 {
				log.WithFields(log.Fields{"snapshotID": snapshot.ID}).Info("Skipping snapshot with zero rewards")
				if snapshot.ID > finalSnapshotId {
					finalSnapshotId = snapshot.ID
				}
				continue
			}

			wg.Add(1)
			semaphore <- struct{}{}
			go func(snapshot SnapshotWithBlock) {
				defer wg.Done()
				defer func() { <-semaphore }()

				if ctx.Err() != nil {
					log.Infof("Goroutine for snapshotID %d skipped due to cancellation", snapshot.ID)
					return
				}

				log.Infof("Processing snapshotID %d", snapshot.ID)
				stakers := l.DisburseRewardsBetweenHolders(snapshot)

				if stakers == 0 {
					mu.Lock()
					if snapshot.ID > finalSnapshotId {
						log.Infof("Updating finalSnapshotId to %d", snapshot.ID)
						finalSnapshotId = snapshot.ID
					}
					mu.Unlock()
					// cancel()
				}
			}(snapshot)
		}
	}

	wg.Wait()

	log.Info("Finished processing past SnapshotTaken events")
	l.lastSnapshotId = finalSnapshotId
	log.Infof("Updating lastSnapshotId to %d", finalSnapshotId)
}
