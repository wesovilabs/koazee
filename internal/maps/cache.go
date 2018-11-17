package maps

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*mapInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *mapInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*mapInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *mapInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type mapInfo struct {
	fnValue      reflect.Value
	fnInputType  reflect.Type
	fnOutputType reflect.Type
	isPtr        bool
}
