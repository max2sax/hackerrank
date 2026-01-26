package array_manipulation

import "log"

// array manipulation performs a series of operations on an array of zeros
// and returns the maximum value in the array after all operations have been performed.
func arrayManipulationOriginal(n int32, queries [][]int32) int64 {
	max := int64(0)
	// this is O(n*m) and is slow for large n and m
	// but serves as a correct reference implementation
	for i := range n {
		// perform operations
		val := int64(0)
		for _, query := range queries {
			a := query[0]
			b := query[1]
			k := query[2]
			if int32(i+1) >= a && int32(i+1) <= b {
				val += int64(k)
			}
		}
		if val > max {
			max = val
		}
	}
	return max
}

// optimized array manipulation using a different approach
func arrayManipulation(n int32, queries [][]int32) int64 {
	// my goal is to implement this function in O(n + m) time complexity
	// where n is the size of the array and m is the number of queries
	// first loop through queries
	arr := make([]int64, n)
	for _, query := range queries {
		if query[0] < 1 {
			log.Printf("query index is out of bounds")
		}
		for i := query[0] - 1; i < query[1]; i++ {
			arr[i] += int64(query[2])
		}
	}

	max := int64(0)
	// then loop through array
	for i := range n {
		// perform operations
		sum := arr[i]
		if sum > max {
			max = sum
		}
	}

	return max
}

func arrayManipulationAI(n int32, queries [][]int32) int64 {
	return 0
}
