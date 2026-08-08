package main

import "fmt"

type Queue struct {
	arr []int
}

func (q *Queue) Enqueue(x int) {
	q.arr = append(q.arr, x)
}

func (q *Queue) Dequeue() int {
	if len(q.arr) == 0 {
		return -1
	}
	val := q.arr[0]
	q.arr = q.arr[1:]
	return val
}

func (q *Queue) Peek() int {
	if len(q.arr) == 0 {
		return -1
	}
	return q.arr[0]
}

// -----------------------------

type MyQueue struct {
	in  []int
	out []int
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			n := len(q.in)
			q.out = append(q.out, q.in[n-1])
			q.in = q.in[:n-1]
		}
	}
	n := len(q.out)
	val := q.out[n-1]
	q.out = q.out[:n-1]
	return val
}

// ---------------------------
// We reuse freed space by overwriting elements that are no longer in the queue using circular indexing.
type CircularQueue struct {
	arr                   []int
	front, size, capacity int
}

func Constructor(k int) CircularQueue {
	return CircularQueue{
		arr:      make([]int, k),
		capacity: k,
	}
}

func (q *CircularQueue) EnQueue(value int) bool {
	if q.size == q.capacity {
		return false
	}
	rear := (q.front + q.size) % q.capacity
	q.arr[rear] = value
	q.size++
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

// ---------------------------------------------

func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{}
	result := []int{}

	for i := 0; i < len(nums); i++ {

		// remove out of window
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// remove smaller elements
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

// ------------------------------------------

func firstNegative(nums []int, k int) []int {
	queue := []int{}
	res := []int{}

	for i := 0; i < len(nums); i++ {

		if nums[i] < 0 {
			queue = append(queue, i)
		}

		if len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}

		if i >= k-1 {
			if len(queue) > 0 {
				res = append(res, nums[queue[0]])
			} else {
				res = append(res, 0)
			}
		}
	}
	return res
}

// ----------------------------------

func generateBinary(n int) []string {
	queue := []string{"1"}
	result := []string{}

	for i := 0; i < n; i++ {
		curr := queue[0]
		queue = queue[1:]

		result = append(result, curr)

		queue = append(queue, curr+"0")
		queue = append(queue, curr+"1")
	}
	return result
}

func main() {

	// ------------------ Basic Queue ------------------
	fmt.Println("---- Simple Queue ----")
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Peek:", q.Peek())       // 10
	fmt.Println("Dequeue:", q.Dequeue()) // 10
	fmt.Println("Dequeue:", q.Dequeue()) // 20
	fmt.Println("Peek:", q.Peek())       // 30

	// ------------------ Queue using 2 Stacks ------------------
	fmt.Println("\n---- Queue using 2 Stacks ----")
	mq := MyQueue{}
	mq.Push(1)
	mq.Push(2)
	mq.Push(3)

	fmt.Println("Pop:", mq.Pop()) // 1
	fmt.Println("Pop:", mq.Pop()) // 2

	// ------------------ Circular Queue ------------------
	fmt.Println("\n---- Circular Queue ----")
	cq := Constructor(3)

	fmt.Println(cq.EnQueue(1)) // true
	fmt.Println(cq.EnQueue(2)) // true
	fmt.Println(cq.EnQueue(3)) // true
	fmt.Println(cq.EnQueue(4)) // false (full)

	fmt.Println(cq.DeQueue())  // true
	fmt.Println(cq.EnQueue(4)) // true

	fmt.Println("Front index value:", cq.arr[cq.front])

	// ------------------ Sliding Window Maximum ------------------
	fmt.Println("\n---- Sliding Window Maximum ----")
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
	// Expected: [3 3 5 5 6 7]

	// ------------------ First Negative in Window ------------------
	fmt.Println("\n---- First Negative in Window ----")
	arr := []int{12, -1, -7, 8, -15, 30, 16, 28}
	fmt.Println(firstNegative(arr, 3))
	// Expected: [-1 -1 -7 -15 -15 0]

	// ------------------ Generate Binary Numbers ------------------
	fmt.Println("\n---- Generate Binary ----")
	fmt.Println(generateBinary(5))
	// Expected: [1 10 11 100 101]
}
