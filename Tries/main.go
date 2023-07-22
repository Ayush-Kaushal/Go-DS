package main

import "fmt"

// AlphabetSize is the size of number of possible childrenin the trie
const AlphabetSize = 26

// Node
type Node struct{
	children [26]*Node
	isEnd bool
}

// Trie
type Trie struct{
	root *Node
}

func InitTrie() *Trie{
	result := &Trie{root: &Node{}}

	return result
}
// Insert wiil take in a word and add it to the trie
func (t * Trie) Insert(w string){
	wordLength := len(w)

	currNode := t.root

	for i := 0; i < wordLength; i++{
		charIndex := w[i] - 'a'
		if currNode.children[charIndex] == nil{
			currNode.children[charIndex] = &Node{}
		}

		currNode = currNode.children[charIndex]
	}

	currNode.isEnd = true

}

// Search will take a word and return true if that word is included in the trie
func (t *Trie) Search(w string) bool{
	wordLength := len(w)

	currNode := t.root

	for i := 0; i < wordLength; i++{
		charIndex := w[i] - 'a'
		if currNode.children[charIndex] == nil{
			return false
		}

		currNode = currNode.children[charIndex]
	}

	return currNode.isEnd
}


func main(){
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	
	for _, v := range toAdd{
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("wizard"))
	
}