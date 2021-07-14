package main

import (
	"fmt"
	"github/ginyy/ginyy"
	"net/http"
)

type Engine struct{}


func main() {
	engine := ginyy.New()
	engine.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	engine.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	engine.Run(":9999")
}

