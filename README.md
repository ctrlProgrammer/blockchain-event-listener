# EVM chain listeners

A way to create event listeners on all the EVM chains with only the topic and the contract address

### Steps

1. Create contracts variable, where you will storage all the contracts addresses to listen. It must be an array on Addresses

2. Create listeners. Here you must create a map with a string key and a Connector Event Listener struct where you will storage the name of the event and the topic of the log in the blockchain

3. Create network using the name of the network, the RPC, the listeners, the contracts and a function to receive all the related events

4. Start reading the related events

In the main.go file you can find a tiny example about how you can start listening events on the Arbitrum Sepolia network

### Team

1. CtrlProgrammer
