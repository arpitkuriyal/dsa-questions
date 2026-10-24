package main

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
