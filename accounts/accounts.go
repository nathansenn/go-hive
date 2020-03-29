package accounts

import (
	"fmt"
	"strconv"

	"github.com/ybbus/jsonrpc"
)

// Chain is used to pass data into unexposed functions.
// When defining a new JSONrpc use the `NewJRPC()` function for Hive api defaults.
// To specify an api endpoint execute `NewJRPC()` with a full URL.
type Chain struct {
	version string
	method  string
	params  interface{}
	id      int
	url     string
}

// NewChain creates an struct with Hive defaults.
// If wish to use a different Hive endpoint (or a different Graphene blockchain
// pass the URL as a parameter. Otherwise leave the paramaters empty.
// If more than one URL is entered, only the first will be used.
// Example:
// hive := NewChain() // this will default to api.hive.blog
func NewChain(url ...string) *Chain {
	chain := &Chain{
		id:  1,
		url: "https://api.hive.blog",
	}

	if len(url) > 0 {
		chain.url = url[0]
	}
	return chain
}

func (c *Chain) getAccountData(inputParams ...interface{}) (*jsonrpc.RPCResponse, error) {
	client := jsonrpc.NewClient(c.url)
	resp, err := client.Call(c.method, jsonrpc.Params(inputParams))

	if err != nil {
		return nil, fmt.Errorf("json rpc call error: %s", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("rpc response returned: %s", resp.Error)
	}
	return resp, nil
}

// GetAccountBandwidth returns the current "forum" average bandwidth for a given account.
// This currently returns "Could not find method" from api.hive.blog
func (c *Chain) GetAccountBandwidth(account string) (string, error) {
	c.method = "get_account_bandwidth"
	c.params = []string{account, "forum"}
	_, err := c.getAccountData()
	if err != nil {
		return "", err
	}
	return "", nil
}

// GetAccountCount returns the current number of accounts on the network.
func (c *Chain) GetAccountCount() (int64, error) {
	c.method = "get_account_count"
	c.params = []string{}

	resp, err := c.getAccountData()
	if err != nil {
		return -1, err
	}

	output, err := resp.GetInt()
	if err != nil {
		return -1, err
	}

	return output, nil
}

// GetAccountHistory returns the history of an account.
// Currently, returns a slice of a slice of interfaces due to mixed data types.
// A better solution is needed and may come once the rest of the methods have been implemented.
// If you know a solution please submit a patch to jr[at]jrswab.com via [git send-email](https://git-send-email.io/)
// or use [git.sr.ht/~jrswab/go-hive/send-email](https://git.sr.ht/~jrswab/go-hive/send-email)
func (c *Chain) GetAccountHistory(acc string, start, limit int) (interface{}, error) {
	c.method = "get_account_history"

	resp, err := c.getAccountData(acc, start, limit)
	if err != nil {
		return nil, err
	}
	arr := [][]interface{}{}
	err = resp.GetObject(&arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetAccountReputation takes  an account name and returns and int of the current reputation.
// When this method returns an error it also sends `-1` as the integer.
func (c *Chain) GetAccountReputation(acc string) (int, error) {
	c.method = "get_account_reputations"

	resp, err := c.getAccountData(acc, 1)
	if err != nil {
		return -1, err
	}

	outMap := make(map[string]string)
	out := []map[string]string{outMap}
	err = resp.GetObject(&out)
	if err != nil {
		return -1, err
	}

	output, err := strconv.Atoi(outMap["reputation"])
	if err != nil {
		return -1, err
	}
	return output, nil

}
