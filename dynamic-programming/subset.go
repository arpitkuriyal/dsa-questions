package main

// Partition Equal Subset Sum (416)
func canPartition(nums []int) bool {
	total := 0
	for _, value := range nums {
		total += value
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, value := range nums {
		for sum := target; sum >= value; sum-- {
			dp[sum] = dp[sum] || dp[sum-value]
		}
	}
	return dp[target]
}
