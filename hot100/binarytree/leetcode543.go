package binarytree

//二叉树的直径 = max(左子树的最大直径，右子树的最大执行，以及过根节点的最大直径)
//其中过根节点的最大直径 可以看做是当前节点为根的左子树节点个数 + 右子树节点个数
//或者看做是2 + 左子树最大链长 + 右子树最大链长 用链长的表示形式更加合理 因此使用这种方式
//当root = nil时，返回-1，这样当到达叶子节点时，计算出来的左右子树最大链长是0
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int //返回值当前子树的最大链长
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		lmax := dfs(root.Left)
		rmax := dfs(root.Right)
		ans = max(ans, 2+lmax+rmax)
		return max(lmax, rmax) + 1
	}
	dfs(root)
	return
}
