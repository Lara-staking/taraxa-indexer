package oracle

import (
	"math/big"
	"strconv"

	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type NodeData = apy_oracle.IApyOracleNodeData

type YieldedValidator struct {
	Account           common.Address
	Rank              uint16
	Rating            uint64
	Yield             string
	Commisson         *uint64
	RegistrationBlock uint64
	PbftCount         uint64
}

type RawValidator struct {
	Address common.Address
	Yield   string
}

func (r *RawValidator) ToYieldedValidator() YieldedValidator {
	return YieldedValidator{
		Account: r.Address,
		Yield:   r.Yield,
	}
}

func (y *YieldedValidator) ToRawValidator() RawValidator {
	return RawValidator{
		Address: y.Account,
		Yield:   y.Yield,
	}
}

func (v *YieldedValidator) ToNodeData(currentBlock uint64) NodeData {
	vRating, from, to := v.calculateRating(currentBlock)
	yield, err := strconv.ParseFloat(v.Yield, 64)
	if err != nil {
		log.Fatalf("Failed to parse yield: %v", err)
	}

	return NodeData{
		Rating:    big.NewInt(int64(vRating.Score)),
		Account:   v.Account,
		Rank:      v.Rank,
		Apy:       uint16(yield * 1000),
		FromBlock: from,
		ToBlock:   to,
	}
}

func (v *YieldedValidator) SaveRating(batch storage.Batch, currentBlock uint64) {
	vRating, _, _ := v.calculateRating(currentBlock)
	model := vRating.ToModel()
	batch.AddToBatch(model, v.Account.Hex(), vRating.BlockHeight)
}

func (validator *YieldedValidator) calculateRating(currentBlock uint64) (vRating storage.ValidatorRating, from, to uint64) {
	commission_float := float64(*validator.Commisson)
	yield_float, err := strconv.ParseFloat(validator.Yield, 64)
	if err != nil {
		log.Fatalf("Failed to parse yield: %v", err)
	}
	commission_percentage := commission_float / float64(100000)
	adjusted_apy := (1 - commission_percentage) * yield_float * 100
	continuity := float64(validator.PbftCount) / float64(currentBlock-validator.RegistrationBlock)

	//w1 * (APY) - (Commission * w2) + w3 * Continuity + w4 * stake
	score := float64(0.4)*adjusted_apy - float64(0.1)*commission_float + float64(0.5)*continuity
	log.WithFields(log.Fields{"validator": validator.Account.String(), "currentBlock": currentBlock, "score": score, "continuity": continuity, "apy": adjusted_apy, "commission": commission_float}).Info("Validator score")

	vRating = storage.ValidatorRating{
		Address:      validator.Account.Hex(),
		BlockHeight:  currentBlock,
		Score:        uint64(score * 1000),
		Adjusted_apy: adjusted_apy,
		Commission:   commission_float,
		Continuity:   continuity,
	}

	return vRating, validator.RegistrationBlock, currentBlock
}
