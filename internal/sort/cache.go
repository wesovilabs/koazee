package sort

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*sortInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *sortInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*sortInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *sortInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type sortInfo struct {
	fnValue      reflect.Value
	fnInput1Type reflect.Type
	fnInput2Type reflect.Type
	fnOutputType reflect.Type
}
