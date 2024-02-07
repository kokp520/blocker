package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	// this would be your generated smart contract bindings

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// address of etherum env
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(client, "20E773d86834dBD6d8aDaeb0eD3b03221e0BE181")


	fmt.Printf("debug ---- %+v -----\n", auth)
	return
	
	//deploying smart contract
	// address, tx, instance, err := api.DeployApi(auth, client)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(address.Hex())

	// _, _ = instance, tx
	// fmt.Println("instance->", instance)
	// fmt.Println("tx->", tx.Hash().Hex())

	// //creating api object to intract with smart contract function
	// conn, err := api.NewApi(common.HexToAddress(address.Hex()), client)
	// if err != nil {
	// 	panic(err)
	// }

	// e := echo.New()

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.GET("/balance", func(c echo.Context) error {
	// 	reply, err := conn.Balance(&bind.CallOpts{})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusOK, reply)
	// })
	// e.GET("/admin", func(c echo.Context) error {
	// 	reply, err := conn.Admin(&bind.CallOpts{})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusOK, reply)
	// })
	// e.POST("/deposite/:amount", func(c echo.Context) error {
	// 	amount := c.Param("amount")
	// 	amt, _ := strconv.Atoi(amount)

	// 	//gets address of account by which amount to be deposite
	// 	var v map[string]interface{}
	// 	err := json.NewDecoder(c.Request().Body).Decode(&v)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	//creating auth object for above account
	// 	auth := getAccountAuth(client, v["accountPrivateKey"].(string))

	// 	reply, err := conn.Deposite(auth, big.NewInt(int64(amt)))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusOK, reply)
	// })
	// e.POST("/withdrawl/:amount", func(c echo.Context) error {
	// 	amount := c.Param("amount")
	// 	amt, _ := strconv.Atoi(amount)

	// 	var v map[string]interface{}
	// 	err := json.NewDecoder(c.Request().Body).Decode(&v)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	auth := getAccountAuth(client, v["accountPrivateKey"].(string))
	// 	// auth.Nonce.Add(auth.Nonce, big.NewInt(int64(1))) //it is use to create next nounce of account if it has to make another transaction

	// 	reply, err := conn.Withdrawl(auth, big.NewInt(int64(amt)))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// // Start server
	// e.Logger.Fatal(e.Start(":1323"))
}

//function to create auth for any account from its private key
func getAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}