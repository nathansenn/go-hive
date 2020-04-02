package accounts

import (
	"fmt"
	"strconv"

	rpc "github.com/ybbus/jsonrpc"
)

// Caller interface is used for testing purposes.
type Caller interface {
	CallRaw(*rpc.RPCRequest) (*rpc.RPCResponse, error)
}

// Chain is used to pass data into unexposed functions.
// When defining a new JSONrpc use the `NewChain()` function for Hive API defaults.
// To specify an api endpoint execute `NewChain()` with a full URL.
type Chain struct {
	version string
	method  string
	params  interface{}
	id      int
	url     string
	client  Caller
}

// AccountData holds the output of the GetAccounts method.
type AccountData struct {
	Active                        map[string]interface{} `json:"active"`
	Balance                       string                 `json:"balance"`
	CanVote                       bool                   `json:"can_vote"`
	CommentCount                  int                    `json:"comment_count"`
	Created                       string                 `json:"created"`
	CurationRewards               int                    `json:"curation_rewards"`
	DelegatedVestingShares        string                 `json:"delegated_vesting_shares"`
	DownVoteManaBar               map[string]interface{} `json:"down_vote_manabar"`
	GuestBloggers                 []string               `json:"guest_bloggers"`
	HbdBalance                    string                 `json:"sbd_balance"`
	HbdSeconds                    string                 `json:"sbd_seconds"`
	HbdSecondsLastUpdate          string                 `json:"sbd_seconds_last_update"`
	HbdLastInterestPayment        string                 `json:"sbd_last_interest_payment"`
	ID                            int                    `json:"id"`
	JSONMetadata                  string                 `json:"json_metadata"`
	LastAccountRecovery           string                 `json:"last_account_recovery"`
	LastAccountUpdate             string                 `json:"last_account_update"`
	LastOwnerUpdate               string                 `json:"last_owner_update"`
	LastPost                      string                 `json:"last_post"`
	LastRootPost                  string                 `json:"last_root_post"`
	LastVoteTime                  string                 `json:"last_vote_time"`
	LifetimeVoteCount             int                    `json:"lifetime_vote_count"`
	MarketHistory                 []interface{}          `json:"market_history"`
	MemoKey                       string                 `json:"memo_key"`
	Mined                         bool                   `json:"mined"`
	Name                          string                 `json:"name"`
	NextVestingWithdraw           string                 `json:"next_vesting_withdraw"`
	OtherHistory                  []interface{}          `json:"other_history"`
	Owner                         map[string]interface{} `json:"owner"`
	PendingClaimedAccounts        int                    `json:"pending_claimed_accounts"`
	PostBandwidth                 int                    `json:"post_bandwidth"`
	PostCount                     int                    `json:"post_count"`
	PostHistory                   []interface{}          `json:"post_history"`
	Posting                       interface{}            `json:"posting"`
	PostingJSONMetadata           string                 `json:"posting_json_metadata"`
	PostingRewards                float64                `json:"posting_rewards"`
	ProxiedVsfVotes               interface{}            `json:"proxied_vsf_votes"`
	Proxy                         string                 `json:"proxy"`
	ReceivedVestingShares         string                 `json:"recived_vesting_shares"`
	RecoveryAccount               string                 `json:"recovery_Account"`
	Reputation                    string                 `json:"reputation"`
	ResetAccount                  string                 `json:"reset_account"`
	RewardHBDBalance              string                 `json:"reward_sbd_balance"`
	RewardHiveBalance             string                 `json:"reward_steem_balance"`
	RewardVestingBalance          string                 `json:"reward_vesting_balance"`
	RewardVestingHive             string                 `json:"reward_vesting_steem"`
	SavingsBalance                string                 `json:"savings_balance"`
	SavingsHbdBalance             string                 `json:"savings_sbd_balance"`
	SavingsHbdSeconds             string                 `json:"savings_sbd_seconds"`
	SavingsHbdSecondsLastUpdate   string                 `json:"savings_sbd_seconds_last_update"`
	SavingsHbdLastInterestPayment string                 `json:"savings_sbd_last_interest_payment"`
	TagsUsage                     []string               `json:"tags_usage"`
	TransferHistory               []interface{}          `json:"transfer_history"`
	ToWithdraw                    int                    `json:"to_withdraw"`
	VestingBalance                string                 `json:"vesting_balance"`
	VestingShares                 string                 `json:"vesting_shares"`
	VestingWithdrawRate           string                 `json:"vesting_withdraw_rate"`
	VoteHistory                   []interface{}          `json:"vote_history"`
	VotingPower                   int                    `json:"voting_power"`
	Withdrawn                     int                    `json:"withdrawn"`
	WithdrawRoutes                int                    `json:"withdraw_routes"`
	WitnessesVotedFor             int                    `json:"witnesses_vote_for"`
	WitnessVotes                  []string               `json:"witness_votes"`
}

// NewChain creates an struct with Hive defaults.
// If wish to use a different Hive endpoint (or a different Graphene blockchain
// pass the URL as a parameter. Otherwise leave the parameters empty.
// If more than one URL is entered, only the first will be used.
// Example:
// hive := NewChain()
func NewChain(url ...string) *Chain {
	chain := &Chain{
		id:     1,
		url:    "https://api.hive.blog",
		client: rpc.NewClient("https://api.hive.blog"),
	}

	if len(url) > 0 {
		chain.url = url[0]
		chain.client = rpc.NewClient(url[0])
	}
	return chain
}

func (c *Chain) getAccountData(inputParams ...interface{}) (*rpc.RPCResponse, error) {
	request := rpc.NewRequest(c.method, inputParams)

	resp, err := c.client.CallRaw(request)
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
func (c *Chain) GetAccountBandwidth(account string) (int64, error) {
	c.method = "get_account_bandwidth"
	c.params = []string{account, "forum"}
	resp, err := c.getAccountData()
	if err != nil {
		return -1, err
	}
	out, err := resp.GetInt()
	if err != nil {
		return -1, err
	}
	return out, nil
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
func (c *Chain) GetAccountHistory(acc string, start, limit int) (interface{}, error) {
	c.method = "get_account_history"

	resp, err := c.getAccountData(acc, start, limit)
	if err != nil {
		return nil, err
	}
	var arr [][]interface{}
	if err = resp.GetObject(&arr); err != nil {
		return nil, err
	}
	return arr, nil
}

// AccountReputation is a struct for receiving data from GetAccountReputation()
type accountReputation struct {
	Account    string `json:"account"`
	Reputation string `json:"reputation"`
}

// GetAccountReputation takes accounts name and returns a slice of type AccountReputation.
// Returns `-1` when error is not nil.
func (c *Chain) GetAccountReputation(acc string) (int, error) {
	c.method = "get_account_reputations"
	data := []accountReputation{}

	resp, err := c.getAccountData(acc, 1)
	if err != nil {
		return -1, err
	}

	if resp.GetObject(&data); err != nil {
		return -1, err
	}

	num, err := strconv.Atoi(data[0].Reputation)
	if err != nil {
		return -1, err
	}
	return num, nil
}

// GetAccounts updates a slice of account data for the accounts passed in.
// At least one account is required.
//Example:
//c := a.NewChain()
//
//accData,err := c.GetAccounts("jrswab")
//if err != nil {
//	fmt.Println(err)
//}
//fmt.Println(data[0].Balance)
func (c *Chain) GetAccounts(acc ...string) (*[]AccountData, error) {
	if len(acc) < 1 {
		return nil, fmt.Errorf("method GetAccounts needs at least one account name")
	}
	c.method = "get_accounts"

	resp, err := c.getAccountData(acc)
	if err != nil {
		return nil, err
	}

	out := []AccountData{}
	err = resp.GetObject(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
