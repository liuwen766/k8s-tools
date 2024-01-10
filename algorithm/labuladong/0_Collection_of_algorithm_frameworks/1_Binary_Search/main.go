package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	countTarget([]int{1}, 1)
}

// 寻找一个数（基本的二分搜索）
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			//right--
			right = mid - 1
		} else {
			//left++
			left = mid + 1
		}
	}
	return -1
}

// 某班级考试成绩按非严格递增顺序记录于整数数组 scores，请返回目标成绩 target 的出现次数。
func countTarget(scores []int, target int) int {
	ans := 0
	left := searchLeft(scores, target)
	right := searchRight(scores, target)
	if left < right {
		ans = right - left + 1
	}
	// ☆ ☆ ☆ ☆ ☆ 【注意边界条件】
	if left == right && left > -1 {
		ans = 1
	}
	return ans
}

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

// 寻找左边界
func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	// ☆☆☆
	if left < 0 || left > len(nums)-1 {
		return -1
	}
	// ☆☆
	if nums[left] == target {
		return left
	}
	return -1
}

// 寻找右边界
func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	// ☆☆☆
	if right < 0 || right > len(nums)-1 {
		return -1
	}
	// ☆☆
	if nums[right] == target {
		return right
	}
	return -1
}
