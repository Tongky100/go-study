package list

//解法1：如果两个链表中存在相交节点，那么依次遍历两根链表中的所有节点一定能找到相交点
//因此使用两个指针p1，p2从两根链表头开始遍历，当p1!=p2时就继续遍历，p1=nil时，转移到另一根链表头继续遍历
//p2同理，如果两根链表不相交最后一定都为nil，假设headA3个节点 headB4个节点，算上各自的空节点一共9个节点
//p1只会交换到headB上一次，不会再从headB上换回headA上
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
