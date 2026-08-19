package main

import "fmt"

// 1. Fundamentals

// Question: Implement a Basic Queue.
// Queue is a basic FIFO queue.
type Queue struct{ arr []int }

func (q *Queue) Enqueue(value int) { q.arr = append(q.arr, value) }
func (q *Queue) Dequeue() int {
	if len(q.arr) == 0 {
		return -1
	}
	value := q.arr[0]
	q.arr = q.arr[1:]
	return value
}
func (q *Queue) Peek() int {
	if len(q.arr) == 0 {
		return -1
	}
	return q.arr[0]
}
func (q *Queue) IsEmpty() bool { return len(q.arr) == 0 }

// Question: Implement a Circular Queue.
// CircularQueue reuses freed positions through circular indexing.
type CircularQueue struct {
	arr                   []int
	front, size, capacity int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{arr: make([]int, capacity), capacity: capacity}
}
func (q *CircularQueue) EnQueue(value int) bool {
	if q.capacity == 0 || q.size == q.capacity {
		return false
	}
	rear := (q.front + q.size) % q.capacity
	q.arr[rear], q.size = value, q.size+1
	return true
}
func (q *CircularQueue) DeQueue() bool {
	if q.size == 0 {
		return false
	}
	q.front = (q.front + 1) % q.capacity
	q.size--
	return true
}
func (q *CircularQueue) Front() int {
	if q.size == 0 {
		return -1
	}
	return q.arr[q.front]
}
func (q *CircularQueue) Rear() int {
	if q.size == 0 {
		return -1
	}
	return q.arr[(q.front+q.size-1)%q.capacity]
}
func (q *CircularQueue) IsEmpty() bool { return q.size == 0 }
func (q *CircularQueue) IsFull() bool  { return q.size == q.capacity }

// Question: Implement a Circular Deque.
// CircularDeque supports insertion and removal from both ends.
type CircularDeque struct {
	arr                   []int
	front, size, capacity int
}

func NewCircularDeque(capacity int) *CircularDeque {
	return &CircularDeque{arr: make([]int, capacity), capacity: capacity}
}
func (d *CircularDeque) InsertFront(value int) bool {
	if d.capacity == 0 || d.size == d.capacity {
		return false
	}
	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.arr[d.front], d.size = value, d.size+1
	return true
}
func (d *CircularDeque) InsertLast(value int) bool {
	if d.capacity == 0 || d.size == d.capacity {
		return false
	}
	rear := (d.front + d.size) % d.capacity
	d.arr[rear], d.size = value, d.size+1
	return true
}
func (d *CircularDeque) DeleteFront() bool {
	if d.size == 0 {
		return false
	}
	d.front = (d.front + 1) % d.capacity
	d.size--
	return true
}
func (d *CircularDeque) DeleteLast() bool {
	if d.size == 0 {
		return false
	}
	d.size--
	return true
}
func (d *CircularDeque) GetFront() int {
	if d.size == 0 {
		return -1
	}
	return d.arr[d.front]
}
func (d *CircularDeque) GetRear() int {
	if d.size == 0 {
		return -1
	}
	return d.arr[(d.front+d.size-1)%d.capacity]
}
func (d *CircularDeque) IsEmpty() bool { return d.size == 0 }
func (d *CircularDeque) IsFull() bool  { return d.size == d.capacity }

// 2. Queue <-> Stack

// Question: Implement a Queue using 2 Stacks.
// MyQueue implements a FIFO queue using two LIFO stacks.
type MyQueue struct{ in, out []int }

func (q *MyQueue) Push(value int) { q.in = append(q.in, value) }
func (q *MyQueue) moveToOut() {
	if len(q.out) > 0 {
		return
	}
	for len(q.in) > 0 {
		last := len(q.in) - 1
		q.out, q.in = append(q.out, q.in[last]), q.in[:last]
	}
}
func (q *MyQueue) Pop() int {
	q.moveToOut()
	if len(q.out) == 0 {
		return -1
	}
	last := len(q.out) - 1
	value := q.out[last]
	q.out = q.out[:last]
	return value
}
func (q *MyQueue) Peek() int {
	q.moveToOut()
	if len(q.out) == 0 {
		return -1
	}
	return q.out[len(q.out)-1]
}
func (q *MyQueue) Empty() bool { return len(q.in) == 0 && len(q.out) == 0 }

// Question: Implement a Stack using Queues.
// StackUsingQueues implements a LIFO stack using one queue. Push is O(n).
type StackUsingQueues struct{ queue []int }

func (s *StackUsingQueues) Push(value int) {
	s.queue = append(s.queue, value)
	for i := 0; i < len(s.queue)-1; i++ {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}
func (s *StackUsingQueues) Pop() int {
	if len(s.queue) == 0 {
		return -1
	}
	value := s.queue[0]
	s.queue = s.queue[1:]
	return value
}
func (s *StackUsingQueues) Top() int {
	if len(s.queue) == 0 {
		return -1
	}
	return s.queue[0]
}
func (s *StackUsingQueues) Empty() bool { return len(s.queue) == 0 }

// 3. Queue + Window

// Question: Find the First Negative in every Window of size k.
func firstNegative(nums []int, k int) []int {
	if k <= 0 || k > len(nums) {
		return []int{}
	}
	negativeIndices := []int{}
	result := make([]int, 0, len(nums)-k+1)
	for i, value := range nums {
		if value < 0 {
			negativeIndices = append(negativeIndices, i)
		}
		if len(negativeIndices) > 0 && negativeIndices[0] <= i-k {
			negativeIndices = negativeIndices[1:]
		}
		if i >= k-1 {
			if len(negativeIndices) == 0 {
				result = append(result, 0)
			} else {
				result = append(result, nums[negativeIndices[0]])
			}
		}
	}
	return result
}

// Question: Find the Sliding Window Maximum.
// maxSlidingWindow uses a decreasing deque of indices, so it runs in O(n).
func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || k > len(nums) {
		return []int{}
	}
	deque := []int{}
	result := make([]int, 0, len(nums)-k+1)
	for i := range nums {
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

// 4. Queue + BFS

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Question: Binary Tree Level Order Traversal.
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result, queue := [][]int{}, []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// Question: Rotting Oranges.
// orangesRotting returns the minutes needed to rot all fresh oranges, or -1.
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	type cell struct{ row, col int }
	queue, fresh := []cell{}, 0
	for row := range grid {
		for col, orange := range grid[row] {
			if orange == 2 {
				queue = append(queue, cell{row, col})
			} else if orange == 1 {
				fresh++
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				nextRow, nextCol := current.row+direction[0], current.col+direction[1]
				if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[nextRow]) || grid[nextRow][nextCol] != 1 {
					continue
				}
				grid[nextRow][nextCol] = 2
				fresh--
				queue = append(queue, cell{nextRow, nextCol})
			}
		}
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

// Question: Number of Islands.
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	type cell struct{ row, col int }
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	islands := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != '1' {
				continue
			}
			islands++
			grid[row][col] = '0'
			queue := []cell{{row, col}}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				for _, direction := range directions {
					nextRow, nextCol := current.row+direction[0], current.col+direction[1]
					if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[nextRow]) || grid[nextRow][nextCol] != '1' {
						continue
					}
					grid[nextRow][nextCol] = '0'
					queue = append(queue, cell{nextRow, nextCol})
				}
			}
		}
	}
	return islands
}

func main() {
	// Question: Implement a Basic Queue.
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	fmt.Println("Basic queue:", queue.Dequeue(), queue.Peek())

	// Question: Implement a Circular Queue.
	circularQueue := NewCircularQueue(3)
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	circularQueue.EnQueue(3)
	circularQueue.DeQueue()
	circularQueue.EnQueue(4)
	fmt.Println("Circular queue:", circularQueue.Front(), circularQueue.Rear())

	// Question: Implement a Circular Deque.
	deque := NewCircularDeque(3)
	deque.InsertLast(2)
	deque.InsertFront(1)
	deque.InsertLast(3)
	fmt.Println("Circular deque:", deque.GetFront(), deque.GetRear())

	// Question: Implement a Queue using 2 Stacks.
	stackQueue := MyQueue{}
	stackQueue.Push(1)
	stackQueue.Push(2)
	fmt.Println("Queue using stacks:", stackQueue.Pop(), stackQueue.Peek())

	// Question: Implement a Stack using Queues.
	queueStack := StackUsingQueues{}
	queueStack.Push(1)
	queueStack.Push(2)
	fmt.Println("Stack using queues:", queueStack.Pop(), queueStack.Top())

	// Question: Find the First Negative in every Window of size k.
	fmt.Println("First negatives:", firstNegative([]int{12, -1, -7, 8, -15, 30, 16, 28}, 3))

	// Question: Find the Sliding Window Maximum.
	fmt.Println("Window maximums:", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

	// Question: Binary Tree Level Order Traversal.
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
	fmt.Println("Level order:", levelOrder(tree))

	// Question: Rotting Oranges.
	fmt.Println("Rotting oranges:", orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))

	// Question: Number of Islands.
	fmt.Println("Number of islands:", numIslands([][]byte{{'1', '1', '0', '0'}, {'1', '0', '0', '1'}, {'0', '0', '1', '1'}}))
}
