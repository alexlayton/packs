package packs

import (
	"sort"
)

func Calculate(count int, packSizes []int) Packs {

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
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	minNode := solvedNode(count, packSizes[len(packSizes)-1])

	// We got lucky and can just solve it with the smallest pack
	if minNode.depth == 1 {
		return minNode
	}

	minCost := minNode.cost(count)

	queue := []node{newNode()}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.solved(count) {
			if node.cost(count) <= minCost {
				return node
			}
		} else {
			for _, packSize := range packSizes {
				next := node.copy()
				next.add(packSize)
				queue = append(queue, next)
			}
		}
	}

	return minNode
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

func solvedNode(count, packSize int) node {
	n := count / packSize
	if count%packSize != 0 {
		n += 1
	}
	node := newNode()
	node.values[packSize] = n
	node.total = n * packSize
	node.depth = n
	return node
}
