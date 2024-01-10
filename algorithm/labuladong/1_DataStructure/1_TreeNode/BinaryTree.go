package main

import "fmt"

func main() {
	fmt.Println("数据结构——二叉树")
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
