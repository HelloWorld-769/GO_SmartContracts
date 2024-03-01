package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"smartcontracts/api"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func main() {
	// address of etherum env
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	auth := getAccountAuth(client, "cf9d8a7f77148b8724ec8aafb931e30acdbefe60d34b954e2a0e5f879b13a7ef")

	fmt.Println("Auth creation sucessful")

	deployedContractAddress, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(deployedContractAddress.Hex())

	_, _ = instance, tx
	fmt.Println("instance->", instance)
	fmt.Println("tx->", tx.Hash().Hex())

	conn, err := api.NewApi(common.HexToAddress(deployedContractAddress.Hex()), client)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/balance", func(c *gin.Context) {
		// usecase
		reply, err := conn.Balance(&bind.CallOpts{})
		fmt.Println("balance")
		if err != nil {
			fmt.Println("Error is", err)
			return
		}
		c.JSON(http.StatusOK, reply)
	})

	r.GET("/admin", func(c *gin.Context) {
		reply, err := conn.Admin(&bind.CallOpts{})
		if err != nil {
			fmt.Println("Error is", err)

		}
		c.JSON(http.StatusOK, reply)
	})

	r.POST("/deposite/:amount", func(ctx *gin.Context) {
		amount := ctx.Param("amount")

		amt, _ := strconv.Atoi(amount)

		//gets address of account by which amount to be deposite
		var v map[string]interface{}
		err := json.NewDecoder(ctx.Request.Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		//creating auth object for above account
		auth := getAccountAuth(client, v["0x001d5bdf425217c8cb2b7597512d371e5465f3076a126679885aa83d7d31c456"].(string))

		reply, err := conn.Deposite(auth, big.NewInt(int64(amt)))
		if err != nil {
			fmt.Println("Error is", err)
		}
		ctx.JSON(http.StatusOK, reply)
	})

	r.POST("/withdrawl/:amount", func(c *gin.Context) {
		amount := c.Param("amount")
		amt, _ := strconv.Atoi(amount)

		var v map[string]interface{}
		err := json.NewDecoder(c.Request.Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		auth := getAccountAuth(client, v["0x001d5bdf425217c8cb2b7597512d371e5465f3076a126679885aa83d7d31c456"].(string))
		// auth.Nonce.Add(auth.Nonce, big.NewInt(int64(1))) //it is use to create next nounce of account if it has to make another transaction

		reply, err := conn.Withdrawl(auth, big.NewInt(int64(amt)))
		if err != nil {
			fmt.Println("Error is", err)
		}
		c.JSON(http.StatusOK, reply)
	})

	r.Run(":1323")
}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("chainID=", chainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	fmt.Println("auth=", auth)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
