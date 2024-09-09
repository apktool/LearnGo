package main

import (
	"app"
	"fmt"
	"net/http"
)

/**
 * curl -X GET "http://localhost:8080/get?name=apktool"
 * curl -X GET "http://localhost:8080/hi/apktool"
 * curl -X POST "http://localhost:8080/post" -d 'username=apktool&password=123456'
 * curl -X POST "http://localhost:8080/ping/apktool" -d 'name=apktool'
 * curl -X GET "http://localhost:8080/v1/hello"
 * curl -X GET "http://localhost:8080/v2/hello"
 */

func main() {
	engine := app.New()
	engine.Use(app.Logger1())
	engine.Use(app.Logger2())
	engine.Get("/get", func(c *app.Context) {
		c.String(http.StatusOK, "get %s, you're at %s\n", c.Form("name"), c.Path)
	})

	engine.Get("/hi/:name", func(c *app.Context) {
		c.String(http.StatusOK, "hi %s, you're at %s\n", c.Form("name"), c.Path)
	})

	engine.Post("/post", func(c *app.Context) {
		c.Json(http.StatusOK, app.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	engine.Post("/ping/:name", func(c *app.Context) {
		c.Json(http.StatusOK, app.H{
			"pong": c.PostForm("name"),
		})
	})

	v1 := engine.Group("/v1")
	{
		v1.Get("/hello", func(c *app.Context) {
			c.String(http.StatusOK, "hello v1, you're at %s\n", c.Path)
		})
	}

	v2 := engine.Group("/v2")
	{
		v2.Get("/hello", func(c *app.Context) {
			c.String(http.StatusOK, "hello v2, you're at %s\n", c.Path)
		})
	}

	err := engine.Run(":8080")
	if err != nil {
		fmt.Printf("main.go, %s", err)
	}
}
