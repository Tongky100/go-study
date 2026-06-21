package list

//解法1 递归求解
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head //base case 不足两个节点无法翻转
	}

	node1, node2 := head.Next, head.Next.Next
	node1.Next = head
	head.Next = swapPairs(node2)
	return node1
}

//解法2 迭代求解
func swapPairs2(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	//     return head //base case 不足两个节点无法翻转
	// } 接了dummyNode，就算head是空链也满足
	dummy := &ListNode{Next: head}
	node0, node1 := dummy, head
	for node1 != nil && node1.Next != nil {
		node2, node3 := node1.Next, node1.Next.Next //不能写成node2.Next因为拿的是node2的旧值
		node0.Next = node2
		node2.Next = node1
		node1.Next = node3
		node0 = node1
		node1 = node3
	}
	return dummy.Next
}
