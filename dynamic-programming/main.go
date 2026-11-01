package main

import "fmt"

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

// Decode Ways (91)
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	prev2, prev1 := 1, 1
	for i := 1; i < len(s); i++ {
		ways := 0
		if s[i] != '0' {
			ways += prev1
		}
		twoDigit := int(s[i-1]-'0')*10 + int(s[i]-'0')
		if twoDigit >= 10 && twoDigit <= 26 {
			ways += prev2
		}
		prev2, prev1 = prev1, ways
	}
	return prev1
}

// Word Break (139)
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	maxLength := 0
	for _, word := range wordDict {
		words[word] = true
		maxLength = max(maxLength, len(word))
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for end := 1; end <= len(s); end++ {
		for start := max(0, end-maxLength); start < end; start++ {
			if dp[start] && words[s[start:end]] {
				dp[end] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// Unique Paths (62)
func uniquePaths(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([]int, n)
	for col := range dp {
		dp[col] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[col] += dp[col-1]
		}
	}
	return dp[n-1]
}

// Longest Common Subsequence (1143)
func longestCommonSubsequence(text1, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		previousDiagonal := 0
		for j := 1; j <= len(text2); j++ {
			old := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = previousDiagonal + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			previousDiagonal = old
		}
	}
	return dp[len(text2)]
}

// Longest Increasing Subsequence (300)
func lengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))
	for _, value := range nums {
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < value {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == len(tails) {
			tails = append(tails, value)
		} else {
			tails[left] = value
		}
	}
	return len(tails)
}

// Longest Palindromic Substring (5)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	expand := func(left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}
	for center := range s {
		left1, right1 := expand(center, center)
		left2, right2 := expand(center, center+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

// Palindromic Substrings (647)
func countSubstrings(s string) int {
	count := 0
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
	}
	for center := range s {
		expand(center, center)
		expand(center, center+1)
	}
	return count
}

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

func main() {
	fmt.Println("Climbing Stairs:", climbStairs(5))
	fmt.Println("House Robber:", rob([]int{1, 2, 3, 1}))
	fmt.Println("House Robber II:", robCircular([]int{2, 3, 2}))
	fmt.Println("Maximum Product Subarray:", maxProduct([]int{2, 3, -2, 4}))
	fmt.Println("Coin Change:", coinChange([]int{1, 2, 5}, 11))
	fmt.Println("Decode Ways:", numDecodings("226"))
	fmt.Println("Word Break:", wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println("Unique Paths:", uniquePaths(3, 7))
	fmt.Println("Longest Common Subsequence:", longestCommonSubsequence("abcde", "ace"))
	fmt.Println("Longest Increasing Subsequence:", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println("Longest Palindromic Substring:", longestPalindrome("babad"))
	fmt.Println("Palindromic Substrings:", countSubstrings("aaa"))
	fmt.Println("Partition Equal Subset Sum:", canPartition([]int{1, 5, 11, 5}))
}
