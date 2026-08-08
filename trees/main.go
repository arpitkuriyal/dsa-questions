package main

import (
	"fmt"
)

// TreeNode Definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. Level Order Traversal (BFS)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var level []int

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}

	return res
}

// 2. Maximum Depth (DFS)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// 3. Diameter of Binary Tree
func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left+right > maxDia {
			maxDia = left + right
		}

		if left > right {
			return left + 1
		}
		return right + 1
	}

	dfs(root)
	return maxDia
}

// 4. Invert Binary Tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 5. Lowest Common Ancestor
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

// 6. Validate BST
func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	return validate(node.Left, min, &node.Val) &&
		validate(node.Right, &node.Val, max)
}

// 7. Path Sum
func hasPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return target == root.Val
	}

	target -= root.Val

	return hasPathSum(root.Left, target) ||
		hasPathSum(root.Right, target)
}

// Helper: Build Sample Tree
func buildTree() *TreeNode {
	return &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
}

func main() {
	root := buildTree()

	fmt.Println("1. Level Order:", levelOrder(root))
	fmt.Println("2. Max Depth:", maxDepth(root))
	fmt.Println("3. Diameter:", diameterOfBinaryTree(root))

	inverted := invertTree(root)
	fmt.Println("4. Inverted Tree Level Order:", levelOrder(inverted))

	// rebuild tree for correct LCA test
	root = buildTree()
	p := root.Right.Left  // 15
	q := root.Right.Right // 7
	lca := lowestCommonAncestor(root, p, q)
	fmt.Println("5. LCA of 15 and 7:", lca.Val)

	// BST example
	bst := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println("6. Is Valid BST:", isValidBST(bst))

	fmt.Println("7. Path Sum (target=30):", hasPathSum(root, 30))
}
