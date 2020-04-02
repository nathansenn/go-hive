package accounts

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"git.sr.ht/~jrswab/go-hive/accounts/mocks"
	"github.com/stretchr/testify/mock"
	rpc "github.com/ybbus/jsonrpc"
)

func TestNewChain(t *testing.T) {
	type args struct {
		url []string
	}
	tests := []struct {
		name string
		args args
		want *Chain
	}{
		{
			name: "No url passed in",
			args: args{
				url: []string{},
			},
			want: &Chain{
				id:     1,
				url:    "https://api.hive.blog",
				client: rpc.NewClient("https://api.hive.blog"),
			},
		},
		{
			name: "url passed in",
			args: args{
				url: []string{"test.URL"},
			},
			want: &Chain{
				id:     1,
				url:    "test.URL",
				client: rpc.NewClient("test.URL"),
			},
		},
		{
			name: "more than one url passed in",
			args: args{
				url: []string{"test.URL", "another.str"},
			},
			want: &Chain{
				id:     1,
				url:    "test.URL",
				client: rpc.NewClient("test.URL"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChain(tt.args.url...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_GetAccountBandwidth(t *testing.T) {
	mockCall := new(mocks.Caller)
	var number json.Number
	number = "1111"
	output := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Result:  number,
		ID:      0,
	}
	output2 := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Error:   &rpc.RPCError{Code: 500, Message: "some error"},
		ID:      0,
	}
	mockCall.On("CallRaw", mock.Anything).Return(output, nil).Once()
	mockCall.On("CallRaw", mock.Anything).Return(nil, fmt.Errorf("fake error message")).Once()
	mockCall.On("CallRaw", mock.Anything).Return(output2, nil).Once()

	type fields struct {
		version string
		method  string
		params  interface{}
		id      int
		url     string
		client  Caller
	}
	type args struct {
		account string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Return bandwidth via proper usage",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{"jrswab"},
			want:    1111,
			wantErr: false,
		},
		{
			name: "Get call error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{"jrswab"},
			want:    -1,
			wantErr: true,
		},
		{
			name: "Get responce error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{"jrswab"},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chain{
				version: tt.fields.version,
				method:  tt.fields.method,
				params:  tt.fields.params,
				id:      tt.fields.id,
				url:     tt.fields.url,
				client:  tt.fields.client,
			}
			got, err := c.GetAccountBandwidth(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.GetAccountBandwidth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Chain.GetAccountBandwidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_GetAccountCount(t *testing.T) {
	mockCall := new(mocks.Caller)
	var number json.Number
	number = "1111"
	output := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Result:  number,
		ID:      0,
	}
	output2 := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Error:   &rpc.RPCError{Code: 500, Message: "some error"},
		ID:      0,
	}
	mockCall.On("CallRaw", mock.Anything).Return(output, nil).Once()
	mockCall.On("CallRaw", mock.Anything).Return(nil, fmt.Errorf("fake error message")).Once()
	mockCall.On("CallRaw", mock.Anything).Return(output2, nil).Once()

	type fields struct {
		version string
		method  string
		params  interface{}
		id      int
		url     string
		client  Caller
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{
			name: "Return a count via proper usage",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			want:    1111,
			wantErr: false,
		},
		{
			name: "Get call error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "Get responce error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chain{
				version: tt.fields.version,
				method:  tt.fields.method,
				params:  tt.fields.params,
				id:      tt.fields.id,
				url:     tt.fields.url,
				client:  tt.fields.client,
			}
			got, err := c.GetAccountCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.GetAccountCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Chain.GetAccountCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_GetAccountHistory(t *testing.T) {
	mockCall := new(mocks.Caller)
	res := make([][]interface{}, 0)
	output := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Result:  res,
		ID:      0,
	}
	output2 := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Error:   &rpc.RPCError{Code: 500, Message: "some error"},
		ID:      0,
	}
	mockCall.On("CallRaw", mock.Anything).Return(output, nil).Once()
	mockCall.On("CallRaw", mock.Anything).Return(nil, fmt.Errorf("fake error message")).Once()
	mockCall.On("CallRaw", mock.Anything).Return(output2, nil).Once()

	type fields struct {
		version string
		method  string
		params  interface{}
		id      int
		url     string
		client  Caller
	}
	type args struct {
		acc   string
		start int
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Return a count via proper usage",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{acc: "jrswab", start: 1000, limit: 1},
			want:    make([][]interface{}, 0),
			wantErr: false,
		},
		{
			name: "Get call error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{acc: "jrswab", start: 1000, limit: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get responce error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{acc: "jrswab", start: 1000, limit: 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chain{
				version: tt.fields.version,
				method:  tt.fields.method,
				params:  tt.fields.params,
				id:      tt.fields.id,
				url:     tt.fields.url,
				client:  tt.fields.client,
			}
			got, err := c.GetAccountHistory(tt.args.acc, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.GetAccountHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain.GetAccountHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_GetAccountReputation(t *testing.T) {
	mockCall := new(mocks.Caller)
	accRep := &accountReputation{Account: "jrswab", Reputation: "1111"}
	output := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Result:  []interface{}{accRep},
		ID:      0,
	}
	output2 := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Error:   &rpc.RPCError{Code: 500, Message: "some error"},
		ID:      0,
	}
	mockCall.On("CallRaw", mock.Anything).Return(output, nil).Once()
	mockCall.On("CallRaw", mock.Anything).Return(nil, fmt.Errorf("fake error message")).Once()
	mockCall.On("CallRaw", mock.Anything).Return(output2, nil).Once()

	type fields struct {
		version string
		method  string
		params  interface{}
		id      int
		url     string
		client  Caller
	}
	type args struct {
		acc string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Testing method logic.",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args: args{
				acc: "jrswab",
			},
			want:    1111,
			wantErr: false,
		},
		{
			name: "Get call error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{acc: "jrswab"},
			want:    -1,
			wantErr: true,
		},
		{
			name: "Get responce error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{acc: "jrswab"},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chain{
				version: tt.fields.version,
				method:  tt.fields.method,
				params:  tt.fields.params,
				id:      tt.fields.id,
				url:     tt.fields.url,
				client:  tt.fields.client,
			}
			got, err := c.GetAccountReputation(tt.args.acc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.GetAccountReputation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Chain.GetAccountReputation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChain_GetAccounts(t *testing.T) {
	mockCall := new(mocks.Caller)
	accMock := &AccountData{ID: 1111, Mined: false, Name: "jrswab"}
	accMock2 := &AccountData{ID: 2222, Mined: false, Name: "hiveio"}
	output := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Result:  []interface{}{accMock},
		ID:      0,
	}

	output2 := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Result:  []interface{}{accMock2, accMock},
		ID:      0,
	}

	output3 := &rpc.RPCResponse{
		JSONRPC: "2.0",
		Error:   &rpc.RPCError{Code: 500, Message: "some error"},
		ID:      0,
	}

	mockCall.On("CallRaw", mock.Anything).Return(output, nil).Once()
	mockCall.On("CallRaw", mock.Anything).Return(output2, nil).Once()
	mockCall.On("CallRaw", mock.Anything).Return(nil, fmt.Errorf("fake error message")).Once()
	mockCall.On("CallRaw", mock.Anything).Return(output3, nil).Once()

	type fields struct {
		version string
		method  string
		params  interface{}
		id      int
		url     string
		client  Caller
	}
	type args struct {
		acc []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]AccountData
		wantErr bool
	}{
		{
			name: "Get single account",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{[]string{"jrswab"}},
			want:    &[]AccountData{*accMock},
			wantErr: false,
		},
		{
			name: "Get two accounts",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{[]string{"jrswab", "hiveio"}},
			want:    &[]AccountData{*accMock2, *accMock},
			wantErr: false,
		},
		{
			name: "Get call error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{[]string{"jrswab", "hiveio"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get responce error message",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{[]string{"jrswab", "hiveio"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty args",
			fields: fields{
				version: "2.0",
				method:  "get_account_bandwidth",
				params:  nil,
				id:      0,
				url:     "https://api.hive.blog",
				client:  mockCall,
			},
			args:    args{[]string{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chain{
				version: tt.fields.version,
				method:  tt.fields.method,
				params:  tt.fields.params,
				id:      tt.fields.id,
				url:     tt.fields.url,
				client:  tt.fields.client,
			}
			got, err := c.GetAccounts(tt.args.acc...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain.GetAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain.GetAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
