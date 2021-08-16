package trees

import (
	"github.com/Data-Structures-Golang/pkg/utils"
	"testing"
)

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
}

func assertTrue(test *testing.T, word string) {
	test.Errorf("expected to be true, received false, %v", word)
}

func assertFalse(test *testing.T) {
	test.Errorf("expected to be false, received true")
}

func TestTrie_InsertSearch(t *testing.T) {
	var err error
	trie := NewTrie()
	words := []string{"home", "find", "absolutely", "Golang"}
	for _, word := range words {
		if err = trie.Insert(word); err != nil {
			assertError(t, err)
		}
	}
	for _, word := range words {
		var ok bool
		ok, err = trie.Search(word)
		if err != nil {
			assertError(t, err)
		}
		if !ok {
			assertTrue(t, word)
		}
	}
	noWord := "abc2d"
	if err = trie.Insert(noWord); err != nil {
		switch err.(type) {
		case *utils.InvalidAlphabeticWord:
		default:
			assertError(t, err)
		}
	} else {
		t.Errorf((&utils.InvalidAlphabeticWord{Word: noWord}).Error())
	}

	exists, err := trie.Search(noWord)
	if err != nil {
		switch err.(type) {
		case *utils.InvalidAlphabeticWord:
		default:
			assertError(t, err)
		}
	} else {
		t.Errorf((&utils.InvalidAlphabeticWord{Word: noWord}).Error())
	}
	if exists {
		assertFalse(t)
	}

}
