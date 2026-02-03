package gameoftwostacks

import (
	"fmt"
	"log"
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
		log.Printf("[%d,%d]: ", i, j)
		nextA := maxSum + 1
		nextB := maxSum + 1
		if i < maxI {
			nextA = a[i]
		}
		if j < maxJ {
			nextB = b[j]
		}
		nextSumA := curSum + nextA
		nextSumB := curSum + nextB
		if nextSumA > maxSum && nextSumB > maxSum {
			break
		}

		log.Printf("nextA: %d, nextB: %d; nextSumA: %d, nextSumB: %d; ", nextA, nextB, nextSumA, nextSumB)
		i++
		curSum = nextSumA
		if nextSumB < nextSumA {
			i--
			j++
			curSum = nextSumB
		}
		// if  {
		// 	fmt.Printf("choosing a : %d", a[i])
		// 	curSum += a[i]
		// 	i++
		// } else if j < maxJ && (i == maxI || b[j] < a[i]) {
		// 	fmt.Printf("choosing b: %d", b[j])
		// 	curSum += b[j]
		// 	j++
		// }
		log.Printf(", CurSum: %d\n", curSum)
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
