package trees

import (
	"fmt"
	"github.com/Data-Structures-Golang/pkg/utils"
)

// prefix tree
type trie struct {
	root *trieNode
}

type trieNode struct {
	complete bool
	value    rune
	children [26]*trieNode
	level    int
}

func NewTrie() Trie {
	return &trie{
		root: &trieNode{
			complete: false,
			value:    0,
			level:    0,
		},
	}
}

type Trie interface {
	// Insert inserts a word to the trie
	Insert(word string) error
	// Search verifies if the word exists
	Search(word string) (exists bool, err error)
}

// Insert inserts a word to the trie
func (t *trie) Insert(word string) (err error) {
	if t == nil {
		return &utils.StructIsNilException{FuncName: "Insert", DataStructure: "Prefix Tree"}
	}
	head := t.root
	var endOfWord bool
	for index, r := range word {
		if r >= 'a' && r <= 'z' {
			r -= 32
		}
		if r < 'A' || r > 'Z' {
			return &utils.InvalidAlphabeticWord{Word: word}
		}
		if index == len(word)-1 {
			endOfWord = true
		}
		fmt.Println("insert:", r, string(r), int(r-65))
		head.children[int(r)-65] = &trieNode{
			complete: endOfWord,
			level:    index + 1,
			value:    r,
		}
		head = head.children[int(r)-65]
	}
	return
}

// Search verifies if the word exists
func (t *trie) Search(word string) (bool, error) {
	if t == nil {
		return false, &utils.StructIsNilException{FuncName: "Search", DataStructure: "Prefix Tree"}
	}
	head := t.root
	for index, r := range word {
		if head == nil {
			break
		}
		if r >= 'a' && r <= 'z' {
			r -= 32
		}
		if r < 'A' || r > 'Z' {
			return false, &utils.InvalidAlphabeticWord{Word: word}
		}
		if head.children[int(r)-65] == nil || head.children[int(r)-65].value != r {
			break
		}
		if index == len(word)-1 && head.children[int(r)-65].complete {
			return true, nil
		}
		head = head.children[int(r)-65]
	}
	return false, nil
}
