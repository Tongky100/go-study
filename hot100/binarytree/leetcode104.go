package binarytree

//dfs(root)表示以root为根节点的二叉树最大深度，显然dfs(root) = 1 + max(dfs(root.Left),dfs(root.Right))
//解法1 自底向上
func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + max(dfs(root.Left), dfs(root.Right))
	}
	return dfs(root)
}

//解法2 自顶向下
func maxDepth2(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		ans = max(ans, depth)
		dfs(root.Left, depth)
		dfs(root.Right, depth)
	}
	dfs(root, 0)
	return
}
