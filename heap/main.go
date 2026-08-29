package main

import (
	"container/heap"
	"fmt"
)

// -----------------------------------------------------------------------------
// 1. Build a Min Heap from Scratch
// -----------------------------------------------------------------------------

// MinHeap stores the smallest value at index 0.
type MinHeap struct {
	data []int
}

// Size returns the number of values in the heap.
func (h *MinHeap) Size() int {
	return len(h.data)
}

// Peek returns the smallest value without removing it.
func (h *MinHeap) Peek() int {
	if len(h.data) == 0 {
		return 0
	}

	return h.data[0]
}

// Push adds a value while preserving the min-heap property.
func (h *MinHeap) Push(value int) {
	h.data = append(h.data, value)

	child := len(h.data) - 1
	for child > 0 {
		parent := (child - 1) / 2

		if h.data[parent] <= h.data[child] {
			break
		}

		h.data[parent], h.data[child] = h.data[child], h.data[parent]

		child = parent
	}
}

// Pop removes and returns the smallest value.
func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		return 0
	}

	min := h.data[0]
	lastIndex := len(h.data) - 1

	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	parent := 0

	for {
		left := 2*parent + 1
		right := 2*parent + 2

		if left >= len(h.data) {
			break
		}

		smallerChild := left

		if right < len(h.data) && h.data[right] < h.data[left] {
			smallerChild = right
		}

		if h.data[parent] <= h.data[smallerChild] {
			break
		}

		h.data[parent], h.data[smallerChild] = h.data[smallerChild], h.data[parent]
		parent = smallerChild
	}
	return min
}

// 2. Last Stone Weight
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, s := range stones {
		heap.Push(h, s)
	}

	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)

		if a != b {
			heap.Push(h, a-b)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}

// 3. Kth Largest Element in an Array
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func findKthLargest(nums []int, k int) int {
	h := &IntMinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

// 4. Top K Frequent Elements
type Freq struct {
	val  int
	freq int
}

type IntMinHeapFreq []Freq

func (h IntMinHeapFreq) Len() int           { return len(h) }
func (h IntMinHeapFreq) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h IntMinHeapFreq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeapFreq) Push(x interface{}) {
	*h = append(*h, x.(Freq))
}

func (h *IntMinHeapFreq) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{}
	for _, n := range nums {
		freqMap[n]++
	}

	h := &IntMinHeapFreq{}
	heap.Init(h)

	for val, freq := range freqMap {
		heap.Push(h, Freq{val, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := []int{}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(Freq).val)
	}
	return res
}

// 5. K Closest Numbers to a Target
type Pair struct {
	val  int
	diff int
}

type MaxHeapPair []Pair

func (h MaxHeapPair) Len() int           { return len(h) }
func (h MaxHeapPair) Less(i, j int) bool { return h[i].diff > h[j].diff }
func (h MaxHeapPair) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapPair) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeapPair) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func kClosestNumbers(nums []int, k int, target int) []int {
	h := &MaxHeapPair{}
	heap.Init(h)

	for _, num := range nums {
		diff := abs(num - target)
		heap.Push(h, Pair{num, diff})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := []int{}
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(Pair).val)
	}
	return res
}

// 6. K Closest Points to the Origin
type Point struct {
	x, y int
	dist int
}

type MaxHeapPoint []Point

func (h MaxHeapPoint) Len() int           { return len(h) }
func (h MaxHeapPoint) Less(i, j int) bool { return h[i].dist > h[j].dist }
func (h MaxHeapPoint) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapPoint) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeapPoint) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func kClosestPoints(points [][]int, k int) [][]int {
	h := &MaxHeapPoint{}
	heap.Init(h)

	for _, p := range points {
		dist := p[0]*p[0] + p[1]*p[1]
		heap.Push(h, Point{p[0], p[1], dist})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := [][]int{}
	for h.Len() > 0 {
		p := heap.Pop(h).(Point)
		res = append(res, []int{p.x, p.y})
	}
	return res
}

// 7. Merge K Sorted Arrays
type Node struct {
	val int
	row int
	idx int
}

type IntMinHeapNode []Node

func (h IntMinHeapNode) Len() int           { return len(h) }
func (h IntMinHeapNode) Less(i, j int) bool { return h[i].val < h[j].val }
func (h IntMinHeapNode) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeapNode) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *IntMinHeapNode) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func mergeKArrays(arr [][]int) []int {
	h := &IntMinHeapNode{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > 0 {
			heap.Push(h, Node{arr[i][0], i, 0})
		}
	}

	res := []int{}
	for h.Len() > 0 {
		node := heap.Pop(h).(Node)
		res = append(res, node.val)

		if node.idx+1 < len(arr[node.row]) {
			heap.Push(h, Node{
				arr[node.row][node.idx+1],
				node.row,
				node.idx + 1,
			})
		}
	}
	return res
}

// 8. Merge K Sorted Lists
type ListNode struct {
	Val  int
	Next *ListNode
}

type IntMinHeapList []*ListNode

func (h IntMinHeapList) Len() int           { return len(h) }
func (h IntMinHeapList) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IntMinHeapList) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeapList) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntMinHeapList) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &IntMinHeapList{}
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

// -----------------------------------------------------------------------------
// Hard
// -----------------------------------------------------------------------------

// 9. Task Scheduler
func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int)
	for _, t := range tasks {
		freq[t]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for _, v := range freq {
		heap.Push(h, v)
	}

	time := 0

	for h.Len() > 0 {
		temp := []int{}
		cycle := n + 1

		for cycle > 0 && h.Len() > 0 {
			val := heap.Pop(h).(int)
			if val-1 > 0 {
				temp = append(temp, val-1)
			}
			time++
			cycle--
		}

		for _, t := range temp {
			heap.Push(h, t)
		}

		if h.Len() == 0 {
			break
		}

		time += cycle
	}

	return time
}

// 10. Median from a Data Stream
type MedianFinder struct {
	left  *MaxHeap
	right *IntMinHeap
}

func Constructor() MedianFinder {
	l := &MaxHeap{}
	r := &IntMinHeap{}
	heap.Init(l)
	heap.Init(r)
	return MedianFinder{l, r}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.left, num)

	heap.Push(m.right, heap.Pop(m.left))

	if m.right.Len() > m.left.Len() {
		heap.Push(m.left, heap.Pop(m.right))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.left.Len() > m.right.Len() {
		return float64((*m.left)[0])
	}
	return float64((*m.left)[0]+(*m.right)[0]) / 2.0
}

// -----------------------------------------------------------------------------
// Helpers and Examples
// -----------------------------------------------------------------------------

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	practiceHeap := &MinHeap{}

	practiceHeap.Push(5)
	practiceHeap.Push(2)
	practiceHeap.Push(8)
	practiceHeap.Push(1)

	fmt.Println("1. Build a Min Heap from Scratch")
	fmt.Println("Heap:        ", practiceHeap.data)   // Expected: [1 2 8 5]
	fmt.Println("Size:        ", practiceHeap.Size()) // Expected: 4
	fmt.Println("Minimum:     ", practiceHeap.Peek()) // Expected: 1
	fmt.Println("Removed min: ", practiceHeap.Pop())  // Expected: 1
	fmt.Println("After pop:   ", practiceHeap.data)   // Expected: [2 5 8]

	fmt.Println("\n2. Last Stone Weight:", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))

	fmt.Println("3. Kth Largest:", findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println("4. Top K Frequent:", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println("7. Merge K Arrays:", mergeKArrays([][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}))
}
