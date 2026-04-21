package binarytree

//解法1 两个数组
func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	cur := []*TreeNode{root} //使用cur切片装当前层的节点，可以通过len(cur)统计出当前层节点的个数
	for len(cur) > 0 {       //队列中还有节点就要继续遍历
		nxt := []*TreeNode{}          //装下一行节点的切片
		vals := make([]int, len(cur)) //装当前行节点值
		for i, node := range cur {
			vals[i] = node.Val
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
		ans = append(ans, vals)
	}
	return
}

//解法2 一个数组
func levelOrder2(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root} //q是java写法中的队列
	for len(q) > 0 {       //队列中还有节点就要继续遍历
		n := len(q)
		vals := make([]int, n) //装当前行节点值
		for i := range vals {  //只遍历当前行的所有节点，注意不是i< len(q),因为q是动态增加的
			node := q[0] //获取队首元素
			q = q[1:]    //队首元素出队
			vals[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return
}
