package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//二叉树先序遍历的顺序是根 左 右 中序遍历的顺序是左 根 右
//假设当前子树对应的先序、中序区间分别是[preStart,preEnd] [inStart,inEnd]
//那么root.Val = preorder[preStart] 同时我们可以在inorder中找到root对应的索引idx
//那么[inStart,idx-1] [idx+1,inEnd]分别是以root为根的中序遍历左、右子树对应的节点范围
//通过中序遍历的分割可以算出左子树节点个数即 idx - 1 - inStart + 1 = idx - inStart
//那么先序遍历的左子树区间为[preStart+1,preStart+idx-inStart] 右子树区间为[preStart+idx-inStart+1,preEnd]
//至此，问题可以划分为更小的子问题，适合使用递归求解，当先序、中序任意子树对应的区间出现索引越界时，则表示树已遍历完
//中序中节点值和下标索引的关系可以提前存储到一个map中
func buildTree(preorder []int, inorder []int) *TreeNode {
	//memo := map[int]int{}
	memo := make(map[int]int)
	for idx, val := range inorder {
		memo[val] = idx
	}

	var dfs func(preStart int, preEnd int, inStart int, inEnd int) *TreeNode
	dfs = func(preStart int, preEnd int, inStart int, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil //base case 树遍历完了
		}
		rootVal := preorder[preStart]
		idx := memo[rootVal]
		leftSize := idx - inStart
		return &TreeNode{
			Val:   rootVal,
			Left:  dfs(preStart+1, preStart+leftSize, inStart, idx-1),
			Right: dfs(preStart+leftSize+1, preEnd, idx+1, inEnd),
		}
		/**
		  或者写成
		  root := &TreeNode{
		      Val : rootVal,
		      Left : dfs(preStart+1, preStart+leftSize, inStart, idx-1),
		      Right : dfs(preStart+leftSize+1, preEnd, idx+1, inEnd),
		  }
		  return root
		  不能写成，语法中直接返回创建的对象即可
		  return root := &TreeNode{
		      Val : rootVal,
		      Left : dfs(preStart+1, preStart+leftSize, inStart, idx-1),
		      Right : dfs(preStart+leftSize+1, preEnd, idx+1, inEnd),
		  }
		  针对变量赋值返回是允许的，假设a是int，可以写成return a=1 即对a赋值后返回a，类比java理解
		  **/
	}
	return dfs(0, len(preorder)-1, 0, len(preorder)-1)
}
