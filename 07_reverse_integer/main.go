package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	var result string
	if x < 0 {
		result += "-"
		x = 0 - x
	}

	result += reverseString(strconv.Itoa(x))
	x, _ = strconv.Atoi(result)

	// especial value
	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	return x
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverse(1534236469))
}
