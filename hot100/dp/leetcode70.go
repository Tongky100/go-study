package dp

//本质上就是斐波那契数列问题，dfs(i) = dfs(i-1) + dfs(i-2)，n=2时，由于每次只能走1步或2步，因此dfs(2)最多只能到dfs(0)
//而dfs(1)只能由dfs(0)变化因此，当i<=1时，直接返回1，也就是base case

//解法1 记忆化递归
func climbStairs(n int) int {
	memo := make([]int, n+1) //创建一个长度为n+1的切片，注意不是容量，这里不能使用数组因为n是变量
	var dfs func(int) int    //定义匿名函数dfs

	//本身使用了闭包
	//函数内部定义的函数 + 引用了外部函数的变量 = 闭包
	//匿名函数外的memo和dfs都属于外部变量
	dfs = func(i int) int { //语法上返回值可以写成(res int)能通过预设的返回值变量名更清晰的显示
		if i <= 1 {
			return 1
		}
		//这里用指针是为了提升访问效率，如果后文使用memo[i]，实际上每次都要通过 内存起始地址 + i * 元素大小 获取值
		//这样就需要多次计算memo[i]了
		p := &memo[i]
		if *p == 0 {
			//说明缓存里没有需要计算
			*p = dfs(i-1) + dfs(i-2) //通过指针修改值
		}
		return *p //包含了命中缓存的情况
	}
	return dfs(n)
}

//测试不使用指针时的速度，结果上和使用指针一样，但是明显感觉到执行的时候慢了
func climbStairs2(n int) int {
	memo := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		//不使用指针p直接获取memo中的缓存值
		p := memo[i]
		if p == 0 {
			p = dfs(i-1) + dfs(i-2)
		}
		//注意！！！go中没有return memo[i] = p这种写法
		memo[i] = p
		return memo[i]
	}
	return dfs(n)
}

//解法2 1:1改写成递推(自底向上计算)
func climbStairs3(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

//解法3 1:1改写成递推后进行空间压缩
//f[i]只依赖f[i-1]和f[i-2]前两个状态，因此可以用2个变量进行滚动更新
func climbStairs4(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		newF := f0 + f1 //这种写法还是java语法习惯，可以直接进行多个变量赋值
		//滚动更新
		f0 = f1
		f1 = newF
	}
	return f1
}

func climbStairs5(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}
aaa