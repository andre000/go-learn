package property

import (
	"errors"
	"strings"
)

const maxNumber = 3999

// ErrBigInt error given when trying to convert numbers greater than 3999
var ErrBigInt = errors.New("❌ number cannot be greater than 3999")

// RomanNumeral list of numbers and its correnspondent in roman
type RomanNumeral struct {
	Value  uint16
	Symbol string
}

// RomanNumerals list do RomanNumeral
type RomanNumerals []RomanNumeral

var romanNumerals = RomanNumerals{
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

// ValueOf translates given symbol to uint16
func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

// ToRoman transforms int to roman
func ToRoman(number uint16) (string, error) {
	err := checkNumber(number)

	if err != nil {
		return "", err
	}

	var result strings.Builder

	for _, roman := range romanNumerals {
		for number >= roman.Value {
			result.WriteString(roman.Symbol)
			number -= roman.Value
		}
	}
	return result.String(), nil
}

// ToArabic transforms roman to int
func ToArabic(roman string) (total uint16) {
	total = 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]
			value := romanNumerals.ValueOf(symbol, nextSymbol)

			if value != 0 {
				total += value
				i++
				continue
			}
		}

		total += romanNumerals.ValueOf(symbol)
	}
	return
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}

func checkNumber(arabic uint16) error {
	if arabic > maxNumber {
		return ErrBigInt
	}
	return nil
}
