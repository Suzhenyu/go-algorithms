// 213. 打家劫舍 II
// https://leetcode.cn/problems/house-robber-ii/

package leetcode

func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums...)
	}
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
