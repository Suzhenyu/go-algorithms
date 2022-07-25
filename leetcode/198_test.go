package leetcode

import "testing"

func TestRob1(t *testing.T) {
	inputs := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
	}
	excepteds := []int{0, 1, 2, 4, 12}

	for index, input := range inputs {
		if excepted, output := excepteds[index], rob1(input); excepted != output {
			t.Errorf("Input: %v, excepted: %v, output: %v.\n", input, excepted, output)
		}
	}
}
