package leetcode

import "testing"

func TestRob2(t *testing.T) {
	inputs := [][]int{
		{2, 3, 2},
		{1, 2, 3, 1},
		{1, 2, 3},
	}
	excepteds := []int{3, 4, 3}

	for index, input := range inputs {
		if excepted, output := excepteds[index], rob2(input); excepted != output {
			t.Errorf("Input: %v, excepted: %v, output: %v.\n", input, excepted, output)
		}
	}
}
