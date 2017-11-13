package client

import (
	"errors"

	"github.com/ybbus/jsonrpc"
)

var ErrEmptyResponse = errors.New("received empty response")

// MopidyClient allows sending RPC commands to a Mopidy server
type MopidyClient struct {
	cli *jsonrpc.RPCClient
}

// NewMopidyClient returns a new mopidy client for the given endpoint
func NewMopidyClient(endpoint string) *MopidyClient {
	return &MopidyClient{
		cli: jsonrpc.NewRPCClient(endpoint),
	}
}

// Call calls a method on Mopidy and returns the result
func (cli *MopidyClient) Call(method string, params interface{}, res interface{}) error {
	var resp *jsonrpc.RPCResponse
	var err error

	if params != nil {
		resp, err = cli.cli.Call(method, params)
	} else {
		resp, err = cli.cli.Call(method)
	}

	if err != nil {
		return err
	}

	if resp.Result == nil && res != nil {
		return ErrEmptyResponse
	}

	if res != nil {
		return resp.GetObject(&res)
	}

	return nil
}
