package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// key is rune and value is index
	letters := make(map[rune]int)
	// convert string to slice of runes
	runes := []rune(s)
	n, maxLength := len(s), 0

	for i, j := 0, 0; j < n; j++ {
		r := runes[j]
		if pos, ok := letters[r]; ok {
			// only move i forward
			if pos >= i {
				i = pos + 1
			}
		}

		if nowlength := j - i + 1; nowlength > maxLength {
			maxLength = nowlength
		}

		letters[r] = j
	}

	return maxLength
}

func main() {
	TesText := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(TesText))
}
