package main

import (
	"fmt"
	combinationsum "leetcode-go/combinationsum"
	threesum "leetcode-go/threesum"
	twosum "leetcode-go/twosum"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twosum.TwoSum(nums, target)
	fmt.Printf("Two Sum - Input: nums = %v, target = %d\n", nums, target)
	fmt.Printf("Two Sum - Output: %v\n", result)

	nums3 := []int{-1, 0, 1, 2, -1, -4}
	result3 := threesum.ThreeSum(nums3)
	fmt.Printf("Three Sum - Input: nums = %v\n", nums3)
	fmt.Printf("Three Sum - Output: %v\n", result3)

	candidates := []int{2, 3, 6, 7}
	target2 := 7
	result2 := combinationsum.CombinationSum(candidates, target2)
	fmt.Printf("Combination Sum - Input: candidates = %v, target = %d\n", candidates, target2)
	fmt.Printf("Combination Sum - Output: %v\n", result2)
}
