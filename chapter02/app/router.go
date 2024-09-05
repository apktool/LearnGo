package app

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Router struct {
	roots    map[string]*Node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*Node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
		}
	}
	return parts
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)

	key := fmt.Sprintf("%s-%s", method, pattern)
	r.handlers[key] = handler
}

func (r *Router) getRoute(method string, pattern string) *Node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}

	searchParts := parsePattern(pattern)
	return root.Search(searchParts, 0)
}

func (r *Router) handle(c *Context) {
	node := r.getRoute(c.Method, c.Path)
	if node != nil {
		key := fmt.Sprintf("%s-%s", c.Method, node.pattern)
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
