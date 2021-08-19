package simple_test

import (
	"github.com/Evolt0/design_pattern/factory/simple"
	"reflect"
	"testing"
)

func TestNewIConfigParser(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want simple.IConfigParser
	}{
		{
			name: "json",
			args: args{data: "json"},
			want: &simple.JsonConfigParser{},
		},
		{
			name: "yaml",
			args: args{data: "yaml"},
			want: &simple.YamlConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simple.NewIConfigParser(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
