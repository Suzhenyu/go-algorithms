package leetcode

import "testing"

func TestMaximalSquare(t *testing.T) {
	inputs := [][][]byte{
		{{'0'}},
		{{'0', '1'}, {'1', '0'}},
		{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
	}
	excepteds := []int{
		0, 1, 4,
	}

	for index, input := range inputs {
		if output, excepted := maximalSquare(input), excepteds[index]; output != excepted {
			t.Errorf("Input: %v, excepted: %v, output: %v.\n", input, excepted, output)
		}
	}
}
