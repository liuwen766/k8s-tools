package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	list1 := &ListNode{Val: 1}
	list2 := &ListNode{Val: 3}
	list3 := &ListNode{Val: 5}

	list1.Next = list2
	list2.Next = list3

	list4 := &ListNode{Val: 2}
	list5 := &ListNode{Val: 4}
	list6 := &ListNode{Val: 6}

	list4.Next = list5
	list5.Next = list6

	printListNode(list1)

	ans1 := mergeTwoLists(list1, list4)
	printListNode(ans1)

	ans2 := middleNode(list1)
	printListNode(ans2)

}

func printListNode(lists *ListNode) {
	for lists != nil {
		fmt.Print(lists.Val, "→")
		lists = lists.Next
	}
	fmt.Println()
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1、合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return list1
}

// 5、寻找单链表的中点
func middleNode(head *ListNode) *ListNode {
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}
	return slow
}
