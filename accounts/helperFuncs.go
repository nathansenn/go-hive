package accounts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func getAccountsTestServer(request string) *Accounts {
	a := new(Accounts)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if request == "single" {
			fmt.Fprintln(w, getSingle())
		}
		if request == "multi" {
			fmt.Fprintln(w, getMulti())
		}
	}))

	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&a)
	return a
}

func getTestData(units string) string {
	data := ""
	if units == "single" {
		data = getSingle()
	}
	if units == "multi" {
		data = getMulti()
	}
	return data
}

func getSingle() string {
	return `{
	"jsonrpc": "2.0",
	"result": [
		{
			"id": 222785,
			"name": "jrswab",
			"owner": {
				"weight_threshold": 1,
				"account_auths": [],
				"key_auths": [
					[
						"STM5n9h8zjMxBNrkDPvzW7W3Bqd1Df6GYJCuAtd5XvgvyuKMQj4AC",
						1
					]
				]
			},
			"active": {
				"weight_threshold": 1,
				"account_auths": [],
				"key_auths": [
					[
						"STM5CbcdrFGcvLkEwnkjptX9APAkh43gS8EHWhuHVrjH1qX1KfLGP",
						1
					]
				]
			},
			"posting": {
				"weight_threshold": 1,
				"account_auths": [
					[
						"peakd.app",
						1
					]
				],
				"key_auths": [
					[
						"STM5SSMk2inVEEKKG2tbXPQ4EajHpHqeHkwiWA3JKTzzS1ZRk4FBi",
						1
					]
				]
			},
			"memo_key": "STM6dgyhyDihkPagCgXY3G7A8mNv76UZUBise7JX1LdMy24UfLwKg",
			"json_metadata": "{\"profile\":{\"name\":\"J. R. Swab\",\"about\":\"Hacker, Content Creator, Golang Enthusiast, Linux User, & Eagle Scout | Rom. 10:9\",\"website\":\"http:\\/\\/jrswab.com\",\"location\":\"Pennsylvania, USA\",\"cover_image\":\"https:\\/\\/cdn.steemitimages.com\\/DQma5bng289kProRehA7kmW9GGRyhXrysKRrdUssdWQygBe\\/pixelYoda.png\",\"profile_image\":\"https:\\/\\/cdn.steemitimages.com\\/DQmdcLD5FTzjKmAXnCCJvG3JSRobFTySTzWnwMWf8wKrBkc\\/meJuly2019.jpeg\"}}",
			"posting_json_metadata": "{\"profile\":{\"name\":\"J. R. Swab\",\"about\":\"Christian Blogger & Podcaster Writing About Technology & Self Improvement. \\ud83e\\udd85 Eagle Scout \\ud83d\\udc68\\u200d\\ud83d\\udcbb Programmer \\ud83d\\udcd7 Komencanto Esperantisto.\",\"website\":\"http:\\/\\/jrswab.com\",\"location\":\"Pennsylvania, USA\",\"cover_image\":\"https:\\/\\/jrswab.com\\/images\\/Banner.png\",\"profile_image\":\"https:\\/\\/jrswab.com\\/images\\/me.jpg\"}}",
			"proxy": "",
			"last_owner_update": "2018-09-22T01:22:36",
			"last_account_update": "2020-03-20T15:22:42",
			"created": "2017-06-26T15:26:09",
			"mined": false,
			"recovery_account": "steem",
			"last_account_recovery": "1970-01-01T00:00:00",
			"reset_account": "null",
			"comment_count": 0,
			"lifetime_vote_count": 0,
			"post_count": 2284,
			"can_vote": true,
			"voting_manabar": {
				"current_mana": "9465968451455",
				"last_update_time": 1584923910
			},
			"downvote_manabar": {
				"current_mana": "2390396073599",
				"last_update_time": 1584923910
			},
			"voting_power": 9899,
			"balance": "1.211 HIVE",
			"savings_balance": "0.000 HIVE",
			"sbd_balance": "4.863 HBD",
			"sbd_seconds": "8243426100",
			"sbd_seconds_last_update": "2020-03-20T15:43:42",
			"sbd_last_interest_payment": "2020-02-29T19:40:12",
			"savings_sbd_balance": "0.000 HBD",
			"savings_sbd_seconds": "0",
			"savings_sbd_seconds_last_update": "2017-11-01T20:48:48",
			"savings_sbd_last_interest_payment": "2017-11-01T20:48:48",
			"savings_withdraw_requests": 0,
			"reward_sbd_balance": "0.000 HBD",
			"reward_steem_balance": "0.000 HIVE",
			"reward_vesting_balance": "182.115824 VESTS",
			"reward_vesting_steem": "0.093 HIVE",
			"vesting_shares": "10343155.439830 VESTS",
			"delegated_vesting_shares": "0.000000 VESTS",
			"received_vesting_shares": "0.000000 VESTS",
			"vesting_withdraw_rate": "781571.145431 VESTS",
			"next_vesting_withdrawal": "2020-03-27T15:44:27",
			"withdrawn": 0,
			"to_withdraw": "10160424890596",
			"withdraw_routes": 0,
			"curation_rewards": 199683,
			"posting_rewards": 3687287,
			"proxied_vsf_votes": [
				"76355272671",
				0,
				0,
				0
			],
			"witnesses_voted_for": 20,
			"last_post": "2020-03-21T14:08:36",
			"last_root_post": "2020-03-21T14:08:36",
			"last_vote_time": "2020-03-23T00:38:30",
			"post_bandwidth": 0,
			"pending_claimed_accounts": 0,
			"vesting_balance": "0.000 HIVE",
			"reputation": "45256100425929",
			"transfer_history": [],
			"market_history": [],
			"post_history": [],
			"vote_history": [],
			"other_history": [],
			"witness_votes": [
				"c0ff33a",
				"dsound",
				"holger80",
				"jackmiller",
				"prc",
				"r0nd0n",
				"techcoderx"
			],
			"tags_usage": [],
			"guest_bloggers": []
		}
	],
	"id": 1
}`
}

func getMulti() string {
	return `{
	"jsonrpc": "2.0",
	"result": [
		{
			"id": 222785,
			"name": "jrswab",
			"owner": {
				"weight_threshold": 1,
				"account_auths": [],
				"key_auths": [
					[
						"STM5n9h8zjMxBNrkDPvzW7W3Bqd1Df6GYJCuAtd5XvgvyuKMQj4AC",
						1
					]
				]
			},
			"active": {
				"weight_threshold": 1,
				"account_auths": [],
				"key_auths": [
					[
						"STM5CbcdrFGcvLkEwnkjptX9APAkh43gS8EHWhuHVrjH1qX1KfLGP",
						1
					]
				]
			},
			"posting": {
				"weight_threshold": 1,
				"account_auths": [
					[
						"peakd.app",
						1
					]
				],
				"key_auths": [
					[
						"STM5SSMk2inVEEKKG2tbXPQ4EajHpHqeHkwiWA3JKTzzS1ZRk4FBi",
						1
					]
				]
			},
			"memo_key": "STM6dgyhyDihkPagCgXY3G7A8mNv76UZUBise7JX1LdMy24UfLwKg",
			"json_metadata": "{\"profile\":{\"name\":\"J. R. Swab\",\"about\":\"Hacker, Content Creator, Golang Enthusiast, Linux User, & Eagle Scout | Rom. 10:9\",\"website\":\"http:\\/\\/jrswab.com\",\"location\":\"Pennsylvania, USA\",\"cover_image\":\"https:\\/\\/cdn.steemitimages.com\\/DQma5bng289kProRehA7kmW9GGRyhXrysKRrdUssdWQygBe\\/pixelYoda.png\",\"profile_image\":\"https:\\/\\/cdn.steemitimages.com\\/DQmdcLD5FTzjKmAXnCCJvG3JSRobFTySTzWnwMWf8wKrBkc\\/meJuly2019.jpeg\"}}",
			"posting_json_metadata": "{\"profile\":{\"name\":\"J. R. Swab\",\"about\":\"Christian Blogger & Podcaster Writing About Technology & Self Improvement. \\ud83e\\udd85 Eagle Scout \\ud83d\\udc68\\u200d\\ud83d\\udcbb Programmer \\ud83d\\udcd7 Komencanto Esperantisto.\",\"website\":\"http:\\/\\/jrswab.com\",\"location\":\"Pennsylvania, USA\",\"cover_image\":\"https:\\/\\/jrswab.com\\/images\\/Banner.png\",\"profile_image\":\"https:\\/\\/jrswab.com\\/images\\/me.jpg\"}}",
			"proxy": "",
			"last_owner_update": "2018-09-22T01:22:36",
			"last_account_update": "2020-03-20T15:22:42",
			"created": "2017-06-26T15:26:09",
			"mined": false,
			"recovery_account": "steem",
			"last_account_recovery": "1970-01-01T00:00:00",
			"reset_account": "null",
			"comment_count": 0,
			"lifetime_vote_count": 0,
			"post_count": 2284,
			"can_vote": true,
			"voting_manabar": {
				"current_mana": "9465968451455",
				"last_update_time": 1584923910
			},
			"downvote_manabar": {
				"current_mana": "2390396073599",
				"last_update_time": 1584923910
			},
			"voting_power": 9899,
			"balance": "1.211 HIVE",
			"savings_balance": "0.000 HIVE",
			"sbd_balance": "4.863 HBD",
			"sbd_seconds": "8243426100",
			"sbd_seconds_last_update": "2020-03-20T15:43:42",
			"sbd_last_interest_payment": "2020-02-29T19:40:12",
			"savings_sbd_balance": "0.000 HBD",
			"savings_sbd_seconds": "0",
			"savings_sbd_seconds_last_update": "2017-11-01T20:48:48",
			"savings_sbd_last_interest_payment": "2017-11-01T20:48:48",
			"savings_withdraw_requests": 0,
			"reward_sbd_balance": "0.000 HBD",
			"reward_steem_balance": "0.000 HIVE",
			"reward_vesting_balance": "182.115824 VESTS",
			"reward_vesting_steem": "0.093 HIVE",
			"vesting_shares": "10343155.439830 VESTS",
			"delegated_vesting_shares": "0.000000 VESTS",
			"received_vesting_shares": "0.000000 VESTS",
			"vesting_withdraw_rate": "781571.145431 VESTS",
			"next_vesting_withdrawal": "2020-03-27T15:44:27",
			"withdrawn": 0,
			"to_withdraw": "10160424890596",
			"withdraw_routes": 0,
			"curation_rewards": 199683,
			"posting_rewards": 3687287,
			"proxied_vsf_votes": [
				"76355272671",
				0,
				0,
				0
			],
			"witnesses_voted_for": 20,
			"last_post": "2020-03-21T14:08:36",
			"last_root_post": "2020-03-21T14:08:36",
			"last_vote_time": "2020-03-23T00:38:30",
			"post_bandwidth": 0,
			"pending_claimed_accounts": 0,
			"vesting_balance": "0.000 HIVE",
			"reputation": "45256100425929",
			"transfer_history": [],
			"market_history": [],
			"post_history": [],
			"vote_history": [],
			"other_history": [],
			"witness_votes": [
				"c0ff33a",
				"dsound",
				"holger80",
				"jackmiller",
				"prc",
				"r0nd0n",
				"techcoderx"
			],
			"tags_usage": [],
			"guest_bloggers": []
		},
		{
			"id": 1370484,
			"name": "hiveio",
			"owner": {
				"weight_threshold": 1,
				"account_auths": [],
				"key_auths": [
					[
						"STM65PUAPA4yC4RgPtGgsPupxT6yJtMhmT5JHFdsT3uoCbR8WJ25s",
						1
					]
				]
			},
			"active": {
				"weight_threshold": 1,
				"account_auths": [],
				"key_auths": [
					[
						"STM69zfrFGnZtU3gWFWpQJ6GhND1nz7TJsKBTjcWfebS1JzBEweQy",
						1
					]
				]
			},
			"posting": {
				"weight_threshold": 1,
				"account_auths": [],
				"key_auths": [
					[
						"STM6vJmrwaX5TjgTS9dPH8KsArso5m91fVodJvv91j7G765wqcNM9",
						1
					]
				]
			},
			"memo_key": "STM7wrsg1BZogeK7X3eG4ivxmLaH69FomR8rLkBbepb3z3hm5SbXu",
			"json_metadata": "",
			"posting_json_metadata": "{\"profile\":{\"name\":\"Hive.io\",\"about\":\"Official account for the Hive Blockchain\n A blockchain built for the community, by the community. #HiveIsAlive\",\"website\":\"https:\\/\\/hive.io\",\"profile_image\":\"https:\\/\\/files.steempeak.com\\/file\\/steempeak\\/hiveio\\/qj0i0okQ-image.png\",\"version\":2,\"cover_image\":\"https:\\/\\/files.steempeak.com\\/file\\/steempeak\\/steempeak\\/7lrbii0D-image.png\"}}",
			"proxy": "",
			"last_owner_update": "1970-01-01T00:00:00",
			"last_account_update": "2020-03-17T23:42:30",
			"created": "2020-03-06T12:22:48",
			"mined": false,
			"recovery_account": "steempeak",
			"last_account_recovery": "1970-01-01T00:00:00",
			"reset_account": "null",
			"comment_count": 0,
			"lifetime_vote_count": 0,
			"post_count": 3,
			"can_vote": true,
			"voting_manabar": {
				"current_mana": "3917361311663",
				"last_update_time": 1584487782
			},
			"downvote_manabar": {
				"current_mana": "979340327915",
				"last_update_time": 1584487782
			},
			"voting_power": 0,
			"balance": "0.006 HIVE",
			"savings_balance": "0.000 HIVE",
			"sbd_balance": "0.000 HBD",
			"sbd_seconds": "0",
			"sbd_seconds_last_update": "1970-01-01T00:00:00",
			"sbd_last_interest_payment": "1970-01-01T00:00:00",
			"savings_sbd_balance": "0.000 HBD",
			"savings_sbd_seconds": "0",
			"savings_sbd_seconds_last_update": "1970-01-01T00:00:00",
			"savings_sbd_last_interest_payment": "1970-01-01T00:00:00",
			"savings_withdraw_requests": 0,
			"reward_sbd_balance": "0.000 HBD",
			"reward_steem_balance": "0.000 HIVE",
			"reward_vesting_balance": "0.000000 VESTS",
			"reward_vesting_steem": "0.000 HIVE",
			"vesting_shares": "0.000000 VESTS",
			"delegated_vesting_shares": "0.000000 VESTS",
			"received_vesting_shares": "3917361.311663 VESTS",
			"vesting_withdraw_rate": "0.000000 VESTS",
			"next_vesting_withdrawal": "1969-12-31T23:59:59",
			"withdrawn": 0,
			"to_withdraw": "0",
			"withdraw_routes": 0,
			"curation_rewards": 0,
			"posting_rewards": 0,
			"proxied_vsf_votes": [
				0,
				0,
				0,
				0
			],
			"witnesses_voted_for": 0,
			"last_post": "2020-03-21T17:05:33",
			"last_root_post": "2020-03-21T17:05:33",
			"last_vote_time": "1970-01-01T00:00:00",
			"post_bandwidth": 0,
			"pending_claimed_accounts": 0,
			"vesting_balance": "0.000 HIVE",
			"reputation": "21024881578520",
			"transfer_history": [],
			"market_history": [],
			"post_history": [],
			"vote_history": [],
			"other_history": [],
			"witness_votes": [],
			"tags_usage": [],
			"guest_bloggers": []
		}
	],
	"id": 1
}`
}
