package main

import "fmt"

func main() {
	fmt.Println("动态规划解题思路：明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义")
}

// 1、斐波那契数解法1——暴力递归
func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	res := fib1(n-1) + fib1(n-2)
	return res
}

// 1、斐波那契数解法2——带备忘录的递归解法
func fib2(n int) int {
	bak := make([]int, n+1)
	return bib2withBak(bak, n)
}

func bib2withBak(bak []int, n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if bak[n] != 0 {
		return bak[n]
	}
	return bib2withBak(bak, n-1) + bib2withBak(bak, n-2)
}

// 1、斐波那契数解法3——带备忘录的迭代解法
func fib3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
