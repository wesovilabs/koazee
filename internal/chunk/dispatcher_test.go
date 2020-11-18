package chunk

import (
	"reflect"
	"testing"

	"github.com/wesovilabs/koazee/utils"
)

func Test_dispatch(t *testing.T) {
	type args struct {
		items     reflect.Value
		itemsType reflect.Type
		size      uint
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 reflect.Value
	}{
		{
			name: "string",
			args: args{
				items:     reflect.ValueOf(utils.ArrayString),
				itemsType: reflect.TypeOf(utils.ValueString),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]string{utils.ArrayString[0:2:2], utils.ArrayString[2:4:4], utils.ArrayString[4:5:5]}),
		},
		{
			name: "*string",
			args: args{
				items:     reflect.ValueOf(utils.ArrayStringPtr),
				itemsType: reflect.TypeOf(utils.ValueStringPtr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*string{utils.ArrayStringPtr[0:2:2], utils.ArrayStringPtr[2:4:4], utils.ArrayStringPtr[4:5:5]}),
		},
		{
			name: "bool",
			args: args{
				items:     reflect.ValueOf(utils.ArrayBool),
				itemsType: reflect.TypeOf(utils.ValueBool),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]bool{utils.ArrayBool[0:2:2], utils.ArrayBool[2:4:4], utils.ArrayBool[4:5:5]}),
		},
		{
			name: "*bool",
			args: args{
				items:     reflect.ValueOf(utils.ArrayBoolPtr),
				itemsType: reflect.TypeOf(utils.ValueBoolPtr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*bool{utils.ArrayBoolPtr[0:2:2], utils.ArrayBoolPtr[2:4:4], utils.ArrayBoolPtr[4:5:5]}),
		},
		{
			name: "int",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt),
				itemsType: reflect.TypeOf(utils.ValueInt),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]int{utils.ArrayInt[0:2:2], utils.ArrayInt[2:4:4], utils.ArrayInt[4:5:5]}),
		},
		{
			name: "*int",
			args: args{
				items:     reflect.ValueOf(utils.ArrayIntPtr),
				itemsType: reflect.TypeOf(utils.ValueIntPtr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*int{utils.ArrayIntPtr[0:2:2], utils.ArrayIntPtr[2:4:4], utils.ArrayIntPtr[4:5:5]}),
		},
		{
			name: "int8",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt8),
				itemsType: reflect.TypeOf(utils.ValueInt8),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]int8{utils.ArrayInt8[0:2:2], utils.ArrayInt8[2:4:4], utils.ArrayInt8[4:5:5]}),
		},
		{
			name: "*int8",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt8Ptr),
				itemsType: reflect.TypeOf(utils.ValueInt8Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*int8{utils.ArrayInt8Ptr[0:2:2], utils.ArrayInt8Ptr[2:4:4], utils.ArrayInt8Ptr[4:5:5]}),
		},
		{
			name: "int16",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt16),
				itemsType: reflect.TypeOf(utils.ValueInt16),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]int16{utils.ArrayInt16[0:2:2], utils.ArrayInt16[2:4:4], utils.ArrayInt16[4:5:5]}),
		},
		{
			name: "*int16",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt16Ptr),
				itemsType: reflect.TypeOf(utils.ValueInt16Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*int16{utils.ArrayInt16Ptr[0:2:2], utils.ArrayInt16Ptr[2:4:4], utils.ArrayInt16Ptr[4:5:5]}),
		},
		{
			name: "int32",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt32),
				itemsType: reflect.TypeOf(utils.ValueInt32),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]int32{utils.ArrayInt32[0:2:2], utils.ArrayInt32[2:4:4], utils.ArrayInt32[4:5:5]}),
		},
		{
			name: "*int32",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt32Ptr),
				itemsType: reflect.TypeOf(utils.ValueInt32Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*int32{utils.ArrayInt32Ptr[0:2:2], utils.ArrayInt32Ptr[2:4:4], utils.ArrayInt32Ptr[4:5:5]}),
		},
		{
			name: "int64",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt64),
				itemsType: reflect.TypeOf(utils.ValueInt64),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]int64{utils.ArrayInt64[0:2:2], utils.ArrayInt64[2:4:4], utils.ArrayInt64[4:5:5]}),
		},
		{
			name: "*int64",
			args: args{
				items:     reflect.ValueOf(utils.ArrayInt64Ptr),
				itemsType: reflect.TypeOf(utils.ValueInt64Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*int64{utils.ArrayInt64Ptr[0:2:2], utils.ArrayInt64Ptr[2:4:4], utils.ArrayInt64Ptr[4:5:5]}),
		},
		{
			name: "uint",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint),
				itemsType: reflect.TypeOf(utils.ValueUint),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]uint{utils.ArrayUint[0:2:2], utils.ArrayUint[2:4:4], utils.ArrayUint[4:5:5]}),
		},
		{
			name: "*uint",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUintPtr),
				itemsType: reflect.TypeOf(utils.ValueUintPtr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*uint{utils.ArrayUintPtr[0:2:2], utils.ArrayUintPtr[2:4:4], utils.ArrayUintPtr[4:5:5]}),
		},
		{
			name: "uint8",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint8),
				itemsType: reflect.TypeOf(utils.ValueUint8),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]uint8{utils.ArrayUint8[0:2:2], utils.ArrayUint8[2:4:4], utils.ArrayUint8[4:5:5]}),
		},
		{
			name: "*uint8",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint8Ptr),
				itemsType: reflect.TypeOf(utils.ValueUint8Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*uint8{utils.ArrayUint8Ptr[0:2:2], utils.ArrayUint8Ptr[2:4:4], utils.ArrayUint8Ptr[4:5:5]}),
		},
		{
			name: "uint16",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint16),
				itemsType: reflect.TypeOf(utils.ValueUint16),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]uint16{utils.ArrayUint16[0:2:2], utils.ArrayUint16[2:4:4], utils.ArrayUint16[4:5:5]}),
		},
		{
			name: "*uint16",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint16Ptr),
				itemsType: reflect.TypeOf(utils.ValueUint16Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*uint16{utils.ArrayUint16Ptr[0:2:2], utils.ArrayUint16Ptr[2:4:4], utils.ArrayUint16Ptr[4:5:5]}),
		},
		{
			name: "uint32",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint32),
				itemsType: reflect.TypeOf(utils.ValueUint32),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]uint32{utils.ArrayUint32[0:2:2], utils.ArrayUint32[2:4:4], utils.ArrayUint32[4:5:5]}),
		},
		{
			name: "*uint32",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint32Ptr),
				itemsType: reflect.TypeOf(utils.ValueUint32Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*uint32{utils.ArrayUint32Ptr[0:2:2], utils.ArrayUint32Ptr[2:4:4], utils.ArrayUint32Ptr[4:5:5]}),
		},
		{
			name: "uint64",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint64),
				itemsType: reflect.TypeOf(utils.ValueUint64),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]uint64{utils.ArrayUint64[0:2:2], utils.ArrayUint64[2:4:4], utils.ArrayUint64[4:5:5]}),
		},
		{
			name: "*uint64",
			args: args{
				items:     reflect.ValueOf(utils.ArrayUint64Ptr),
				itemsType: reflect.TypeOf(utils.ValueUint64Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*uint64{utils.ArrayUint64Ptr[0:2:2], utils.ArrayUint64Ptr[2:4:4], utils.ArrayUint64Ptr[4:5:5]}),
		},
		{
			name: "float32",
			args: args{
				items:     reflect.ValueOf(utils.ArrayFloat32),
				itemsType: reflect.TypeOf(utils.ValueFloat32),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]float32{utils.ArrayFloat32[0:2:2], utils.ArrayFloat32[2:4:4], utils.ArrayFloat32[4:5:5]}),
		},
		{
			name: "*float32",
			args: args{
				items:     reflect.ValueOf(utils.ArrayFloat32Ptr),
				itemsType: reflect.TypeOf(utils.ValueFloat32Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*float32{utils.ArrayFloat32Ptr[0:2:2], utils.ArrayFloat32Ptr[2:4:4], utils.ArrayFloat32Ptr[4:5:5]}),
		},
		{
			name: "float64",
			args: args{
				items:     reflect.ValueOf(utils.ArrayFloat64),
				itemsType: reflect.TypeOf(utils.ValueFloat64),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]float64{utils.ArrayFloat64[0:2:2], utils.ArrayFloat64[2:4:4], utils.ArrayFloat64[4:5:5]}),
		},
		{
			name: "*float64",
			args: args{
				items:     reflect.ValueOf(utils.ArrayFloat64Ptr),
				itemsType: reflect.TypeOf(utils.ValueFloat64Ptr),
				size:      2,
			},
			want:  true,
			want1: reflect.ValueOf([][]*float64{utils.ArrayFloat64Ptr[0:2:2], utils.ArrayFloat64Ptr[2:4:4], utils.ArrayFloat64Ptr[4:5:5]}),
		},
		{
			name: "complex",
			args: args{
				items:     reflect.ValueOf(utils.Person{}),
				itemsType: reflect.TypeOf([]utils.Person{}),
				size:      2,
			},
			want:  false,
			want1: reflect.ValueOf(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := dispatch(tt.args.items, tt.args.itemsType, tt.args.size)
			if got != tt.want {
				t.Errorf("dispatch() got = %v, want %v", got, tt.want)
			}
			if got1 == reflect.ValueOf(nil) && tt.want1 == reflect.ValueOf(nil) {
				return
			}
			if !reflect.DeepEqual(got1.Interface(), tt.want1.Interface()) {
				t.Errorf("dispatch() got1 = %v, want %v", got1.Interface(), tt.want1.Interface())
			}
		})
	}
}
