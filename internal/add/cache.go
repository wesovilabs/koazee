package add

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]*addInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, itemType reflect.Type, item *addInfo) {
	if _, ok := cache[itemsType]; !ok {
		cache[itemsType] = make(map[reflect.Type]*addInfo)
	}
	cache[itemsType][itemType] = item
}

func (c *cacheType) get(itemsType, funcType reflect.Type) *addInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcType]; ok {
			return val2
		}
	}
	return nil
}

type addInfo struct {
	itemType  *reflect.Type
}
