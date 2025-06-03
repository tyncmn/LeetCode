package main

func IntToRoman(num int) string {
	total := ""
	numerals := []struct {
		value  int
		symbol string
	}{
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
	for _, numeral := range numerals {
		for num >= numeral.value {
			total += numeral.symbol
			num -= numeral.value
		}
	}
	if num != 0 {
		return ""
	}
	return total
}
