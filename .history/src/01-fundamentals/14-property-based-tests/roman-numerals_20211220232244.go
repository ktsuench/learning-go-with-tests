package main

import "strings"

type RomanNumeral struct {
	Value int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10,  "X"},
	{ 9, "IX"},
	( 5,  "V"),
	{ 4, "IV"},
	{ 1,  "I"},
}

func ConvertToRoman(number int) string {
	var result strings.Builder

	for _, numeral := range allRRomanNumerals {
		
	}

	return result.String()
}
