package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if reflect.TypeOf(field) == reflect.Type{string} {
			fn(field.String())
		}
	}
}
