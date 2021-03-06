package main

import "strings"

func ConvertToRoman(number int) string {
	var result strings.Builder

	for i := number; i > 0; i-- {
		if i == 5 {
			result.WriteString("V")
			break
		}
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
