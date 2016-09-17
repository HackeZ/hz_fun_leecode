package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var result = make([]string, numRows+1)
	sLength := len(s)
	var i, j int
	for i < sLength {
		for j = 0; j < numRows && i < sLength; j++ {
			result[j] += string(s[i])
			i++
		}
		for j = numRows - 2; j >= 1 && i < sLength; j-- {
			result[j] += string(s[i])
			i++
		}
	}

	var zipzag string
	for _, v := range result {
		zipzag += v
	}
	return zipzag
}

func main() {
	fmt.Println(convert("A", 2))
}
