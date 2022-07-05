package leetcode

import "testing"

func TestProductExceptSelf(t *testing.T) {
	inputs := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
	}
	excepted := [][]int{
		nil,
		nil,
		{2, 1},
		{6, 3, 2},
		{24, 12, 8, 6},
	}

	for index, input := range inputs {
		if output := productExceptSelf(input); !equal(output, excepted[index]) {
			t.Errorf("Input: %v, excepted: %v, output: %v", input, excepted[index], output)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}
