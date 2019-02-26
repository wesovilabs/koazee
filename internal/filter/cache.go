package filter

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*filterInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *filterInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*filterInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *filterInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type filterInfo struct {
	fnValue   reflect.Value
	fnIn1Type reflect.Type
	fnIn2Type reflect.Type
}
