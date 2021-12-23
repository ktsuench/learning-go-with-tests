package main

import "go/types"

func Walk(x interface{}, fn func(input string)) {
	for _, v := range x {
		switch x:=typeof x {
		case v:

		default:
			// do nothing
		}
	}
}
