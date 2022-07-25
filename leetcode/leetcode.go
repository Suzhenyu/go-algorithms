package leetcode

import "math"

func min(nums ...int) int {
	result := math.MaxInt
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}

func max(nums ...int) int {
	result := math.MinInt
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}
