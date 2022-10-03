package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to a geth node (when using Infura, you need to use your own API key)
	conn, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/INPUT_YOUR_INFURA_KEY_HERE")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	contractAddress := common.HexToAddress("PLEASE_INPUT_ERC20_TOKEN_ADDRESS_HERE")
	token, err := NewErc20(contractAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	// Access token properties
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)

	if err != nil {
		log.Fatal(err)
	}

	fromAddresses := []common.Address{}

	toAddresses := []common.Address{
		common.HexToAddress("PLEASE_INPUT_TO_ADDRESSES_HERE"),
	}

	transferEvents := make(chan *Erc20Transfer)

	sub, err := token.WatchTransfer(nil, transferEvents, fromAddresses, toAddresses)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case te := <-transferEvents:
			fmt.Printf("%s -> %s : %v", te.From, te.To, te.Value)
		}
	}
}
