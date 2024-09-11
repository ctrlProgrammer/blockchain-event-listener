package core

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ConnectorEventListener struct {
	Topic common.Hash
	Name  string
}

type Connector struct {
	RPC                *string
	ConnectorClient    *ethclient.Client
	ConnectorContracts []common.Address
	ConnectorListeners map[string]ConnectorEventListener
	ConnectorCallback  func(event *ConnectorEventListener, log types.Log)
}

func (x *Connector) InvalidConnectionError() error {
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

func (x *Connector) IsValidConnection() bool {
	return x.RPC != nil && x.ConnectorClient != nil
}

func (x *Connector) SetConnectorContracts(contracts []common.Address) error {
	if !x.IsValidConnection() {
		return x.InvalidConnectionError()
	}

	x.ConnectorContracts = contracts

	return nil
}

func (x *Connector) SetConnectorListeners(events map[string]ConnectorEventListener) error {
	if !x.IsValidConnection() {
		return x.InvalidConnectionError()
	}

	x.ConnectorListeners = events

	return nil
}

func (x *Connector) SetConnectorCallback(callback func(event *ConnectorEventListener, log types.Log)) error {
	if !x.IsValidConnection() {
		return x.InvalidConnectionError()
	}

	x.ConnectorCallback = callback

	return nil
}

func (x *Connector) ConnectWithEvents() error {
	if !x.IsValidConnection() {
		return x.InvalidConnectionError()
	}

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
				event := x.validateEvent(vLog.Topics[0])

				if event != nil {
					x.evaluateEvent(vLog, event)
				}
			}
		}
	}(i, &wg)

	wg.Wait()

	return nil
}

func (x *Connector) validateEvent(topic common.Hash) *ConnectorEventListener {
	for _, v := range x.ConnectorListeners {
		if v.Topic == topic {
			return &v
		}
	}

	return nil
}

func (x *Connector) evaluateEvent(log types.Log, event *ConnectorEventListener) {
	if x.ConnectorCallback != nil {
		x.ConnectorCallback(event, log)
	}
}

func NewConnector(rpc *string) (*Connector, error) {
	connector := Connector{
		RPC: rpc,
	}

	err := connector.CreateConnection()

	if err != nil {
		return nil, err
	}

	return &connector, nil
}
