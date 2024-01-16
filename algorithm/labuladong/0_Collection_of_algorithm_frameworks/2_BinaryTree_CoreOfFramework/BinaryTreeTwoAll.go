package main

import "fmt"

/*
动归/DFS/回溯算法都可以看做二叉树问题的扩展，只是它们的关注点不同：
- 动态规划算法属于分解问题的思路，它的关注点在整棵「子树」。
- 回溯算法属于遍历的思路，它的关注点在节点间的「树枝」。
- DFS 算法属于遍历的思路，它的关注点在单个「节点」。
*/

func main() {
	fmt.Println("动态规划关注整棵「子树」，回溯算法关注节点间的「树枝」，DFS 算法关注单个「节点」。")
}

// 第一个例子，给你一棵二叉树，请你用分解问题的思路写一个 count 函数，计算这棵二叉树共有多少个节点。
// 这就是动态规划分解问题的思路，它的着眼点永远是结构相同的整个子问题，类比到二叉树上就是「子树」
// count 返回给定二叉树的节点总数。
func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := count(root.Left)
	rightCount := count(root.Right)
	return leftCount + rightCount + 1
}

// 第二个例子，给你一棵二叉树，请你用遍历的思路写一个 traverse 函数，打印出遍历这棵二叉树的过程。
// 这就是回溯算法遍历的思路，它的着眼点永远是在节点之间移动的过程，类比到二叉树上就是「树枝」
// 回溯算法秒杀所有排列-组合-子集问题
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("从节点 %s 进入节点 %s", root, root.Left)
	traverse(root.Left)
	fmt.Printf("从节点 %s 回到节点 %s", root.Left, root)

	fmt.Printf("从节点 %s 进入节点 %s", root, root.Right)
	traverse(root.Right)
	fmt.Printf("从节点 %s 回到节点 %s", root.Right, root)
}

// 第三个例子，我给你一棵二叉树，请你写一个 traverse 函数，把这棵二叉树上的每个节点的值都加一。
// 这就是 DFS 算法遍历的思路，它的着眼点永远是在单一的节点上，类比到二叉树上就是处理每个「节点」。
// 一文秒杀所有岛屿题目
func traverse1(root *TreeNode) {
	if root == nil {
		return
	}
	// 遍历过的每个节点的值加一
	root.Val++
	traverse1(root.Left)
	traverse1(root.Right)
}
