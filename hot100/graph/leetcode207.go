package graph

//题目的本质是图中找环

//解法1 使用dfs找环
/***比如存在下述三种图 3->2->1  4->5->3  6->7
                                           7->6
  第一条路径标记了3->2->1 不会成环
  第二条路径遍历到3时 因为标记过3不会成环因此执行会很快
  第三条路径有环
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	vis := make([]bool, numCourses)  //全局判断哪些点已经走过
	path := make([]bool, numCourses) //判断一条路径上哪些点已经走过 用于找环
	//构造图
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		//根据题目描述 p[1]是前置课
		g[p[1]] = append(g[p[1]], p[0])
	}

	var dfs func(int)
	var hasCycle bool
	dfs = func(cur int) {
		if path[cur] {
			//cur在当前路径走过了成环了
			hasCycle = true
			return
		}
		if vis[cur] {
			return //cur这个点在所有路径中已经走过了
		}
		vis[cur], path[cur] = true, true //go不支持链式赋值，豆包经典胡说
		for _, nei := range g[cur] {
			dfs(nei)
		}
		path[cur] = false //回溯
	}
	for i := range numCourses {
		if !vis[i] {
			dfs(i)
		}
	}
	return !hasCycle
}

//解法2 使用bfs图的拓扑排序 先计算所有点的入度，a节点指向b节点，那么b节点的入度为1
//显然所有入度为0的节点就是课程起点
func canFinish2(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	g := make([][]int, numCourses)
	//构造图时，同时计算每个节点的入度
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
		indegree[p[0]]++
	}
	cur := []int{}
	for i, v := range indegree {
		if v == 0 {
			cur = append(cur, i)
		}
	}
	numCourses -= len(cur) //如果能拓扑排序最后所有课都被上完了
	for len(cur) > 0 {
		nxt := []int{}
		for _, c := range cur {
			for _, nei := range g[c] {
				indegree[nei]--
				if indegree[nei] == 0 {
					//又一门课被上完了
					numCourses--
					nxt = append(nxt, nei)
				}
			}
		}
		cur = nxt
	}
	return numCourses == 0
}

//解法3 使用三色染色法求解 类似jvm垃圾回收算法的染色思路
/**
对于每个节点x，都定义三种颜色值（状态值）：
0：节点x尚未被访问到。
1：节点x正在访问中，dfs(x) 尚未结束。
2：节点x已经完全访问完毕。注意这还说明从 x出发无法找到环。所以当我们遇到状态值为2的节点x时，无需递归 x。

假设图为：
 1  ->  2
 ↓
 0

⚠误区：不能只用两种状态表示节点「没有访问过」和「访问过」。例如上图，我们先 dfs(0)，再 dfs(1)，此时1的邻居0已经访问过，但这并不能表示此时就找到了环。
**/
func canFinish3(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}

	colors := make([]int, numCourses)
	//返回true表示找到了环
	var dfs func(int) bool
	dfs = func(x int) bool {
		colors[x] = 1 //标记x正在访问中
		for _, nei := range g[x] {
			//colors[nei] == 1 表示找到了环
			//colors[nei] == 0 表示没访问过nei，继续递归访问nei
			//colors[nei] == 2 重复访问nei只会重蹈覆辙，跳过
			if colors[nei] == 1 || colors[nei] == 0 && dfs(nei) {
				return true //找到了环
			}
		}
		colors[x] = 2 //x完全访问完毕，从x出发未找到环
		return false
	}

	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false //有环
		}
	}
	return true //无环
}
