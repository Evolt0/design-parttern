package for_go

import (
	"reflect"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *Builder
	}{
		{
			name: "builder for go",
			args: args{
				opts: []Option{
					func(builder *Builder) {
						builder.data = 10
					},
				},
			},
			want: &Builder{
				data: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuilder(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
