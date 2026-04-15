package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
go中树的结构体使用指针的解释：
1. 结构体不能嵌套自己本身（会无限递归）
如果写成：
type TreeNode struct {
  Left TreeNode // 错误！
}
TreeNode包含Left TreeNode → 这个 Left 又包含 Left → 无限套娃
2. 指针才能表示「空节点」
二叉树叶子节点的 Left/Right 是 nil只有指针能是 nil，结构体值类型不能为 nil
3. 节省内存 + 共享结构
指针只存地址（8 字节）如果存结构体本身，会复制整个树，内存爆炸
java中虽然没有指针的定义，但是对象的引用本质上就是指针
下面定义的方法中如果入参是结构体而非指针，用的是值拷贝，那么会拷贝整棵树，如果题目需要对原树进行节点变更
传入结构体无法实现，因为变更都落在复制出来的树上
*/
// 解法1 使用递归求解 为了方便少写变量 把题干加了个返回值参数
func inorderTraversal(root *TreeNode) (res []int) {
	var dfs func(*TreeNode)      //先声明变量（此时是nil，但有地址）
	dfs = func(root *TreeNode) { //再把匿名函数赋值给它
		if root == nil {
			return //空树或者遍历完树了
		}
		//闭包调用dfs
		dfs(root.Left) //语法糖不用写成 (*root).Left
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return
}

// 解法2 使用栈模拟递归过程
func inorderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{} //这里没有使用make创建切片因为不知道节点个数
	//树不为空且栈中还有节点时就需要继续遍历
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root) //先把当前节点压入栈
			root = root.Left            //然后继续遍历左子树
		} else {
			//左子树为空了
			root = stack[len(stack)-1]   //获取栈顶元素
			stack = stack[:len(stack)-1] //栈顶元素出栈
			res = append(res, root.Val)
			root = root.Right //继续遍历右子树
		}
	}
	return
}
