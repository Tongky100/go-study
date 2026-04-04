package dp

/**
完全背包问题，定义dfs(i,j)为使用前i种硬币凑成j时所使用的硬币的最少数量
使用当前硬币，dfs(i,j) = dfs(i,j-i)+1
不使用当前硬币，dfs(i,j) = dfs(i-1,j)
二者之间的最小值即为所求，当i<0时，j=0返回0，否则返回amount+1
**/
//解法1 记忆化递归
func coinChange(coins []int, amount int) int {
	n := len(coins)
	invalid := amount + 1
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, invalid)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	//注意写成 var dfs = func(int,int,[]int) int就错了，因为这是赋值语句而不是声明一个函数变量
	//给函数赋值必须带函数体
	var dfs func(int, int, []int) int
	dfs = func(i, j int, coins []int) int {
		if i < 0 {
			if j == 0 {
				return 0
			}
			return invalid
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if coins[i] > j {
			//没法选
			*p = dfs(i-1, j, coins)
		} else {
			//可选可不选
			*p = min(dfs(i-1, j, coins), 1+dfs(i, j-coins[i], coins))
		}
		return *p
	}
	ans := dfs(n-1, amount, coins)
	if ans == invalid {
		return -1
	}
	return ans
}

//利用defer闭包完成memo赋值，整个写法会显得更漂亮点
func coinChange2(coins []int, amount int) int {
	n := len(coins)
	invalid := amount + 1
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, invalid)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	//注意写成 var dfs = func(int,int,[]int) int就错了，因为这是赋值语句而不是声明一个函数变量
	//给函数赋值必须带函数体
	var dfs func(int, int) int
	//coins和memo本质是一样的都是函数外变量，都可以不用传
	//使用带返回值变量名这种写法，必须带()号
	//在匿名函数外部，提前创建了一个名为 res 的返回值变量，专门用来接收这个函数的返回结果。
	dfs = func(i, j int) (res int) {
		if i < 0 {
			if j == 0 {
				return 0
			}
			return invalid
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }() //利用defer闭包赋值
		if coins[i] > j {
			//没法选
			return dfs(i-1, j)
		} else {
			//可选可不选
			return min(dfs(i-1, j), 1+dfs(i, j-coins[i]))
		}
	}
	ans := dfs(n-1, amount)
	if ans == invalid {
		return -1
	}
	return ans
}

//解法2 1:1改写成递归
func coinChange3(coins []int, amount int) int {
	n := len(coins)
	invalid := amount + 1
	f := make([][]int, n+1) //对应dfs(i<0)时，多开1位
	for i := range f {
		f[i] = make([]int, invalid)
	}
	for j := range f[0] {
		f[0][j] = invalid //对应dfs(-1,j>0)
	}
	f[0][0] = 0 //对应dfs(0,0)
	for i, x := range coins {
		for c := 0; c <= amount; c++ {
			if c < x {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], 1+f[i+1][c-x])
			}
		}
	}
	ans := f[n][amount]
	if ans == invalid {
		return -1
	}
	return ans
}

//解法3 1:1改写成递归后进行空间压缩 f[i+1][c] 只和 上一行f[i][c] 以及 左边f[i=1][c-x]的状态有关
func coinChange4(coins []int, amount int) int {
	invalid := amount + 1
	f := make([]int, invalid)
	for i := range f {
		f[i] = invalid //初始化时对应dfs(-1,j>0)
	}
	f[0] = 0                  //对应dfs(0,0)
	for _, x := range coins { //注意想跳过索引不用必须这么写，写成for x := range coins的话，x是索引，不是coins[i]的值
		//完全背包内圈循环遍历方向是正方向
		for c := x; c <= amount; c++ {
			f[c] = min(f[c], 1+f[c-x])
		}
	}
	ans := f[amount]
	if ans == invalid {
		return -1
	}
	return ans
}
