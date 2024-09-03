package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
		Path:     r.URL.Path,
		Method:   r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Response.WriteHeader(code)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.Status(code)
	c.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Response.Write([]byte(fmt.Sprintf(format, value...)))
}

func (c *Context) Json(code int, obj interface{}) {
	c.Status(code)
	c.Response.Header().Set("Content-Type", "application/json; charset=utf-8")
	encoder := json.NewEncoder(c.Response)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Response, err.Error(), 500)
	}
}
