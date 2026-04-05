package dp

/**
划分型dp问题：定义状态为 dfs(i)，表示能否把前缀 s[:i]（表示 s[0] 到 s[i−1] 这段子串）划分成若干段，使得每段都在wordDict中
注意！！！i比实际取到的字符右边界多1，这样设计有两点好处：1 适配字符串切片包左不好有 2 方便计算使用子串的长度假设取的子串是[j,i)
长度就是i-j，和使用索引计算长度(使用的字符个数)得到的结果一样 i-1-j+1
枚举 s[:i] 最后一段的长度：
长为 1，即子串s[i−1:i]，如果它在wordDict中，那么问题变成：能否把前缀 s[:i−1] 划分成若干段，使得每段都在 wordDict 中，即 dfs(i−1)。
长为 2，即子串s[i−2:i]，如果它在wordDict中，那么问题变成：能否把前缀 s[:i−1] 划分成若干段，使得每段都在 wordDict 中，即 dfs(i−2)。
设wordDict中字符串的最长长度为 maxLen，枚举的上限不超过 maxLen，因为更长的子串必然不在 wordDict 中。

枚举 j=i−1,i−2,i−3,…,max(i−maxLen,0)，只要其中一个 j 满足s[j:i]wordDict中且dfs(j)=true，那么dfs(i)就是true。
**/

//解法1 记忆化递归
func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	//go中的bool和java中的Boolean不一样，没有null false true三种状态
	memo := make([]int8, n+1) //需要存-1 1 0 三种状态
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int8
	dfs = func(i int) (res int8) {
		if i == 0 {
			return 1
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && dfs(j) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(n) == 1

}

//解法2 1:1改写成递推
func wordBreak2(s string, wordDict []string) bool {
	maxLen := 0
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
		maxLen = max(maxLen, len(w))
	}

	n := len(s)
	//使用递推时则可以说过bool类型记录状态
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && f[j] {
				f[i] = true
				break //找到一个切法内圈循环就可以停了
			}
		}
	}
	return f[n]
}
