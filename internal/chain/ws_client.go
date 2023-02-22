package chain

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

// WsClient is a struct that connects to a Taraxa node.
type WsClient struct {
	rpc *rpc.Client
	ctx context.Context
}

// NewWsClient creates a new instance of the WsClient struct.
func NewWsClient(url string) (*WsClient, error) {
	ctx := context.Background()
	client, err := rpc.DialWebsocket(ctx, url, "")
	if err != nil {
		return nil, err
	}
	return &WsClient{rpc: client, ctx: ctx}, nil
}

// Call calls an RPC method on the chain
func (client *WsClient) Call(method string, args ...interface{}) (res map[string]interface{}) {
	err := client.rpc.Call(&res, method, args...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func (client *WsClient) GetBlockByNumber(number uint64) (blk *block) {
	blk = new(block)
	err := client.rpc.Call(blk, "eth_getBlockByNumber", fmt.Sprintf("0x%x", number), false)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func (client *WsClient) GetTransactionByHash(hash string) (trx *transaction) {
	trx = new(transaction)
	err := client.rpc.Call(trx, "eth_getTransactionByHash", hash)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func (client *WsClient) AddTransactionReceiptData(trx *transaction) {
	err := client.rpc.Call(&trx, "eth_getTransactionReceipt", trx.Hash)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (client *WsClient) GetPbftBlockWithDagBlocks(period uint64) (pbftWithDags pbftBlockWithDags) {
	err := client.rpc.Call(&pbftWithDags, "taraxa_getScheduleBlockByPeriod", fmt.Sprintf("0x%x", period))
	if err != nil {
		log.Fatal("GetPbftBlockWithDagBlocks error ", err.Error())
	}
	return
}

func (client *WsClient) GatDagBlockByHash(hash string) (dag dagBlock) {
	err := client.rpc.Call(&dag, "taraxa_getDagBlockByHash", hash, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

// Close disconnects from the node
func (client *WsClient) Close() {
	client.rpc.Close()
}