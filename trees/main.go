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

// 1. Preorder Traversal
func preorderTraversal(root *TreeNode) []int {
	result := []int{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// 2. Inorder Traversal
func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// 3. Level Order Traversal (BFS)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
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
		result = append(result, level)
	}

	return result
}

// 4. Maximum Depth (DFS)
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

// 5. Same Tree
func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

// 6. Symmetric Tree
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isMirror func(*TreeNode, *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}

		return left.Val == right.Val &&
			isMirror(left.Left, right.Right) &&
			isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}

// 7. Invert Binary Tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 8. Balanced Binary Tree
func isBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.Left)
		rightHeight := height(node.Right)
		if leftHeight == -1 || rightHeight == -1 || leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			return -1
		}

		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	return height(root) != -1
}

// 9. Path Sum
func hasPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return target == root.Val
	}

	target -= root.Val
	return hasPathSum(root.Left, target) || hasPathSum(root.Right, target)
}

// 10. Diameter of Binary Tree
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.Left)
		rightHeight := height(node.Right)
		if leftHeight+rightHeight > maxDiameter {
			maxDiameter = leftHeight + rightHeight
		}
		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	height(root)
	return maxDiameter
}

// 11. Validate BST
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

// 12. Lowest Common Ancestor
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

	fmt.Println("1. Preorder Traversal:", preorderTraversal(root))
	fmt.Println("2. Inorder Traversal:", inorderTraversal(root))
	fmt.Println("3. Level Order:", levelOrder(root))
	fmt.Println("4. Max Depth:", maxDepth(root))
	fmt.Println("5. Same Tree:", isSameTree(root, buildTree()))

	symmetric := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println("6. Is Symmetric:", isSymmetric(symmetric))

	inverted := invertTree(buildTree())
	fmt.Println("7. Inverted Tree Level Order:", levelOrder(inverted))

	fmt.Println("8. Is Balanced:", isBalanced(root))
	fmt.Println("9. Path Sum (target=30):", hasPathSum(root, 30))
	fmt.Println("10. Diameter:", diameterOfBinaryTree(root))

	bst := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println("11. Is Valid BST:", isValidBST(bst))

	p := root.Right.Left  // 15
	q := root.Right.Right // 7
	lca := lowestCommonAncestor(root, p, q)
	fmt.Println("12. LCA of 15 and 7:", lca.Val)
}
