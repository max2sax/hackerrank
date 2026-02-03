package gameoftwostacks

import (
	"fmt"
)

func twoStacks(maxSum int32, a []int32, b []int32) int32 {
	// Write your code here
	var i, j int
	maxI := len(a)
	maxJ := len(b)
	var selections int32
	curSum := int32(0)
	for {
		fmt.Printf("[%d,%d]: ", i, j)
		if i < maxI && (j == maxJ || a[i] <= b[j]) {
			fmt.Printf("choosing a : %d", a[i])
			curSum += a[i]
			i++
		} else if j < maxJ && (i == maxI || b[j] < a[i]) {
			fmt.Printf("choosing b: %d", b[j])
			curSum += b[j]
			j++
		}
		fmt.Printf(", CurSum: %d\n", curSum)
		if curSum > maxSum {
			break
		}
		if i == maxI && j == maxJ {
			break
		}
		selections++
	}
	return selections
}

func twoStacksNew(maxSum int32, a []int32, b []int32) int32 {
	// Write your code here
	var i, j int
	maxI := len(a)
	maxJ := len(b)
	var selections int32
	curSum := int32(0)
	for {
		fmt.Printf("[%d,%d]: ", i, j)
		if i < maxI && (j == maxJ || a[i] <= b[j]) {
			fmt.Printf("choosing a : %d", a[i])
			curSum += a[i]
			i++
		} else if j < maxJ && (i == maxI || b[j] < a[i]) {
			fmt.Printf("choosing b: %d", b[j])
			curSum += b[j]
			j++
		}
		fmt.Printf(", CurSum: %d\n", curSum)
		if curSum > maxSum {
			break
		}
		if i == maxI && j == maxJ {
			break
		}
		selections++
	}
	return selections
}
