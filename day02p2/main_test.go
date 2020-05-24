package main

import "testing"

func TestExecute(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{99}, []int{99}},
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for i, test := range tests {
		execute(test.input)
		for j := 0; j < len(test.input); j++ {
			if test.input[j] != test.output[j] {
				t.Errorf("test #%d after execute = %v, want %v", i, test.input, test.output)
				break
			}
		}
	}
}
