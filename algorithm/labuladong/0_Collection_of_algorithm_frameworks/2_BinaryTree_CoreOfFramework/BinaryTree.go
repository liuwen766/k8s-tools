package main

import "fmt"

/*
二叉树题目的递归解法可以分两类思路，
第一类是遍历一遍二叉树得出答案，
第二类是通过分解问题计算出答案，
这两类思路分别对应着 回溯算法核心框架 和 动态规划核心框架。
*/
func main() {
	fmt.Println("二叉树题目的递归解法可以分两类思路，" +
		"第一类是遍历一遍二叉树得出答案，" +
		"第二类是通过分解问题计算出答案，" +
		"这两类思路分别对应着 回溯算法核心框架 和 动态规划核心框架。")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1、二叉树的最大深度——通过分解问题计算出答案——动态规划核心框架
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

// 1、二叉树的最大深度——遍历一遍二叉树得出答案——回溯算法核心框架
var maxAns int

func maxDepth2(root *TreeNode) int {

	return maxAns
}

// 2、二叉树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
func preorderTraversal(root *TreeNode) []int {

	return nil
}
