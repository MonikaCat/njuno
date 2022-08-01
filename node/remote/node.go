package remote

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/MonikaCat/njuno/node"

	remoteConfig "github.com/MonikaCat/njuno/node/remote/config"
	httpclient "github.com/tendermint/tendermint/rpc/client/http"
	jsonrpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
)

var (
	_ node.Node = &Node{}
)

// Node implements a wrapper around both a Tendermint RPCConfig client and a
// chain SDK REST client that allows for essential data queries.
type Node struct {
	ctx      context.Context
	codec    codec.Marshaler
	client   *httpclient.HTTP
	RPCNode  string // RPC node
	RESTNode string // REST node
}

// NewNode allows to build a new Node instance
func NewNode(cfg *remoteConfig.Details, codec codec.Marshaler) (*Node, error) {
	httpClient, err := jsonrpcclient.DefaultHTTPClient(cfg.RPC.Address)
	if err != nil {
		return nil, err
	}

	// Tweak the transport
	httpTransport, ok := (httpClient.Transport).(*http.Transport)
	if !ok {
		return nil, fmt.Errorf("invalid HTTP Transport: %T", httpTransport)
	}
	httpTransport.MaxConnsPerHost = cfg.RPC.MaxConnections

	rpcClient, err := httpclient.NewWithClient(cfg.RPC.Address, "/websocket", httpClient)
	if err != nil {
		return nil, err
	}

	err = rpcClient.Start()
	if err != nil {
		return nil, err
	}

	return &Node{
		ctx:      context.Background(),
		codec:    codec,
		client:   rpcClient,
		RPCNode:  cfg.RPC.Address,
		RESTNode: cfg.REST.Address,
	}, nil
}

// Stop implements node.Node
func (cp *Node) Stop() {
	err := cp.client.Stop()
	if err != nil {
		panic(fmt.Errorf("error while stopping proxy: %s", err))
	}
}
