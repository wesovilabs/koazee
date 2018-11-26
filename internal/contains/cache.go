package contains

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*containsInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *containsInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*containsInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *containsInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type containsInfo struct {
}
