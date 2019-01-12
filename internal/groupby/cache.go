package groupby

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*groupByInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcType reflect.Type, item *groupByInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*groupByInfo)
	}
	cache[itemsType][funcType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *groupByInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type groupByInfo struct {
	fnValue      reflect.Value
	fnInputType  reflect.Type
	fnOutputType reflect.Type
	isPtr        bool
	hasError     bool
}
