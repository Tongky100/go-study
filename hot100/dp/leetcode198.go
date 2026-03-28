package dp

/**dfs(i)表示偷前i+1([0,i])个房子能获得的最高金额
   因此有两种选：
   1、偷当前房子，dfs(i) = num[i] + dfs(i-2)
   2、不偷当前房子。dfs(i) = dfs(i-1)
   二者之中的最大值即为所求
**/
//解法1 记忆化递归
func rob(nums []int) int {
	//memo := [len(nums)]int 也没法通过数组来创建memo，因为len(nums)无法再编译期获得结果，因此还是变量
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0 //base case
		}

		p := &memo[i]
		if *p != -1 {
			return *p //命中缓存
		}

		res := max(dfs(i-1), nums[i]+dfs(i-2))
		*p = res
		return *p
	}
	return dfs(n - 1)
}

//解法2 1:1改写成递推 dfs中i=0时，i-2是-2因此需要多补两位，即nums中的索引0对应的是f中的2，nums中的索引n-1，对应的是f中的n+1
//以达到往右错2位的效果
func rob2(nums []int) int {
	n := len(nums)
	f := make([]int, n+2)
	for i, x := range nums {
		f[i+2] = max(f[i+1], x+f[i])
	}
	return f[n+1]
}

//解法3 1:1改写成递推后进行空间压缩，因为f[i]只和f[i-1]和f[i-2]有关
func rob3(nums []int) int {
	f0, f1 := 0, 0
	for _, x := range nums {
		//不用担心f0的值被f1覆盖取不到旧值(等号右侧max计算时使用的是f0的旧值)
		//go的规定：等号右边所有值，全部先算完，再统一赋值给左边
		f0, f1 = f1, max(f1, f0+x)
	}
	return f1
}
