package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"smartcontracts/voting"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545") // Connect to Ethereum node
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to Ethereum node", err)
	}

	auth := getAccountAuth(client, "8ae7b3e1bb13eac8a728c94941175239bc136f9085ea5c105cb7cc43b205c719")

	// contractAddress := common.HexToAddress("0x8496E804ba5Da430bd5aA13bb44f8dEdfECCfD97") // Address of deployed contract

	deployedContractAddress, tx, instance, err := voting.DeployVoting(auth, client)
	if err != nil {
		fmt.Println("Unable to bind to deployed instance of contract", err)
	}

	_ = tx
	_ = instance

	fmt.Println("Contract deployed to ", deployedContractAddress.Hex())

	n, err := instance.GetVoteCount(&bind.CallOpts{}, "red")
	if err != nil {
		fmt.Println("Error getting vote count", err)
	}
	fmt.Println("Vote count for red is ", n)

	opts, err := instance.VoteForColor(auth, "red")
	if err != nil {
		fmt.Println("Error voting for red", err)
		return
	}
	fmt.Println("Vote for red", opts)

	n, err = instance.GetVoteCount(&bind.CallOpts{}, "red")
	if err != nil {
		fmt.Println("Error getting vote count", err)
	}
	fmt.Println("Vote count for adsad red is ", n)

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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get latest block: %v", err)
	}

	maxFeePerGas := new(big.Int).Add(header.BaseFee, big.NewInt(1e9)) // 1 Gwei
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasFeeCap = maxFeePerGas
	auth.GasTipCap = big.NewInt(1e9) // 1 Gwei

	return auth
}
