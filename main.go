package main

import (
	"net/http"
	"tinygin"
)

func main() {
	// day 2
	// r := tinygin.New()
	// r.GET("/", func(c *tinygin.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello tinygin</h1>")
	// })
	// r.GET("/hello", func(c *tinygin.Context) {
	// 	// expect /hello?name=tinyginktutu
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// })
	//
	// r.POST("/login", func(c *tinygin.Context) {
	// 	c.JSON(http.StatusOK, tinygin.H{
	// 		"username": c.PostForm("username"),
	// 		"password": c.PostForm("password"),
	// 	})
	// })
	//
	// r.Run(":9999")

	// day 3
	r := tinygin.New()
	r.GET("/", func(c *tinygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello tinygin</h1>")
	})

	r.GET("/hello", func(c *tinygin.Context) {
		// expect /hello?name=tinyginktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *tinygin.Context) {
		// expect /hello/tinyginktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *tinygin.Context) {
		c.JSON(http.StatusOK, tinygin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
