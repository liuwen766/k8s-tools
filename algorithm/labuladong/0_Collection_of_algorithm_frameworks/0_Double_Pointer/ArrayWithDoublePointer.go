package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 9, 9, 10, 15, 15, 18}
	sum := twoSum(arr, 9)
	fmt.Println("arr:", arr, "sum:", sum)
}

// 1、两数之和【找出这两个数的下标】
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{-1, -1}
	}
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}

// 同类题，找出这两个数
func twosum(price []int, target int) []int {
	left := 0
	right := len(price) - 1
	for left < right {
		if price[left]+price[right] == target {
			return []int{price[left], price[right]}
		} else if price[left]+price[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

// 2、删除有序数组中的重复项，返回删除后数组的新长度。
func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			// ☆☆ 注意是先slow++
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	// ☆☆ 注意是slow + 1
	return slow + 1
}

// 3、移除元素
func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			// ☆☆ 注意是后slow++
			slow++
		}
		fast++
	}
	// ☆☆ 注意是slow
	return slow
}

// 4、移动零
func moveZeroes(nums []int) {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// 将后面置为0
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 5、反转字符串
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}

// 6、最长回文子串
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {

		s1 := getMaxLen(s, i, i)
		s2 := getMaxLen(s, i, i+1)

		if len(res) > len(s1) {
			res = res
		} else {
			res = s1
		}

		if len(res) > len(s2) {
			res = res
		} else {
			res = s2
		}

	}
	return res
}

func getMaxLen(s string, left int, right int) string {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

// 7、删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {

	// ☆ 注意边界条件
	if head == nil {
		return nil
	}

	slow := head
	fast := head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
