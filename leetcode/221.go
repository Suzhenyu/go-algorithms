package leetcode

func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for side := maxSide + 1; j+side-1 < len(matrix[i]); side++ {
				if !isSquare(matrix, i, j, side) {
					break
				}
				maxSide = max(maxSide, side)
			}
		}
	}
	return maxSide * maxSide
}

func isSquare(matrix [][]byte, x, y, side int) bool {
	for i := x; i < x+side; i++ {
		for j := y; j < y+side; j++ {
			if i >= len(matrix) || j >= len(matrix[i]) || matrix[i][j] != '1' {
				return false
			}
		}
	}
	return true
}
