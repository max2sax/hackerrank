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
	curMax := int32(0)
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
		log.Printf("nextA: %d, nextB: %d; nextSumA: %d, nextSumB: %d; ", nextA, nextB, nextSumA, nextSumB)
		if nextSumA > maxSum && nextSumB > maxSum {
			break
		}
		if nextA > curMax {
			curMax = nextA
		}
		if nextB > curMax {
			curMax = nextB
		}

		hasSelection := false
		if nextA < curMax && nextSumA <= maxSum {
			log.Printf("choosing a : %d", nextA)
			curSum = nextSumA
			nextSumB += nextA
			i++
			selections++
			hasSelection = true
		}
		if nextB < curMax && nextSumB <= maxSum {
			log.Printf("choosing b: %d", nextB)
			curSum = nextSumB
			j++
			hasSelection = true
			selections++
		}
		log.Printf(", CurSum: %d\n", curSum)
		if !hasSelection || curSum > maxSum {
			break
		}

		if i == maxI && j == maxJ {
			break
		}
	}
	return selections
}

func twoStacksArrayOfSums(maxSum int32, a []int32, b []int32) int32 {
	// Write your code here	var i, j int
	maxI := len(a)
	maxJ := len(b)

	sums := make([]map[int32]bool, maxI+maxJ+1)
	prevSums := map[int32]bool{0: true}
	// todo: optimize to avoid storing all sums, don't need duplicates
	// and break early when all sums in a layer exceed maxSum
	var iA, iB, curMax int32
	for range sums {
		newSumsMap := map[int32]bool{}
		for prevSum := range prevSums {
			if iA < int32(maxI) {
				newSum := prevSum + a[iA]
				if newSum <= maxSum {
					newSumsMap[newSum] = true
				}
			}
			if iB < int32(maxJ) {
				newSum := prevSum + b[iB]
				if newSum <= maxSum {
					newSumsMap[newSum] = true
				}
			}
		}
		if len(newSumsMap) == 0 {
			break
		}
		curMax++
		iA++
		iB++
		prevSums = newSumsMap
	}
	// sumsloop:
	// 	for _, s := range sums {
	// 		log.Printf("sums: %v\n", s)
	// 		for _, v := range s {
	// 			if v <= maxSum {
	// 				log.Printf("found valid sum: %d\n", v)
	// 				curMax++
	// 				continue sumsloop
	// 			}
	// 		}
	// 		break
	// 	}
	return curMax
}

type sumKey struct {
	iA int // index in stack A, 0 --> iA
	iB int // index in stack B, 0 --> iB
	// window size is iA + iB
}

func twoStacksSlidingWindow(maxSum int32, a []int32, b []int32) (selections int32) {
	// Write your code here
	maxI := len(a)
	maxJ := len(b)
	curSum := int32(0)

	memoizedSumsA := make([]int32, 1, maxI+1)
	memoizedSumsB := make([]int32, 1, maxJ+1)
	for i := 0; i < maxI; i++ {
		curSum += a[i]
		if curSum > maxSum {
			maxI = i
			break
		}
		memoizedSumsA = append(memoizedSumsA, curSum)
	}
	curSum = 0
	for i := 0; i < maxJ; i++ {
		curSum += b[i]
		if curSum > maxSum {
			maxJ = i
			break
		}
		memoizedSumsB = append(memoizedSumsB, curSum)
	}
	// maxI--
	// maxJ--
	suma := memoizedSumsA[maxI]
	sumb := memoizedSumsB[maxJ]
	if suma+sumb <= maxSum {
		// we can pick everything from both stacks
		return int32(maxI + maxJ) // maybe +2
	}
	maxWindowSize := maxI + maxJ
	// todo: when changing window size, calculate new pivot like binary search
	// for now just increment by 1
	// we're gauranteed that the window will fit within stacka + stackb
	// but because stacka might be a different size than stackb we can't just start in the middle
	// so max window size is len(stacka) + len(stackb)
	windowSize := 1
	if suma <= maxSum {
		selections = int32(maxI)
		windowSize = maxI
	}
	if sumb <= maxSum && maxJ > maxI {
		selections = int32(maxJ)
		windowSize = maxJ
	}
windowLoop:
	for windowSize > 0 && windowSize <= maxWindowSize {
		iA := windowSize
		var iB int
		// get start position for a:
		if iA > maxI {
			iA = maxI
			iB = windowSize - maxI
		}
		// if we can take all values in windowSize without exceeding maxSum, increase windowSize
		// otherwise, decrease windowSize and try again
		// start all the way in a and slide into b until we reach the window size
		for iB <= windowSize && iB <= maxJ {
			// gauranteed that iB <= maxJ
			sumA := memoizedSumsA[iA]
			sumB := memoizedSumsB[iB]

			totalSum := sumA + sumB
			if totalSum <= maxSum {
				selections = int32(windowSize)
				// increase window size
				windowSize++
				continue windowLoop
			}
			// slide window and try again:
			// increment a, and decrement b
			iA--
			iB++
		}
		break
	}

	return selections
}
