package main

// Climbing Stairs (70)
func climbStairs(n int) int {
	if n <= 2 {
		if n <= 0 {
			return 0
		}
		return n
	}
	prev2, prev1 := 1, 2
	for step := 3; step <= n; step++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}

// House Robber (198)
func rob(nums []int) int {
	prev2, prev1 := 0, 0
	for _, value := range nums {
		prev2, prev1 = prev1, max(prev1, prev2+value)
	}
	return prev1
}

// House Robber II (213)
func robCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
