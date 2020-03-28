package accounts

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

// JSONrpc is used to pass data into unexposed functions.
// When defining a new JSONrpc use the `NewJRPC()` function for Hive api defaults.
// To specify data pass a &JSONrpc{} with the desired data into `NewJRPC()`.
type JSONrpc struct {
	version string
	method  string
	params  interface{}
	id      int
	url     string
}

// NewJRPC creates an JSONrpc struct with defaults.
// If wish to use a different Hive endpoint pass the URL
// as a parameter. Otherwise leave the paramaters empty.
func NewJRPC(url ...string) *JSONrpc {
	jrpc := &JSONrpc{
		version: "2.0",
		method:  "",
		params:  make([]string, 0),
		id:      1,
		url:     "https://api.hive.blog",
	}

	if len(url) > 0 {
		jrpc.url = url[0]
	}
	return jrpc
}

func (jrpc *JSONrpc) getAccountData() (*jsonrpc.RPCResponse, error) {
	client := jsonrpc.NewClient(jrpc.url)
	resp, err := client.Call(jrpc.method, jrpc.params)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("%s", resp.Error)
	}
	return resp, nil
}

// GetAccountBandwidth returns the current "forum" bandwidth for a given account.
// This currently returns "Could not find method" from api.hive.blog
func (jrpc *JSONrpc) GetAccountBandwidth(account string) (string, error) {
	jrpc.method = "get_account_bandwidth"
	jrpc.params = []string{account, "forum"}
	_, err := jrpc.getAccountData()
	if err != nil {
		return "", err
	}
	return "", nil
}
