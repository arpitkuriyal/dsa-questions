package main

import (
	"fmt"
	"sort"
)

// 1. Maximum Subarray (Kadane's Algorithm)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curr := nums[0]

	for i := 1; i < len(nums); i++ {
		curr = max(nums[i], curr+nums[i])
		maxSum = max(maxSum, curr)
	}
	return maxSum
}

// 2. Jump Game (Can reach end)
func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
	}
	return true
}

// 3. Jump Game II (Min jumps to reach end)
func jump(nums []int) int {
	jumps := 0    // How many jumps we've taken.
	end := 0      // How many jumps we've taken.
	farthest := 0 // While scanning the current range, what's the farthest place I can reach with the next jump?

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == end {
			jumps++
			end = farthest
		}
	}
	return jumps
}

// 4. Merge Intervals
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for _, curr := range intervals[1:] {
		last := res[len(res)-1]

		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
			res[len(res)-1] = last
		} else {
			res = append(res, curr)
		}
	}
	return res
}

// 5. Assign Cookies
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

// 6. Best Time to Buy and Sell Stock
func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else {
			profit = max(profit, p-minPrice)
		}
	}
	return profit
}

func main() {
	fmt.Println("Max Subarray:", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println("Can Jump:", canJump([]int{2, 3, 1, 1, 4}))

	fmt.Println("Min Jumps:", jump([]int{2, 3, 1, 1, 4}))

	fmt.Println("Merge Intervals:",
		merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
	)

	fmt.Println("Assign Cookies:",
		findContentChildren([]int{1, 2, 3}, []int{1, 1}),
	)

	fmt.Println("Max Profit:",
		maxProfit([]int{7, 1, 5, 3, 6, 4}),
	)
}
