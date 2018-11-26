package drop

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*dropInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, itemType reflect.Type, item *dropInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*dropInfo)
	}
	cache[itemsType][itemType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *dropInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type dropInfo struct {
	itemValue reflect.Value
	itemType  *reflect.Type
}
