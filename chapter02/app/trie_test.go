package app

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	lang1 := []string{"", ":lang", "intro"}
	lang2 := []string{"", ":lang", "tutorial"}
	lang3 := []string{"", ":lang", "doc"}
	about := []string{"", ":about"}
	p1 := []string{"", "p", "blog"}
	p2 := []string{"", "p", "related"}

	var root Node = Node{part: "", children: make([]*Node, 0), isWild: false}
	root.Insert(lang1, 0)
	root.Insert(lang2, 0)
	root.Insert(lang3, 0)
	root.Insert(about, 0)
	root.Insert(p1, 0)
	root.Insert(p2, 0)

	ok := reflect.DeepEqual(len(root.children), 1)
	ok = ok && reflect.DeepEqual(len(root.children[0].children), 3)
	ok = ok && reflect.DeepEqual(len(root.children[0].children[0].children), 3)
	ok = ok && reflect.DeepEqual(root.children[0].children[1].part, ":about")
	if !ok {
		t.Fatal("test failed")
	}

	query := []string{"", "p", "blog"}
	node := root.Search(query, 0)
	fmt.Println(node.String())
	ok = reflect.DeepEqual(node.part, "blog")
	if !ok {
		t.Fatal("test failed")
	}

	query = []string{"", "p", "app"}
	node = root.Search(query, 0)
	if node != nil {
		t.Fatal("test failed")
	}

}
