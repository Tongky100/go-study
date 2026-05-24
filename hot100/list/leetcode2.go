package list

//解法1 使用迭代求解
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}
	p, remain, v1, v2 := dummy, 0, 0, 0
	for l1 != nil || l2 != nil || remain > 0 {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next

		} else {
			v2 = 0
		}
		node := &ListNode{
			Val: (v1 + v2 + remain) % 10,
		}
		remain = (v1 + v2 + remain) / 10
		p.Next = node
		p = p.Next
	}
	return dummy.Next
}

//解法2 对解法1中的迭代写法做优化
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	//使用指针才能和nil比较，p.Next 自动 = (*p).Next是go做的隐式解引用
	//解引用是因为Next并不是指针的属性，而是指针指向对象的属性
	p := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 { //从题目数据范围知所有节点值都是非负数
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		//p.Next = &ListNode{Val: sum % 10}//写在同一行不需要逗号
		p.Next = &ListNode{
			Val: sum % 10,
		}
		carry = sum / 10
		p = p.Next
	}
	return dummy.Next
}

//解法3 使用递归求解，创建新的节点
func addTwoNumbers3(node1 *ListNode, node2 *ListNode) *ListNode {
	var dfs func(*ListNode, *ListNode, int) *ListNode
	dfs = func(l1, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil && carry == 0 {
			return nil //base case
		}

		s := carry
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		//按结构体的字段顺序赋值可以省略参数名,但是要对所有参数都进行赋值
		//return &ListNode{s % 10, dfs(l1, l2, s / 10)}
		return &ListNode{Val: s % 10, Next: dfs(l1, l2, s/10)}
	}
	return dfs(node1, node2, 0)
}

//解法4 使用递归求解，不创建新节点在l1节点上进行覆盖
func addTwoNumbers4(node1 *ListNode, node2 *ListNode) *ListNode {
	var dfs func(*ListNode, *ListNode, int) *ListNode
	dfs = func(l1, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil {
			if carry != 0 {
				return &ListNode{Val: carry} //一行写完，最后一个赋值的字段不用加逗号
			}
			return nil //base case
		}

		if l1 == nil {
			//交换l1和l2的位置，能到这个if分支，说明l2一定不为nil
			l1, l2 = l2, l1
		}

		s := carry + l1.Val
		if l2 != nil { //在l1，l2没交换，也就是都不为空的时候，s也要加上l2节点的值
			s += l2.Val
			l2 = l2.Next
		}

		l1.Val = s % 10 //原地修改l1的值
		l1.Next = dfs(l1.Next, l2, s/10)
		return l1
	}
	return dfs(node1, node2, 0)
}
