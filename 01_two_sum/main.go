package main

import "log"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	answer := twoSum(nums, target)
	for v := range answer {
		if v != 1 && v != 0 {
			log.Fatal("False!")
		}
	}
	log.Println("Accept")
}

func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		secondTarget := target - v
		for i := 0; i < len(nums) && i != k; i++ {
			if nums[i] == secondTarget {
				return []int{k, i}
			}
		}
	}
	return nil
}
