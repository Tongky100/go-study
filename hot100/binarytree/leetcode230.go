package binarytree

//利用BST中序遍历是升序数组的性质求解
//解法1 使用额外变量记录结果
func kthSmallest(root *TreeNode, k int) int {
	var found bool
	var dfs func(*TreeNode)
	var ans int

	dfs = func(root *TreeNode) {
		if found || root == nil {
			return //base case已经找到了剪枝不再继续查找 或者 遍历完子树了
		}
		dfs(root.Left) //左
		//中
		k--
		if k == 0 {
			ans = root.Val
			found = true
		}
		dfs(root.Right) //右
	}
	dfs(root)
	return ans
}

//解法2 对解法1的优化不使用found变量记录已经找到了结果，而是使用k==0进行判断
func kthSmallest2(root *TreeNode, k int) int {
	//var found bool 可以通过k来记录当k是0时，表示已经找到
	var dfs func(*TreeNode)
	var ans int

	dfs = func(root *TreeNode) {
		if k == 0 || root == nil {
			return //base case已经找到了剪枝不再继续查找 或者 遍历完子树了
		}
		dfs(root.Left) //左
		//中
		k--
		if k == 0 {
			ans = root.Val
		}
		dfs(root.Right) //右
	}
	dfs(root)
	return ans
}

//解法3 不使用额外变量，通过dfs的返回值记录是否找到了第k个节点。
//由于题目数据范围已经表明了树中节点值非负，因此如果dfs返回的值是非负数表示已经找到，可以提前剪枝返回
//如果返回的值是-1则继续查找
func kthSmallest3(root *TreeNode, k int) int {
	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1 //base case 遍历完子树没找到，需要继续查找
		}
		l := dfs(root.Left) //左
		if l != -1 {
			return l //在子树中已经找到了第k个节点的值
		}
		//中
		k--
		if k == 0 {
			return root.Val
		}
		return dfs(root.Right) //右
	}
	return dfs(root)
}
