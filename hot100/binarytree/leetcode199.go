package binarytree

//解法1 使用递归求解 优先遍历右子树，当len(ans)==当前遍历层数时表示需要将右子树节点加入到结果集中
func rightSideView(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return //base case 树遍历完了
		}
		if len(ans) == level {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(root, 0)
	return //这里写return ans也是一样的，已经定义了返回值参数名为ans
}

//解法2 使用bfs求解 只需要把每层最后一个节点值加入到结果集中即可
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		ans = append(ans, cur[len(cur)-1].Val)
		nxt := make([]*TreeNode, 0) //第二种创建切片的写法
		for _, node := range cur {
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
	}
	return
}
