package main

import (
	"app"
	"fmt"
	"net/http"
)

/**
 * curl -X GET "http://localhost:8080/get?name=apktool"
 * curl -X POST "http://localhost:8080/post" -d 'username=apktool&password=123456'
 */

func main() {
	engine := app.New()
	engine.Get("/get", func(c *app.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	engine.Post("/post", func(c *app.Context) {
		c.Json(http.StatusOK, app.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	err := engine.Run(":8080")
	if err != nil {
		fmt.Printf("main.go, %s", err)
	}
}
