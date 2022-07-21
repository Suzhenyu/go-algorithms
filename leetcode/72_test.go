package leetcode

import "testing"

func TestMinDistance(t *testing.T) {
	inputs := [][2]string{
		{"horse", "ros"},
		{"intention", "execution"},
	}
	excepteds := []int{
		3,
		5,
	}

	for index, input := range inputs {
		if excepted, output := excepteds[index], minDistance(input[0], input[1]); excepted != output {
			t.Errorf("Input: %v, excepted: %v, output: %v.\n", input, excepted, output)
		}
	}
}
