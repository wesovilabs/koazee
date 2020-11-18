package stream_test

import (
	"reflect"
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"github.com/wesovilabs/koazee/utils"
)

func TestStream_Chunk(t *testing.T) {
	type args struct {
		size uint
	}
	tests := []struct {
		name string
		s    stream.Stream
		args args
		want stream.Stream
	}{
		{
			name: "normal",
			s:    koazee.StreamOf(utils.ArrayString).Chunk(2),
			args: args{size: 2},
			want: koazee.StreamOf([][]string{utils.ArrayString[0:2:2], utils.ArrayString[2:4:4], utils.ArrayString[4:5:5]}),
		},
		{
			name: "before",
			s:    koazee.StreamOf(utils.ArrayString).Map(func(v string) string { return v + "a" }).Chunk(2),
			args: args{size: 2},
			want: koazee.StreamOf([][]string{{utils.ArrayString[0] + "a", utils.ArrayString[1] + "a"}, {utils.ArrayString[2] + "a", utils.ArrayString[3] + "a"}, {utils.ArrayString[4] + "a"}}),
		},
		{
			name: "after",
			s:    koazee.StreamOf(utils.ArrayString).Chunk(2).Map(func(v []string) []string { return append(v, "a") }),
			args: args{size: 2},
			want: koazee.StreamOf([][]string{{utils.ArrayString[0], utils.ArrayString[1], "a"}, {utils.ArrayString[2], utils.ArrayString[3], "a"}, {utils.ArrayString[4], "a"}}),
		},
		{
			name: "beforeAndAfter",
			s:    koazee.StreamOf(utils.ArrayString).Map(func(v string) string { return v + "a" }).Chunk(2).Map(func(v []string) []string { return append(v, "b") }),
			args: args{size: 2},
			want: koazee.StreamOf([][]string{{utils.ArrayString[0] + "a", utils.ArrayString[1] + "a", "b"}, {utils.ArrayString[2] + "a", utils.ArrayString[3] + "a", "b"}, {utils.ArrayString[4] + "a", "b"}}),
		},
		{
			name: "emptyAfter",
			s:    koazee.StreamOf([]string{}).Chunk(2).Map(func(v []string) []string { return append(v, "a") }),
			args: args{size: 2},
			want: koazee.StreamOf([][]string{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s; !reflect.DeepEqual(got.Out().Val(), tt.want.Out().Val()) {
				t.Errorf("Stream.Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
