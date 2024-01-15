package main

import "fmt"

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func main() {
	fmt.Println("爬楼梯共有多少种爬法：", climbStairs(36))
}

// 运行超时
func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 带备忘录/迭代的解法
func climbStairs1(n int) int {
	meme := make([]int, n+1)
	for i := 0; i < len(meme); i++ {
		meme[i] = -1
	}
	return climb2(meme, n)
}

func climb1(meme []int, n int) int {
	if n <= 2 {
		return n
	}
	if meme[n] != -1 {
		return meme[n]
	}
	return climb1(meme, n-1) + climb1(meme, n-2)
}

func climb2(meme []int, n int) int {
	if n <= 2 {
		return n
	}
	meme[0] = 0
	meme[1] = 1
	meme[2] = 2
	for i := 3; i < len(meme); i++ {
		meme[i] = meme[i-1] + meme[i-2]
	}
	return meme[n]
}
