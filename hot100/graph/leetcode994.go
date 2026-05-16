package graph

//本题求最小扩散分钟，且腐烂橘子的起点不止一个。更适合使用BFS求解
//解法1 使用bfs广度优先搜索求解
func orangesRotting(grid [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	cur := [][]int{} //存当前腐烂的橙子有哪些
	var fresh int    //计算一共有多少个新鲜橘子
	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				fresh++
			}
			if col == 2 {
				cur = append(cur, []int{i, j})
			}
		}
	}
	var ans int
	//腐烂的橘子队列不为空时，继续做bfs
	for fresh > 0 && len(cur) > 0 {
		nxt := [][]int{}
		var x, y int
		for _, v := range cur {
			//遍历每个腐烂的橘子
			for _, dir := range dirs {
				x, y = v[0]+dir[0], v[1]+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					fresh--
					grid[x][y] = 2 //防止橘子重复入队
					nxt = append(nxt, []int{x, y})
				}
			}
		}
		ans++
		cur = nxt
	}
	if fresh > 0 {
		return -1 //有橘子没有被感染
	}
	return ans
}
