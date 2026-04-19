package binarytree

//解法1 自底向上翻转
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root //空树或已经遍历完子树
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	//翻转
	root.Right = left
	root.Left = right
	return root
}

//解法2 自上向下翻转
func invertTree2(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		left := root.Left
		root.Left = root.Right
		root.Right = left
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return root
}
