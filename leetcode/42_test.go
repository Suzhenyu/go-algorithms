package leetcode

import "testing"

func TestTrap(t *testing.T) {
	inputs := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
	}
	excepted := []int{
		6,
		9,
	}

	for index, input := range inputs {
		if output := trap(input); output != excepted[index] {
			t.Errorf("Input: %v, excepted: %d, output: %d", input, excepted[index], output)
		}
	}
}
