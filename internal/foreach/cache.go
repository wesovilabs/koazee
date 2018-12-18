package foreach

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*forEachInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcInType reflect.Type, item *forEachInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*forEachInfo)
	}
	cache[itemsType][funcInType] = item
}

func (c *cacheType) get(itemsType, funcInType reflect.Type) *forEachInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcInType]; ok {
			return val2
		}
	}
	return nil
}

type forEachInfo struct {
	fnValue  reflect.Value
	hasError bool
}
