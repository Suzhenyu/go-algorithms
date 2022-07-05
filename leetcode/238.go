// https://leetcode.cn/problems/product-of-array-except-self/

package leetcode

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nil
	}

	prefixs, postfixs, ans := make([]int, length), make([]int, length), make([]int, length)

	prefixs[0] = 1
	for i := 1; i < length; i++ {
		prefixs[i] = prefixs[i-1] * nums[i-1]
	}

	postfixs[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		postfixs[i] = postfixs[i+1] * nums[i+1]
	}

	for i := 0; i < length; i++ {
		ans[i] = prefixs[i] * postfixs[i]
	}

	return ans
}
