package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//解法1 后续遍历求解 flatten函数定义就是把以root为根的二叉树链表化
func flatten(root *TreeNode) {
	if root == nil {
		return //base case 树为空 或 已遍历完
	}

	flatten(root.Left)
	flatten(root.Right)

	//左子树和右子树整合到一起
	right := root.Right //记录右子树
	root.Right = root.Left
	root.Left = nil //左子树设置为空
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
	return
}

//解法2 后序遍历求解 借助额外指针头插法
func flatten2(root *TreeNode) {
	var head *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		//先遍历右子树这样每次都能保证左子树节点会插在右子树节点前，所以是头插
		dfs(node.Right)
		dfs(node.Left)
		node.Left = nil
		node.Right = head //头插法在链表中插入节点
		head = node
	}
	dfs(root)
}
