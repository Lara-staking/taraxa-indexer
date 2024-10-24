package indexer

import (
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	log "github.com/sirupsen/logrus"
)

type Indexer struct {
	Client                      chain.Client
	oracle                      *oracle.Oracle
	storage                     storage.Storage
	config                      *common.Config
	retry_time                  time.Duration
	consistency_check_available bool
}

func (i *Indexer) Run(url string, s storage.Storage, c *common.Config, o *oracle.Oracle) {
	if o == nil {
		log.Fatal("Oracle is nil")
	}
	i.oracle = o
	for {
		err := i.run()
		f := i.storage.GetFinalizationData()
		log.WithError(err).WithField("period", f.PbftCount).Error("Error occurred. Restart indexing")
		time.Sleep(i.retry_time)
	}
}

func NewIndexer(url string, s storage.Storage, c *common.Config) (i *Indexer) {
	i = new(Indexer)
	i.retry_time = 5 * time.Second
	i.storage = s
	i.config = c
	// connect is retrying to connect every retry_time
	i.connect(url)
	log.Info("Indexer Instance initialized")
	return i
}

func (i *Indexer) connect(url string) {
	var err error
	var client *chain.WsClient
	for {
		client, err = chain.NewWsClient(url)
		i.Client = client
		if err == nil {
			break
		}
		log.WithError(err).Error("Can't connect to chain")
		time.Sleep(i.retry_time)
	}

	version, err := i.Client.GetVersion()
	if err != nil || !chain.CheckProtocolVersion(version) {
		log.WithFields(log.Fields{"version": version, "minimum": chain.MinimumProtocolVersion}).Fatal("Unsupported protocol version")
	}
	i.init()
	_, stats_err := i.Client.GetChainStats()
	i.consistency_check_available = (stats_err == nil)
	if !i.consistency_check_available {
		log.WithError(stats_err).Warn("Method for consistency check isn't available")
	}
}

func (i *Indexer) init() {
	genesis_blk, err := i.Client.GetBlockByNumber(0)
	if err != nil {
		log.WithError(err).Fatal("GetBlockByNumber error")
		return
	}

	remote_hash := storage.GenesisHash(genesis_blk.Hash)
	db_clean := false
	if i.storage.GenesisHashExist() {
		local_hash := i.storage.GetGenesisHash()
		if local_hash != remote_hash {
			log.WithFields(log.Fields{"local_hash": local_hash, "remote_hash": remote_hash}).Warn("Genesis changed, cleaning DB and restarting indexing")
			if err := i.storage.Clean(); err != nil {
				log.WithError(err).Warn("Error during storage cleaning")
			}
			db_clean = true
		}
	} else {
		db_clean = true
	}

	chain_genesis, err := i.Client.GetGenesis()
	if err != nil {
		log.WithError(err).Fatal("GetGenesis error")
	}
	i.config.Chain = chain_genesis.ToChainConfig()

	if err != nil {
		log.WithError(err).Fatal("Failed to get current block")
	}

	// Process genesis if db is clean
	if db_clean {
		genesis := MakeGenesis(i.storage, i.Client, i.oracle, chain_genesis, remote_hash)
		// Genesis hash and finalized period(0) is set inside
		genesis.process()
	}

}

func (i *Indexer) sync() error {
	// start processing blocks from the next one
	start := i.storage.GetFinalizationData().PbftCount + 1
	end, p_err := i.Client.GetLatestPeriod()
	if p_err != nil {
		return p_err
	}
	if start >= end {
		return nil
	}
	queue_limit := min(end-start, i.config.SyncQueueLimit)
	log.WithFields(log.Fields{"start": start, "end": end, "queue_limit": queue_limit}).Info("Syncing: started")
	sq := MakeSyncQueue(start, queue_limit, i.Client)

	go sq.Start()
	prev := time.Now()
	for {
		bd := sq.PopNext()
		if bd == nil {
			if sq.GetCurrent() > end {
				break
			}
			continue
		}
		bc := MakeBlockContext(i.storage, i.Client, i.oracle, i.config)
		dc, tc, err := bc.process(bd)
		if err != nil {
			return err
		}

		if bd.Pbft.Number%100 == 0 {
			log.WithFields(log.Fields{"period": bd.Pbft.Number, "elapsed_ms": time.Since(prev).Milliseconds()}).Info("Syncing: block applied")
			prev = time.Now()
		}
		log.WithFields(log.Fields{"period": bd.Pbft.Number, "dags": dc, "trxs": tc}).Debug("Syncing: block processed")
	}
	log.WithFields(log.Fields{"period": end}).Info("Syncing: finished")
	return nil
}

func (i *Indexer) run() error {
	err := i.sync()
	if err != nil {
		return err
	}
	log.Info("Syncing: finished, subscribing to new blocks")
	ch, sub, err := i.Client.SubscribeNewHeads()
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case blk := <-ch:
			finalized_period := i.storage.GetFinalizationData().PbftCount
			if blk.Number != finalized_period+1 {
				err := i.sync()
				if err != nil {
					return err
				}
				continue
			}
			if blk.Number < finalized_period {
				log.WithFields(log.Fields{"finalized": finalized_period, "received": blk.Number}).Warn("Received block number is lower than finalized. Node was reset?")
				continue
			}
			// We need to get block from API one more time if we doesn't have it in object from subscription
			var bd *chain.BlockData
			if blk.Transactions == nil {
				bd, err = chain.GetBlockData(i.Client, blk.Number)
			} else {
				bd, err = chain.GetBlockDataFromPbft(i.Client, &blk)

			}
			if err != nil {
				return err
			}

			bc := MakeBlockContext(i.storage, i.Client, i.oracle, i.config)
			dc, tc, err := bc.process(bd)
			if err != nil {
				return err
			}
			// perform consistency check on blocks from subscription
			if i.consistency_check_available {
				i.consistencyCheck(bc.finalized)
			}
			if blk.Number%100 == 0 {
				log.WithFields(log.Fields{"period": blk.Number, "dags": dc, "trxs": tc}).Info("Block processed")
			} else {
				log.WithFields(log.Fields{"period": blk.Number, "dags": dc, "trxs": tc}).Debug("Block processed")
			}
		}
	}
}

func (i *Indexer) consistencyCheck(finalized *storage.FinalizationData) {
	remote_stats, stats_err := i.Client.GetChainStats()
	if stats_err == nil {
		finalized.Check(remote_stats)
	}
}
