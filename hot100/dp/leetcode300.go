package dp

import (
	"slices"
	"sort"
)

/**
LIS经典问题，dfs(i)表示以nums[i]结尾的最长子序列长度
枚举[0,i-1]中的每个索引j，如果nums[j]<nums[i],那么dfs(i) = 1+dfs(j)
所有可能的dfs(i)中的最大值即为所求
*/

// 解法1 记忆化递归
func lengthOfLIS(nums []int) int {
	n := len(nums)
	//0可以作为缓存值，因为最大值最小都是1
	memo := make([]int, n)
	var dfs func(int) int
	//每次调用 dfs 函数时，默认初始值都是 0
	//可以在函数体内通过res = 1对res重新初始化值
	//go 1.20+以上的版本 可以在对dfs赋值时直接声明res的值dfs = func(i int) (res int = 1)
	dfs = func(i int) (res int) {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != 0 {
			return *p
		}
		defer func() { *p = res }()
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				//里面不能写成1+dfs(j),因为我们初始化的返回值变量res是0，如果nums[i]比他前面的每个nums[j]都小
				//那么res是0，不会被更新，因为+1操作应该放在外层
				res = max(res, dfs(j))
			}
		}
		return res + 1
	}
	var ans int
	//不确定以哪个nums[i]结尾的LIS最大，因此还要遍历一次
	for i := range n { //遍历0到n-1
		ans = max(ans, dfs(i))
	}
	return ans
}

// 对解法1进行代码优化，显得更简洁，外层函数可加返回值变量减少一次内部变量的声明
func lengthOfLIS2(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		// if i < 0 {
		//     return 0
		// } 可以省略因为下面for循环遍历索引的下限就是1
		p := &memo[i]
		if *p != 0 {
			return *p
		}
		defer func() { *p = res }()
		for j, x := range nums[:i] { //利用数组的切片遍历
			if x < nums[i] {
				res = max(res, dfs(j))
			}
		}
		return res + 1
	}
	for i := range n {
		ans = max(ans, dfs(i))
	}
	return ans
}

// 解法2 1:1改写成递推
func lengthOfLIS3(nums []int) (ans int) {
	n := len(nums)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i]++
		ans = max(ans, f[i])
	}
	return ans
}

// 对递推解法再次进行优化
func lengthOfLIS4(nums []int) (ans int) {
	n := len(nums)
	f := make([]int, n)
	for i, x := range nums {
		for j, y := range nums[:i] {
			if y < x {
				f[i] = max(f[i], f[j])
			}
		}
		f[i]++
		ans = max(ans, f[i])
	}
	return ans
}

// 对递推解法再次进行优化
func lengthOfLIS5(nums []int) int {
	f := make([]int, len(nums))
	for i, x := range nums {
		for j, y := range nums[:i] {
			if y < x {
				f[i] = max(f[i], f[j])
			}
		}
		f[i]++
	}
	return slices.Max(f)
}

// 解法3 贪心+二分 改写的java中的写法
func lengthOfLIS6(nums []int) int {
	piles := 0
	for _, x := range nums {
		idx := helper(nums, x, piles)
		if idx == piles {
			piles++
		}
		nums[idx] = x
	}
	return piles
}

func helper(nums []int, target int, r int) int {
	l := -1
	//go中没有while循环，得用for循环改写
	for l+1 < r {
		//加法运算和位移运算同级因此要加括号
		mid := l + ((r - l) >> 1)
		if nums[mid] >= target {
			r = mid //右指针左移继续查找比target的大的数(比第一次查找的数要小)
		} else {
			l = mid //找的数小了，左指针右移找更大的数
		}
	}
	return r
}

// 解法3 贪心+二分 使用golang中的内置函数 自己写的二分会比内置函数的二分快一点
// 内置函数 二分源码用的左闭右开
func lengthOfLIS7(nums []int) int {
	g := []int{} //全新空切片（独立、安全）
	//上面初始化写成 g := make([]int,0)也是可以的
	for _, x := range nums {
		//内置函数在g中查找>=x的最小索引
		j := sort.SearchInts(g, x)
		if j == len(g) { //>=x 的 g[j] 不存在
			//这里操作的是切片，不是数组，给切片中越界的索引位置赋值也会报错
			//append函数中第一个参数是切片，第二个参数是往切片中添加的元素，是个可变数组
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}

// 解法3 贪心+二分 使用golang中的内置函数 但是g和nums共享内存(g是基于nums的空切片)
func lengthOfLIS8(nums []int) int {
	g := nums[:0] //基于nums的空切片（共享底层数组）
	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) { //>=x 的 g[j] 不存在
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}
