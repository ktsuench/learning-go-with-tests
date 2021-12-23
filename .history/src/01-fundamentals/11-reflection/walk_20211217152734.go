package main

func Walk(x interface{}, fn func(input string)) {
	for _, v := range x {
		switch t := v.(type) {
		case string:
			fn(t)
		default:
			// do nothing
		}
	}
}
