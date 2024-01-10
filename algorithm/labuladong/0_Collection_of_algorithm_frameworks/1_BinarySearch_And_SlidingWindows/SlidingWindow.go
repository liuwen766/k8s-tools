package main

import "fmt"

func main() {
	fmt.Println("滑动窗口算法")
	s := "bcdedit"
	t := "dec"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(findAnagrams(s, t))
	fmt.Println(checkInclusion(s, t))
	fmt.Println(minWindow(s, t))
}

// 1、无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	return -1
}

// 2、找到字符串中所有字母异位词
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
func findAnagrams(s string, p string) []int {
	return nil
}

// 3、字符串的排列
// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
func checkInclusion(s1 string, s2 string) bool {
	return false
}

// 4、最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s string, t string) string {
	return ""
}
