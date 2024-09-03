package app

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(*Context)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	r.handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
