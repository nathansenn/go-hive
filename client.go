package gohive

import (
	"fmt"

	rpc "github.com/ybbus/jsonrpc"
)

// Caller interface is used for testing purposes.
type Caller interface {
	CallRaw(*rpc.RPCRequest) (*rpc.RPCResponse, error)
}

// Client is used to pass data into unexposed functions.
// When defining a new JSONrpc use the `NewClient()` function for Hive API defaults.
// To specify an api endpoint execute `NewClient()` with a full URL.
type Client struct {
	URL    string
	Client Caller
}

// NewClient creates an struct with Hive defaults.
// If wish to use a different Hive endpoint (or a different Graphene blockchain
// pass the URL as a parameter. Otherwise leave the parameters empty.
// If more than one URL is entered, only the first will be used.
// Example:
// hive := NewClient()
func NewClient(URL ...string) *Client {
	c := &Client{
		URL:    "https://api.hive.blog",
		Client: rpc.NewClient("https://api.hive.blog"),
	}

	if len(URL) > 0 {
		c.URL = URL[0]
		c.Client = rpc.NewClient(URL[0])
	}
	return c
}

// GetAccountData retrieves the data requested by a method of type Client.
func (c *Client) getAccountData(method string, inputParams ...interface{}) (*rpc.RPCResponse, error) {
	request := rpc.NewRequest(method, inputParams)

	resp, err := c.Client.CallRaw(request)
	if err != nil {
		return nil, fmt.Errorf("json rpc call error: %s", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("rpc response returned: %s", resp.Error)
	}
	return resp, nil
}
