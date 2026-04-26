package binarytree

//平衡BST = 以root为根的BST，左子树和右子树的树高差为1，由于nums已经是有序数组
//因此可以通过中点切割，根据题目数据范围已知BST不为空

//解法1 使用递归闭包求解
func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(int, int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil //base case 没有节点可以构造了
		}
		mid := left + ((right - left) >> 1)
		root := &TreeNode{Val: nums[mid]}
		leftNode := dfs(left, mid-1)
		rightNode := dfs(mid+1, right)
		root.Left = leftNode
		root.Right = rightNode
		return root
	}
	return dfs(0, len(nums)-1)
}

//解法2 灵神轻量级写法
func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil //空树 或 已构造完子树
	}
	m := len(nums) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}
