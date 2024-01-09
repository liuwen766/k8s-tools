package main

import (
	"container/heap"
	"fmt"
)

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

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1、合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{-1, nil}
	p := dummy

	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}

		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

// 2、链表的分解
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
func partition(head *ListNode, x int) *ListNode {

	dummy1 := &ListNode{-1, nil}
	dummy2 := &ListNode{-1, nil}
	p1 := dummy1
	p2 := dummy2

	p := head
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}

		// ☆☆☆☆☆
		tmp := p.Next
		p.Next = nil
		p = tmp
	}

	p1.Next = dummy2.Next

	return dummy1.Next
}

// 3、合并 k 个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{-1, nil}
	p := dummy

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := range lists {
		if lists[i] != nil {
			heap.Push(&pq, lists[i])
		}
	}

	for pq.Len() > 0 {
		pop := heap.Pop(&pq).(*ListNode)
		p.Next = pop
		if pop.Next != nil {
			heap.Push(&pq, pop.Next)
		}
		p = p.Next
	}

	return dummy.Next
}

// 4、寻找单链表的倒数第 k 个节点
func trainingPlan(head *ListNode, cnt int) *ListNode {
	quick := head
	for i := 0; i < cnt; i++ {
		quick = quick.Next
	}
	slow := head
	for quick != nil {
		quick = quick.Next
		slow = slow.Next
	}
	return slow
}

// 同类题：删除链表的倒数第 n 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head

	tmp := trainingPlan(dummy, n+1)

	tmp.Next = tmp.Next.Next

	// ☆☆☆
	return dummy.Next
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

// 6、判断单链表是否包含环
func hasCycle(head *ListNode) bool {
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if quick == slow {
			return true
		}
	}
	return false
}

// 同类题：判断单链表是否包含环并找出环起点
func detectCycle(head *ListNode) *ListNode {
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			break
		}
	}

	// ☆☆
	if quick == nil || quick.Next == nil {
		return nil
	}

	slow = head
	for slow != quick {
		slow = slow.Next
		quick = quick.Next
	}

	return slow
}

// 7、判断两个单链表是否相交并找出交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		// p1前进一位
		if p1 != nil {
			p1 = p1.Next
			//	否则p1到headB
		} else {
			p1 = headB
		}
		// p2前进一位
		if p2 != nil {
			p2 = p2.Next
			//	否则p2到headA
		} else {
			p2 = headA
		}
	}

	return p1
}

// PriorityQueue 优先级队列，Go代码实现最小堆
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
