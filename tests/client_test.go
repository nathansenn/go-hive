package gohive

import (
	"reflect"
	"testing"

	h "github.com/nathansenn/go-hive"
	rpc "github.com/ybbus/jsonrpc"
)

func TestNewClient(t *testing.T) {
	type args struct {
		URL []string
	}
	tests := []struct {
		name string
		args args
		want *h.Client
	}{
		{
			name: "No URL passed in",
			args: args{
				URL: []string{},
			},
			want: &h.Client{
				URL:    "https://api.hive.blog",
				Client: rpc.NewClient("https://api.hive.blog"),
			},
		},
		{
			name: "URL passed in",
			args: args{
				URL: []string{"test.URL"},
			},
			want: &h.Client{
				URL:    "test.URL",
				Client: rpc.NewClient("test.URL"),
			},
		},
		{
			name: "more than one URL passed in",
			args: args{
				URL: []string{"test.URL", "another.str"},
			},
			want: &h.Client{
				URL:    "test.URL",
				Client: rpc.NewClient("test.URL"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := h.NewClient(tt.args.URL...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}
