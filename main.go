package main

import (
	connectors "EVM/EventListener/core"
)

func main() {
	arbitrumRPC := "wss://arb-mainnet.g.alchemy.com/v2/6QUL7GY-FZKzuH6OouFvQHbhuqYugJQ-"
	arbConnector, _ := connectors.NewConnector(arbitrumRPC)
}
