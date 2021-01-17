package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take a work and add it to the trie
func (t *Trie) Insert(word string) {

	wordLength := len(word)
	currentNode := t.root
	// loop to process the word, letter by letter
	for i := 0; i < wordLength; i++ {

		// create index for each letter by subtracting ASCII value for 'a' (97)
		// this makes index of the letter 'a' 0 (zero)
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			// if nil the current letter isn't a known child
			currentNode.children[charIndex] = &Node{}
		}
		// if index isn't empty, advance to it
		currentNode = currentNode.children[charIndex]
	}
	// Reached the end of the word
	currentNode.isEnd = true
}

// Search takes in a word and returns true if the word is in the Trie
func (t *Trie) Search(word string) bool {

	wordLength := len(word)
	currentNode := t.root
	// loop to process the word, letter by letter
	for i := 0; i < wordLength; i++ {

		// create index for each letter by subtracting ASCII value for 'a' (97)
		// this makes index of the letter 'a' 0 (zero)
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			// if nil the the work isn't found, return
			return false
		}
		// if index isn't empty, advance to it
		currentNode = currentNode.children[charIndex]
	}
	// Found the end of the word
	if currentNode.isEnd == true {
		return true
	}

	// No match
	return false
}

func main() {
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

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Printf("Searching for argon ")
	fmt.Println(myTrie.Search("argon"))

	fmt.Printf("Searching for wizard ")
	fmt.Println(myTrie.Search("wizard"))
}
