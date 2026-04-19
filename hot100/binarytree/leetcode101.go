package binarytree

//由题目数据范围可知树一定不为空
func isSymmetric(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(n1, n2 *TreeNode) bool {
		// if n1 == nil {
		//     return n2 == nil
		// }

		// if n2 == nil {
		//     return n1 == nil
		// }
		// 上面多行可以优化为一行
		if n1 == nil || n2 == nil {
			return n1 == n2
		}
		return n1.Val == n2.Val && dfs(n1.Left, n2.Right) && dfs(n1.Right, n2.Left)
	}
	return dfs(root.Left, root.Right)
}
