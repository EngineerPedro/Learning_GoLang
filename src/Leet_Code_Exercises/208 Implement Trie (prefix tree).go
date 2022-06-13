package main

func main() {
	trie := Constructor_02()
	trie.Insert("apple")
	trie.Search("apple")   // return True
	trie.Search("app")     // return False
	trie.StartsWith("app") // return True
	trie.Insert("app")
	trie.Search("app") // return True
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor_02() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}
