package main

import (
	"fmt"
	"math/rand"
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

func main() {
	// -----------------------------
	// MinStack Demo
	// -----------------------------
	fmt.Println("=== MinStack Demo ===")

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
	// Parentheses Check
	// -----------------------------
	fmt.Println("\n=== Parentheses Check ===")

	str1 := "()[]{}"
	str2 := "([)]"

	fmt.Println(str1, "->", parantheseCheck(str1))
	fmt.Println(str2, "->", parantheseCheck(str2))

	// -----------------------------
	// Reverse String
	// -----------------------------
	fmt.Println("\n=== Reverse String ===")

	word := "golang"
	fmt.Println(word, "->", reverseString(word))

	// -----------------------------
	// Next Greater Element
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
}
