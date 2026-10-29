package main

import "fmt"

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
