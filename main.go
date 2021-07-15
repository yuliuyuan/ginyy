package main

import (
	"ginyy"
	"net/http"
)

func main() {
	e := ginyy.New()
	e.GET("/", func(c *ginyy.Context) {
		c.HTML(http.StatusOK, "<h1>hello salix</h1>")
	})

	e.GET("/hello", func(c *ginyy.Context) {
		c.HTML(http.StatusOK, "<h1>hello test</h1>")
	})
	e.Run(":9999")
}

