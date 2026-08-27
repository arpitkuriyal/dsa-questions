package main

import (
	"fmt"
	"slices"
	"sort"
)

// Contains Duplicate
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, n := range nums {
		if set[n] {
			return true
		}
		set[n] = true
	}
	return false
}

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

// Two Sum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if idx, ok := m[target-n]; ok {
			return []int{idx, i}
		}
		m[n] = i
	}
	return nil
}

// Group Anagrams
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

// Encode & Decode Strings (simple join)
func encode(strs []string) string {
	res := ""
	for _, s := range strs {
		res += fmt.Sprintf("%d#%s", len(s), s)
	}
	return res
}

func decode(s string) []string {
	res := []string{}
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		var length int
		fmt.Sscanf(s[i:j], "%d", &length)
		j++
		res = append(res, s[j:j+length])
		i = j + length
	}
	return res
}

// Product Except Self
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}

// Longest Consecutive Sequence
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}

	longest := 0

	for n := range set {
		if !set[n-1] {
			length := 1
			for set[n+length] {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
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

// 3Sum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// Container With Most Water
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := 0

	for l < r {
		h := min(height[l], height[r])
		area := h * (r - l)
		if area > max {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

// Best Time to Buy and Sell Stock
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
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

// Longest Repeating Character Replacement
func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	l := 0
	maxCount := 0
	res := 0

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		maxCount = max(maxCount, count[s[r]])

		if (r-l+1)-maxCount > k {
			count[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

// Subarray Sum Equals K
func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	m := map[int]int{0: 1}

	for _, n := range nums {
		prefixSum += n
		if val, ok := m[prefixSum-k]; ok {
			count += val
		}
		m[prefixSum]++
	}
	return count
}

// Kadane (Max Subarray)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curr := 0

	for _, n := range nums {
		if curr < 0 {
			curr = 0
		}
		curr += n
		if curr > maxSum {
			maxSum = curr
		}
	}
	return maxSum
}

func maxSubArrayProduct(nums []int) int {
	currentMax := nums[0]
	currentMin := nums[0]
	answer := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// A negative number swaps the roles of max and min
		if num < 0 {
			currentMax, currentMin = currentMin, currentMax
		}

		// Either start a new subarray or extend the old one
		currentMax = max(num, currentMax*num)
		currentMin = min(num, currentMin*num)

		answer = max(answer, currentMax)
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Min Size Subarray Sum (Window)
func minSubArrayLen(target int, nums []int) int {
	l := 0
	sum := 0
	res := len(nums) + 1

	for r := 0; r < len(nums); r++ {
		sum += nums[r]

		for sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}

// Move Zeroes
func moveZeroes(nums []int) {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}

// Rotate Array
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// Find Peak Element
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// Majority Element
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// sort Colors
func sortColors(nums []int) {
	low, mid := 0, 0
	high := len(nums) - 1
	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

// Gas Station
func canCompleteCircuit(gas []int, cost []int) int {
	total, curr, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		total += diff
		curr += diff

		if curr < 0 {
			start = i + 1
			curr = 0
		}
	}

	if total < 0 {
		return -1
	}
	return start
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

// Car Fleet
type Car struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := []Car{}

	for i := 0; i < n; i++ {
		time := float64(target-position[i]) / float64(speed[i])
		cars = append(cars, Car{position[i], time})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	currTime := 0.0

	for _, c := range cars {
		if c.time > currTime {
			fleets++
			currTime = c.time
		}
	}
	return fleets
}

func main() {

	fmt.Println("Contains Duplicate:", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println("Valid Anagram:", isAnagram("anagram", "nagaram"))
	fmt.Println("Two Sum:", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Group Anagrams:", groupAnagrams([]string{"eat", "tea", "tan", "ate"}))
	fmt.Println("Top K Frequent:", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	encoded := encode([]string{"hello", "world"})
	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decode(encoded))

	fmt.Println("Product Except Self:", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println("Longest Consecutive:", longestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	fmt.Println("Valid Palindrome:", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("3Sum:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("Max Area:", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	fmt.Println("Max Profit:", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("Longest Substring:", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("Character Replacement:", characterReplacement("AABABBA", 1))

	fmt.Println("Subarray Sum =", subarraySum([]int{1, 1, 1}, 2))
	fmt.Println("Kadane =", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println("Min Subarray Length =", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println("Move Zeroes =", nums)

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	fmt.Println("Rotate =", arr)

	fmt.Println("Peak Index =", findPeakElement([]int{1, 2, 3, 1}))

	fmt.Println("Majority =", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))

	fmt.Println("Gas Station =", canCompleteCircuit(
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5, 1, 2},
	))

	fmt.Println("Merge Intervals =", merge([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))

	fmt.Println("Car Fleet =", carFleet(
		12,
		[]int{10, 8, 0, 5, 3},
		[]int{2, 4, 1, 1, 3},
	))
}
