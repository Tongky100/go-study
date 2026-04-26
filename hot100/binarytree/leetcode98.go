package binarytree

import "math"

//根据BST的性质，以root为根的BST，其左子树节点值小于root.Val，其右子树节点值大于root.Val
//由题目范围知树不为空

//解法1 前序遍历
func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, mn int, mx int) bool {
		if root == nil {
			return true //base case 树遍历完了
		}
		// //当前节点root的值在(mn,mx)范围内视为合法
		// if root.Val >= mx || root.Val <= mn {
		//     return false
		// }也可以合并到下面

		//递归遍历左子树和右子树
		return root.Val > mn && root.Val < mx && dfs(root.Left, mn, root.Val) && dfs(root.Right, root.Val, mx)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

//解法2 中序遍历 利用BST中序遍历后是有序数组的性质
func isValidBST2(root *TreeNode) bool {
	var dfs func(*TreeNode) bool
	pre := math.MinInt
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true //base case 树遍历完了
		}

		if !dfs(root.Left) { //左
			return false
		}

		if pre >= root.Val {
			return false //中
		}

		pre = root.Val
		//右
		return dfs(root.Right)
	}
	return dfs(root)
}

//解法3 后序遍历返回每个root左子树的最大值和右子树的最小值
func isValidBST3(root *TreeNode) bool {
	var dfs func(*TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			//base case 树遍历完了
			//对立取值是因为会比较lmax是否小于root.Val 以及 root.Val是否小于rMin
			return math.MaxInt, math.MinInt
		}
		x := root.Val
		lMin, lMax := dfs(root.Left)
		rMin, rMax := dfs(root.Right)

		if x <= lMax || x >= rMin {
			return math.MinInt, math.MaxInt //也可以在递归完左子树之后立刻判断，如果发现不是二叉搜索树，就不用递归右子树了
		}
		return min(x, lMin), max(x, rMax) //使用max和min 是因为考虑了子树为空的情况
	}
	_, mx := dfs(root)
	return mx != math.MaxInt
}
