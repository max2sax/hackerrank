package bfs

import (
	"fmt"
	"slices"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name          string
		startNode     int32
		numberOfNodes int32
		edges         [][]int32
		expected      []int32
	}{
		{"testcase7-1", 1, 5, [][]int32{{1, 2}, {1, 3}, {3, 4}}, []int32{6, 6, 12, -1}},
		{"testcase0-1", 1, 4, [][]int32{{1, 2}, {1, 3}}, []int32{6, 6, -1}},
		{"testCase0-2", 2, 3, [][]int32{{2, 3}}, []int32{-1, 6}},
		{"testcase3-1", 3, 10, [][]int32{{3, 1}, {10, 1}, {10, 1}, {3, 1}, {1, 8}, {5, 2}},
			[]int32{6, -1, -1, -1, -1, -1, 12, -1, 12}},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			out := bfs(testCase.numberOfNodes, testCase.edges, testCase.startNode)
			exp := testCase.expected
			if len(exp) != len(out) {
				t.Errorf("lengths don't match. expected %v but got %v", exp, out)
				return
			}
			if !slices.Equal(exp, out) {
				for i := range exp {
					if out[i] == exp[i] {
						out[i] = 0
						exp[i] = 0
					}
				}
				t.Errorf("incorrect results. expected %v but got %v", exp, out)
			}
		})
	}
}

func TestBFSinput4(t *testing.T) {
	for i, testCase := range input {
		t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
			out := bfs(testCase.numberNodes, testCase.edges, testCase.startNode)
			expected := output[i].expected
			if len(out) != len(output[i].expected) {
				t.Errorf("lengths are different")
				return
			}
			if !slices.Equal(out, output[i].expected) {
				t.Fail()
				for i := range out {
					if out[i] == expected[i] {
						out[i] = 0
						expected[i] = 0
					}
				}
			}
			t.Log(expected)
			t.Log(out)
		})
	}
}
