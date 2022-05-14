package task0208

// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and
// retrieve keys in a dataset of strings. There are various applications of this data structure, such
// as autocomplete and spellchecker.
// Implement the Trie class:
// Trie() Initializes the trie object.
//   void insert(String word) Inserts the string word into the trie.
//   boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
//   boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix,
//      and false otherwise.
//
// Constraints:
//   1 <= word.length, prefix.length <= 2000
//   word and prefix consist only of lowercase English letters.
//   At most 3 * 10^4 calls in total will be made to insert, search, and startsWith.
//
// Results:
//   Runtime: 52 ms, faster than 90.77% of Go online submissions for Implement Trie (Prefix Tree).
//   Memory Usage: 14.5 MB, less than 74.95% of Go online submissions for Implement Trie (Prefix Tree).
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (_this *Trie) Insert(word string) {
	curr := _this
	var idx rune
	for _, chr := range word {
		idx = chr - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (_this *Trie) Search(word string) bool {
	curr := _this
	var idx rune
	for _, chr := range word {
		idx = chr - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

func (_this *Trie) StartsWith(word string) bool {
	curr := _this
	var idx rune
	for _, chr := range word {
		idx = chr - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}
