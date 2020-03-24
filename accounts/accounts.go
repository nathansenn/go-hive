package accounts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Accounts hold the data retevied by the GetAccount and GetAccounts functions.
type Accounts struct {
	Jsonrpc string `json:"jsonrpc"`
	Data    []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Owner struct {
			WeightThreshold int             `json:"weight_threshold"`
			AccountAuths    []interface{}   `json:"account_auths"`
			KeyAuths        [][]interface{} `json:"key_auths"`
		} `json:"owner"`
		Active struct {
			WeightThreshold int             `json:"weight_threshold"`
			AccountAuths    []interface{}   `json:"account_auths"`
			KeyAuths        [][]interface{} `json:"key_auths"`
		} `json:"active"`
		Posting struct {
			WeightThreshold int             `json:"weight_threshold"`
			AccountAuths    [][]interface{} `json:"account_auths"`
			KeyAuths        [][]interface{} `json:"key_auths"`
		} `json:"posting"`
		MemoKey             string `json:"memo_key"`
		JSONMetadata        string `json:"json_metadata"`
		PostingJSONMetadata string `json:"posting_json_metadata"`
		Proxy               string `json:"proxy"`
		LastOwnerUpdate     string `json:"last_owner_update"`
		LastAccountUpdate   string `json:"last_account_update"`
		Created             string `json:"created"`
		Mined               bool   `json:"mined"`
		RecoveryAccount     string `json:"recovery_account"`
		LastAccountRecovery string `json:"last_account_recovery"`
		ResetAccount        string `json:"reset_account"`
		CommentCount        int    `json:"comment_count"`
		LifetimeVoteCount   int    `json:"lifetime_vote_count"`
		PostCount           int    `json:"post_count"`
		CanVote             bool   `json:"can_vote"`
		VotingManabar       struct {
			CurrentMana    string `json:"current_mana"`
			LastUpdateTime int    `json:"last_update_time"`
		} `json:"voting_manabar"`
		DownvoteManabar struct {
			CurrentMana    string `json:"current_mana"`
			LastUpdateTime int    `json:"last_update_time"`
		} `json:"downvote_manabar"`
		VotingPower                   int           `json:"voting_power"`
		Balance                       string        `json:"balance"`
		SavingsBalance                string        `json:"savings_balance"`
		SbdBalance                    string        `json:"sbd_balance"`
		SbdSeconds                    string        `json:"sbd_seconds"`
		SbdSecondsLastUpdate          string        `json:"sbd_seconds_last_update"`
		SbdLastInterestPayment        string        `json:"sbd_last_interest_payment"`
		SavingsSbdBalance             string        `json:"savings_sbd_balance"`
		SavingsSbdSeconds             string        `json:"savings_sbd_seconds"`
		SavingsSbdSecondsLastUpdate   string        `json:"savings_sbd_seconds_last_update"`
		SavingsSbdLastInterestPayment string        `json:"savings_sbd_last_interest_payment"`
		SavingsWithdrawRequests       int           `json:"savings_withdraw_requests"`
		RewardSbdBalance              string        `json:"reward_sbd_balance"`
		RewardSteemBalance            string        `json:"reward_steem_balance"`
		RewardVestingBalance          string        `json:"reward_vesting_balance"`
		RewardVestingSteem            string        `json:"reward_vesting_steem"`
		VestingShares                 string        `json:"vesting_shares"`
		DelegatedVestingShares        string        `json:"delegated_vesting_shares"`
		ReceivedVestingShares         string        `json:"received_vesting_shares"`
		VestingWithdrawRate           string        `json:"vesting_withdraw_rate"`
		NextVestingWithdrawal         string        `json:"next_vesting_withdrawal"`
		Withdrawn                     int           `json:"withdrawn"`
		ToWithdraw                    string        `json:"to_withdraw"`
		WithdrawRoutes                int           `json:"withdraw_routes"`
		CurationRewards               int           `json:"curation_rewards"`
		PostingRewards                int           `json:"posting_rewards"`
		ProxiedVsfVotes               []interface{} `json:"proxied_vsf_votes"`
		WitnessesVotedFor             int           `json:"witnesses_voted_for"`
		LastPost                      string        `json:"last_post"`
		LastRootPost                  string        `json:"last_root_post"`
		LastVoteTime                  string        `json:"last_vote_time"`
		PostBandwidth                 int           `json:"post_bandwidth"`
		PendingClaimedAccounts        int           `json:"pending_claimed_accounts"`
		VestingBalance                string        `json:"vesting_balance"`
		Reputation                    string        `json:"reputation"`
		TransferHistory               []interface{} `json:"transfer_history"`
		MarketHistory                 []interface{} `json:"market_history"`
		PostHistory                   []interface{} `json:"post_history"`
		VoteHistory                   []interface{} `json:"vote_history"`
		OtherHistory                  []interface{} `json:"other_history"`
		WitnessVotes                  []string      `json:"witness_votes"`
		TagsUsage                     []interface{} `json:"tags_usage"`
		GuestBloggers                 []interface{} `json:"guest_bloggers"`
	} `json:"result"`
}

// JSONrpc is used to pass data into unexposed functions.
// When defining a new JSONrpc use the `NewJRPC()` function for Hive api defaults.
// To specify data pass a &JSONrpc{} with the desired data into `NewJRPC()`.
type JSONrpc struct {
	Version string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  [][]string `json:"params"`
	ID      int        `json:"id"`
	url     string
}

// NewJRPC creates an JSONrpc struct with defaults.
// If wish to use a different Hive endpoint pass the URL
// as a parameter. Otherwise leave the paramaters empty.
func NewJRPC(url ...string) *JSONrpc {
	jrpc := &JSONrpc{
		Version: "2.0",
		Method:  "",
		Params:  make([][]string, 0),
		ID:      1,
		url:     "https://api.hive.blog",
	}

	if len(url) > 0 {
		jrpc.url = url[0]
	}
	return jrpc
}

// GetAccount takes an account name and performs the needed actions to return a pointer to a
// struct.
//
//Example:
//import "github.com/jrswab/go-hive/accounts"
//c := accounts.NewJRPC()
//accStruct, err := c.GetAcount("jrswab")
//     if err != nil {
//        // handle  error
//     }
//fmt.Println(accStruct.Data[0].Name) // Prints out the account name that was passed in.
func (jrpc *JSONrpc) GetAccount(account string) (*Accounts, error) {
	prep := []string{account}
	params := [][]string{prep}

	jrpc.Method = "get_accounts"
	jrpc.Params = params
	raw, err := jrpc.getAccountData()
	if err != nil {
		return nil, err
	}

	accountData := &Accounts{}
	err = json.Unmarshal(raw, accountData)
	if err != nil {
		return nil, err
	}
	return accountData, nil
}

// GetAccounts takes a slice of account names and performs the needed actions to return
// a pointer to a struct with a slice of structs. The slice of structs contain the accounts
// requested via the slice of strings provided.
//
//Example:
//import "github.com/jrswab/go-hive/accounts"
//c := accounts.NewJRPC()
//accountNames := []string{"jrswab", "hiveio"}
//accStructs, err := c.GetAccounts(accountNames)
//     if err != nil {
//        // handle  error
//     }
//fmt.Println(accStruct.Data[0].Name) // Prints out the first account name that was passed in.
func (jrpc *JSONrpc) GetAccounts(accList []string) (*Accounts, error) {
	params := [][]string{accList}

	jrpc.Method = "get_accounts"
	jrpc.Params = params
	raw, err := jrpc.getAccountData()
	if err != nil {
		return nil, err
	}

	accountsData := &Accounts{}
	err = json.Unmarshal(raw, accountsData)
	if err != nil {
		return nil, err
	}
	return accountsData, nil
}

func (jrpc *JSONrpc) getAccountData() ([]byte, error) {
	bytes, err := json.Marshal(jrpc)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(string(bytes))
	req, err := http.NewRequest("POST", jrpc.url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
