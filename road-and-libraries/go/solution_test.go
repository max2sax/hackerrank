package solution

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	cities       [][]int32
	numberCities int32
	costLibrary  int32
	costRoad     int32
	expected     int64
}{
	// {[][]int32{{1, 2}, {1, 3}, {1, 4}}, 5, 6, 1, 15},
	// {[][]int32{{1, 3}, {2, 5}, {6, 7}, {3, 4}}, 8, 6, 1, 28},
	// {[][]int32{{1, 3}, {2, 5}, {6, 7}, {3, 4}, {3, 5}, {2, 4}, {5, 6}, {5, 8}}, 8, 6, 1, 13},
	{[][]int32{{6, 7}, {5, 8}, {9, 10}, {8, 9}, {7, 8}}, 10, 5, 1, 30},
	// {[][]int32{{1, 3}, {2, 5}, {5, 3}, {3, 4}, {3, 5}, {12, 14}, {11, 13}, {10, 12}, {12, 13}, {10, 14},
	// 	{11, 15}, {5, 8}, {6, 9}, {7, 9}, {2, 10}, {9, 8}},
	// 	15, 6, 1, 20},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test#%d", i), func(t *testing.T) {
			actual := roadsAndLibraries(tc.numberCities, tc.costLibrary, tc.costRoad, tc.cities)
			if actual != tc.expected {
				t.Fatalf("expected %d but got %d", tc.expected, actual)
			}
		})
	}
}

func TestInput02(t *testing.T) {
	// r := 6
	// for i := r; i < r+1; i++ {
	for i := range input02 {
		tc := input02[i]
		t.Run(fmt.Sprintf("Test#%d:c-%d,lib-%d,rd-%d", i, tc.numberCities, tc.costLibrary, tc.costRoad), func(t *testing.T) {
			actual := roadsAndLibraries(tc.numberCities, tc.costLibrary, tc.costRoad, tc.cities)
			expected := output02[i]
			if actual != expected {
				t.Fatalf("expected %d but got %d", expected, actual)
			}
		})
	}
}

func TestInput02DFS(t *testing.T) {
	// r := 0
	// for i := r; i < r+1; i++ {
	for i := range input02 {
		tc := input02[i]
		t.Run(fmt.Sprintf("Test#%d:c-%d,lib-%d,rd-%d", i, tc.numberCities, tc.costLibrary, tc.costRoad), func(t *testing.T) {
			actual := roadsAndLibrariesDFS(tc.numberCities, tc.costLibrary, tc.costRoad, tc.cities)
			expected := output02[i]
			if actual != expected {
				t.Fatalf("expected %d but got %d", expected, actual)
			}
		})
	}
}
