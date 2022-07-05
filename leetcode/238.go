// https://leetcode.cn/problems/product-of-array-except-self/

package leetcode

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nil
	}

	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ans[i] = 1
	}

	prefix, suffix := 1, 1
	for i := 0; i < length; i++ {
		ans[i] *= prefix
		ans[length-1-i] *= suffix
		prefix *= nums[i]
		suffix *= nums[length-1-i]
	}

	return ans
}
