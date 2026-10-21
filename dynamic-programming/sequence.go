package main

// Maximum Product Subarray (152)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxEnding, minEnding, result := nums[0], nums[0], nums[0]
	for _, value := range nums[1:] {
		if value < 0 {
			maxEnding, minEnding = minEnding, maxEnding
		}
		maxEnding = max(value, maxEnding*value)
		minEnding = min(value, minEnding*value)
		result = max(result, maxEnding)
	}
	return result
}

// Coin Change (322)
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for value := 1; value <= amount; value++ {
		dp[value] = amount + 1
		for _, coin := range coins {
			if coin > 0 && coin <= value {
				dp[value] = min(dp[value], dp[value-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
