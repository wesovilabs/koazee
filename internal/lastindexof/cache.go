package lastindexof

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*lastIndexOfInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *lastIndexOfInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*lastIndexOfInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *lastIndexOfInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type lastIndexOfInfo struct {
}
