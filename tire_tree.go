package golang_segment

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type TireTree struct {
	root map[int32]*TireNode
}

func NewTireTree() *TireTree {
	ins := &TireTree{
		root: make(map[int32]*TireNode),
	}
	return ins
}

func (this *TireTree) LoadDict(dictpath string) {
	f, err := os.Open(dictpath)
	if err != nil {
		log.Println("open file:" + dictpath + "error !")
		return
	}
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		word := string(line)
		this.insert(word)
	}
	f.Close()
}

func (this *TireTree) insert(word string) {
	words := strings.Split(word, " ")
	word = words[0]
	runes := []rune(string(word))
	var stopIdx = len(runes) - 1
	var current rune = runes[0]
	var root *TireNode = this.getRoot(current)
	for i := 1; i <= stopIdx; i++ {
		current = runes[i]
		seg := root.AddChild(current)
		root = seg
	}
}

func (this *TireTree) getRoot(value int32) *TireNode {
	node := this.root[value]
	if node == nil {
		node = &TireNode{
			value:    value,
			children: make(map[int32]*TireNode),
		}
		this.root[value] = node
	}
	return node
}

func (this *TireTree) Search(sentence string) []string {
	runes := []rune(sentence)
	words := make([]string, 0)
	length := len(runes)
	var current int32
	var root *TireNode = nil
	var word string = ""
	for i:=0; i<length; i++ {
		current = runes[i]
		if root == nil {
			word = ""
			word += string(current)
			root = this.root[current]
		} else {
			node := root.getChild(current)
			if node != nil {
				word += string(current)
			} else {
				words = append(words, word)
				i--
			}
			root = node
		}
	}
	return words
}