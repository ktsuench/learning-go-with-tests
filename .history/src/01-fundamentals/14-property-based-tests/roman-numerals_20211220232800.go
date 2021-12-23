package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
}

func ConvertToRoman(number int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for number >= numeral.Value {
			result.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	return result.String()
}
