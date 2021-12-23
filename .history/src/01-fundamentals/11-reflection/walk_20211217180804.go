package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice:
		numOfValues = val.Len()
		walk(val.Index(j).Interface(), fn)

	case reflect.Struct:
		numOfValues = val.NumField()
		walk(val.Field(i).Interface(), fn)

	}

	for i := 0; i < numOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
