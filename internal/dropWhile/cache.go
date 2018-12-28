package dropWhile

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*dropWhileInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *dropWhileInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*dropWhileInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *dropWhileInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type dropWhileInfo struct {
	fnValue     reflect.Value
	fnInputType reflect.Type
}
