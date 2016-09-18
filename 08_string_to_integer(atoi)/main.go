package main

import "fmt"

func myAtoi(str string) int {
	if str == "" {
		return 0
	}

	byt := []byte(str)
	var x int

	var k, sign, qt0 = 1, 0, 1
	for i := len(byt) - 1; i >= 0 && sign < 2; i-- {
		if byt[i] == ' ' {
			continue
		}
		if byt[i] == '-' || byt[i] == '+' {
			sign++
			if byt[i] == '-' {
				qt0 = -1
			}
			continue
		}
		if byt[i] < '0' || byt[i] > '9' {
			x, k = 0, 1
			continue
		}
		v := byt[i] - '0'
		x += int(v) * k
		k *= 10
	}

	if sign == 2 {
		return 0
	}
	return x * qt0
}

func main() {
	fmt.Println(myAtoi("   +0 123"))
}
