package main

import (
	"fmt"
	"sort"
)

// Contains Duplicate
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, n := range nums {
		if set[n] {
			return true
		}
		set[n] = true
	}
	return false
}

// Two Sum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if idx, ok := m[target-n]; ok {
			return []int{idx, i}
		}
		m[n] = i
	}
	return nil
}

// Product Except Self
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}

// Longest Consecutive Sequence
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}

	longest := 0

	for n := range set {
		if !set[n-1] {
			length := 1
			for set[n+length] {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

// 3Sum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// Container With Most Water
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := 0

	for l < r {
		h := min(height[l], height[r])
		area := h * (r - l)
		if area > max {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

// Best Time to Buy and Sell Stock
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}

// Kadane (Max Subarray)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curr := 0

	for _, n := range nums {
		if curr < 0 {
			curr = 0
		}
		curr += n
		if curr > maxSum {
			maxSum = curr
		}
	}
	return maxSum
}

func maxSubArrayProduct(nums []int) int {
	currentMax := nums[0]
	currentMin := nums[0]
	answer := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// A negative number swaps the roles of max and min
		if num < 0 {
			currentMax, currentMin = currentMin, currentMax
		}

		// Either start a new subarray or extend the old one
		currentMax = max(num, currentMax*num)
		currentMin = min(num, currentMin*num)

		answer = max(answer, currentMax)
	}

	return answer
}

// Min Size Subarray Sum (Window)
func minSubArrayLen(target int, nums []int) int {
	l := 0
	sum := 0
	res := len(nums) + 1

	for r := 0; r < len(nums); r++ {
		sum += nums[r]

		for sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}

// Move Zeroes
func moveZeroes(nums []int) {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}

// Rotate Array
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// Find Peak Element
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// Majority Element
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// sort Colors
func sortColors(nums []int) {
	low, mid := 0, 0
	high := len(nums) - 1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		default:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

// Gas Station
func canCompleteCircuit(gas []int, cost []int) int {
	total, curr, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		total += diff
		curr += diff

		if curr < 0 {
			start = i + 1
			curr = 0
		}
	}

	if total < 0 {
		return -1
	}
	return start
}

// Car Fleet
type Car struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := []Car{}

	for i := 0; i < n; i++ {
		time := float64(target-position[i]) / float64(speed[i])
		cars = append(cars, Car{position[i], time})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	currTime := 0.0

	for _, c := range cars {
		if c.time > currTime {
			fleets++
			currTime = c.time
		}
	}
	return fleets
}

// Find the Duplicate Number (Floyd's cycle detection)
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// Spiral Matrix
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}
	return result
}

func main() {

	fmt.Println("Contains Duplicate:", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println("Two Sum:", twoSum([]int{2, 7, 11, 15}, 9))
	runHardExamples()

	fmt.Println("Product Except Self:", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println("Longest Consecutive:", longestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	fmt.Println("3Sum:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("Max Area:", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	fmt.Println("Max Profit:", maxProfit([]int{7, 1, 5, 3, 6, 4}))

	fmt.Println("Kadane =", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println("Min Subarray Length =", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println("Move Zeroes =", nums)

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	fmt.Println("Rotate =", arr)

	fmt.Println("Peak Index =", findPeakElement([]int{1, 2, 3, 1}))

	fmt.Println("Majority =", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))

	fmt.Println("Gas Station =", canCompleteCircuit(
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5, 1, 2},
	)) // revision

	fmt.Println("Car Fleet =", carFleet(
		12,
		[]int{10, 8, 0, 5, 3},
		[]int{2, 4, 1, 1, 3},
	)) // revision
	fmt.Println("Find Duplicate Number =", findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println("Spiral Matrix =", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))

}
