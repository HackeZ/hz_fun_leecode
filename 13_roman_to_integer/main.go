package main

import "fmt"

func romanToInt(s string) int {
	roman := make(map[byte]int)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	result := roman[s[0]]
	prev := result
	for i := 1; i < len(s); i++ {
		curr := roman[s[i]]
		if curr > prev {
			result += curr - 2*prev
		} else {
			result += curr
		}
		prev = curr
	}
	return result
}
func main() {
	fmt.Println(romanToInt("MCMXCVI"))
}
