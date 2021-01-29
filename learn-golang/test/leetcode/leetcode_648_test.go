package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * 648. 单词替换
 *
 */
func Test_leetcode_648(t *testing.T) {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords([]string{"c", "a", "b"}, "aadsfasf absbs bbab cadsfafs"))
}

type Node struct {
	end  int
	next map[string]*Node
}

func (root *Node) insert(word string) {
	if word == "" {
		return
	}
	node := root
	length := len(word)
	for i := 0; i < length; i++ {
		c := word[i : i+1]
		if _, ok := node.next[c]; !ok {
			//不包含
			node.next[c] = &Node{
				end:  0,
				next: map[string]*Node{},
			}
		}
		node = node.next[c]
	}
	node.end++
}

func (root *Node) prefix(s string) string {
	//s的最短前缀
	node := root
	l := len(s)
	for i := 0; i < l; i++ {
		if _, ok := node.next[s[i:i+1]]; !ok {
			return s
		}
		node = node.next[s[i:i+1]]

		if node.end > 0 {
			return s[0 : i+1]
		}
	}
	return s
}

func replaceWords(dictionary []string, sentence string) string {
	// 把词根放入前缀树，在树上查找每个单词的最短词根
	length := len(dictionary)
	root := &Node{next: map[string]*Node{}}
	for i := 0; i < length; i++ {
		root.insert(dictionary[i])
	}

	s := strings.Split(sentence, " ")

	l := len(s)
	for i := 0; i < l; i++ {
		res := root.prefix(s[i])
		s[i] = res
	}
	return strings.Join(s, " ")
}
