// https://leetcode.cn/problems/product-of-array-except-self/

package leetcode

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nil
	}

	ans := make([]int, length)
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	R := 1
	for i := length - 2; i >= 0; i-- {
		R *= nums[i+1]
		ans[i] *= R
	}

	return ans
}
