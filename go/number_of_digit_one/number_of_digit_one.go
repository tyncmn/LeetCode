package numberofdigitone

import (
	"strconv"
)

func NumberOfDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	var numberOfDigitOne int = 0

	for i := 1; i <= n; i++ {
		str := strconv.Itoa(i)
		for _, ch := range str {
			if ch == '1' {
				numberOfDigitOne++
			}
		}
	}
	return numberOfDigitOne
}
