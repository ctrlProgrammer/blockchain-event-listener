package main

import (
	connectors "EVM/EventListener/core"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func listenEvent(event *connectors.ConnectorEventListener, log types.Log) {
	fmt.Println("New log call " + event.Name)
}

func main() {
	listeners := make(map[string]connectors.ConnectorEventListener)

	listeners["USDTransfer"] = connectors.ConnectorEventListener{
		Name:  "USDTransfer",
		Topic: common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
	}

	arbitrumRPC := "wss://arb-mainnet.g.alchemy.com/v2/6QUL7GY-FZKzuH6OouFvQHbhuqYugJQ-"
	arbConnector, _ := connectors.NewConnector(arbitrumRPC)

	arbConnector.SetConnectorCallback(listenEvent)
	arbConnector.SetConnectorListeners(listeners)
	arbConnector.SetConnectorContracts([]common.Address{common.HexToAddress("0xB8981C1E85f5Acfbf1760Cb4DB3933526d8a269e")})

}
