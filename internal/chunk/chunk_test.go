package chunk

import (
	"reflect"
	"testing"

	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/utils"
)

func TestChunk_Run(t *testing.T) {
	tests := []struct {
		name  string
		op    *Chunk
		want  reflect.Value
		want1 *errors.Error
	}{
		{
			name: "default",
			op: &Chunk{
				ItemsType:  reflect.TypeOf(utils.ValueString),
				ItemsValue: reflect.ValueOf(utils.ArrayString),
				Size:       2,
			},
			want:  reflect.ValueOf([][]string{utils.ArrayString[0:2:2], utils.ArrayString[2:4:4], utils.ArrayString[4:5:5]}),
			want1: nil,
		},
		{
			name: "complex",
			op: &Chunk{
				ItemsType: reflect.TypeOf(utils.Person{}),
				ItemsValue: reflect.ValueOf([]utils.Person{
					{"John", "Doe", 23, true},
					{"Jane", "Doe", 22, false},
					{"Jim", "Doe", 56, true},
					{"Jamie", "Doe", 89, false},
				}),
				Size: 2,
			},
			want: reflect.ValueOf([][]utils.Person{
				{
					{"John", "Doe", 23, true},
					{"Jane", "Doe", 22, false},
				},
				{
					{"Jim", "Doe", 56, true},
					{"Jamie", "Doe", 89, false},
				},
			}),
			want1: nil,
		},
		{
			name: "just",
			op: &Chunk{
				ItemsType:  reflect.TypeOf(utils.ValueString),
				ItemsValue: reflect.ValueOf(utils.ArrayString[0:2:2]),
				Size:       2,
			},
			want:  reflect.ValueOf([][]string{utils.ArrayString[0:2:2]}),
			want1: nil,
		},
		{
			name: "large",
			op: &Chunk{
				ItemsType:  reflect.TypeOf(utils.ValueString),
				ItemsValue: reflect.ValueOf(utils.ArrayString[0:1:1]),
				Size:       2,
			},
			want:  reflect.ValueOf([][]string{utils.ArrayString[0:1:1]}),
			want1: nil,
		},
		{
			name: "empty",
			op: &Chunk{
				ItemsType:  reflect.TypeOf(utils.ValueString),
				ItemsValue: reflect.ValueOf([]string{}),
				Size:       2,
			},
			want:  reflect.ValueOf([][]string{}),
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.op.Run()
			if !reflect.DeepEqual(got.Interface(), tt.want.Interface()) {
				t.Errorf("Chunk.Run() got = %v, want %v", got, tt.want)
			}
			if got1 != nil || tt.want1 != nil {
				t.Errorf("Chunk.Run() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
