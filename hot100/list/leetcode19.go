package list

//数学问题 以示例1为例，从节点1走到链表末尾(NIL节点)恰好要走一个链表长度也就是走len步
//我们先使用p1指针从节点1走n步，然后使用另一个指针p2从节点1开始走，p1走了n步后再走到nil，走了len-n步
//而p2也走了len-n步，转化成索引就是所求，因为倒数第n个节点，就是整数第len-n个节点对应的索引，示例1中
//倒数第2个节点就是正数第4个节点，对应的索引是3,由于我们要删除这个节点，因此我们应该找倒数第n+1个节点
//修改倒数第n+1个节点的Next指针，也就是把节点3和节点5拼接在一起
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	p1, p2 := dummy, dummy
	for ; n+1 > 0; n-- {
		p1 = p1.Next
	}
	for ; p1 != nil; p1 = p1.Next {
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}

//解法2 第二次for循环时查找倒数第n+1个节点
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	for ; n > 0; n-- {
		p1 = p1.Next //先走N步，为了找倒数n个节点
	}
	for p1.Next != nil { //少遍历一次就是找的倒数第n+1个节点了
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next
}
