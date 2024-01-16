package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 8, 2, 9, 2, 4, 6, 2, 8, 4, 6, 9, 7, 3}
	fmt.Println("快速排序就是个二叉树的前序遍历:", sortArray1(arr))
	fmt.Println("归并排序就是个二叉树的后序遍历:", sortArray2(arr))
}

// 1、快速排序就是个二叉树的前序遍历
func sortArray1(nums []int) []int {
	return nums
}

// 2、归并排序就是个二叉树的后序遍历
func sortArray2(nums []int) []int {
	return nil
}
