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
	var pre, cur, nxt *ListNode
	pre = nil
	cur = head
	nxt = head
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
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
// 1 2 3 4 5
// 3 2 1 4 5
func reverseN(head *ListNode, n int) *ListNode {
	var successor *ListNode
	if n == 1 {
		successor = head.Next
		return head
	}

	newHead := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return newHead
}

// 3、反转链表——反转从位置 left 到位置 right 的链表节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		return reverseN(head, right)
	}
	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 4、反转链表——K个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a := head
	b := head
	for i := 0; i < k; i++ {
		// ☆☆  说明不用反转了，直接返回头结点即可
		if b == nil {
			return head
		}
		b = b.Next
	}

	// 先反转前k个，即a到b个
	newHead := reverseAandB(a, b)
	// 再递归
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseAandB(a *ListNode, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	cur = a
	nxt = a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
