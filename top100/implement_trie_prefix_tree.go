package top100


// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// 前缀树，wiki： https://zh.wikipedia.org/wiki/Trie

// 一开始next用map，后来发现有限状态机优势就是状态有限，可以用长度为26的数组表示
// 查询效率会高很多
// 这种数据结构前缀可以有多种形式，包括数组/hash/链表，各有优缺点
// TODO: 另外还有优化的实现，双数组Trie，三数组Trie
type node struct {
	val byte
	next [27]*node
}

type Trie struct {
	node *node
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		node: &node{
			next: [27]*node{},
		},
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	tmp := this.node
	for i := range word {
		if tmp.next[word[i]-'a'] == nil {
			tmp.next[word[i]-'a'] = &node{
				val: word[i],
				next: [27]*node{},
			}
		}
		tmp = tmp.next[word[i]-'a']
	}
	tmp.next[26] = &node{}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	tmp := this.node
	for i := range word {
		if tmp.next[word[i]-'a'] == nil {
			return false
		}
		tmp = tmp.next[word[i]-'a']
	}
	return tmp.next[26] != nil
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	tmp := this.node
	for i := range prefix {
		if tmp.next[prefix[i]-'a'] == nil {
			return false
		}
		tmp = tmp.next[prefix[i]-'a']
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */