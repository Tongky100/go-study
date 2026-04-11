package dp

import (
	"math"
	"slices"
)

/*
和最大子数组和是一类问题，涉及到负数记录dp状态时要考虑负数*负数=正数
因此定义mx[i]表示以nums[i]结尾的最大子数组乘积，mn[i]表示以nums[i]结尾的最小子数组乘积
如果nums[i]>=0,mx[i] = max(mx[i-1]*nums[i],nums[i]) mn[i] = min(mn[i-1]*nums[i],nums[i])
如果nums[i]<0,mx[i] = max(mn[i-1]*nums[i],nums[i]) mn[i] = min(mx[i-1]*nums[i],nums[i])
所有mx中的最大值即为所求
*/
//解法1 递推
func maxProduct(nums []int) int {
	n := len(nums)
	mx := make([]int, n)
	mn := make([]int, n)
	var ans int = math.MinInt
	for i, v := range nums {
		if i == 0 { //v=0时，也可以写在这里特判
			mx[i], mn[i] = v, v //不支持 mx[i]=mn[i]=v这种语法
			ans = max(ans, mx[i])
			continue
		}
		if v >= 0 {
			mx[i] = max(mx[i-1]*v, v) //和之前的最大值拼在一起，或者自立门户，兼容v=0
			mn[i] = min(mn[i-1]*v, v) //和之前的最小值拼在一起，或者自立门户，兼容v=0
		} else {
			mx[i] = max(mn[i-1]*v, v) //同上
			mn[i] = min(mx[i-1]*v, v) //同上
		}
		ans = max(ans, mx[i])
	}
	return ans
}

func maxProduct2(nums []int) int {
	n := len(nums)
	mx := make([]int, n)
	mn := make([]int, n)
	var ans int = math.MinInt
	for i, v := range nums {
		if i == 0 || v == 0 { //v=0时，也可以写在这里特判
			mx[i], mn[i] = v, v //不支持 mx[i]=mn[i]=v这种语法
			ans = max(ans, mx[i])
			continue
		}
		if v > 0 {
			mx[i] = max(mx[i-1]*v, v) //和之前的最大值拼在一起，或者自立门户，兼容v=0
			mn[i] = min(mn[i-1]*v, v) //和之前的最小值拼在一起，或者自立门户，兼容v=0
		} else {
			mx[i] = max(mn[i-1]*v, v) //同上
			mn[i] = min(mx[i-1]*v, v) //同上
		}
		ans = max(ans, mx[i])
	}
	return ans
}

// 解法2 递推优化写法，为了不考虑奇偶性，我们可以每次都用nums[i]和mx[i-1]*nums[i]以及mn[i-1]*nums[i]做比较
// 因为结果只能从这3个数中算出，同时也可以不用手动计算ans，通过slices包的Max方法可直接计算，但是要额外一次遍历
// go中max min函数本身就支持传入多个值
func maxProduct3(nums []int) int {
	n := len(nums)
	mx := make([]int, n)
	mn := make([]int, n)
	mx[0], mn[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		v := nums[i]
		mx[i] = max(v, mx[i-1]*v, mn[i-1]*v)
		mn[i] = min(v, mx[i-1]*v, mn[i-1]*v)
	}
	return slices.Max(mx)
}

// 解法3 递推后的空间压缩优化，实际上只需要用到mx[i-1]和mn[i-1]这两个状态的值，因此可以用两个变量做空间优化
func maxProduct4(nums []int) int {
	mx, mn := 1, 1
	ans := math.MinInt
	for _, v := range nums {
		mx, mn = max(v, mx*v, mn*v),
			min(v, mx*v, mn*v) //同时给两个值赋值，不用引入第三变量记录mx上一个状态的旧值规避覆盖问题
		ans = max(ans, mx)
	}
	return ans
}
