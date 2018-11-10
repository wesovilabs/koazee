package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := errors.InvalidType(stream.OpCodeForEach, "error")
	assert.Equal(
		t,
		err,
		stream.Error(err).Out().Err(),
	)
}

func TestNew(t *testing.T) {
	stream := stream.New([]int{2, 3})
	count, _ := stream.Count()
	assert.Equal(t, 2, count)
	assert.Equal(t, []int{2, 3}, stream.Out().Val())
}

func TestOutput_Bool(t *testing.T) {
	assert.Equal(
		t,
		false,
		stream.New([]bool{false, true}).First().Bool(),
	)
	assert.Equal(
		t,
		false,
		stream.New([]int{2, 3}).First().Bool(),
	)
}

func TestOutput_String(t *testing.T) {
	assert.Equal(
		t,
		"Stop",
		stream.New([]string{"Stop", "hunting"}).First().String(),
	)
	assert.Equal(
		t,
		"",
		stream.New([]int{2, 3}).First().String(),
	)
}

func TestOutput_Int(t *testing.T) {
	assert.Equal(
		t,
		3,
		stream.New([]int{3, 4}).First().Int(),
	)
	assert.Equal(
		t,
		0,
		stream.New([]bool{false, false}).First().Int(),
	)
}

func TestOutput_Int8(t *testing.T) {
	assert.Equal(
		t,
		int8(3),
		stream.New([]int8{3, 4}).First().Int8(),
	)
	assert.Equal(
		t,
		int8(0),
		stream.New([]bool{false, false}).First().Int8(),
	)
}

func TestOutput_Int16(t *testing.T) {
	assert.Equal(
		t,
		int16(3),
		stream.New([]int16{3, 4}).First().Int16(),
	)
	assert.Equal(
		t,
		int16(0),
		stream.New([]bool{false, false}).First().Int16(),
	)
}

func TestOutput_Int32(t *testing.T) {
	assert.Equal(
		t,
		int32(3),
		stream.New([]int32{3, 4}).First().Int32(),
	)
	assert.Equal(
		t,
		int32(0),
		stream.New([]bool{false, false}).First().Int32(),
	)
}

func TestOutput_Int64(t *testing.T) {
	assert.Equal(
		t,
		int64(3),
		stream.New([]int64{3, 4}).First().Int64(),
	)
	assert.Equal(
		t,
		int64(0),
		stream.New([]bool{false, false}).First().Int64(),
	)
}

func TestOutput_Uint(t *testing.T) {
	assert.Equal(
		t,
		uint(3),
		stream.New([]uint{3, 4}).First().Uint(),
	)
	assert.Equal(
		t,
		uint(0),
		stream.New([]bool{false, false}).First().Uint(),
	)
}

func TestOutput_Uint8(t *testing.T) {
	assert.Equal(
		t,
		uint8(3),
		stream.New([]uint8{3, 4}).First().Uint8(),
	)
	assert.Equal(
		t,
		uint8(0),
		stream.New([]bool{false, false}).First().Uint8(),
	)
}

func TestOutput_Uint16(t *testing.T) {
	assert.Equal(
		t,
		uint16(3),
		stream.New([]uint16{3, 4}).First().Uint16(),
	)
	assert.Equal(
		t,
		uint16(0),
		stream.New([]bool{false, false}).First().Uint16(),
	)
}

func TestOutput_Uint32(t *testing.T) {
	assert.Equal(
		t,
		uint32(3),
		stream.New([]uint32{3, 4}).First().Uint32(),
	)
	assert.Equal(
		t,
		uint32(0),
		stream.New([]bool{false, false}).First().Uint32(),
	)
}

func TestOutput_Uint64(t *testing.T) {
	assert.Equal(
		t,
		uint64(3),
		stream.New([]uint64{3, 4}).First().Uint64(),
	)
	assert.Equal(
		t,
		uint64(0),
		stream.New([]bool{false, false}).First().Uint64(),
	)
}

func TestOutput_Float32(t *testing.T) {
	assert.Equal(
		t,
		float32(3.23),
		stream.New([]float32{3.23, 4}).First().Float32(),
	)
	assert.Equal(
		t,
		float32(0.00),
		stream.New([]bool{false, false}).First().Float32(),
	)
}

func TestOutput_Float64(t *testing.T) {
	assert.Equal(
		t,
		float64(3.23),
		stream.New([]float64{3.23, 4}).First().Float64(),
	)
	assert.Equal(
		t,
		float64(0.00),
		stream.New([]bool{false, false}).First().Float64(),
	)
}

func TestOutput_Err(t *testing.T) {
	assert.Equal(
		t,
		errors.ItemsNil(stream.OpCodeFirst, "It can not be taken an element from a nil stream"),
		stream.New(nil).First().Err(),
	)
}

func TestOutput_Val(t *testing.T) {
	assert.Equal(
		t,
		[]string{"I", "don't", "eat", "my", "friends"},
		stream.New([]string{"I", "don't", "eat", "my", "friends"}).Out().Val(),
	)
}
