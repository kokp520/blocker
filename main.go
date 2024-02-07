package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)


var testUrl = "https://127.0.0.1:8545"

func main(){
    client, err := ethclient.DialContext(context.Background(), testUrl)
    if err != nil {
        log.Fatal("ERROR %+v", err)
    }
    // fmt.Printf("start clinet :%+v", client)

    defer client.Close()
    
    block, err := client.BlockByNumber(context.Background(), nil)
    if err != nil{
        log.Fatal("error to get a block %v", err)
    }

    fmt.Println("block number %v", block.Number())
}