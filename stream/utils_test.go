package stream

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObj struct {
	name string
}

func Test_isPointer(t *testing.T) {
	assert.Equal(
		t,
		true,
		isPointer(reflect.ValueOf(&testObj{})),
	)
	assert.Equal(
		t,
		false,
		isPointer(reflect.ValueOf(testObj{})),
	)
}

func Test_equalsValues(t *testing.T) {

	assert.Equal(
		t,
		true,
		equalsValues(reflect.ValueOf(&testObj{"John"}), reflect.ValueOf(&testObj{"John"})),
	)
	assert.Equal(
		t,
		false,
		equalsValues(reflect.ValueOf(&testObj{"David"}), reflect.ValueOf(&testObj{"John"})),
	)

	assert.Equal(
		t,
		true,
		equalsValues(reflect.ValueOf(4), reflect.ValueOf(4)),
	)
	assert.Equal(
		t,
		false,
		equalsValues(reflect.ValueOf(4), reflect.ValueOf(3)),
	)
}

func Test_equalsPtr(t *testing.T) {
	assert.Equal(
		t,
		true,
		equalsPtr(reflect.ValueOf(&testObj{"John"}), reflect.ValueOf(&testObj{"John"})),
	)
	assert.Equal(
		t,
		false,
		equalsPtr(reflect.ValueOf(&testObj{"David"}), reflect.ValueOf(&testObj{"John"})),
	)
}

func Test_equalsValue(t *testing.T) {
	assert.Equal(
		t,
		true,
		equalsValue(reflect.ValueOf(4), reflect.ValueOf(4)),
	)
	assert.Equal(
		t,
		false,
		equalsValue(reflect.ValueOf(4), reflect.ValueOf(3)),
	)
}
