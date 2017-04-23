package main

import "fmt"

const ALPHABET_SIZE = 26

type Trie struct {
	IsWord   bool
	children []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var pNode Trie
	pNode.children = make([]*Trie, 26)
	return pNode
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, w := range word {
		if this.children[w-'a'] == nil {
			tmp := Constructor()
			this.children[w-'a'] = &tmp
		}
		this = this.children[w-'a']
	}
	this.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, w := range word {
		if this.children[w-'a'] != nil {
			this = this.children[w-'a']
		} else {
			return false
		}
	}
	return this.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, w := range prefix {
		if this.children[w-'a'] != nil {
			this = this.children[w-'a']
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
	obj.Insert("abcabc")
	obj.Insert("abcfdsafds")
	obj.Insert("abgfsd")
	obj.Insert("cabgfsd")
	fmt.Println(obj.Search("abgfsd"))
	fmt.Println(obj.Search("abg"))
	fmt.Println(obj.Search("abcabcab"))
	fmt.Println(obj.StartsWith("ab"))
	fmt.Println(obj.StartsWith("ca"))
}
