package main

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
