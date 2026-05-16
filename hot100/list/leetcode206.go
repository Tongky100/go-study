package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法1 双指针反转链表
func reverseList(head *ListNode) *ListNode {
	// var prev *ListNode
	// cur := head
	// 因为nil不能做类型推断 写成prev, cur := nil, head是错的
	//var prev, cur *ListNode = nil, head 下面这两种写法都可以
	prev, cur := (*ListNode)(nil), head
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}

// 解法2 递归反转链表
func reverseList2(head *ListNode) *ListNode {
	//base case 链表是空链或者遍历到了最后一个节点
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next)
	head.Next.Next = head //反转
	head.Next = nil       //原来head向后的指针断开
	return node           //一直返回新链表的头
}
