package dp

/**
[1]
[1,1]
[1,2,1]
[1,3,3,1]
[1,4,6,4,1]
​
杨辉三角：每个数字 = 它上一行正对位置的数字 + 它上一行正对位置左侧位置数字
且每一行首尾数字都是1，假设当前行是i(恰好是最后一个元素的索引)，那么当前行一共有i+1个元素
除了首尾元素外，需要计算的元素的索引j范围是[1,i-1]
**/
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1) //每行元素个数不等
		ans[i][0], ans[i][i] = 1, 1
		//ans[i][0] = ans[i][i] = 1 豆包说支持连续赋值语法，实际在vscode中测试不支持
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
