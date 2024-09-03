package leetcode

import "sort"

/*
	208: Implement Trie (Prefix Tree)
*/

type MyTrie struct {
	tr     [][]int
	isWord []bool
	root   int
	id     int
}

func MyTrieConstructor() MyTrie {
	trie := MyTrie{}
	return trie
}

func (this *MyTrie) Newchild() []int {
	arr := make([]int, 26)
	for i := 0; i < 26; i++ {
		arr[i] = -1
	}
	return arr
}

func (this *MyTrie) Init() {
	this.root = 0
	this.id = 1
	this.tr = [][]int{this.Newchild()}
	this.isWord = []bool{false}
}

func (this *MyTrie) NewNode() int {
	this.tr = append(this.tr, this.Newchild())
	this.isWord = append(this.isWord, false)
	newNode := this.id
	this.id++
	return newNode
}

func (this *MyTrie) Add(s string) {
	u := this.root
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if this.tr[u][c] == -1 {
			this.tr[u][c] = this.NewNode()
		}
		u = this.tr[u][c]
	}
	this.isWord[u] = true
}

func (this *MyTrie) Find(s string) int {
	u := this.root
	for i := 0; i < len(s) && u != -1; i++ {
		c := s[i] - 'a'
		u = this.tr[u][c]
	}
	return u
}

type Trie struct {
	tr MyTrie
}

func Constructor() Trie {
	tr := MyTrieConstructor()
	tr.Init()
	return Trie{
		tr: tr,
	}
}

func (this *Trie) Insert(word string) {
	this.tr.Add(word)
}

func (this *Trie) Search(word string) bool {
	node := this.tr.Find(word)
	return node != -1 && this.tr.isWord[node]
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.tr.Find(prefix)
	return node != -1
}
