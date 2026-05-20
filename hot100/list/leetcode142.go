package list

//经典问题环形链表找环的起点，先使用快慢指针判断链表是否有环，无环返回null即可
//如果有环，如果slow指针走了k步，那么fast指针就走了2k步，多走的k步恰好就是环的长度
//假设换的起点是m，slow指针从链表起点走到环的起点走了k-m步，而k-m恰好也是fast指针以一倍速
//从相遇点走到环起点的距离，因此当slow和fast相遇时，让slow从链表起点再次出发，slow和fast指针
//每次可走一步，再次相遇时即找到了环的起点
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
