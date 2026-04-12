package dp

//子集型01背包问题
/**
s表示nums中所有元素的和，如果s是奇数则无法划分成两个元素和相等的子集
否则问题转换成使用nums的前n个元素能否恰好凑出s/2，每个元素至多选择一次，是标准的01背包问题
dfs(i,x)表示使用前i个元素能否组成x
如果nums[i]<x,dfs(i,x) = dfs(i-1,x) 只能不选nums[i]
如果nums[i]>=x,dfs(i,x) = dfs(i-1,x) || dfs(i-1,x-nums[i]) 可以选nums[i]也可以不选
*/
//解法1 记忆化递归
func canPartition(nums []int) bool {
	//s := slices.Sum(nums)没有这个方法
	var s int
	for _, v := range nums {
		s += v
	}
	if s&1 == 1 {
		return false
	}
	s /= 2
	n := len(nums)
	memo := make([][]int8, n) //用-1表示没有计算过
	for i := range memo {
		memo[i] = make([]int8, s+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) (res bool) {
		if i < 0 {
			if j == 0 {
				return true
			}
			return false
		}
		p := &memo[i][j]
		defer func() {
			if res {
				*p = 1
			} else {
				*p = 0
			}
		}()
		if *p != -1 {
			return *p == 1 //表示能恰好凑成
		}
		if nums[i] > j {
			return dfs(i-1, j)
		}
		return dfs(i-1, j) || dfs(i-1, j-nums[i])
	}
	return dfs(n-1, s)
}

//解法2 1:1改写成递推
func canPartition2(nums []int) bool {
	//s := slices.Sum(nums)没有这个方法
	var s int
	for _, v := range nums {
		s += v
	}
	if s&1 == 1 {
		return false
	}
	s /= 2
	n := len(nums)
	f := make([][]bool, n+1) //多开一位，对应dfs(i<0)
	for i := range f {
		f[i] = make([]bool, s+1)
	}
	f[0][0] = true
	for i, x := range nums {
		for y := range s + 1 {
			f[i+1][y] = f[i][y] || y >= x && f[i][y-x]
		}
	}
	return f[n][s]
}

//解法3 1:1改写成递推后进行空间压缩
//f[i+1][y] 只和f[i][y] 和 f[i][y-x]有关
//为了保证第i行的状态不被覆盖 01背包内圈for循环的遍历倒序
func canPartition3(nums []int) bool {
	//s := slices.Sum(nums)没有这个方法
	var s int
	for _, v := range nums {
		s += v
	}
	if s&1 == 1 {
		return false
	}
	s /= 2
	f := make([]bool, s+1) //多开一位，对应dfs(i<0)
	f[0] = true
	for _, x := range nums {
		for y := s; y >= x; y-- {
			// |=左侧是不能选当前x包含了可以选但是不选 右侧是选当前x
			f[y] = f[y] || f[y-x] //bool类型在go中不支持|=写法，int类型可以
		}
	}
	return f[s]
}

//解法4 在解法3写法上更进一步优化 使用额外变量s2计算nums的前缀和，截止到元素i，能凑成的最大和一定是min(s2,s)
//当f[s]时提前返回true做剪枝
func canPartition4(nums []int) bool {
	//s := slices.Sum(nums)没有这个方法
	var s int
	for _, v := range nums {
		s += v
	}
	if s&1 == 1 {
		return false
	}
	s /= 2
	f := make([]bool, s+1) //多开一位，对应dfs(i<0)
	f[0] = true
	s2 := 0
	for _, x := range nums {
		s2 = min(s2+x, s)
		for y := s2; y >= x; y-- {
			// |=左侧是不能选当前x包含了可以选但是不选 右侧是选当前x
			f[y] = f[y] || f[y-x] //bool类型在go中不支持|=写法，int类型可以
			if f[s] {
				return true
			}
		}
	}
	return false
}
