package rewards

import (
	"math/big"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var multiplier = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)
var percentage_multiplier = big.NewInt(10000)
var YieldFractionDecimalPrecision = big.NewInt(1e+6)

func (r *Rewards) calculateCurrentYield(current_total_tara_supply *big.Int) *big.Int {
	// Current yield = (max supply - current total supply) / current total supply
	current_yield := big.NewInt(0).Sub(r.config.Chain.Hardforks.AspenHf.MaxSupply, current_total_tara_supply)
	current_yield.Mul(current_yield, YieldFractionDecimalPrecision)
	current_yield.Div(current_yield, current_total_tara_supply)

	return current_yield
}

func GetMultipliedYield(reward, stake *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(reward, multiplier)
	r.Div(r, stake)

	return r
}

func GetValidatorsYield(rewards map[string]*big.Int, validators *Validators) []storage.ValidatorYield {
	ret := make([]storage.ValidatorYield, 0, len(validators.validators))
	for v_addr, v := range validators.validators {
		if v.TotalStake.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		if rewards[v_addr] != nil {
			ret = append(ret, storage.ValidatorYield{Validator: v_addr, Yield: GetMultipliedYield(rewards[v_addr], v.TotalStake)})
		}
	}

	return ret
}

func GetYieldForInterval(yields_sum, blocks_per_year *big.Int, elem_count int64) float64 {
	res := big.NewInt(0)
	res.Mul(yields_sum, blocks_per_year)
	res.Mul(res, percentage_multiplier)
	res.Div(res, big.NewInt(int64(elem_count)))
	res.Div(res, multiplier)

	ret := float64(res.Uint64())
	ret /= float64(percentage_multiplier.Uint64())
	return ret
}
func (r *Rewards) processIntervalYield(batch storage.Batch) {
	sum := big.NewInt(0)
	storage.ProcessIntervalData(r.storage, r.blockNum-r.config.TotalYieldSavingInterval, func(key []byte, o storage.MultipliedYield) (stop bool) {
		sum.Add(sum, o.Yield)
		batch.Remove([]byte(key))
		return false
	})

	yield := GetYieldForInterval(sum, r.config.Chain.BlocksPerYear, int64(r.config.TotalYieldSavingInterval))
	log.WithFields(log.Fields{"total_yield": yield}).Info("processIntervalYield")
	batch.AddSingleKey(&storage.Yield{Yield: common.FormatFloat(yield)}, storage.FormatIntToKey(r.blockNum))
}

func (r *Rewards) processValidatorsIntervalYield(batch storage.Batch) {
	start := uint64(0)
	if r.blockNum > r.config.ValidatorsYieldSavingInterval {
		start = r.blockNum - r.config.ValidatorsYieldSavingInterval
	}

	sum_by_validator := make(map[string]*big.Int)

	storage.ProcessIntervalData(r.storage, start, func(key []byte, o storage.ValidatorsYield) (stop bool) {
		for _, y := range o.Yields {
			if sum_by_validator[y.Validator] == nil {
				sum_by_validator[y.Validator] = big.NewInt(0)
			}
			sum_by_validator[y.Validator].Add(sum_by_validator[y.Validator], y.Yield)
		}
		batch.Remove(key)
		return false
	})

	log.WithFields(log.Fields{"validators": len(sum_by_validator)}).Info("processValidatorsIntervalYield")
	yields := make([]oracle.RawValidator, 0, len(sum_by_validator))
	for val, sum := range sum_by_validator {
		yield := GetYieldForInterval(sum, r.config.Chain.BlocksPerYear, int64(r.config.ValidatorsYieldSavingInterval))
		log.WithFields(log.Fields{"validator": val, "yield": yield}).Info("processValidatorsIntervalYield")
		batch.Add(&storage.Yield{Yield: common.FormatFloat(yield)}, val, r.blockNum)
		if yield > 0 {
			yields = append(yields, oracle.RawValidator{Yield: common.FormatFloat(yield), Address: ethcommon.HexToAddress(val)})
		}
	}
	// currentBlock := r.storage.GetFinalizationData().PbftCount
	// chainHead, err := r.oracle.Eth.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.WithError(err).Fatal("Failed to get chain head")
	// }
	// if currentBlock < chainHead.Number.Uint64()-50 {
	// 	log.WithFields(log.Fields{"currentBlock": currentBlock, "chainHead": chainHead.Number.Uint64()}).Warn("Current block not close to chain head, skipping Oracle push")
	// 	return
	// }
	go func() {
		nonFullCommissionValidators := []string{
			"0x386fd3885c05b30927be3f3cca08dfbc20fa7159",
			"0x17d442e3e16e8df269e5970e7099cf37a433f300",
			"0x87acc7c55916285646e6509fcb9154a8b34cda25",
			"0x4f1c855b7c23f0632bafa1db178610c852d01c25",
			"0xf302a2808c8b0998992ed59ee5210c95f8c6cb52",
			"0x523c52ea0d5180fb9e47572b889cff04e8758a89",
		}
		// remove all full commission validators from the list
		filteredYields := []oracle.RawValidator{}
		for _, v := range yields {
			if slices.Contains(nonFullCommissionValidators, strings.ToLower(v.Address.Hex())) {
				filteredYields = append(filteredYields, v)
			}
		}

		log.WithFields(log.Fields{"filteredYields": len(filteredYields)}).Info("filteredYields has been filtered")

		// sort yields by yield
		sort.Slice(filteredYields, func(i, j int) bool {
			yieldI, _ := strconv.ParseFloat(filteredYields[i].Yield, 32)
			yieldJ, _ := strconv.ParseFloat(filteredYields[j].Yield, 32)
			return yieldI > yieldJ
		})
		r.oracle.PushValidators(filteredYields)
	}()
}
