package client

import (
	"errors"

	"github.com/ybbus/jsonrpc"
)

// ErrEmptyResponse is returned if the caller of `MopidyClient.Call` expected a result but Mopidy returned an empty response
var ErrEmptyResponse = errors.New("received empty response")

// MopidyClient allows sending RPC commands to a Mopidy server
type MopidyClient interface {
	Call(method string, params interface{}, res interface{}) error
	CallResult(method string, params interface{}) (*jsonrpc.RPCResponse, error)
}

type mopidyClient struct {
	cli *jsonrpc.RPCClient
}

// NewMopidyClient returns a new mopidy client for the given endpoint
func NewMopidyClient(endpoint string) MopidyClient {
	return &mopidyClient{
		cli: jsonrpc.NewRPCClient(endpoint),
	}
}

// Call calls a method on Mopidy and returns the result
func (cli *mopidyClient) Call(method string, params interface{}, res interface{}) error {
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

	// If the caller was expecting a result, we are going to return an error
	if resp.Result == nil && res != nil {
		return ErrEmptyResponse
	}

	if res != nil {
		return resp.GetObject(&res)
	}

	return nil
}

func (cli *mopidyClient) CallResult(method string, params interface{}) (*jsonrpc.RPCResponse, error) {
	var resp *jsonrpc.RPCResponse
	var err error

	if params != nil {
		resp, err = cli.cli.Call(method, params)
	} else {
		resp, err = cli.cli.Call(method)
	}

	return resp, err
}
