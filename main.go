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

	contracts := []common.Address{common.HexToAddress("0xB8981C1E85f5Acfbf1760Cb4DB3933526d8a269e")}

	listeners := make(map[string]core.ConnectorEventListener)

	listeners["USDTransfer"] = core.ConnectorEventListener{
		Name:  "USDTransfer",
		Topic: common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
	}

	arbNetwork, err := core.NewNetwork("Arbitrum Sepolia", os.Getenv("SEPOLIA_RPC"), listeners, contracts, listenEvent)

	if err != nil {
		log.Fatal(err)
		return
	}

	arbNetwork.Start()

	select {}
}
