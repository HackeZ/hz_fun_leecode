package main

import "fmt"

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left, right := 0, len(height)-1
	maxA := 0
	for left < right {
		nowA := 0
		if height[left] < height[right] {
			nowA = (right - left) * height[left]
			left++
		} else {
			nowA = (right - left) * height[right]
			right--
		}
		if nowA > maxA {
			maxA = nowA
		}
	}
	return maxA
}

func main() {
	x := []int{1, 2, 3, 4, 6, 3, 4, 5}
	fmt.Println(maxArea(x))
}
