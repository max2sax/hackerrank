package coinchange

import (
	"fmt"
	"testing"
)

func TestMakeChange(t *testing.T) {
	tests := []struct {
		coins    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 6}, 10, 9},
		{[]int{1, 2}, 3, 2},
		{[]int{2, 5, 3, 6}, 10, 5},
		{[]int{1, 2, 3}, 4, 4},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("#%d,%d-%v", i, tc.target, tc.coins), func(t *testing.T) {
			result := makeChange(tc.coins, tc.target)
			if result != tc.expected {
				t.Errorf("wrong answer. expected %d but got %d", tc.expected, result)
			}
		})
	}
}
