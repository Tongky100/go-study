package binarytree

//经典题目LCA 在root的左、右子树中分别查找p和q，找到了就返回
//如果root的左、右子树返回结果不为空，那么root就是LCA
//解法1 使用二叉树后序遍历求解
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root //base case 树遍历完或找到p或q了
	}
	//递归在左右子树中查找
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l == nil {
		return r
	} else {
		return l
	}
}
