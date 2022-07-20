package leetcode

import "testing"

func TestNumSquares(t *testing.T) {
	inputs := []int{12, 13}
	excepteds := []int{3, 2}

	for index, input := range inputs {
		if output, excepted := numSquares(input), excepteds[index]; output != excepted {
			t.Errorf("Input: %v, excepted: %v, output: %v.\n", input, excepted, output)
		}
	}
}
