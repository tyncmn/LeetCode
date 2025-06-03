package main

func RomanToInt(s string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	total := 0
	for i := n - 1; i >= 0; i-- {
		value := romanNumerals[rune(s[i])]
		if i < n-1 && value < romanNumerals[rune(s[i+1])] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}
