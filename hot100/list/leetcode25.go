package list

//解法1 使用递归求解
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head //base case
	}
	p := head
	//查找链表中是否存在k个节点，不足ke个节点时不进行翻转直接返回
	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	newHead := revsereK(head, k)
	head.Next = reverseKGroup(head.Next, k)
	return newHead
}

var successor *ListNode

func revsereK(head *ListNode, k int) *ListNode {
	if k == 1 {
		successor = head.Next
		return head
	}
	k--
	node := revsereK(head.Next, k)
	head.Next.Next = head
	head.Next = successor
	return node
}

//解法2 迭代求解
func reverseKGroup2(head *ListNode, k int) *ListNode {
	//统计节点个数
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{Next: head}
	// p0, pre := dummy, nil 短变量声明 := 有强制规则 左侧至少要有一个全新定义的变量，同时Go无法从nil自动推导指针类型。
	p0, cur := dummy, head
	var pre *ListNode
	//k个一组翻转
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next //nxt是短变量，只在循环体内有效
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		nxt := p0.Next     //记录每次翻转前的头节点
		p0.Next.Next = cur //翻转前的头节点是现在的尾节点和后续链表进行拼接
		p0.Next = pre      //p0和翻转后的新头节点进行拼接
		p0 = nxt           //p0放到下次翻转链表头节点的前一个节点上，语义上一直充当dummy的角色
	}
	return dummy.Next
}
