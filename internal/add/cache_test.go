package add

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAdd_cacheAdd(t *testing.T) {
	cache = cacheType{}
	itemsType := reflect.TypeOf("")
	itemType := reflect.TypeOf("")
	info := &addInfo{
		itemType:  &itemType,
	}
	cache.add(itemsType, itemType, info)
	assert.Equal(t, 1, len(cache))
}

func TestAdd_cacheGetUnExistingElement(t *testing.T) {
	cache = cacheType{}
	itemsType := reflect.TypeOf("")
	itemType := reflect.TypeOf("")
	assert.Nil(t, cache.get(itemsType, itemType))
}

func TestAdd_cacheGetExistingElement(t *testing.T) {
	itemsType := reflect.TypeOf("")
	itemType := reflect.TypeOf("")
	info := &addInfo{
		itemType:  &itemType,
	}
	cache.add(itemsType, itemType, info)
	assert.Equal(t, info, cache.get(itemsType, itemType))
}
