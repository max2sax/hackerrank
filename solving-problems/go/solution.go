package solvingproblems

import (
	"fmt"
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
		lastSolvedRating := v[0]
		for _, rating := range v[1:] {
			diff := rating - int32(lastSolvedRating)
			negDiff := int32(lastSolvedRating) - rating
			if diff < 0 {
				diff = negDiff
			}
			if diff < k {
				unsolved = append(unsolved, rating)
				continue
			}
			lastSolvedRating = rating
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

type t struct {
	index int
	val   int32
}
type pair struct {
	start t
	end   t
}

func problemSolvingRecursiveGraph(k int32, v []int32) int32 {
	// any consecutive problem less than k apart cannot be solved on the same day
	// solving something on one day, could create unsolvable adjacencies the next day
	// options, pick subarrays to solve, pick individual elements to solve, skipping around
	// always have to pick from left to right(lowest index to highest index)
	// let's try creating subarays to solve and then grouping the subarrays until it's not possible
	if len(v) < 2 {
		return 1
	}
	minusK := k * -1
	days := make([][]int32, 0, len(v)) // treating days like a stack?
	days = append(days, []int32{v[0]})
	fmt.Println(v)
	fmt.Println(" ________")
	for _, rating := range v[1:] {
		foundDay := false
		// check if the last solved problem in each day is < k
		for i, d := range days {
			lastDayRating := d[len(d)-1]
			diff := lastDayRating - rating
			if diff < k && diff > minusK {
				// too close, check next day
				continue
			}
			// found a day, so add it here
			days[i] = append(d, rating)
			foundDay = true
			break
		}
		if foundDay {
			continue
		}
		days = append(days, []int32{rating})
	}
	for _, d := range days {
		fmt.Println(d)
	}
	return int32(len(days))
}

func problemSolvingSubArrays(k int32, v []int32) int32 {
	// any consecutive problem less than k apart cannot be solved on the same day
	// solving something on one day, could create unsolvable adjacencies the next day
	// options, pick subarrays to solve, pick individual elements to solve, skipping around
	// always have to pick from left to right(lowest index to highest index)
	// let's try creating subarays to solve and then grouping the subarrays until it's not possible
	if len(v) < 2 {
		return 1
	}
	minusK := k * -1
	days := make([][]int32, 0, len(v)) // treating days like a stack?
	days = append(days, []int32{v[0]})
	fmt.Println(v)
	fmt.Println(" ________")
	problemGroups := make([]pair, 0)

	// build the groups

	firstPair := pair{start: t{0, v[0]}}
	for i := range v[:len(v)-1] {
		curRating := v[i]
		nextRating := v[i+1]
		diff := curRating - nextRating
		if diff < k && diff > minusK {
			// too close, check next day
			firstPair.end = t{i, v[i]}
			problemGroups = append(problemGroups, firstPair)

			firstPair.start = t{i + 1, v[i+1]}
		}
	}
	// can't join 2 groups next to each other,
	// but it might be possible to join groups not next to each other
	for _, p := range problemGroups {
		fmt.Println("pairx : ", p)
	}
	for _, d := range days {
		fmt.Println(d)
	}
	return int32(len(problemGroups))
}
