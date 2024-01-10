package main

import "fmt"

func main() {
	fmt.Println("")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1、二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + getMax(leftDepth, rightDepth)
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2、二叉树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
func preorderTraversal(root *TreeNode) []int {
	
	return nil
}
