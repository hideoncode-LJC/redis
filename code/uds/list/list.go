package list

type any int

type ListNode struct {
	// 前置节点
	pre *ListNode
	// 后驱节点
	next *ListNode
	// 指向任何类型的指针
	value *any
}

type List struct {
	// head
	head *ListNode
	// tail
	tail *ListNode
	// dup 节点值复制
	dup func(*ListNode, *ListNode)
	// free 释放节点
	free func(*ListNode)
	// match
	match func(*ListNode, *ListNode) int
	// length 64bit整数
	len uint64
}

