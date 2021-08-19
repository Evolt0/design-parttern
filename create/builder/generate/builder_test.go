package generate

import (
	"reflect"
	"testing"
)

func TestBuilder_Build(t *testing.T) {
	type fields struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{name: "build",
			fields: fields{
				data: 10,
			},
			want: &Config{data: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				data: tt.fields.data,
			}
			if got := b.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
