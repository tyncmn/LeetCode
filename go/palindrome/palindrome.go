package palindrome

import "fmt"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	var a int = 0
	var i int = 1
	for temp > 9 {
		i *= 10
		temp /= 10
	}
	fmt.Printf("i = %d\n", i)
	temp = x
	fmt.Println(temp)
	for i > 0 {
		a += (temp % 10) * i
		fmt.Printf("a = %d\n", a)
		temp /= 10
		i /= 10
	}
	return a == x
}
