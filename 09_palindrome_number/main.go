package main

import "fmt"

func isPalindrome(x int) bool {
	// negative number
	if x < 0 {
		return false
	}
	// digit
	if x < 10 {
		return true
	}
	// tens digit
	if x%10 == 0 {
		return false
	}
	// hundreds digit
	if x < 100 && x%11 == 0 {
		return true
	}
	// thousand digit
	if x < 1000 && ((x/100)*10+x%10)%11 == 0 {
		return true
	}

	// actual logic
	v := x % 10
	x /= 10
	for x-v > 0 {
		v = v*10 + x%10
		x /= 10
	}
	// if len(x) is odd, Handle middle digit
	if v > x {
		v /= 10
	}

	if v == x {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(123321))
}
