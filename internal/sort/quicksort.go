package sort


import "reflect"

func quickSort(items reflect.Value, itemsType reflect.Type, function interface{}) reflect.Value {
	less := reflect.MakeSlice(reflect.SliceOf(itemsType), 0, items.Len())
	middle := reflect.MakeSlice(reflect.SliceOf(itemsType), 0, items.Len())
	more := reflect.MakeSlice(reflect.SliceOf(itemsType), 0, items.Len())
	fn := reflect.ValueOf(function)
	if items.Len() == 0 {
		return less
	}
	if items.Len() == 1 {
		less = reflect.Append(less, items.Index(0))
		return less
	}
	m := items.Index(0)
	for index := 0; index < items.Len(); index++ {
		item := items.Index(index)
		argv := make([]reflect.Value, 2)
		argv[0] = m
		argv[1] = item
		result := fn.Call(argv)[0]
		switch result.Int() {
		case 1:
			less = reflect.Append(less, item)
		case 0:
			middle = reflect.Append(middle, item)
		case -1:
			more = reflect.Append(more, item)
		}
	}
	less = quickSort(reflect.ValueOf(less.Interface()), itemsType, function)
	more = quickSort(reflect.ValueOf(more.Interface()), itemsType, function)
	less = reflect.AppendSlice(less, middle)
	less = reflect.AppendSlice(less, more)
	return less
}
