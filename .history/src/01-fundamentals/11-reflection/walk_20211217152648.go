package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	for _, v := range x {
		switch reflect.TypeOf(x) {
		case string:

		default:
			// do nothing
		}
	}
}
