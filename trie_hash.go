package main

import "fmt"

type Trie struct {
	IsWord   bool
	children map[int]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var pNode Trie
	pNode.children = make(map[int]*Trie)
	return pNode
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, w := range word {
		if this.children[int(w-'a')] == nil {
			tmp := Constructor()
			this.children[int(w-'a')] = &tmp
		}
		this = this.children[int(w-'a')]
	}
	this.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, w := range word {
		if this.children[int(w-'a')] != nil {
			this = this.children[int(w-'a')]
		} else {
			return false
		}
	}
	return this.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, w := range prefix {
		if this.children[int(w-'a')] != nil {
			this = this.children[int(w-'a')]
		} else {
			return false
		}
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
func main() {
	obj := Constructor()
	obj.Insert("abcabc...fdsfas")
	obj.Insert("abcfdsafds")
	obj.Insert("abgfsd")
	obj.Insert("cabgfsd")
	fmt.Println(obj.Search("abgfsd"))
	fmt.Println(obj.Search("abcabc...fdsfas"))
	fmt.Println(obj.Search("abcabcab"))
	fmt.Println(obj.StartsWith("ab"))
	fmt.Println(obj.StartsWith("ca"))
}
