package accounts

import (
	"reflect"
	"testing"
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
				id:      1,
				url:     "https://api.hive.blog",
			},
		},
		{
			name: "url passed in",
			args: args{
				url: []string{"test.URL"},
			},
			want: &Chain{
				id:      1,
				url:     "test.URL",
			},
		},
		{
			name: "more than one url passed in",
			args: args{
				url: []string{"test.URL", "another.str"},
			},
			want: &Chain{
				id:      1,
				url:     "test.URL",
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
