// https://leetcode.cn/problems/edit-distance/

package leetcode

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	if n*m == 0 {
		return n + m
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 边界状态初始化
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			left := dp[i][j-1] + 1
			down := dp[i-1][j] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down += 1
			}
			dp[i][j] = min(left, down, left_down)
		}
	}

	return dp[n][m]
}
