package binarytree

/**解法1 使用二叉树后序遍历求解 统计以root为根的所有可能路径和然后和targetSum进行计数比较*/
func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(root *TreeNode) []int
	var ans int
	dfs = func(root *TreeNode) (res []int) {
		if root == nil {
			return //返回空切片res
		}

		l := dfs(root.Left)
		r := dfs(root.Right)

		for _, v := range l {
			v1 := v + root.Val
			if v1 == targetSum {
				ans++
			}
			res = append(res, v1)
		}

		for _, v := range r {
			v1 := v + root.Val
			if v1 == targetSum {
				ans++
			}
			res = append(res, v1)
		}

		if root.Val == targetSum {
			ans++
		}
		res = append(res, root.Val) //根节点自己也是路径
		return                      //这里不写return语句也会返回一个切片res
	}
	dfs(root)
	return ans
}

/**解法2 使用前缀和进行优化求解，以示例1为例10->5->3->3只看这条路径，是单链表完全可以使用前缀和求解
只需记录路径上已出现sum的次数用于后序计算*/
func pathSum2(root *TreeNode, targetSum int) int {
	cnt := map[int]int{0: 1} //初始化空树的值用于后续计算
	var ans int
	var dfs func(root *TreeNode, s int)
	dfs = func(root *TreeNode, s int) {
		if root == nil {
			return //base case 树遍历完了
		}
		s += root.Val
		/**
		  如果targetSum不为0，先更新cnt还是ans哪个都行，如果targetSum为0，只能先更新ans
		  假设tagetSum为0，现在s=1，先更新cnt[1]=1，那么计算cnt[s-targetSum]时发现,cnt[1]存在，ans中多加了1，显然不对
		  s是其他值时，也同样会ans+1,每个节点单独组成全路径时，都会错误的是ans+=1，显然不对，因此要先更新ans
		  **/
		if v, ok := cnt[s-targetSum]; ok {
			ans += v
		}
		if v, ok := cnt[s]; ok {
			cnt[s] = v + 1
		} else {
			cnt[s] = 1
		}
		dfs(root.Left, s)
		dfs(root.Right, s)
		//回溯，左子树已对前缀和做的计数不能影响右子树，这个方法本身就是数组中套用过来的，只能线性的计算不能分叉
		cnt[s] = cnt[s] - 1
	}
	dfs(root, 0)
	return ans
}

/**解法3 对解法2做优化
从go中的map获取一个不存在的key时，不会报错、不会返回nil，而是直接返回该值类型的「零值 / 默认值」
比如map[int]int，获取一个不存在的key时，结果是0*/
func pathSum3(root *TreeNode, targetSum int) int {
	cnt := map[int]int{0: 1} //初始化空树的值用于后续计算
	var ans int
	var dfs func(root *TreeNode, s int)
	dfs = func(root *TreeNode, s int) {
		if root == nil {
			return //base case 树遍历完了
		}
		s += root.Val
		/**
		  如果targetSum不为0，先更新cnt还是ans哪个都行，如果targetSum为0，只能先更新ans
		  假设tagetSum为0，现在s=1，先更新cnt[1]=1，那么计算cnt[s-targetSum]时发现,cnt[1]存在，ans中多加了1，显然不对
		  s是其他值时，也同样会ans+1,每个节点单独组成全路径时，都会错误的是ans+=1，显然不对，因此要先更新ans
		  **/
		ans += cnt[s-targetSum]
		cnt[s]++
		dfs(root.Left, s)
		dfs(root.Right, s)
		//回溯，左子树已对前缀和做的计数不能影响右子树，这个方法本身就是数组中套用过来的，只能线性的计算不能分叉
		cnt[s]--
	}
	dfs(root, 0)
	return ans
}
