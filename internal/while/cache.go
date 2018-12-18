package while

import "reflect"

type cacheType map[reflect.Type]map[reflect.Type]map[reflect.Type]*whileInfo

var cache = cacheType{}

func (c *cacheType) add(itemsType, funcInType1, funcInType2 reflect.Type, item *whileInfo) {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcInType1]; ok {
			if _, ok := val2[funcInType2]; !ok {
				val2[funcInType2] = item
			}
		} else {
			val[funcInType1] = make(map[reflect.Type]*whileInfo)
			val[funcInType1][funcInType2] = item
		}
	} else {
		cache[itemsType] = make(map[reflect.Type]map[reflect.Type]*whileInfo)
		cache[itemsType][funcInType1] = make(map[reflect.Type]*whileInfo)
		cache[itemsType][funcInType1][funcInType2] = item
	}
}

func (c *cacheType) get(itemsType, funcInType1, funcInType2 reflect.Type) *whileInfo {
	if val, ok := cache[itemsType]; ok {
		if val2, ok := val[funcInType1]; ok {
			if val3, ok := val2[funcInType2]; ok {
				return val3
			}
		}
	}
	return nil
}

type whileInfo struct {
	fnValue        reflect.Value
	fnInputType    reflect.Type
	doFnValue      reflect.Value
	doFnInputType  reflect.Type
	doFnOutputType reflect.Type
}
