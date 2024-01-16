package main

import "fmt"

// 回溯算法秒杀所有排列-组合-子集问题
func main() {
	fmt.Println("回溯算法秒杀所有排列-组合-子集问题")
	num := []int{1, 2, 3, 4, 5}
	fmt.Println("子集问题：", subsets(num))
	fmt.Println("组合问题：", combine(5, 3))
	fmt.Println("排列问题：", permute(num))
}

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	return nil
}
