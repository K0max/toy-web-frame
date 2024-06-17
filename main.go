package main

import (
	"log"
	"net/http"
	"time"
	"tinygin"
)

func onlyForV2() tinygin.HandlerFunc {
	return func(c *tinygin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

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
	// r := tinygin.New()
	// r.GET("/", func(c *tinygin.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello tinygin</h1>")
	// })
	//
	// r.GET("/hello", func(c *tinygin.Context) {
	// 	// expect /hello?name=tinyginktutu
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// })
	//
	// r.GET("/hello/:name", func(c *tinygin.Context) {
	// 	// expect /hello/tinyginktutu
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	// })
	//
	// r.GET("/assets/*filepath", func(c *tinygin.Context) {
	// 	c.JSON(http.StatusOK, tinygin.H{"filepath": c.Param("filepath")})
	// })
	//
	// r.Run(":9999")

	// day 4
	// r := tinygin.New()
	// r.GET("/", func(c *tinygin.Context) {
	// 	c.HTML(http.StatusOK, "<h1>index</h1>")
	// })
	// v1 := r.Group("/v1")
	// v1.GET("/", func(c *tinygin.Context) {
	// 	c.HTML(http.StatusOK, "<h1>htllo</h1>")
	// })
	// v1.GET("/hello", func(ctx *tinygin.Context) {
	// 	ctx.String(http.StatusOK, "Hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	// })
	//
	// v2 := r.Group("/v2")
	// v2.GET("/hello/:name", func(ctx *tinygin.Context) {
	// 	ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	// })
	//
	// v2.POST("/login", func(ctx *tinygin.Context) {
	// 	ctx.JSON(http.StatusOK, tinygin.H{
	// 		"username": ctx.PostForm("username"),
	// 		"password": ctx.PostForm("password"),
	// 	})
	// })
	// r.Run(":9999")

	// day5
	r := tinygin.New()
	r.Use(tinygin.Logger()) // global midlleware
	r.GET("/", func(c *tinygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello tinygin</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *tinygin.Context) {
			// expect /hello/tinyginktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
