package main

// simple binary search
func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// Lower Bound = first index where arr[index] >= target
// If target not present, returns position where target can be inserted.
func lowerBound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	ans := len(arr)
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// Upper Bound = first index where arr[index] > target
// If none found, returns len(arr)
func upperBound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	ans := len(arr)
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func firstOccurance(arr []int, target int) int {
	left, right := 0, len(arr)-1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			ans = mid
			right = mid - 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func lastOccurance(arr []int, target int) int {
	left, right := 0, len(arr)-1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			ans = mid
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// Find count of element in sorted Array
func CountElement(arr []int, target int) int {
	firstIdx := firstOccurance(arr, target)
	if firstIdx == -1 {
		return 0
	}

	lastIdx := lastOccurance(arr, target)
	return lastIdx - firstIdx + 1
}

func searchInsertPos(arr []int, target int) int {
	left, right := 0, len(arr)-1
	ans := len(arr)
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func floorAndCeil(arr []int, target int) (int, int) {
	// floor
	floor := func(arr []int, target int) int {
		ans := -1
		left, right := 0, len(arr)-1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] <= target {
				ans = arr[mid]
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return ans
	}

	// ceil
	ceil := func(arr []int, target int) int {
		ans := -1
		left, right := 0, len(arr)-1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] >= target {
				ans = arr[mid]
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return ans
	}

	return floor(arr, target), ceil(arr, target)
}

func searchRotated(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		// left half sorted
		if arr[left] <= arr[mid] {
			if target >= arr[left] && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right half sorted
			if target > arr[mid] && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// Find minimum element in rotated sorted array
func findMin(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] > arr[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left]
}
