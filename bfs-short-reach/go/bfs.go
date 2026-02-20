package bfs

type node struct {
	//   id int32
	nextNodes       []int32
	distanceToStart int32
	visited         bool
}

func bfs(n int32, edges [][]int32, s int32) []int32 {
	// Write your code here
	// assume that the start node is the first index, and next node is second index
	// and that the we can't traverse backwards
	nodes := make(map[int32]node)
	for i := range n {
		nodes[i+1] = node{nextNodes: []int32{}, distanceToStart: -1}
	}
	for _, edge := range edges {
		parent := edge[0]
		child := edge[1]
		n := nodes[parent]
		n.nextNodes = append(n.nextNodes, child)
		nodes[parent] = n
		n = nodes[child]
		n.nextNodes = append(n.nextNodes, parent)
		nodes[child] = n
	}

	// traverse nodes and add distances
	startNode := nodes[s]
	startNode.visited = true
	nodes[s] = startNode
	queue := []int32{}
	for _, n := range startNode.nextNodes {
		node := nodes[n]
		node.distanceToStart = 0
		nodes[n] = node
		queue = append(queue, n)
	}
	pointer := 0
	level := 1
	for pointer < len(queue) {
		queueLength := len(queue)
		for _, i := range queue[pointer:queueLength] {
			node := nodes[i]
			if node.visited {
				continue
			}
			node.distanceToStart = int32(6 * level)
			node.visited = true
			nodes[i] = node
			for _, j := range node.nextNodes {
				nextNode := nodes[j]
				if nextNode.visited {
					continue
				}
				queue = append(queue, j)
			}
		}
		level++
		pointer = queueLength
	}

	distances := make([]int32, 0, n)
	for i := range n {
		if i+1 == s {
			continue
		}
		node := nodes[i+1]
		distances = append(distances, node.distanceToStart)
	}

	return distances
}
