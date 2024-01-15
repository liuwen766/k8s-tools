package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("动态规划解题思路：明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义")
	fmt.Println("斐波那契数:", fib1(15))
	fmt.Println("凑零钱问题:", coinChange1([]int{1, 2, 5}, 16))
}

/*
下面通过 斐波那契数列问题 和 凑零钱问题 来详解动态规划的基本原理。
前者主要是让你明白什么是重叠子问题，
后者主要举集中于如何列出状态转移方程。
*/

//-----------------斐波那契数列问题--------------------

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

// 1、斐波那契数解法3——带备忘录的迭代解法——空间复杂度O(n)
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

// 1、斐波那契数解法4——迭代解法——空间复杂度O(1)
func fib4(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp1 := 0
	dp2 := 1
	for i := 2; i <= n; i++ {
		dpi := dp1 + dp2
		dp1 = dp2
		dp2 = dpi
	}
	return dp2
}

//--------------------凑零钱问题-------------------------

func coinChange1(coins []int, amount int) int {
	return dp(coins, amount)
}

// 该解法会超时
func dp(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	res := math.MaxInt32

	for i := range coins {
		subProblem := dp(coins, amount-coins[i])
		if subProblem == -1 {
			continue
		}
		res = getMin(res, subProblem+1)
	}

	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// 带有备忘录的解法
func coinChange2(coins []int, amount int) int {
	meme := make([]int, amount+1)
	return meme[0]
}
