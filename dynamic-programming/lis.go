package main

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
