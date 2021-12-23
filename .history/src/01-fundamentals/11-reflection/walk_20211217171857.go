package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Ptr:
			walk(field.Elem() fn)
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
