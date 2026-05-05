package graph

/**经典问题小岛数，每个grid[i][j] = '1'的点作为起点，计数后开始沿上下左右四个方向遍历，
遇到值为1的点将其变为0后继续沿着四个方向遍历，直至没有店可以遍历，值从1变为0的点不需要回溯
因为遍历grid[i][j]其他起点时不会再遍历到，因此不会重复计算**/
//解法1 使用dfs深度优先搜索
func numIslands(grid [][]byte) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])

	var inArea func(int, int) bool //判断点（x,y）是否在grid范围内
	inArea = func(x int, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, x int, y int) {
		if !inArea(x, y) {
			return //不在grid内直接返回
		}
		if grid[x][y] != '1' {
			return //这快岛已经被填成水了，没法走,反之重复走
		}
		grid[x][y] = '0' //填成别的不为1的数组也行和上面防止重复走要保持同步
		for _, dir := range dirs {
			dfs(grid, x+dir[0], y+dir[1])
		}
	}

	var ans int
	for i := range grid { //只写一个变量时获取的是索引
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

//解法2 使用dfs深度优先搜索 对dfs函数赋值时做优化
func numIslands2(grid [][]byte) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])

	var inArea func(int, int) bool //判断点（x,y）是否在grid范围内
	inArea = func(x int, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	var dfs func(int, int)
	dfs = func(x, y int) { //grid不用传了，本身就是闭包函数，x和y同一类型可以简化写法
		if !inArea(x, y) {
			return //不在grid内直接返回
		}
		if grid[x][y] != '1' {
			return //这快岛已经被填成水了，没法走,反之重复走
		}
		grid[x][y] = '0' //填成别的不为1的数组也行和上面防止重复走要保持同步
		for _, dir := range dirs {
			dfs(x+dir[0], y+dir[1])
		}
	}

	var ans int
	for i := range grid { //只写一个变量时获取的是索引
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

//解法3 使用dfs深度优先搜索 更轻量级的写法，不额外引入二维切片dirs和inArea函数速度变快很多
func numIslands3(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)
	dfs = func(x, y int) { //grid不用传了，本身就是闭包函数，x和y同一类型可以简化写法
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '1' {
			return //base case不在grid内或岛已经被填过了直接返回
		}

		grid[x][y] = '0' //填成别的不为1的数组也行和上面防止重复走要保持同步

		//上下左右四个方向手写
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	var ans int
	for i, row := range grid { //只写一个变量时获取的是索引
		for j, c := range row {
			if c == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
