package indexer

import (
	"math/big"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/rewards"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type blockContext struct {
	Storage      storage.Storage
	Batch        storage.Batch
	Config       *common.Config
	Client       chain.Client
	block        *chain.Block
	dags         []chain.DagBlock
	transactions []chain.Transaction
	accounts     storage.Accounts
	blockFee     *big.Int
	addressStats *storage.AddressStatsMap
	finalized    *storage.FinalizationData
}

func MakeBlockContext(s storage.Storage, client chain.Client, config *common.Config) *blockContext {
	var bc blockContext
	bc.Storage = s
	bc.Batch = s.NewBatch()
	bc.Client = client
	bc.Config = config
	bc.accounts = bc.Storage.GetAccounts()
	bc.blockFee = big.NewInt(0)

	bc.addressStats = storage.MakeAddressStatsMap()
	bc.finalized = s.GetFinalizationData()

	return &bc
}

func (bc *blockContext) commit() {
	bc.Batch.SetFinalizationData(bc.finalized)
	bc.addressStats.AddToBatch(bc.Batch)
	bc.Batch.CommitBatch()

	metrics.StorageCommitCounter.Inc()
}

func (bc *blockContext) process(raw chain.Block) (dags_count, trx_count uint64, err error) {
	start_processing := time.Now()
	bc.block = &raw

	tp := common.MakeThreadPool()
	tp.Go(func() { bc.updateValidatorStats(bc.block) })
	tp.Go(common.MakeTaskWithoutParams(bc.processDags, &err).Run)
	tp.Go(common.MakeTask(bc.processTransactions, raw.Transactions, &err).Run)
	votes := new(chain.VotesResponse)
	tp.Go(common.MakeTaskWithResult(bc.Client.GetPreviousBlockCertVotes, bc.block.Number, votes, &err).Run)

	validators := make([]chain.Validator, 0)
	tp.Go(common.MakeTaskWithResult(bc.Client.GetValidatorsAtBlock, bc.block.Number, &validators, &err).Run)

	tp.Wait()
	if err != nil {
		return
	}

	totalReward := common.ParseStringToBigInt(raw.TotalReward)

	r := rewards.MakeRewards(bc.Storage, bc.Batch, bc.Config, bc.block, bc.blockFee, validators)
	blockFee := r.Process(totalReward, bc.dags, bc.transactions, *votes)
	if blockFee != nil {
		bc.accounts.AddToBalance(common.DposContractAddress, blockFee)
	}

	bc.accounts.AddToBalance(common.DposContractAddress, totalReward)

	if bc.block.Number%1000 == 0 {
		bc.checkIndexedBalances()
	}
	bc.Batch.SaveAccounts(bc.accounts)

	dags_count = uint64(len(bc.dags))
	trx_count = bc.block.TransactionCount
	bc.finalized.TrxCount += trx_count
	bc.finalized.DagCount += dags_count
	bc.finalized.PbftCount++

	pbft_author_index := bc.addressStats.GetAddress(bc.Storage, bc.block.Author).AddPbft(bc.block.Timestamp)
	log.WithFields(log.Fields{"author": bc.block.Author, "hash": bc.block.Hash}).Debug("Saving PBFT block")
	bc.Batch.AddToBatch(bc.block, bc.block.Author, pbft_author_index)

	// If stats is available check for consistency
	remote_stats, stats_err := bc.Client.GetChainStats()
	if stats_err == nil {
		bc.finalized.Check(remote_stats)
	}
	bc.commit()
	r.AfterCommit()

	metrics.Save(start_processing, dags_count, trx_count, bc.finalized)

	return
}

func (bc *blockContext) checkIndexedBalances() {
	tp := common.MakeThreadPool()
	for _, account := range bc.accounts {
		tp.Go(func() {
			address := account.Address
			balance := account.Balance
			b, get_err := bc.Client.GetBalanceAtBlock(address, bc.block.Number)
			if get_err != nil {
				log.WithError(get_err).WithField("address", address).Warn("GetBalanceAtBlock error for address")
				return
			}
			chain_balance := common.ParseStringToBigInt(b)
			if balance.Cmp(chain_balance) != 0 {
				log.WithFields(log.Fields{"address": address, "balance": balance, "chain_balance": chain_balance}).Error("Balance check failed")
			}
		})
	}
	tp.Wait()
}

func (bc *blockContext) updateValidatorStats(block *chain.Block) {
	tn, _ := goment.Unix(int64(block.Timestamp))
	weekStats := bc.Storage.GetWeekStats(int32(tn.ISOWeekYear()), int32(tn.ISOWeek()))
	weekStats.AddPbftBlock(block.GetModel())
	bc.Batch.UpdateWeekStats(weekStats)
}

func (bc *blockContext) processDags() (err error) {
	dag_blocks, err := bc.Client.GetPeriodDagBlocks(bc.block.Number)
	if err != nil {
		log.WithError(err).Fatal("GetPeriodDagBlocks error")
	}
	bc.dags = make([]chain.DagBlock, len(dag_blocks))
	for i, dag := range dag_blocks {
		bc.dags[i] = dag
		bc.saveDag(&dag)
	}
	return
}

func (bc *blockContext) saveDag(dag *chain.DagBlock) {
	log.WithFields(log.Fields{"sender": dag.Sender, "hash": dag.Hash}).Trace("Saving DAG block")
	dag_index := bc.addressStats.GetAddress(bc.Storage, dag.Sender).AddDag(dag.GetModel().Timestamp)
	bc.Batch.AddToBatch(dag.GetModel(), dag.Sender, dag_index)
}
