package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for j := 0; j < val.Len(); j++ {
			walk(val.Index(j).Interface(), fn)
		}
		return
	}

		switch field.Kind() {
		case reflect.Struct:
			for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		case reflect.String:
			fn(field.String())
		}
	
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
