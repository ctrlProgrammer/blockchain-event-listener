package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Network struct {
	Name      string
	RPC       *string
	listeners map[string]ConnectorEventListener
	contracts []common.Address
	callback  func(log types.Log, event *ConnectorEventListener, network *Network)
	Connector *Connector
}

func (network *Network) CreateCallback(event *ConnectorEventListener, log types.Log) {
	if network.callback == nil {
		return
	}

	network.callback(log, event, network)
}

func (network *Network) Start() error {
	err := network.Connector.ConnectWithEvents()

	if err != nil {
		return err
	}

	return nil
}

func NewNetwork(name string, rpc string, listeners map[string]ConnectorEventListener, contracts []common.Address, callback func(log types.Log, event *ConnectorEventListener, network *Network)) (*Network, error) {
	connector, err := NewConnector(&rpc)

	if err != nil {
		return nil, err
	}

	net := Network{
		Name:      name,
		RPC:       &rpc,
		listeners: listeners,
		contracts: contracts,
		callback:  callback,
		Connector: connector,
	}

	err = net.Connector.SetConnectorCallback(net.CreateCallback)

	if err != nil {
		return nil, err
	}

	err = net.Connector.SetConnectorContracts(contracts)

	if err != nil {
		return nil, err
	}

	err = net.Connector.SetConnectorListeners(listeners)

	if err != nil {
		return nil, err
	}

	return &net, nil
}
