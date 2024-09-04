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
	r.roots[method].Insert(parts, 0)

	key := fmt.Sprintf("%s-%s", method, pattern)
	r.handlers[key] = handler
}

func (r *Router) getRoute(method string, pattern string) (*Node, map[string]string) {
	searchParts := parsePattern(pattern)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	node := root.Search(searchParts, 0)
	return node, nil
}

func (r *Router) handle(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
