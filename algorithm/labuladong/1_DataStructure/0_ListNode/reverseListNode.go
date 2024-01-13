package main

import "fmt"

func main() {
	fmt.Println("反转链表!")
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1、反转链表1
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	return pre
}

// 递归解法
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 2、反转链表——反转链表的前n个节点
func reverseN(head *ListNode, n int) *ListNode {
	return head
}

// 3、反转链表——反转从位置 left 到位置 right 的链表节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	return head
}

// 4、反转链表——K个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	return head
}
