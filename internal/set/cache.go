package set

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*setInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *setInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*setInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *setInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type setInfo struct {
	itemType *reflect.Type
	index    int
}
