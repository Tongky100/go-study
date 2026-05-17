package list

//从题目数据量范围提示来看链表不为空

//解法1 先查找链表中点，找到链表中点后，反转后半段链表，然后比较每个节点的值是否相等
func isPalindrome(head *ListNode) bool {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true

}

//查找链表中点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

//解法2 利用全局指针递归求解
func isPalindrome2(head *ListNode) bool {
	left := head //从左向右遍历的全局指针
	var dfs func(*ListNode) bool
	dfs = func(head *ListNode) bool {
		if head == nil {
			return true //base case 遍历到头了
		}
		ans := dfs(head.Next)
		if !ans {
			return false //剪枝提速
		}
		if left.Val != head.Val {
			ans = false
		}
		left = left.Next
		return ans
	}
	return dfs(head)
}

//解法3 更优雅的递归写法
func isPalindrome3(head *ListNode) bool {
	left := head //从左向右遍历的全局指针
	var dfs func(*ListNode) bool
	dfs = func(right *ListNode) bool {
		if right.Next != nil && !dfs(right.Next) {
			return false //子链存在且不是回文链表时返回false
		}
		if left.Val != right.Val {
			return false //首尾节点值不一样时返回false
		}
		left = left.Next //左指针右移
		return true
	}
	return dfs(head)
}
