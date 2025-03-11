package lara

import (
	"context"
	"strings"

	multicall_contract "github.com/Taraxa-project/taraxa-indexer/abi/multicall"
	"github.com/Taraxa-project/taraxa-indexer/internal/transact"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	// estimate gas first

	gas, err := state.estimateGasForMulticall(multicall)
	if err != nil {
		log.Errorf("Failed to estimate gas: %v", err)
		return nil, err
	}

	if gas > state.gasLimit {
		log.Infof("Gas limit exceeded, breaking into chunks")
		// break the multicall into smaller chunks
		chunks, err := state.splitMulticall(multicall, state.gasLimit)
		if err != nil {
			log.Errorf("Failed to split multicall: %v", err)
			return nil, err
		}
		for _, chunk := range chunks {
			receipt, err := state.Multicall(chunk)
			if err != nil {
				return nil, err
			}
			log.Infof("Processed chunk with receipt: %v", receipt)
		}
	} else {
		tx, err := state.contract.Aggregate(state.signer, multicall)
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
	return nil, nil
}

// Define the splitMulticall function
func (state *MulticallContract) splitMulticall(multicall []multicall_contract.MulticallCall, gasLimit uint64) ([][]multicall_contract.MulticallCall, error) {
	var chunks [][]multicall_contract.MulticallCall
	var currentChunk []multicall_contract.MulticallCall
	var currentGas uint64

	for _, call := range multicall {
		// Estimate gas for the current call (this is a placeholder, replace with actual gas estimation logic)
		callGas, err := state.estimateGasForCall(call)
		if err != nil {
			return nil, err
		}

		if currentGas+callGas > gasLimit {
			// If adding this call exceeds the gas limit, start a new chunk
			chunks = append(chunks, currentChunk)
			currentChunk = []multicall_contract.MulticallCall{}
			currentGas = 0
		}

		currentChunk = append(currentChunk, call)
		currentGas += callGas
	}

	// Add the last chunk if it's not empty
	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}
	log.Infof("Split multicall into %d chunks", len(chunks))

	return chunks, nil
}

func (state *MulticallContract) estimateGasForMulticall(multicall []multicall_contract.MulticallCall) (uint64, error) {
	abi, err := abi.JSON(strings.NewReader(multicall_contract.MulticallContractABI))
	if err != nil {
		log.Errorf("Failed to parse ABI: %v", err)
		return 0, err
	}
	data, err := abi.Pack("aggregate", multicall)
	if err != nil {
		log.Errorf("Failed to pack data: %v", err)
		return 0, err
	}
	deployAddress := common.HexToAddress(state.deploymentAddress)
	gas, err := state.Eth.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &deployAddress,
		Data: data,
	})
	if err != nil {
		log.Errorf("Failed to estimate gas: %v", err)
		return 0, err
	}
	return gas, nil
}

func (state *MulticallContract) estimateGasForCall(call multicall_contract.MulticallCall) (uint64, error) {
	abi, err := abi.JSON(strings.NewReader(multicall_contract.MulticallContractABI))
	if err != nil {
		return 0, err
	}
	data, err := abi.Pack("aggregate", call)
	if err != nil {
		return 0, err
	}
	deployAddress := common.HexToAddress(state.deploymentAddress)
	gas, err := state.Eth.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &deployAddress,
		Data: data,
	})
	if err != nil {
		return 0, err
	}
	return gas, nil
}
