package core

type Network struct {
	Name string
	RPC  *string
}

func NewNetwork(name string, rpc string) *Network {
	net := Network{
		Name: name,
		RPC:  &rpc,
	}

	return &net
}
