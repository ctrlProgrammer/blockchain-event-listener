package connectors

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ConnectorEventListener struct {
	Name  string
	Topic common.Hash
}

type Connector struct {
	RPC                *string
	ConnectorClient    *ethclient.Client
	ConnectorContracts []common.Address
	ConnectorListeners map[string]ConnectorEventListener
	ConnectorCallback  func(eventName string, log types.Log)
}

func (x *Connector) invalidConnectionError() error {
	return errors.New("invalid connection")
}

func (x *Connector) CreateConnection() error {
	if x.RPC == nil {
		return errors.New("invalid rpc connection")
	}

	client, err := ethclient.Dial(*x.RPC)

	if err != nil {
		return err
	}

	x.ConnectorClient = client

	return nil
}

func (x *Connector) isValidConnection() bool {
	return x.RPC != nil && x.ConnectorClient != nil
}

func (x *Connector) GetBalance(address common.Address) (*big.Int, error) {
	if !x.isValidConnection() {
		return nil, x.invalidConnectionError()
	}

	balance, err := x.ConnectorClient.BalanceAt(context.Background(), address, nil)

	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (x *Connector) SetConnectorContracts(contracts []common.Address) {
	x.ConnectorContracts = contracts
}

func (x *Connector) SetConnectorListeners(events map[string]ConnectorEventListener) {
	x.ConnectorListeners = events
}

func (x *Connector) SetConnectorCallback(callback func(eventName string, log types.Log)) {
	x.ConnectorCallback = callback
}

func (x *Connector) ConnectWithEvents() {
	query := ethereum.FilterQuery{
		Addresses: x.ConnectorContracts,
	}

	logs := make(chan types.Log)

	sub, err := x.ConnectorClient.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	var i int

	go func(i int, wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			select {
			case <-sub.Err():
				fmt.Println("Error reading blockchain logs... passing")
			case vLog := <-logs:
				if x.isValidEventTopic(vLog.Topics[0]) {
					x.evaluateEvent(vLog)
				}
			}
		}
	}(i, &wg)

	wg.Wait()
}

func (x *Connector) isValidEventTopic(topic common.Hash) bool {
	for _, v := range x.ConnectorListeners {
		if v.Topic == topic {
			return true
		}
	}

	return false
}

func (x *Connector) evaluateEvent(log types.Log) {

}

func NewConnector(rpc string) (*Connector, error) {
	connector := Connector{
		RPC: &rpc,
	}

	err := connector.CreateConnection()

	if err != nil {
		return nil, err
	}

	return &connector, nil
}
