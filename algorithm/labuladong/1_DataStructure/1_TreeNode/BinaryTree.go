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

// 1、从前序与中序遍历序列构造二叉树
func buildTreeFromPreorderAndInorder(preorder []int, inorder []int) *TreeNode {
	return build1(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build1(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]

	rootIndex := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIndex = i
		}
	}
	leftSize := rootIndex - inStart

	root := &TreeNode{rootVal, nil, nil}

	root.Left = build1(preorder, preStart+1, preStart+leftSize, inorder, inStart, inStart+leftSize-1)

	root.Right = build1(preorder, preStart+leftSize+1, preEnd, inorder, rootIndex+1, inEnd)

	return root
}

// 2、从中序与后序遍历序列构造二叉树
func buildTreeFromInorderAndPostOrder(inorder []int, postorder []int) *TreeNode {
	return build2(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build2(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
	// ☆☆☆边界条件
	if inStart > inEnd {
		return nil
	}

	// 1、先找到根节点
	rootVal := postorder[postEnd]

	// 2、区分左右子树
	rootIndex := -1
	for i := range inorder {
		if inorder[i] == rootVal {
			rootIndex = i
		}
	}
	// ☆☆
	leftSize := rootIndex - inStart

	// 3、构造根节点
	root := &TreeNode{rootVal, nil, nil}
	// 4、构造左子树
	root.Left = build2(inorder, inStart, rootIndex-1, postorder, postStart, postStart+leftSize-1)
	// 5、构造右子树
	root.Right = build2(inorder, rootIndex+1, inEnd, postorder, postStart+leftSize, postEnd-1)

	return root
}
