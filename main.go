package main

import (
	"fmt"
	twosum "leetcode-go/twosum"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twosum.TwoSum(nums, target)
	fmt.Printf("Two Sum - Input: nums = %v, target = %d\n", nums, target)
	fmt.Printf("Two Sum - Output: %v\n", result)
}
