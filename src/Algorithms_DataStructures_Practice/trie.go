package main

import "fmt"

//Aplhabetsiz is the number of possible charactersin the trie
const Alphabetsize = 26

type Node_Trie struct {
	children [26]*Node_Trie
	isEnd    bool
}

//Trie
type Trie struct {
	root *Node_Trie
}

//InitTrie will create a new Trie
func IniitTrie() *Trie {
	result := &Trie{root: &Node_Trie{}}
	return result
}

//Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node_Trie{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true

}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return false
	}
	return false
}

func main() {
	trie := IniitTrie()
	fmt.Println(trie.root)
}
