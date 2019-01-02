package indexesof

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*indexesOfInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *indexesOfInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*indexesOfInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *indexesOfInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type indexesOfInfo struct {
}
