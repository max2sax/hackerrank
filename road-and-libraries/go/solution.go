package solution

import "fmt"

/*
 * Complete the 'roadsAndLibraries' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n        - number of cities
 *  2. INTEGER c_lib    - cost to build a library
 *  3. INTEGER c_road   - cost to build a road
 *  4. 2D_INTEGER_ARRAY cities - how the cities are connected
 */

func roadsAndLibraries(n int32, costLib int32, costRoad int32, cities [][]int32) int64 {
	// Write your code here

	// if the cost to build a library is < cost to build roads we want to do that in every city
	if costLib < costRoad {
		return int64(n) * int64(costLib)
	}

	// things we do know:
	// - if there are cities that can't be connected by a road, then we must build a library there
	// - The goal is to minimize the total cost.
	// that means when it is always preferable to build roads
	// because if we split a group of connected cities, it will cost
	// more to build two libraries than it would to connect the groups.

	// this means the first step is to create distinct groups of cities.
	// start by assuming each city is it's own group
	cityGroups := make(map[int32]*int32, len(cities))
	groupIds := int32(0)
	for _, possibleRoad := range cities {
		fmt.Println("cityGroups: ", cityGroups)
		cityA := possibleRoad[0]
		cityB := possibleRoad[1]
		groupA := cityGroups[cityA]
		groupB := cityGroups[cityB]
		if groupA == nil && groupB == nil {
			// belongs to no group, so create a new one
			groupIds++
			newGroup := groupIds
			cityGroups[cityA] = &newGroup
			cityGroups[cityB] = &newGroup
			continue
		}
		if groupA == nil {
			// implies && groupB != nil
			cityGroups[cityA] = groupB
			continue
		}
		if groupB == nil {
			// implies && groupA != nil
			cityGroups[cityB] = groupA
			continue
		}
		// implies groupA && groupB!= nil, but are different groups:
		// choose the smallest value always and overwrite the group pointed at,
		// which should theoretically update all cities, effectively merging the groups
		if *groupA < *groupB {
			*groupB = *groupA
			cityGroups[cityB] = groupA
			continue
		}
		*groupA = *groupB
		cityGroups[cityB] = groupA
	}
	fmt.Println("cityGroups: ", cityGroups)

	// merge city groups to create a set of groups, we really only need to count elements in each set
	citiesPerGroup := make(map[int32]int32, len(cityGroups))
	for _, groupId := range cityGroups {
		currentCount := citiesPerGroup[*groupId]
		citiesPerGroup[*groupId] = currentCount + 1
	}
	fmt.Println("citiesPerGroup: ", citiesPerGroup)
	// each map in the array represents the set of cities that can be connected
	// my initial naive assumption is to say build a library in each set and a set of roads to connect all cities
	// in the set
	totalCost := int64(0)
	unconnectedCities := n
	for _, count := range citiesPerGroup {
		numCities := count
		costForSet := (int64(numCities-1) * int64(costRoad)) + int64(costLib)
		totalCost += costForSet
		unconnectedCities -= int32(numCities)
	}

	// need to handle how unconnected cities and add libraries to them
	totalCost += (int64(unconnectedCities) * int64(costLib))

	return totalCost
}

func roadsAndLibraries2(n int32, costLib int32, costRoad int32, cities [][]int32) int64 {
	// Write your code here

	// if the cost to build a library is < cost to build roads we want to do that in every city
	if costLib < costRoad {
		return int64(n) * int64(costLib)
	}

	// things we do know:
	// - if there are cities that can't be connected by a road, then we must build a library there
	// - The goal is to minimize the total cost.
	// that means when it is always preferable to build roads
	// because if we split a group of connected cities, it will cost
	// more to build two libraries than it would to connect the groups.

	// this means the first step is to create distinct groups of cities.
	// start by assuming each city is it's own group
	cityGroups := make(map[int32]*int32, len(cities))
	groupIds := int32(0)
	numberGroups := int32(0)
	for _, possibleRoad := range cities {
		// if int32(len(cityGroups)) == n {
		// 	break
		// }
		fmt.Println("cityGroups: ", cityGroups)
		cityA := possibleRoad[0]
		cityB := possibleRoad[1]
		groupA := cityGroups[cityA]
		groupB := cityGroups[cityB]
		if groupA == nil && groupB == nil {
			// belongs to no group, so create a new one
			numberGroups++
			groupIds++
			newGroup := groupIds
			cityGroups[cityA] = &newGroup
			cityGroups[cityB] = &newGroup
			continue
		}
		if groupA == nil {
			// implies && groupB != nil
			cityGroups[cityA] = groupB
			continue
		}
		if groupB == nil {
			// implies && groupA != nil
			cityGroups[cityB] = groupA
			continue
		}
		// implies groupA && groupB!= nil, but are different groups:
		// choose the smallest value always and overwrite the group pointed at,
		// which should theoretically update all cities, effectively merging the groups
		if *groupA == *groupB {
			continue
		}
		numberGroups--
		if *groupA < *groupB {
			*groupB = *groupA
			continue
		}
		*groupA = *groupB
	}
	fmt.Println("cityGroups: ", cityGroups)

	// merge city groups to create a set of groups, we really only need to count elements in each set
	// citiesPerGroup := make(map[int32]int32, len(cityGroups))
	// for _, groupId := range cityGroups {
	// 	currentCount := citiesPerGroup[*groupId]
	// 	citiesPerGroup[*groupId] = currentCount + 1
	// }
	// fmt.Println("citiesPerGroup: ", citiesPerGroup)
	// each map in the array represents the set of cities that can be connected
	// my initial naive assumption is to say build a library in each set and a set of roads to connect all cities
	// in the set
	connectedCities := int32(len(cityGroups))
	unconnectedCities := n - connectedCities
	numberRoads := connectedCities - numberGroups

	// need to handle how unconnected cities and add libraries to them
	totalCost := int64(numberRoads) * int64(costRoad)
	totalCost += (int64(unconnectedCities+numberGroups) * int64(costLib))

	return totalCost
}
func roadsAndLibrariesDFS(n int32, costLib int32, costRoad int32, cities [][]int32) int64 {
	// if the cost to build a library is < cost to build roads we want to do that in every city
	if costLib < costRoad {
		return int64(n) * int64(costLib)
	}
	adjacentCities := make([][]int32, n+1)
	for _, possibleRoad := range cities {
		cityA := possibleRoad[0]
		cityB := possibleRoad[1]
		adjacentCities[cityA] = append(adjacentCities[cityA], cityB)
		adjacentCities[cityB] = append(adjacentCities[cityB], cityA)
	}
	// visit all the connected nodes:
	citiesVisited := make(map[int]bool, n+1)
	connectedCities := make([][]int32, 0)
	for i, connected := range adjacentCities[1:] {
		if int32(len(citiesVisited)) == n {
			break
		}
		cityId := i + 1
		if citiesVisited[cityId] {
			continue
		}
		citiesVisited[cityId] = true
		citiesInGroup := []int32{int32(cityId)}
		nextCities := connected
		for len(nextCities) > 0 {
			newCity := nextCities[0]
			if citiesVisited[int(newCity)] {
				nextCities = nextCities[1:]
				continue
			}
			citiesInGroup = append(citiesInGroup, int32(newCity))
			citiesVisited[int(newCity)] = true
			nextCities = append(adjacentCities[newCity], nextCities[1:]...)
		}
		connectedCities = append(connectedCities, citiesInGroup)
	}
	totalCost := int64(0)
	fmt.Printf("number of groups: %d", len(connectedCities))
	for _, count := range connectedCities {
		numCities := len(count)
		costForSet := (int64(numCities-1) * int64(costRoad)) + int64(costLib)
		totalCost += costForSet
	}

	unconnectedCities := n - int32(len(citiesVisited))
	// need to handle how unconnected cities and add libraries to them
	totalCost += (int64(unconnectedCities) * int64(costLib))
	return totalCost
}
