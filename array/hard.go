package main

import (
	"fmt"
	"slices"
	"sort"
)

// Group Anagrams O(n * k log k)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		arr := []rune(s)
		slices.Sort(arr)
		key := string(arr)
		m[key] = append(m[key], s)
	}
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// optimal O(n*k)
func fastgroupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, word := range strs {
		// A new empty count array for every word
		var count [26]int

		for _, ch := range word {
			count[ch-'a']++
		}

		// Go checks whether this exact array key
		// already exists in the map.
		groups[count] = append(groups[count], word)
	}

	var result [][]string

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Top K Frequent Elements
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range count {
		buckets[freq] = append(buckets[freq], num)
	}

	res := []int{}
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, buckets[i]...)
	}
	return res[:k]
}

// Longest Repeating Character Replacement
func characterReplacement(s string, k int) int {
	count := [26]int{}
	left := 0
	longest := 0
	mostCommon := 0
	for right := 0; right < len(s); right++ {
		letter := s[right] - 'A'
		count[letter]++

		if count[letter] > mostCommon {
			mostCommon = count[letter]
		}

		windowLength := right - left + 1
		replacementsNeeded := windowLength - mostCommon

		if replacementsNeeded > k {
			leftLetter := s[left] - 'A'
			count[leftLetter]--
			left++
		}

		windowLength = right - left + 1
		if windowLength > longest {
			longest = windowLength
		}
	}
	return longest
}

// Subarray Sum Equals K // tricky hard
func subarraySum(nums []int, k int) int {
	seen := map[int]int{0: 1}
	prefix := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		prefix = prefix + nums[i]
		needed := prefix - k
		ans += seen[needed] // Add the number of matching earlier prefixes.
		seen[prefix]++      // Record the current prefix.
	}
	return ans
}

// Merge Intervals
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for _, curr := range intervals[1:] {
		last := res[len(res)-1]

		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			res = append(res, curr)
		}
	}
	return res
}

func runHardExamples() {
	fmt.Println("Group Anagrams:", groupAnagrams([]string{"eat", "tea", "tan", "ate"}))
	fmt.Println("Top K Frequent:", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println("Character Replacement:", characterReplacement("AABABBA", 1))
	fmt.Println("Subarray Sum:", subarraySum([]int{1, 1, 1}, 2))
	fmt.Println("Merge Intervals:", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
