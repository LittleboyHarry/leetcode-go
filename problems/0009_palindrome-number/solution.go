package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(x int) int {
	r := 0
	for x != 0 {
		r *= 10
		r += x % 10
		x /= 10
	}
	return r
}
