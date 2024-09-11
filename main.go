package main

import (
	core "EVM/EventListener/core"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joho/godotenv"
)

func listenEvent(log types.Log, event *core.ConnectorEventListener, network *core.Network) {
	fmt.Println("New log call " + event.Name + " from " + network.Name)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	contracts := []common.Address{}

	listeners := make(map[string]core.ConnectorEventListener)

	arbNetwork, err := core.NewNetwork("Arbitrum Sepolia", os.Getenv("SEPOLIA_RPC"), listeners, contracts, listenEvent)

	if err != nil {
		log.Fatal(err)
		return
	}

	arbNetwork.Start()

	select {}
}
