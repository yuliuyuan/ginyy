package ginyy

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req * http.Request)

//Engine implement the interface of ServeHTTP
type Engine struct{
	router map[string]HandlerFunc
}

//
func New() *Engine{
	return &Engine{router: make( map[string]HandlerFunc )}
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handlerFunc
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc){
	engine.addRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc){
	engine.addRoute("POST", pattern, handlerFunc)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	key := req.Method + "-" + req.URL.Path
	if HandlerFunc, ok := engine.router[key]; ok{
		HandlerFunc(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
