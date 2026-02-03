package gameoftwostacks

import (
	"fmt"
	"testing"
)

func TestTwoStacks(t *testing.T) {
	maxT := len(tests)

	maxT = 5
	for i := 4; i < maxT; i++ {
		test := tests[i]
		tName := fmt.Sprintf("%d:  ", i)
		t.Run(tName, func(t *testing.T) {
			expected := outputs01[i]
			result := twoStacks(test.maxSum, test.a, test.b)
			if result != expected {
				t.Errorf("twoStacks(%d, \n%v, \n%v) \n= %d; want %d", test.maxSum, test.a, test.b, result, expected)
			}
		})
	}
}
