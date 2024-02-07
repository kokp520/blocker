package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var privateKeyString1 = "2468e962df77168d366c261b792d534be0cc778b96032e28993fb405037e4d58"
var privateKeyString0 = "02f539cddb68bb83335828675829229539d50737b327412713332dd18c83eaa9"
var rpcURL = "http://127.0.0.1:7545"

func main() {
	// 解析私钥1
	privateKey1, err := crypto.HexToECDSA(privateKeyString1)
	if err != nil {
		log.Fatalf("Failed to parse private key 1: %v", err)
	}

	// 解析私钥0
	privateKey0, err := crypto.HexToECDSA(privateKeyString0)
	if err != nil {
		log.Fatalf("Failed to parse private key 0: %v", err)
	}

	// 连接到以太坊节点
	client, err := rpc.DialContext(context.Background(), rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	ethClient := ethclient.NewClient(client)

	// 获取发送者地址（私钥1对应的地址）
	fromAddress1 := crypto.PubkeyToAddress(privateKey1.PublicKey)

	// 获取接收者地址（私钥0对应的地址）
	toAddress0 := crypto.PubkeyToAddress(privateKey0.PublicKey)

	// 获取私钥1对应地址的nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress1)
	if err != nil {
		log.Fatalf("Failed to retrieve nonce: %v", err)
	}

	// 获取Gas Price
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve gas price: %v", err)
	}

	// 构建转账交易
	value := big.NewInt(10) // 转账金额，单位为wei
	gasLimit := uint64(21000) // 交易的燃气限制
	tx := types.NewTransaction(nonce, toAddress0, value, gasLimit, gasPrice, nil)

	// 对交易进行签名（使用私钥1进行签名）
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve network ID: %v", err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey1)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// 发送交易

	// todo 現在這遍 Failed to send transaction: Invalid signature v value 20240207 這邊先擱置
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s", signedTx.Hash().Hex())
}
