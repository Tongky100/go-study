package list

//经典问题链表找环，可以转化为追击问题
//即快慢指针同时从head出发向后遍历，快指针每次走两步(实际上大于等于2都可以，有环一定会相遇)
//慢指针每次走一步，如果存在环两指针会相遇，否则无环
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		//先走后比较，不然slow和fast在起点时一定相等
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
