package lara

import (
	"context"
	"math/big"

	multicall_contract "github.com/Taraxa-project/taraxa-indexer/abi/multicall"
	"github.com/Taraxa-project/taraxa-indexer/internal/transact"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type MulticallContract struct {
	Eth               *ethclient.Client
	signer            *bind.TransactOpts
	deploymentAddress string
	chainID           int
	contract          *multicall_contract.MulticallContract
	gasLimit          uint64
}

func MakeMulticallContract(rpc *ethclient.Client, signing_key, deployment_address string, chainID int, gasLimit uint64) *MulticallContract {
	state := new(MulticallContract)
	state.Eth = rpc
	state.signer = transact.MakeSigner(signing_key, chainID)
	state.deploymentAddress = deployment_address
	state.chainID = chainID
	state.gasLimit = gasLimit
	contract, err := multicall_contract.NewMulticallContract(common.HexToAddress(deployment_address), rpc)
	if err != nil {
		log.Fatalf("Failed to create multicall contract: %v", err)
	}
	state.contract = contract
	return state
}

func (state *MulticallContract) Multicall(multicall []multicall_contract.MulticallCall) (*types.Receipt, error) {
	if len(multicall) == 0 {
		log.Infof("No calls to multicall")
		return nil, nil
	}
	opts := &bind.TransactOpts{
		From:     state.signer.From,
		Signer:   state.signer.Signer,
		GasLimit: state.gasLimit,
		GasPrice: big.NewInt(110000000100),
		Context:  context.Background(),
	}
	tx, err := state.contract.Aggregate(opts, multicall)
	if err != nil {
		log.Errorf("Failed to aggregate multicall: %v", err)
		return nil, err
	}
	// wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), state.Eth, tx)
	if err != nil {
		log.Errorf("Failed to get transaction receipt for tx %s: %v", tx.Hash().Hex(), err)
		return nil, err
	}
	log.WithFields(log.Fields{"txhash": tx.Hash().Hex()}).Infof("Multicall executed")
	return receipt, nil
}
