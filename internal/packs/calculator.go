package packs

import (
	"sort"
)

// Calculate the combination of packs needed to fulfill an order
func Calculate(count int, packSizes []int) Packs {

	if len(packSizes) == 0 || count == 0 {
		return Packs{}
	}

	tmpPackSizes := append([]int{}, packSizes...)

	node := internalCalculate(count, tmpPackSizes)

	packs := Packs{}
	for k, v := range node.values {
		packs = append(packs, Pack{Size: k, Count: v})
	}

	sort.Sort(packs)

	return packs
}

func internalCalculate(count int, packSizes []int) node {

	// Reverse sort all the pack sizes so we search with the biggest pack sizes first
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	// Create a naive solution which is just trying to use the largest packs which
	// we'll try and beat
	minNode := naiveNode(count, packSizes)

	// We got lucky and can just solve it with the smallest pack
	if minNode.depth == 1 {
		return minNode
	}

	// Calculate the cost - the amount that it goes past 0 by
	minCost := minNode.cost(count)

	// Here we perform a bfs until we find a solution
	queue := []node{newNode()}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Pop the queue and check if the current node has fulfilled the order
		if node.solved(count) {

			// If we can beat the cost of fulfilling the order with the minimum
			// pack size then I think we can assert the following;
			// 1. We're sending out no more items than necessary
			// 2. Doing reverse bfs means we find a solution with the least
			// 	  possible packs
			nodeCost := node.cost(count)
			if nodeCost == 0 || nodeCost <= minCost {
				return node
			}

		} else {

			// If we cant fulfill the order add new nodes to the queue with
			// each possible pack size to see if they can fulfill it
			for _, packSize := range packSizes {
				next := node.copy()
				next.add(packSize)
				queue = append(queue, next)
			}

		}
	}

	// We return the worst case here, but this should never be reached as worst
	// case the bfs will return the same result
	return minNode
}

func naiveNode(count int, packSizes []int) node {
	i, n, curr := 0, newNode(), count

	for i < len(packSizes) {
		packSize := packSizes[i]
		if curr >= packSize {
			nPacks := curr / packSize
			for j := 0; j < nPacks; j++ {
				n.add(packSize)
			}
			curr = curr % packSize
		}
		i += 1
	}

	if curr != 0 {
		n.add(packSizes[len(packSizes)-1])
	}

	return n
}

type node struct {
	total, depth int
	values       map[int]int
}

func newNode() node {
	return node{
		total:  0,
		depth:  0,
		values: map[int]int{},
	}
}

func (n *node) add(packSize int) {
	n.values[packSize] = n.values[packSize] + 1
	n.total += packSize
	n.depth += 1
}

func (n node) solved(count int) bool {
	return count-n.total <= 0
}

func (n node) cost(count int) int {
	diff := count - n.total
	if diff < 0 {
		return -diff
	}
	return diff
}

func (n node) copy() node {
	new := newNode()
	new.total = n.total
	new.depth = n.depth
	for k, v := range n.values {
		new.values[k] = v
	}
	return new
}
