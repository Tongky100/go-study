package list

//解法1 使用递归求解
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

//解法2 使用双指针求解
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return dummy.Next
}
