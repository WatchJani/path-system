package main

import (
	"time"
)

const (
	Dir   string = "dir"
	Empty        = ""
	End          = "end"
)

type Node struct {
	List []Object
	Next map[string]*Node
}

type Object struct {
	ObjType string
	Name    string
	Created time.Time
}

func NewObject(name, objType string) Object {
	return Object{
		ObjType: objType,
		Name:    name,
		Created: time.Now(),
	}
}

func EmptyNode() *Node {
	return &Node{
		Next: make(map[string]*Node),
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: EmptyNode(),
	}
}

func (t *Trie) NewDir(path string) {
	current, start := t.root, 0

	//create all dir
	//go os.mkDirAll(path)

	for index := 0; index < len(path); index++ {
		if path[index] == '/' {
			path := path[start:index]
			start = index + 1

			if _, ok := current.Next[path]; ok {
				current = current.Next[path] //switch to next node, to next path
			} else { //is not exist this dir, create them
				//Create real new dir

				current.Next[path] = EmptyNode()
				current.List = append(current.List, NewObject(path, Dir))
				current = current.Next[path] //next node
			}
		}
	}
}

func (t *Trie) PrintDirTree() {
	current := t.root
	stack := []*Node{current}

	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, obj := range current.List {
			if obj.ObjType == Dir {
				// fmt.Println(obj.Name)
				stack = append(stack, current.Next[obj.Name])
			}
		}
	}
}

func main() {
	dir := NewTrie()

	dir.NewDir("user/public/main/")
	dir.NewDir("user/public/tom/")

	dir.PrintDirTree()
}
