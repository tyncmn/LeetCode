package reverse_integer

func Reverse(x int) int {
	var result int = 0
	var isMinus bool = false

	if x < 0 {
		isMinus = true
		x = -x
	}
	for x != 0 {
		digit := x % 10
		x /= 10
		result = result*10 + digit
	}
	if isMinus {
		result = -result
	}
	if result < -2147483648 || result > 2147483647 {
		return 0
	}
	return result
}
