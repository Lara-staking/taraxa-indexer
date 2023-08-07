package migration

import (
	"bytes"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

type OldTransaction struct {
	BlockNumber      models.Counter         `json:"blockNumber"`
	Calldata         *models.CallData       `json:"calldata,omitempty" rlp:"nil"`
	From             models.Address         `json:"from"`
	GasPrice         models.Counter         `json:"gasPrice"`
	GasUsed          models.Counter         `json:"gasUsed"`
	Hash             models.Hash            `json:"hash"`
	Input            string                 `json:"input"`
	Nonce            models.Counter         `json:"nonce"`
	Status           bool                   `json:"status"`
	Timestamp        models.Timestamp       `json:"timestamp"`
	To               models.Address         `json:"to"`
	TransactionIndex models.Counter         `json:"transactionIndex"`
	Type             models.TransactionType `json:"type"`
	Value            string                 `json:"value"`
}

// RemoveSenderMigration is a migration that removes the Sender attribute from the Dag struct.
type RemoveNonceTxIndexAddFeeMigration struct {
	id string
}

func (m *RemoveNonceTxIndexAddFeeMigration) GetId() string {
	return m.id
}

// Apply is the implementation of the Migration interface for the RemoveSenderMigration.
func (m *RemoveNonceTxIndexAddFeeMigration) Apply(s *pebble.Storage) error {
	const MAX_TX_THRESHOLD = 1000
	batch := s.NewBatch()
	var last_key []byte

	for {
		count := 0
		s.ForEachFromKey([]byte(pebble.GetPrefix(models.Transaction{})), last_key, func(key, res []byte) (stop bool) {
			// don't apply it second time on the same key and don't apply it on total supply as its starting with the same prefix
			if bytes.Equal(key, last_key) {
				return false
			}
			var o OldTransaction
			err := rlp.DecodeBytes(res, &o)
			if err != nil {
				if err.Error() == "rlp: too few elements for migration.OldTransaction" {
					return false
				}
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error decoding Transaction")
			}
			tx := models.Transaction{
				BlockNumber: o.BlockNumber,
				Calldata:    o.Calldata,
				From:        o.From,
				GasCost:     o.GasPrice * o.GasUsed,
				Hash:        o.Hash,
				Input:       o.Input,
				Status:      o.Status,
				Timestamp:   o.Timestamp,
				To:          o.To,
				Type:        o.Type,
				Value:       o.Value,
			}
			err = batch.AddToBatchFullKey(&tx, key)
			if err != nil {
				log.WithFields(log.Fields{"migration": m.id, "error": err}).Fatal("Error adding Transaction to batch")
			}

			last_key = key
			count++
			return count == MAX_TX_THRESHOLD
		})
		batch.CommitBatch()
		batch = s.NewBatch()
		if count < MAX_TX_THRESHOLD {
			break
		}
	}

	return nil
}