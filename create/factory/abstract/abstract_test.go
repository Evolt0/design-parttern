package abstract

import (
	"reflect"
	"testing"
)

func TestJsonConfigParserFactory_CreateParser(t *testing.T) {
	tests := []struct {
		name string
		want IConfigParser
	}{
		{
			name: "json",
			want: &JsonConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &JsonConfigParserFactory{}
			if got := i.CreateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
