package time_map

import (
	"reflect"
	"testing"
)

func Test_runTest(t *testing.T) {
	type args struct {
		commands []string
		payloads [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{commands: []string{"TimeMap", "set", "get", "get", "set", "get", "get"},
				payloads: [][]string{{}, {"foo", "bar", "1"}, {"foo", "1"}, {"foo", "3"}, {"foo", "bar2", "4"}, {"foo", "4"}, {"foo", "5"}},
			},
			want: []string{"null", "null", "bar", "bar", "null", "bar2", "bar2"},
		},
		{
			name: "Example2",
			args: args{commands: []string{"TimeMap", "set", "set", "get", "get", "get", "get", "get"},
				payloads: [][]string{{}, {"love", "high", "10"}, {"love", "low", "20"}, {"love", "5"}, {"love", "10"}, {"love", "15"}, {"love", "20"}, {"love", "25"}},
			},
			want: []string{"null", "null", "null", "", "high", "high", "low", "low"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runTest(tt.args.commands, tt.args.payloads); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
