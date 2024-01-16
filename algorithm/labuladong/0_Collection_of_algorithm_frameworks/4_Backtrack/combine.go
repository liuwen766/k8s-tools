package main

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
func combine(n int, k int) [][]int {
	var res [][]int
	var track []int
	bak(1, n, k, track, &res)
	return res
}

func bak(start int, n int, k int, track []int, res *[][]int) {
	if len(track) == k {
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		bak(i+1, n, k, track, res)
		track = (track)[:len(track)-1]
	}
}
