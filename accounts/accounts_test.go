package accounts

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewJRPC(t *testing.T) {
	type args struct {
		url []string
	}
	tests := []struct {
		name string
		args args
		want *JSONrpc
	}{
		{
			name: "No url passed in",
			args: args{
				url: []string{},
			},
			want: &JSONrpc{
				Version: "2.0",
				Method:  "",
				Params:  make([][]string, 0),
				ID:      1,
				url:     "https://api.hive.blog",
			},
		},
		{
			name: "url passed in",
			args: args{
				url: []string{"ts.URL"},
			},
			want: &JSONrpc{
				Version: "2.0",
				Method:  "",
				Params:  make([][]string, 0),
				ID:      1,
				url:     "ts.URL",
			},
		},
		{
			name: "more than one url passed in",
			args: args{
				url: []string{"ts.URL", "another.str"},
			},
			want: &JSONrpc{
				Version: "2.0",
				Method:  "",
				Params:  make([][]string, 0),
				ID:      1,
				url:     "ts.URL",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJRPC(tt.args.url...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONrpc_GetAccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getTestData("single"))
	}))
	defer ts.Close()

	type fields struct {
		Version string
		Method  string
		Params  [][]string
		ID      int
		url     string
	}
	type args struct {
		account string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Accounts
		wantErr bool
	}{
		{
			name: "Test correct call.",
			fields: fields{
				Version: "2.0",
				Method:  "",
				Params:  [][]string{},
				ID:      1,
				url:     ts.URL,
			},
			args:    args{"jrswab"},
			want:    runTestServer("single"),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jrpc := &JSONrpc{
				Version: tt.fields.Version,
				Method:  tt.fields.Method,
				Params:  tt.fields.Params,
				ID:      tt.fields.ID,
				url:     tt.fields.url,
			}
			got, err := jrpc.GetAccount(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONrpc.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONrpc.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONrpc_GetAccounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getTestData("multi"))
	}))
	defer ts.Close()

	type fields struct {
		Version string
		Method  string
		Params  [][]string
		ID      int
		url     string
	}
	type args struct {
		accList []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Accounts
		wantErr bool
	}{
		{
			name: "Test correct call.",
			fields: fields{
				Version: "2.0",
				Method:  "",
				Params:  [][]string{},
				ID:      1,
				url:     ts.URL,
			},
			args:    args{[]string{"jrswab", "hiveio"}},
			want:    runTestServer("multi"),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jrpc := &JSONrpc{
				Version: tt.fields.Version,
				Method:  tt.fields.Method,
				Params:  tt.fields.Params,
				ID:      tt.fields.ID,
				url:     tt.fields.url,
			}
			got, err := jrpc.GetAccounts(tt.args.accList)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONrpc.GetAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONrpc.GetAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
