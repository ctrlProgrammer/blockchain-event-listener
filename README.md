# EVM chain listeners

A way to create event listeners on all the EVM chains with only the topic and the contract address

### Steps

1. Create contracts variable, where you will storage all the contracts addresses to listen. It must be an array on Addresses

```
contracts := []common.Address{common.HexToAddress("0xB8981C1E85f5Acfbf1760Cb4DB3933526d8a269e")}
```

2. Create listeners. Here you must create a map with a string key and a Connector Event Listener struct where you will storage the name of the event and the topic of the log in the blockchain

```
listeners := make(map[string]core.ConnectorEventListener)

listeners["USDTRANSFER] = core.ConnectorEventListener{
    Name: "USDTRANSFER",
    Topic: common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}
```

3. Create network using the name of the network, the RPC, the listeners, the contracts and a function to receive all the related events

4. Start reading the related events

In the main.go file you can find a tiny example about how you can start listening events on the Arbitrum Sepolia network

### Team

1. CtrlProgrammer
