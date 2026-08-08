package main

import (
	"fmt"
	"sort"
)

// 1. Subsets (Include / Exclude)
func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)

	dfs = func(i int, curr []int) {
		if i == len(nums) {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		// include
		curr = append(curr, nums[i])
		dfs(i+1, curr)

		// exclude
		curr = curr[:len(curr)-1]
		dfs(i+1, curr)
	}

	dfs(0, []int{})
	return res
}

// 2. Subsets II (skip duplicates)
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var dfs func(int, []int)

	dfs = func(start int, curr []int) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			dfs(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, []int{})
	return res
}

// 3. Permutations
func permute(nums []int) [][]int {
	var res [][]int

	var dfs func(int)
	dfs = func(start int) {
		if start == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	dfs(0)
	return res
}

// 4. Permutations II (with duplicates)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	used := make([]bool, len(nums))

	var dfs func([]int)
	dfs = func(curr []int) {
		if len(curr) == len(nums) {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			curr = append(curr, nums[i])
			dfs(curr)
			curr = curr[:len(curr)-1]
			used[i] = false
		}
	}

	dfs([]int{})
	return res
}

// 5. Combination Sum (reuse allowed)
func combinationSum(nums []int, target int) [][]int {
	var res [][]int

	var dfs func(int, int, []int)
	dfs = func(i, sum int, curr []int) {
		if sum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}
		if i >= len(nums) || sum > target {
			return
		}

		// pick
		curr = append(curr, nums[i])
		dfs(i, sum+nums[i], curr)

		// skip
		curr = curr[:len(curr)-1]
		dfs(i+1, sum, curr)
	}

	dfs(0, 0, []int{})
	return res
}

// 6. Combination Sum II (no reuse)
func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int

	var dfs func(int, int, []int)
	dfs = func(start, sum int, curr []int) {
		if sum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			if sum+nums[i] > target {
				break
			}

			curr = append(curr, nums[i])
			dfs(i+1, sum+nums[i], curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, 0, []int{})
	return res
}

// 7. Palindrome Partitioning
func partition(s string) [][]string {
	var res [][]string

	var isPal func(string) bool
	isPal = func(str string) bool {
		i, j := 0, len(str)-1
		for i < j {
			if str[i] != str[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var dfs func(int, []string)
	dfs = func(start int, curr []string) {
		if start == len(s) {
			temp := make([]string, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for i := start; i < len(s); i++ {
			sub := s[start : i+1]
			if isPal(sub) {
				curr = append(curr, sub)
				dfs(i+1, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}

	dfs(0, []string{})
	return res
}

// 8. Generate Parentheses
func generateParenthesis(n int) []string {
	var res []string

	var dfs func(string, int, int)
	dfs = func(curr string, open, close int) {
		if len(curr) == 2*n {
			res = append(res, curr)
			return
		}

		if open < n {
			dfs(curr+"(", open+1, close)
		}
		if close < open {
			dfs(curr+")", open, close+1)
		}
	}

	dfs("", 0, 0)
	return res
}

// 9. Word Search (Grid DFS)
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var dfs func(int, int, int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[i] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		found := dfs(r+1, c, i+1) ||
			dfs(r-1, c, i+1) ||
			dfs(r, c+1, i+1) ||
			dfs(r, c-1, i+1)

		board[r][c] = temp
		return found
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

// 10. N-Queens
func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var isValid func(int, int) bool
	isValid = func(r, c int) bool {
		for i := 0; i < r; i++ {
			if board[i][c] == 'Q' {
				return false
			}
		}
		for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		for i, j := r-1, c+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			temp := make([]string, n)
			for i := range board {
				temp[i] = string(board[i])
			}
			res = append(res, temp)
			return
		}

		for c := 0; c < n; c++ {
			if isValid(r, c) {
				board[r][c] = 'Q'
				dfs(r + 1)
				board[r][c] = '.'
			}
		}
	}

	dfs(0)
	return res
}

func main() {
	fmt.Println("1. Subsets:", subsets([]int{1, 2}))
	fmt.Println("2. Subsets II:", subsetsWithDup([]int{1, 2, 2}))
	fmt.Println("3. Permutations:", permute([]int{1, 2, 3}))
	fmt.Println("4. Permutations II:", permuteUnique([]int{1, 1, 2}))
	fmt.Println("5. Combination Sum:", combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println("6. Combination Sum II:", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println("7. Palindrome Partition:", partition("aab"))
	fmt.Println("8. Generate Parentheses:", generateParenthesis(3))

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println("9. Word Search:", exist(board, "ABCCED"))

	fmt.Println("10. N-Queens:", solveNQueens(4))
}
