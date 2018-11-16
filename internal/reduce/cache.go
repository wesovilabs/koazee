package reduce

import "reflect"

var cache map[reflect.Type]*reduceInfo

type reduceInfo struct {
	fnType    reflect.Type
	fnIn1Type reflect.Type
	fnIn2Type reflect.Type
	fnOutType reflect.Type
	fnValue   reflect.Value
}
