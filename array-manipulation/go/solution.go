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

// another naive array manipulation using a different approach
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
	// for i := range n {
	// 	// perform operations
	// 	sum := arr[i]
	// 	if sum > max {
	// 		max = sum
	// 	}
	// }

	return max
}

type node struct {
	start int64
	end   int64
	value int64
	next  *node
}

func arrayManipulationLinkedLis(n int32, queries [][]int32) int64 {
	max := int64(0)
	if len(queries) == 0 {
		return max
	}
	if len(queries) == 1 {
		return int64(queries[0][2])
	}
	// this is an attempt to optimize the naive approach
	// by merging overlapping queries
	// into a single query
	// this is still not O(n + m) but should be faster than O(n*m)
	newQueries := &node{
		start: int64(queries[0][0]),
		end:   int64(queries[0][1]),
		value: int64(queries[0][2]),
		next:  nil,
	}
	for _, query := range queries[1:] {
		queryStart := int64(query[0])
		queryEnd := int64(query[1])
		querySum := int64(query[2])

		if queryEnd < newQueries.start {
			// insert before the current node
			newNode := &node{
				start: queryStart,
				end:   queryEnd,
				value: querySum,
				next:  newQueries,
			}
			newQueries = newNode
			continue
		}
		if queryEnd == newQueries.start {
			// extend the current node
			newQueries.start = queryStart
			newQueries.value += querySum
			continue
		}

		curNode := newQueries
		prevNode := newQueries
	nodeloop:
		for curNode != nil {

			curStart := curNode.start
			curEnd := curNode.end

			// let's just handle all cases explicitly:
			if queryEnd < curStart {
				// query starts before  node and ends before the current node
				// create a new node and insert it before the current node
				newNode := &node{
					start: queryStart,
					end:   queryEnd,
					value: querySum,
					next:  curNode,
				}
				newQueries = newNode
				break nodeloop
			}

			if queryEnd < curStart && queryStart > prevNode.end {
				// query starts after previous node and ends before current node
				newNode := &node{
					start: queryStart,
					end:   queryEnd,
					value: querySum,
					next:  curNode,
				}
				prevNode.next = newNode
				break nodeloop
			}

			// query starts before the current node and ends on the current node
			if queryStart < curStart && queryEnd == curStart {
				newNode := &node{
					start: queryStart,
					end:   queryEnd - 1,
					value: querySum,
					next:  curNode,
				}
				if prevNode == nil {
					newQueries = newNode
				} else {
					prevNode.next = newNode
				}
			}

			// check if the query starts inside the current node
			if queryStart == curStart {
				if curEnd == curStart && queryEnd == curStart {
					// single point node
					curNode.value += querySum
					break
				}
				if curEnd == curStart {
					curNode.value += querySum
					queryStart++
					continue
				}
				// currentstart != currentend
				if queryEnd < curEnd {
					// left node
					leftNode := &node{
						start: queryStart,
						end:   queryEnd,
						value: curNode.value + querySum,
						next:  curNode.next,
					}
					// right node
					curNode.start = queryEnd + 1
					// link left node
					if prevNode == nil {
						newQueries = leftNode
					} else {
						prevNode.next = leftNode
					}
					leftNode.next = curNode
					break
				}
				if queryEnd == curEnd {
					curNode.value += querySum
					break
				}
				// queryEnd > curEnd
				curNode.value += querySum
				// set new query start for next iteration
				queryStart = curEnd + 1
				// move to next node
				prevNode = curNode
				curNode = curNode.next
				continue
			}

			if queryStart > curStart && queryStart <= curEnd {
				// left node
				curNode.end = queryStart - 1
				// middle node
				middleNode := &node{
					// middle node
					start: queryStart,
					end:   queryEnd,
					value: curNode.value + querySum,
					next:  curNode.next,
				}
				curNode.next = middleNode
				if queryEnd < curEnd {
					middleNode.next = &node{
						start: queryEnd + 1,
						end:   curEnd,
						value: curNode.value,
						next:  middleNode.next,
					}
					break
				}
				if queryEnd == curEnd {
					// done with this query
					break
				}
				// queryEnd > curEnd
				// set new query start for next iteration
				queryStart = curEnd + 1
				// link middle node
				curNode.next = middleNode
				middleNode.end = curEnd
				// done with this query
				prevNode = curNode
				curNode = middleNode
				// right node will be handled in next iterations
				continue
			}

			// // query starts after the current node and before the next node
			// // then insert after the current node
			// if queryStart > curEnd && (curNode.next == nil || queryEnd < curNode.next.start) {
			// 	newNode := &node{
			// 		start: queryStart,
			// 		end:   queryEnd,
			// 		value: querySum,
			// 		next:  curNode.next,
			// 	}
			// 	curNode.next = newNode
			// 	curNode = newNode
			// 	break
			// }

			// // query starts at the same spot current node and ends anywhere
			// if queryStart == curStart {
			// 	curNode.value += querySum
			// 	// set new query start for next iteration
			// 	queryStart++
			// 	// stay on the current node,
			// 	// the query will start after this and be handled accordingly
			// 	continue
			// }

			// // query starts after current node
			// if queryStart > curEnd {
			// 	// move to next node
			// 	prevNode = curNode
			// 	curNode = curNode.next
			// 	continue
			// }

			// // query starts inside current node and ends after current node
			// // not sure if this is right or covered by another case
			// if queryStart > curStart {
			// 	curNode.end = queryStart - 1
			// 	newNode := &node{
			// 		// middle node
			// 		start: queryStart,
			// 		end:   curEnd,
			// 		value: curNode.value + querySum,
			// 		next:  curNode.next,
			// 	}
			// 	curNode.next = newNode
			// 	// set new query start for next iteration
			// 	queryStart = curEnd + 1
			// }
			// // move to next node
			// prevNode = curNode
			// curNode = curNode.next
		}
	}
	count := 0
	for nq := newQueries; nq != nil; nq = nq.next {
		count++
		log.Printf("node: %+v\n", *nq)
		// perform operations
		sum := nq.value
		if sum > max {
			max = sum
		}
	}
	log.Println("nodes created - ", count)

	return max
}

func arrayManipulationAI(n int32, queries [][]int32) int64 {
	// lets see what AI comes up with,
	// so far I've seen something about a difference array,
	// what could that be?
	//
	// construct a difference array
	arr := make([]int64, n+2)
	for _, query := range queries {
		if query[0] < 1 {
			log.Printf("query index is out of bounds")
		}
		startIndex := query[0]
		endIndex := query[1]
		value := int64(query[2])

		arr[startIndex] += value
		arr[endIndex] -= value
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

func arrayManipulationWithDifferenceArray(n int32, queries [][]int32) int64 {
	// implement the difference array approach
	diffArr := make([]int64, n+1)
	arr := make([]int64, n)
	// perform operations
	for i, query := range queries {
		a := query[0]
		b := query[1]
		k := query[2]
		for i := a - 1; i < b; i++ {
			arr[i] += int64(k)
		}
		log.Printf("query[%d]: %v: %v\n", i, query, arr)
	}

	log.Println("Start: ", diffArr)
	for i, query := range queries {
		start := query[0] - 1
		end := query[1]
		value := int64(query[2])
		diffArr[start] += value
		diffArr[end] -= value
		log.Printf("query[%d]: %v: %v\n", i, query, diffArr)
	}

	max := int64(0)
	current := int64(0)
	for i := int32(0); i < n; i++ {
		current += diffArr[i]
		if current > max {
			max = current
		}
	}

	return max
}
