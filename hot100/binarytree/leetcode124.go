package binarytree

import "math"

/**
拿到题的第一想法是遍历全树，记录每棵子树的节点值之和，然后二次遍历整棵树。到每个节点时有三种切割选择，左、右、负，利用减法可以计算这三部分的值
以上思路的问题在于：一棵子树中的所有节点并不一定都要参与计算。而且遍历的次数越多算法越慢

先明确两个核心概念：
链：从下面的某个节点（不一定是叶子）到当前节点的路径。把这条链的节点值之和，作为dfs的返回值。如果节点值之和是负数，则返回 0（和 0 取最大值）。这个思想和53最大子数组和是一样的，如果左侧子数组的元素和是负数，就不和当前元素拼起来。

直径：等价于由两条（或者一条）链拼成的路径。我们枚举每个node，假设直径在这里「拐弯」，也就是计算由左右两条从下面的某个节点（不一定是叶子，假设当前节点是root，左、右子树的返回值是0时相当于没有取子树的节点。也就是“不一定是叶子节点”）到 node 的链的节点值之和，去更新答案的最大值。

如果所有节点值都是负数，ans = max(ans, l_val + r_val + node.val)等价于ans = max(ans, node.val),符合预期，因为负数+负数会得到更小的负数 **/
// 解法1 使用二叉树后序遍历求解
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(*TreeNode) int //dfs返回以root为根的最大链长
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0 //base case 空树/树遍历完了
		}

		lmax := dfs(root.Left)
		rmax := dfs(root.Right)
		ans = max(ans, root.Val+lmax+rmax)      //这里可以保证lmax和rmax都是大于等于0的
		return max(max(lmax, rmax)+root.Val, 0) //比0小就不需要子树了
	}
	dfs(root)
	return ans
}
