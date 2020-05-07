package fsn

import (
	"errors"
	"math/big"

	"github.com/fsn-dev/crossChain-Bridge/common"
	"github.com/fsn-dev/crossChain-Bridge/common/hexutil"
	"github.com/fsn-dev/crossChain-Bridge/rpc/client"
	"github.com/fsn-dev/crossChain-Bridge/tools/rlp"
	"github.com/fsn-dev/crossChain-Bridge/types"
)

func (b *FsnBridge) GetLatestBlockNumber() (uint64, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result string
	err := client.RpcPost(&result, url, "eth_blockNumber")
	if err != nil {
		return 0, err
	}
	return common.GetUint64FromStr(result)
}

func (b *FsnBridge) GetBlockByHash(blockHash string) (*types.RPCBlock, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result *types.RPCBlock
	err := client.RpcPost(&result, url, "eth_getBlockByHash", blockHash, false)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("block not found")
	}
	return result, nil
}

func (b *FsnBridge) GetTransaction(txHash string) (*types.RPCTransaction, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result *types.RPCTransaction
	err := client.RpcPost(&result, url, "eth_getTransactionByHash", txHash)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("tx not found")
	}
	return result, nil
}

func (b *FsnBridge) GetTransactionReceipt(txHash string) (*types.RPCTxReceipt, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result *types.RPCTxReceipt
	err := client.RpcPost(&result, url, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("tx receipt not found")
	}
	return result, nil
}

func (b *FsnBridge) GetPoolNonce(address string) (uint64, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	account := common.HexToAddress(address)
	var result hexutil.Uint64
	err := client.RpcPost(&result, url, "eth_getTransactionCount", account, "pending")
	return uint64(result), err
}

func (b *FsnBridge) SuggestPrice() (*big.Int, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result hexutil.Big
	err := client.RpcPost(&result, url, "eth_gasPrice")
	if err != nil {
		return nil, err
	}
	return result.ToInt(), nil
}

func (b *FsnBridge) SendSignedTransaction(tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result interface{}
	return client.RpcPost(&result, url, "eth_sendRawTransaction", common.ToHex(data))
}

func (b *FsnBridge) ChainID() (*big.Int, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result hexutil.Big
	err := client.RpcPost(&result, url, "eth_chainId")
	if err != nil {
		return nil, err
	}
	return result.ToInt(), nil
}

func (b *FsnBridge) GetTransactionAndReceipt(txHash string) (*types.RPCTxAndReceipt, error) {
	gateway := b.GatewayConfig
	url := gateway.ApiAddress
	var result *types.RPCTxAndReceipt
	err := client.RpcPost(&result, url, "fsn_getTransactionAndReceipt", txHash)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("tx and receipt not found")
	}
	return result, nil
}
