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
				version: "2.0",
				method:  "",
				params:  make([]string, 0),
				id:      1,
				url:     "https://api.hive.blog",
			},
		},
		{
			name: "url passed in",
			args: args{
				url: []string{"test.URL"},
			},
			want: &JSONrpc{
				version: "2.0",
				method:  "",
				params:  make([]string, 0),
				id:      1,
				url:     "test.URL",
			},
		},
		{
			name: "more than one url passed in",
			args: args{
				url: []string{"test.URL", "another.str"},
			},
			want: &JSONrpc{
				version: "2.0",
				method:  "",
				params:  make([]string, 0),
				id:      1,
				url:     "test.URL",
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

func TestGetAccounts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getTestData("multi"))
	}))
	defer ts.Close()

	type fields struct {
		version string
		method  string
		params  interface{}
		id      int
		url     string
	}
	type args struct {
		account []string
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
				version: "2.0",
				method:  "",
				params:  [][]string{},
				id:      1,
				url:     ts.URL,
			},
			args:    args{[]string{"jrswab", "hiveio"}},
			want:    getAccountsTestServer("multi"),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jrpc := &JSONrpc{
				version: tt.fields.version,
				method:  tt.fields.method,
				params:  tt.fields.params,
				id:      tt.fields.id,
				url:     tt.fields.url,
			}
			got, err := jrpc.GetAccounts(tt.args.account...)
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
