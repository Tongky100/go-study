package dp

import "math"

/**恰好型求最小数量完全背包问题
背包容量是n，物品种类是1,2,3,4...根号下n*/

//解法1 记忆化递归
//题目范围n最大是10000，因此物品容量最值是100,写在外面，多个测试数据之间可以共享，减少计算量
var memo [101][10001]int

//go的特性init函数会在mian函数之前执行，可以用来做初始化
func init() {
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
}

func numSquares(n int) int {
	var dfs func(int, int) int //定义dfs是一个多参数的函数

	dfs = func(i, j int) int {
		//go中没有三元运算符
		if i == 0 {
			if j == 0 {
				return 0
			}
			return math.MaxInt / 2 //这里写成math.MaxInt也可以因为一个数一定可以n个1相加，也就是一定能拆分成完全平方数求和
		}

		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if j < i*i {
			//只能不选
			*p = dfs(i-1, j)
		} else {
			//可以选或不选
			*p = min(dfs(i-1, j), 1+dfs(i, j-i*i))
		}
		return *p
	}

	return dfs(int(math.Sqrt(float64(n))), n)
}

func numSquares2(n int) int {
	//go是一门强类型的语言，如果定义dfs使用了可变参数
	//那么使用匿名函数对dfs赋值时也得使用可变参数
	//但是换成可变参数，dfs又没有定义变量i和j无法在匿名函数中使用
	//因此最佳写法就是在定义func时显式设定两个参数
	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		//go中没有三元运算符
		if i == 0 {
			if j == 0 {
				return 0
			}
			return math.MaxInt / 2 //这里写成math.MaxInt也可以因为一个数一定可以n个1相加，也就是一定能拆分成完全平方数求和
		}

		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if j < i*i {
			//只能不选
			*p = dfs(i-1, j)
		} else {
			//可以选或不选
			*p = min(dfs(i-1, j), 1+dfs(i, j-i*i))
		}
		return *p
	}

	return dfs(int(math.Sqrt(float64(n))), n)
}

//解法2 1:1改写成递推
const N = 10000

var f [101][N + 1]int

//多个init方法都会在main方法之前执行，init方法按照代码的先后顺序执行
func init() {
	for i := 1; i <= N; i++ {
		//对应dfs(i,j>0) = max，dfs(i,j=0)=0
		//f[.][0]默认值就是0，不用刻意初始化
		f[0][i] = math.MaxInt
	}
	for i := 1; i*i <= N; i++ {
		// 这里j从0或1开始遍历是一样的，因为j=0时f[.][0]已经初始化为0了，从逻辑上也无法选完全平方数用于求和
		for j := 1; j <= N; j++ {
			if j < i*i {
				//只能不选
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = min(f[i-1][j], 1+f[i][j-i*i])
			}
		}
	}
}

func numSquares3(n int) int {
	return f[int(math.Sqrt(float64(n)))][n]
}

//解法3 1:1改写成递推后进行空间压缩 f[i][j]只和其左侧还有上行的状态有关，内循环遍历顺序正序，区别于01背包
var f1 [N + 1]int

func init() {
	for i := 1; i <= N; i++ {
		f1[i] = math.MaxInt
	}
	//f[0]是0不用刻意初始化
	for i := 1; i*i <= N; i++ {
		//注意j从i^2开始遍历，这样f1[j]上一个状态的含义就是不选当前i(包含了能选不选)
		for j := i * i; j <= N; j++ {
			f1[j] = min(f1[j], 1+f1[j-i*i]) //不选时,用的就是i-1时的f[j]
		}
	}
}

func numSquares4(n int) int {
	return f1[n]
}
