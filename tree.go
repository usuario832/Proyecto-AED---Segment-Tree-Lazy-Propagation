package main

import (
	"fmt"
	"math"
)

type SegmentTree struct {
	Tree []Node
	Data []float64
	N    int
}

func NewSegmentTree(data []float64) *SegmentTree {
	n := len(data)
	st := &SegmentTree{
		Tree: make([]Node, 4*n),
		Data: data,
		N:    n,
	}
	st.build(1, 0, n-1)
	return st
}

func (st *SegmentTree) build(node, left, right int) {
	if left == right {
		valor := st.Data[left]
		st.Tree[node] = Node{Sum: valor, Min: valor, Max: valor}
		return
	}

	mid := (left + right) / 2
	st.build(node*2, left, mid)
	st.build(node*2+1, mid+1, right)

	leftNode := st.Tree[node*2]
	rightNode := st.Tree[node*2+1]

	st.Tree[node] = Node{
		Sum: leftNode.Sum + rightNode.Sum,
		Min: min(leftNode.Min, rightNode.Min),
		Max: max(leftNode.Max, rightNode.Max),
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (st *SegmentTree) QuerySum(l, r int) float64 {
	return st.querySum(1, 0, st.N-1, l, r)
}

func (st *SegmentTree) querySum(node, left, right, qLeft, qRight int) float64 {
	if qRight < left || qLeft > right {
		return 0
	}
	if qLeft <= left && right <= qRight {
		return st.Tree[node].Sum
	}

	st.push(node, left, right)
	mid := (left + right) / 2

	return st.querySum(node*2, left, mid, qLeft, qRight) +
		st.querySum(node*2+1, mid+1, right, qLeft, qRight)
}

func (st *SegmentTree) QueryMin(l, r int) float64 {
	return st.queryMin(1, 0, st.N-1, l, r)
}

func (st *SegmentTree) queryMin(node, left, right, qLeft, qRight int) float64 {
	if qRight < left || qLeft > right {
		return math.MaxFloat64
	}
	if qLeft <= left && right <= qRight {
		return st.Tree[node].Min
	}

	st.push(node, left, right)
	mid := (left + right) / 2

	leftMin := st.queryMin(node*2, left, mid, qLeft, qRight)
	rightMin := st.queryMin(node*2+1, mid+1, right, qLeft, qRight)

	return min(leftMin, rightMin)
}

func (st *SegmentTree) QueryMax(l, r int) float64 {
	return st.queryMax(1, 0, st.N-1, l, r)
}

func (st *SegmentTree) queryMax(node, left, right, qLeft, qRight int) float64 {
	if qRight < left || qLeft > right {
		return -math.MaxFloat64
	}
	if qLeft <= left && right <= qRight {
		return st.Tree[node].Max
	}

	st.push(node, left, right)
	mid := (left + right) / 2

	leftMax := st.queryMax(node*2, left, mid, qLeft, qRight)
	rightMax := st.queryMax(node*2+1, mid+1, right, qLeft, qRight)

	return max(leftMax, rightMax)
}

func (st *SegmentTree) Update(pos int, value float64) {
	st.update(1, 0, st.N-1, pos, value)
}

func (st *SegmentTree) update(node, left, right, pos int, value float64) {
	if left == right {
		st.Tree[node] = Node{Sum: value, Min: value, Max: value}
		st.Data[pos] = value
		return
	}

	mid := (left + right) / 2
	if pos <= mid {
		st.update(node*2, left, mid, pos, value)
	} else {
		st.update(node*2+1, mid+1, right, pos, value)
	}

	leftNode := st.Tree[node*2]
	rightNode := st.Tree[node*2+1]

	st.Tree[node] = Node{
		Sum: leftNode.Sum + rightNode.Sum,
		Min: min(leftNode.Min, rightNode.Min),
		Max: max(leftNode.Max, rightNode.Max),
	}
}

func (st *SegmentTree) PrintTree() {
	fmt.Println("\n=== Segment Tree ===")
	for i := 1; i < len(st.Tree); i++ {
		node := st.Tree[i]
		if node.Sum == 0 && node.Min == 0 && node.Max == 0 && node.Lazy == 0 {
			continue
		}
		// %.2f formatea el número para que parezca dinero real (ej. 150.50)
		fmt.Printf("[%d] Sum=%.2f Min=%.2f Max=%.2f Lazy=%.2f\n",
			i, node.Sum, node.Min, node.Max, node.Lazy)
	}
}

func (st *SegmentTree) apply(node, left, right int, delta float64) {
	length := float64(right - left + 1)
	st.Tree[node].Sum += delta * length
	st.Tree[node].Min += delta
	st.Tree[node].Max += delta
	st.Tree[node].Lazy += delta
}

func (st *SegmentTree) push(node, left, right int) {
	lazy := st.Tree[node].Lazy
	if lazy == 0 {
		return
	}

	if left != right {
		mid := (left + right) / 2
		st.apply(node*2, left, mid, lazy)
		st.apply(node*2+1, mid+1, right, lazy)
	}

	st.Tree[node].Lazy = 0
}

func (st *SegmentTree) RangeUpdate(l, r int, delta float64) {
	st.rangeUpdate(1, 0, st.N-1, l, r, delta)
}

func (st *SegmentTree) rangeUpdate(node, left, right, qLeft, qRight int, delta float64) {
	if qRight < left || qLeft > right {
		return
	}

	if qLeft <= left && right <= qRight {
		st.apply(node, left, right, delta)
		return
	}

	st.push(node, left, right)
	mid := (left + right) / 2

	st.rangeUpdate(node*2, left, mid, qLeft, qRight, delta)
	st.rangeUpdate(node*2+1, mid+1, right, qLeft, qRight, delta)

	leftNode := st.Tree[node*2]
	rightNode := st.Tree[node*2+1]

	st.Tree[node].Sum = leftNode.Sum + rightNode.Sum
	st.Tree[node].Min = min(leftNode.Min, rightNode.Min)
	st.Tree[node].Max = max(leftNode.Max, rightNode.Max)
}
