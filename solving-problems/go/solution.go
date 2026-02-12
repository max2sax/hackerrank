package solvingproblems

import (
	"log"
)

func problemSolving(k int32, v []int32) int32 {
	// Write your code here
	days := int32(1)
	log.Println("starting v on day 0  :", v)
	// slices.Sort(v)
	// log.Println("starting v on day 1  :", v)
	for {
		unsolved := v[:0]
		prevRating := v[0]
		for _, rating := range v[1:] {
			diff := rating - int32(prevRating)
			negDiff := int32(prevRating) - rating
			if diff < 0 {
				diff = negDiff
			}
			if diff < k {
				unsolved = append(unsolved, rating)
				continue
			}
			prevRating = rating
		}
		if len(unsolved) == 0 {
			break
		}
		v = v[0:len(unsolved)]
		log.Printf("unsolved after day[%d]: %v ", days, unsolved)
		// log.Println("v after first day: ", v)
		days++
	}
	return days
}
