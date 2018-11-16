package reduce

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*reduceInfo

var cache cacheType

func (c *cacheType) add(itemsType, funcType reflect.Type, item *reduceInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*reduceInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *reduceInfo {
	if val, ok := cache[itemsType]; ok {
		if val, ok := val[funcType]; ok {
			return val
		}
	}
	return nil
}

type reduceInfo struct {
	fnType    reflect.Type
	fnIn1Type reflect.Type
	fnIn2Type reflect.Type
	fnOutType reflect.Type
	fnValue   reflect.Value
}
