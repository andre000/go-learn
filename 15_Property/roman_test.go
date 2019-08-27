package property

import (
	"testing"
	"testing/quick"
)

var cases = []struct {
	Description string
	Number      uint16
	Roman       string
}{
	{"should be able to convert between 1 and I", 1, "I"},
	{"should be able to convert between 2 and II", 2, "II"},
	{"should be able to convert between 3 and III", 3, "III"},
	{"should be able to convert between 4 and IV", 4, "IV"},
	{"should be able to convert between 5 and V", 5, "V"},
	{"should be able to convert between 6 and VI", 6, "VI"},
	{"should be able to convert between 7 and VII", 7, "VII"},
	{"should be able to convert between 8 and VIII", 8, "VIII"},
	{"should be able to convert between 9 and IX", 9, "IX"},
	{"should be able to convert between 10 and X", 10, "X"},
	{"should be able to convert between 14 and XIV", 14, "XIV"},
	{"should be able to convert between 18 and XVIII", 18, "XVIII"},
	{"should be able to convert between 20 and XX", 20, "XX"},
	{"should be able to convert between 39 and XXXIX", 39, "XXXIX"},
	{"should be able to convert between 40 and XL", 40, "XL"},
	{"should be able to convert between 47 and XLVII", 47, "XLVII"},
	{"should be able to convert between 49 and XLIX", 49, "XLIX"},
	{"should be able to convert between 50 and L", 50, "L"},
	{"should be able to convert between 91 and XCI", 91, "XCI"},
	{"should be able to convert between 150 and CL", 150, "CL"},
	{"should be able to convert between 412 and CDXII", 412, "CDXII"},
	{"should be able to convert between 501 and DI", 501, "DI"},
	{"should be able to convert between 999 and CMXCIX", 999, "CMXCIX"},
	{"should be able to convert between 1984 and MCMLXXXIV", 1984, "MCMLXXXIV"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			received, _ := ToRoman(test.Number)
			expected := test.Roman

			assertString(t, received, expected)
		})
	}

	t.Run("Should throw an error when trying to convert a number greater than 3999", func(t *testing.T) {
		_, err := ToRoman(5000)
		if err != ErrBigInt {
			t.Error("❌ expected error received none")
		}
	})
}

func TestArabicNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			received := ToArabic(test.Roman)
			expected := test.Number

			assertInt(t, received, expected)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		t.Log("🔨 Testing", arabic)
		roman, _ := ToRoman(arabic)
		fromRoman := ToArabic(roman)
		return fromRoman == arabic
	}

	quickConfig := &quick.Config{
		MaxCount: 1000,
	}

	if err := quick.Check(assertion, quickConfig); err != nil {
		t.Error("❌ failed checks", err)
	}
}

func assertString(t *testing.T, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %q expected %q", received, expected)
	}
}

func assertInt(t *testing.T, received, expected uint16) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %d expected %d", received, expected)
	}
}
