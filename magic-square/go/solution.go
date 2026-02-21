package magicsquares

import "log"

func generateAllValidSquares() (squares [][]int32) {
	// assume 3x3 with 1-9
	// hard code:
	// generate a square, calculate sums, then try again
	// each element is unique, so
	// we know that a magic square has a middle square of 5
	// and sum of 15

	// canonical magic square
	// [2 7 6]
	// [9 5 1]
	// [4 3 8]
	// rotations counter clockwise
	//  [0 1 2]    [2 5 8]    [8 7 6]    [6 3 0]
	//  [3 4 5]    [1 4 7]    [5 4 3]    [7 4 1]
	//  [6 7 8] => [0 3 6] => [2 1 0] => [8 5 2]
	// mirror along diagonals and middle
	// [0 1 2]    [2 1 0]    [6 7 8]    [8 5 2]    [0 3 6]
	// [3 4 5]    [5 4 3]    [3 4 5]    [7 4 1]    [1 4 7]
	// [6 7 8] => [8 7 6] => [0 1 2] => [6 3 0] => [2 5 8]
	variants := [][]int{
		{2, 5, 8, 1, 4, 7, 0, 3, 6},
		{8, 7, 6, 5, 4, 3, 2, 1, 0},
		{6, 3, 0, 7, 4, 1, 8, 5, 2},
		{2, 1, 0, 5, 4, 3, 8, 7, 6},
		{6, 7, 8, 3, 4, 5, 0, 1, 2},
		{8, 5, 2, 7, 4, 1, 6, 3, 0},
		{0, 3, 6, 1, 4, 7, 2, 5, 8},
	}
	canonical := []int32{2, 7, 6, 9, 5, 1, 4, 3, 8}
	squares = append(squares, canonical)
	for _, variant := range variants {
		square := make([]int32, 9)
		for i, j := range variant {
			square[i] = canonical[j]
		}
		squares = append(squares, square)
	}
	return
}

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	allValidSquares := generateAllValidSquares()
	log.Print(allValidSquares)
	// enumerate each array - vertical, horizontal, diagonal
	// then calculate the sume for each and determine the max variation
	// that will determine the minimum cost. I don't think we necessarily need to
	// know the which individual element to change
	// assume a 3x3 array,
	// we know the middle number must be 5 and the magic sum is 15

	flattenedSquare := make([]int32, 9)
	for i := range 3 {
		for j := range 3 {
			idx := int((i * 3) + j)
			flattenedSquare[idx] = s[i][j]
		}
	}
	minDiff := int32(-1)

	for _, square := range allValidSquares {
		var diff int32
		for i := range 9 {
			d := square[i] - flattenedSquare[i]
			if d < 0 {
				d = flattenedSquare[i] - square[i]
			}
			diff += d
		}
		log.Printf("difference for square %v: %d", square, diff)
		if minDiff < 0 || diff < minDiff {
			minDiff = diff
		}
	}
	return int32(minDiff)
}
