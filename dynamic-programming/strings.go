package main

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
