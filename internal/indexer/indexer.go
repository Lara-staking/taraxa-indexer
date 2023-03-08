package indexer

import (
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	log "github.com/sirupsen/logrus"
)

type Indexer struct {
	retry_time time.Duration
	client     *chain.WsClient
	storage    *storage.Storage
}

func MakeAndRun(url string, storage *storage.Storage) {
	i := NewIndexer(url, storage)
	for {
		err := i.run()
		f := i.storage.GetFinalizationData()
		log.WithError(err).WithField("period", f.PbftCount).Error("Error occurred. Restart indexing")
		time.Sleep(i.retry_time)
	}
}

func NewIndexer(url string, storage *storage.Storage) (i *Indexer) {
	var err error
	i = new(Indexer)
	i.retry_time = 5 * time.Second
	i.storage = storage
	for {
		i.client, err = chain.NewWsClient(url)
		if err == nil {
			err = i.init()
		}
		if err == nil {
			break
		}
		log.WithError(err).Error("Can't connect to chain")
		time.Sleep(i.retry_time)
	}
	return
}

func (i *Indexer) init() error {
	genesis_blk, err := i.client.GetBlockByNumber(0)
	if err != nil {
		return err
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
	if !db_clean {
		return nil
	}

	genesis, err := MakeGenesis(i.storage, i.client, remote_hash)
	if err != nil {
		return err
	}
	// Genesis hash and finalized period(0) is set inside
	genesis.process()
	return nil
}

func (i *Indexer) sync() (err error) {
	// start processing blocks from the next one
	start := i.storage.GetFinalizationData().PbftCount + 1
	end, p_err := i.client.GetLatestPeriod()
	if p_err != nil {
		return p_err
	}
	if start >= end {
		return
	}
	log.WithFields(log.Fields{"start": start, "end": end}).Info("Syncing: start")
	prev := time.Now()
	for p := uint64(start); p <= end; p++ {
		blk, b_err := i.client.GetBlockByNumber(p)
		if b_err != nil {
			return b_err
		}
		if p%100 == 0 {
			log.WithFields(log.Fields{"period": p, "elapsed_ms": time.Since(prev).Milliseconds()}).Info("Syncing: block applied")
			prev = time.Now()
		}

		dc, tc, process_err := MakeBlockContext(i.storage, i.client).process(blk)
		if process_err != nil {
			return process_err
		}
		log.WithFields(log.Fields{"period": p, "dags": dc, "trxs": tc}).Debug("Syncing: block processed")
	}
	return
}

func (i *Indexer) run() error {
	err := i.sync()
	if err != nil {
		return err
	}

	ch, sub, err := i.client.SubscribeNewHeads()
	if err != nil {
		return err
	}
	for {
		select {
		case err := <-sub.Err():
			return err
		case sub_blk := <-ch:
			p := chain.ParseInt(sub_blk.Number)
			if p != i.storage.GetFinalizationData().PbftCount+1 {
				err := i.sync()
				if err != nil {
					return err
				}
				continue
			}
			// We need to get block from API one more time because chain isn't returning transactions in this subscription object
			blk, err := i.client.GetBlockByNumber(p)
			if err != nil {
				return err
			}
			dc, tc, err := MakeBlockContext(i.storage, i.client).process(blk)
			if err != nil {
				return err
			}
			log.WithFields(log.Fields{"period": p, "dags": dc, "trxs": tc}).Info("Block processed")
		}
	}
}
