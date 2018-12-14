package indexof

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*indexOfInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *indexOfInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*indexOfInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *indexOfInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type indexOfInfo struct {
}
