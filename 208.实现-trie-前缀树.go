/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children [26]*Trie // 数组下标 0-25 隐式代表了 'a' - 'z'
	isEnd    bool      // 标记这里是否是一个单词的终点
}

// 初始化前缀树对象
func Constructor() Trie {
	return Trie{}
}

// Insert 插入单词
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		idx := ch - 'a' // 转换为下标
		if node.children[idx] == nil {
			node.children[idx] = &Trie{} // 如果没路了，就新建一个Tire
		}
		node = node.children[idx] // 顺着路走
	}
	node.isEnd = true // 标记单词结尾
}

// 查找单词是否存在
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith 查找是否有以前缀 prefix 开头的单词
func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

// 辅助方法： 沿着前缀向下走， 返回最后一个节点
func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil // 只要有一步路断了，说明前缀/单词不存在
		}
		node = node.children[idx]
	}
	return node // 返回导航停止的那个节点
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

