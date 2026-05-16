package graph

//解法1 不封装方法求解 直接实现
type Trie struct {
	root *Node
}

//从题目数据范围可知树中节点值只有小写字母组成，因此只需要开26位的数组
type Node struct {
	isWord   bool
	children [26]*Node
}

func Constructor() Trie {
	return Trie{&Node{}}
}

func (this *Trie) Insert(word string) {
	cur := this.root //获取树的根节点
	for _, c := range word {
		//遍历word中的每一个字符
		c -= 'a' //计算c在前缀树当前层对应的索引
		if cur.children[c] == nil {
			cur.children[c] = &Node{}
		}
		cur = cur.children[c]
	}
	cur.isWord = true //word最后一个字符表明创建完了一个单词
}

func (this *Trie) Search(word string) bool {
	cur := this.root //获取树的根节点
	for _, c := range word {
		//遍历word中的每一个字符
		c -= 'a' //计算c在前缀树当前层对应的索引
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root //获取树的根节点
	for _, c := range prefix {
		//遍历word中的每一个字符
		c -= 'a' //计算c在前缀树当前层对应的索引
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
	}
	return true //推出循环说明前缀一定存在
}

//解法2中变更的内容
//发现word是单词时返回2 发现word是前缀时返回1
func (this *Trie) find(word string) int {
	cur := this.root //获取树的根节点
	for _, c := range word {
		//遍历word中的每一个字符
		c -= 'a' //计算c在前缀树当前层对应的索引
		if cur.children[c] == nil {
			return 0
		}
		cur = cur.children[c]
	}
	if cur.isWord {
		return 2
	}
	return 1
}

func (this *Trie) Search(word string) bool {
	return this.find(word) == 2
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) >= 1 //单词一定是自己的前缀
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
