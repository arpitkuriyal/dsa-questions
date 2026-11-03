package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid Anagram
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

// valid anagram solution 2 with space complexity 0 but worked with lower case character
func validAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}

// Encode & Decode Strings (simple join)
func encode(strs []string) string {
	var encoded strings.Builder

	for _, word := range strs {
		encoded.WriteString(strconv.Itoa(len(word)))
		encoded.WriteByte('#')
		encoded.WriteString(word)
	}

	return encoded.String()
}

func decode(s string) []string {
	res := []string{}
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		j++
		res = append(res, s[j:j+length])
		i = j + length
	}
	return res
}

// Valid Palindrome
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isAlnum(s[l]) {
			l++
			continue
		}
		if !isAlnum(s[r]) {
			r--
			continue
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlnum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	l := 0
	maxLen := 0

	for r := 0; r < len(s); r++ {
		for set[s[r]] {
			delete(set, s[l])
			l++
		}
		set[s[r]] = true
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}

// Permutation in String
func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var need, window [26]int
	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	if need == window {
		return true
	}
	for right := len(s1); right < len(s2); right++ {
		window[s2[right]-'a']++
		window[s2[right-len(s1)]-'a']--
		if need == window {
			return true
		}
	}
	return false
}

// Minimum Window Substring
func minWindow(s, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, missing := 0, len(t)
	bestStart, bestLength := 0, len(s)+1
	for right := 0; right < len(s); right++ {
		if need[s[right]] > 0 {
			missing--
		}
		need[s[right]]--
		for missing == 0 {
			if length := right - left + 1; length < bestLength {
				bestStart, bestLength = left, length
			}
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++
			}
			left++
		}
	}
	if bestLength == len(s)+1 {
		return ""
	}
	return s[bestStart : bestStart+bestLength]
}

func main() {
	fmt.Println("Permutation in String =", checkInclusion("ab", "eidbaooo"))
	fmt.Println("Minimum Window Substring =", minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println("Longest Substring:", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("Valid Palindrome:", isPalindrome("A man, a plan, a canal: Panama"))

	encoded := encode([]string{"hello", "world"})
	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decode(encoded)) // revision

	fmt.Println("Valid Anagram:", isAnagram("anagram", "nagaram"))
}
