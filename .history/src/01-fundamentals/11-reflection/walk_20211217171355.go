package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch k:=field.Kind() {

		case reflect.String:
			fn(field.String())
		}
	}
}
