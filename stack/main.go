package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Question 1: Reverse a String using a Stack.
func reverseString(s string) string {
	stack := []rune{}
	for _, ch := range s {
		stack = append(stack, ch)
	}

	reversed := []rune{}
	for len(stack) > 0 {
		last := len(stack) - 1
		reversed = append(reversed, stack[last])
		stack = stack[:last]
	}
	return string(reversed)
}

// Question 2: Valid Parentheses.
func parenthesesCheck(s string) bool {
	stack := []rune{}
	matchingOpeningBracket := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != matchingOpeningBracket[ch] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

// Question 3: Design a Min Stack.
type MinStack struct {
	arr      []int
	minStack []int
}

func (m *MinStack) push(value int) {
	m.arr = append(m.arr, value)
	if len(m.minStack) == 0 || value <= m.minStack[len(m.minStack)-1] {
		m.minStack = append(m.minStack, value)
	}
}

func (m *MinStack) pop() int {
	if len(m.arr) == 0 {
		return -1
	}
	last := len(m.arr) - 1
	value := m.arr[last]
	m.arr = m.arr[:last]
	if value == m.minStack[len(m.minStack)-1] {
		m.minStack = m.minStack[:len(m.minStack)-1]
	}
	return value
}

func (m *MinStack) top() int {
	if len(m.arr) == 0 {
		return -1
	}
	return m.arr[len(m.arr)-1]
}

func (m *MinStack) getMin() int {
	if len(m.minStack) == 0 {
		return -1
	}
	return m.minStack[len(m.minStack)-1]
}

func (m *MinStack) getRandom() int {
	if len(m.arr) == 0 {
		return -1
	}
	return m.arr[rand.Intn(len(m.arr))]
}

// Question 4: Implement a Queue using 2 Stacks.
type MyQueue struct {
	in  []int
	out []int
}

func (q *MyQueue) push(value int) { q.in = append(q.in, value) }

func (q *MyQueue) moveToOut() {
	if len(q.out) > 0 {
		return
	}
	for len(q.in) > 0 {
		last := len(q.in) - 1
		q.out = append(q.out, q.in[last])
		q.in = q.in[:last]
	}
}

func (q *MyQueue) pop() int {
	q.moveToOut()
	if len(q.out) == 0 {
		return -1
	}
	last := len(q.out) - 1
	value := q.out[last]
	q.out = q.out[:last]
	return value
}

func (q *MyQueue) peek() int {
	q.moveToOut()
	if len(q.out) == 0 {
		return -1
	}
	return q.out[len(q.out)-1]
}

func (q *MyQueue) isEmpty() bool { return len(q.in) == 0 && len(q.out) == 0 }

// Question 5: Next Greater Element.
func nextGreater(nums []int) []int {
	answer := make([]int, len(nums))
	for i := range answer {
		answer[i] = -1
	}

	stack := []int{} // Indices whose next greater value has not been found.
	for i := range nums {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[index] = nums[i]
		}
		stack = append(stack, i)
	}
	return answer
}

// Question 6: Daily Temperatures.
func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := []int{} // Indices of decreasing temperatures.

	for i := range temperatures {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[index] = i - index
		}
		stack = append(stack, i)
	}
	return answer
}

type priceSpan struct {
	price int
	span  int
}

// Question 7: Online Stock Span.
type StockSpanner struct{ stack []priceSpan }

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

// Question 8: Car Fleet.
func carFleet(target int, position, speed []int) int {
	cars := make([]car, len(position))
	for i := range position {
		cars[i] = car{position: position[i], time: float64(target-position[i]) / float64(speed[i])}
	}
	sort.Slice(cars, func(i, j int) bool { return cars[i].position < cars[j].position })

	fleets := 0
	slowestTime := 0.0
	for i := len(cars) - 1; i >= 0; i-- {
		if cars[i].time > slowestTime {
			fleets++
			slowestTime = cars[i].time
		}
	}
	return fleets
}

// Question 9: Largest Rectangle in Histogram.
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
	// Question 1: Reverse a String using a Stack.
	fmt.Println("Reverse golang:", reverseString("golang"))

	// Question 2: Valid Parentheses.
	fmt.Println("Valid ()[]{}:", parenthesesCheck("()[]{}"))
	fmt.Println("Valid ([)]:", parenthesesCheck("([)]"))

	// Question 3: Design a Min Stack.
	minStack := MinStack{}
	minStack.push(10)
	minStack.push(20)
	minStack.push(2)
	fmt.Println("Min stack top/min/random:", minStack.top(), minStack.getMin(), minStack.getRandom())

	// Question 4: Implement a Queue using 2 Stacks.
	queue := MyQueue{}
	queue.push(10)
	queue.push(20)
	fmt.Println("Queue using stacks:", queue.pop(), queue.peek())

	// Question 5: Next Greater Element.
	fmt.Println("Next greater:", nextGreater([]int{2, 1, 3, 4}))

	// Question 6: Daily Temperatures.
	fmt.Println("Daily temperatures:", dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))

	// Question 7: Online Stock Span.
	spanner := StockSpanner{}
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	spans := make([]int, 0, len(prices))
	for _, price := range prices {
		spans = append(spans, spanner.next(price))
	}
	fmt.Println("Stock span:", spans)

	// Question 8: Car Fleet.
	fmt.Println("Car fleet:", carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))

	// Question 9: Largest Rectangle in Histogram.
	fmt.Println("Largest rectangle:", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
