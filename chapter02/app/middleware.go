package app

import (
	"fmt"
	"time"
)

func Logger1() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		fmt.Printf("[%d] Logger1: %s in %v\n", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func Logger2() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		fmt.Printf("[%d] Logger2: %s in %v\n", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
