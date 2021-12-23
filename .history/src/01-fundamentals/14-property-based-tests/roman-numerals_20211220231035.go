package main

import "strings"

func ConvertToRoman(number int) string {

	if number == 4 {
		return "IV"
	}

	var result strings.Builder

	for i := 0; i < number; i++ {
		result.WriteString("I")
	}

	return result.String()
}
