package main

import "strings"

type (
	RomanNumeral struct {
		Value  int
		Symbol string
	}
	romanNumerals []RomanNumeral
	windowedRoman string
)

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			// build the two character string
			potentialNumber := []byte{symbol, nextSymbol}

			// get the value of the two character string
			value := allRomanNumerals.ValueOf(potentialNumber...)

			if value != 0 {
				total += value
				i++ // move past this char for the next loop
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (total int) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}

	return
}

func couldBeSubtractive(index int, curSymbol uint8, roman string) bool {
	isSubtractiveSymbol := curSymbol == 'I' || curSymbol == 'X' || curSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
