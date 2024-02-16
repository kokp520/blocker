package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// index 0
var privateKeyString = "dd64527f564b4a86e896deb659c34b449ca1b67aecf2198b0ef4845c22cdccfa"
var rpcURL = "http://127.0.0.1:8545"
// index 1
var recipient_address = "18f57E1eD403A39C664dde2ca6386D50cBe45E30"

func main() {
	// get 私鑰 ecc 橢圓加密xyz
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// connect 到以太節點
	client, err := rpc.DialContext(context.Background(), rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	ethClient := ethclient.NewClient(client)

	// get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}

	// get address from sending person
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// retrieve nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to retrieve nonce: %v", err)
	}
	value := big.NewInt(10) // xxx ETH
	gasLimit := uint64(21000) // gas fee
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve gas price: %v", err)
	}

	toAddress := common.HexToAddress(recipient_address) // 接收者地址

	// create transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// sign for transacation
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to retrieve network ID: %v", err)
	}


	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// send transacation -> finish
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s", signedTx.Hash().Hex())
}
