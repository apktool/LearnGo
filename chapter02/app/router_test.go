package app

import (
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/blog"), []string{"p", "blog"})

	if !ok {
		t.Fatal("test failed")
	}
}

func TestRoute(t *testing.T) {
	var router *Router = newRouter()
	router.addRoute("GET", "/p/blog", func(c *Context) {})
	router.addRoute("GET", "/p/about", func(c *Context) {})
	router.addRoute("GET", "/p/ping", func(c *Context) {})

	node, _ := router.getRoute("GET", "/p/blog")
	ok := reflect.DeepEqual(node.part, "blog")
	if !ok {
		t.Fatal("test failed")
	}
}
