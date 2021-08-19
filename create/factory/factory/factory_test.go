package factory_test

import (
	"github.com/Evolt0/design_pattern/factory/factory"
	"reflect"
	"testing"
)

func TestNewIConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want factory.IConfigParserFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &factory.JsonConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &factory.YamlConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory.NewIConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
