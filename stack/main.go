package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type MinStack struct {
	arr      []int
	minStack []int
}

func (m *MinStack) push(val int) {
	m.arr = append(m.arr, val)

	if len(m.minStack) == 0 || val <= m.minStack[len(m.minStack)-1] {
		m.minStack = append(m.minStack, val)
	}
}

func (m *MinStack) pop() {
	val := m.arr[len(m.arr)-1]
	m.arr = m.arr[:len(m.arr)-1]

	if val == m.minStack[len(m.minStack)-1] {
		m.minStack = m.minStack[:len(m.minStack)-1]
	}

}

func (m *MinStack) top() int {
	return m.arr[len(m.arr)-1]
}

func (m *MinStack) getMin() int {
	return m.minStack[len(m.minStack)-1]
}

func (m *MinStack) GetRandom() int {
	idx := rand.Intn(len(m.arr))
	return m.arr[idx]
}

// MyQueue implements a FIFO queue using two LIFO stacks.
type MyQueue struct {
	in  []int
	out []int
}

func (q *MyQueue) push(val int) {
	q.in = append(q.in, val)
}

func (q *MyQueue) moveToOut() {
	if len(q.out) > 0 {
		return
	}

	for len(q.in) > 0 {
		top := q.in[len(q.in)-1]
		q.in = q.in[:len(q.in)-1]
		q.out = append(q.out, top)
	}
}

func (q *MyQueue) pop() int {
	q.moveToOut()
	top := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return top
}

func (q *MyQueue) peek() int {
	q.moveToOut()
	return q.out[len(q.out)-1]
}

func (q *MyQueue) empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func parantheseCheck(s string) bool {
	stack := []rune{}

	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != mp[ch] {
				return false
			}
		}
	}

	return len(stack) == 0
}

func reverseString(s string) string {
	stack := []rune{}

	for _, ch := range s {
		stack = append(stack, ch)
	}

	ans := []rune{}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, top)
	}

	return string(ans)
}

func nextGreater(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	for i := range ans {
		ans[i] = -1
	}

	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[idx] = nums[i]
		}

		stack = append(stack, i)
	}

	return ans
}

func dailyTemperatures(t []int) []int {
	n := len(t)
	ans := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && t[i] > t[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[idx] = i - idx
		}
		stack = append(stack, i)
	}

	return ans
}

type priceSpan struct {
	price int
	span  int
}

// StockSpanner returns the span of consecutive prices less than or equal to
// the latest price.
type StockSpanner struct {
	stack []priceSpan
}

func (s *StockSpanner) next(price int) int {
	span := 1
	for len(s.stack) > 0 && s.stack[len(s.stack)-1].price <= price {
		span += s.stack[len(s.stack)-1].span
		s.stack = s.stack[:len(s.stack)-1]
	}

	s.stack = append(s.stack, priceSpan{price: price, span: span})
	return span
}

type car struct {
	position int
	time     float64
}

// carFleet returns the number of fleets that reach target.
func carFleet(target int, position, speed []int) int {
	cars := make([]car, len(position))
	for i := range position {
		cars[i] = car{
			position: position[i],
			time:     float64(target-position[i]) / float64(speed[i]),
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position < cars[j].position
	})

	fleets := 0
	maxTime := 0.0
	for i := len(cars) - 1; i >= 0; i-- {
		if cars[i].time > maxTime {
			fleets++
			maxTime = cars[i].time
		}
	}

	return fleets
}

// largestRectangleArea returns the largest rectangle that can be formed in a
// histogram, using a monotonic increasing stack of bar indices.
func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		currentHeight := 0
		if i < len(heights) {
			currentHeight = heights[i]
		}

		for len(stack) > 0 && currentHeight < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			leftBoundary := -1
			if len(stack) > 0 {
				leftBoundary = stack[len(stack)-1]
			}

			width := i - leftBoundary - 1
			if area := height * width; area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

func main() {
	// -----------------------------
	// Basic Stack
	// -----------------------------
	fmt.Println("=== Basic Stack ===")
	fmt.Println("Reverse golang:", reverseString("golang"))
	fmt.Println("Valid ()[]{}:", parantheseCheck("()[]{}"))
	fmt.Println("Valid ([)]:", parantheseCheck("([)]"))

	// -----------------------------
	// Stack + Design
	// -----------------------------
	fmt.Println("\n=== MinStack Demo ===")

	s := MinStack{}

	s.push(10)
	s.push(20)
	s.push(30)

	fmt.Println("Stack:", s.arr)
	fmt.Println("MinStack:", s.minStack)
	fmt.Println("Top:", s.top())
	fmt.Println("Min:", s.getMin())

	s.pop()
	s.pop()

	s.push(2)
	s.push(9)
	s.push(3)
	s.push(1)

	fmt.Println("After Operations:")
	fmt.Println("Stack:", s.arr)
	fmt.Println("MinStack:", s.minStack)
	fmt.Println("Top:", s.top())
	fmt.Println("Min:", s.getMin())
	fmt.Println("Random:", s.GetRandom())

	// -----------------------------
	// Queue Using Stacks
	// -----------------------------
	fmt.Println("\n=== Queue Using Stacks ===")
	q := MyQueue{}
	q.push(10)
	q.push(20)
	q.push(30)
	fmt.Println("Peek:", q.peek())
	fmt.Println("Pop:", q.pop())
	fmt.Println("Empty:", q.empty())

	// -----------------------------
	// Monotonic Stack
	// -----------------------------
	fmt.Println("\n=== Next Greater Element ===")

	nums := []int{2, 1, 3, 4}
	fmt.Println(nums, "->", nextGreater(nums))

	// -----------------------------
	// Daily Temperatures
	// -----------------------------
	fmt.Println("\n=== Daily Temperatures ===")

	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(temp, "->", dailyTemperatures(temp))

	// -----------------------------
	// Stock Span
	// -----------------------------
	fmt.Println("\n=== Stock Span ===")
	spanner := StockSpanner{}
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	spans := make([]int, 0, len(prices))
	for _, price := range prices {
		spans = append(spans, spanner.next(price))
	}
	fmt.Println(prices, "->", spans)

	// -----------------------------
	// Car Fleet
	// -----------------------------
	fmt.Println("\n=== Car Fleet ===")
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))

	// -----------------------------
	// Largest Rectangle in Histogram
	// -----------------------------
	fmt.Println("\n=== Largest Rectangle in Histogram ===")
	fmt.Println([]int{2, 1, 5, 6, 2, 3}, "->", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
