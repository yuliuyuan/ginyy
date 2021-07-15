package ginyy

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

type router struct {
	routes map[string]*node //k为http的方法：get、put
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{routes: make(map[string]*node) ,handlers: make(map[string]HandlerFunc )}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) error{
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler

	parts := parsePattern(pattern)
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = &node{}
	}

	//过滤错误url
	if len(parts) == 0 {
		if pattern == "/" {
			r.routes[method].pattern = "/"
			return nil
		} else {
			return errors.New("error url: " + pattern)
		}
	}
	r.routes[method].insert(pattern, parts, 1)
	return nil
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {

	if path == "/" {
		if r.routes[method] == nil{
			return nil, nil
		} else if( r.routes[method].pattern == "/" ){
			return r.routes[method], nil
		} else {
			return nil, nil
		}
	}

	searchParts := parsePattern(path)

	params := make(map[string]string) //保存请求路径中的参数

	routes, ok := r.routes[method]
	if ok {
		n := routes.search(searchParts, 1)
		if n != nil {
			parts := parsePattern(n.pattern)
			for index, part := range parts {
				if part[0] == ':' {
					params[part[1:]] = searchParts[index]
				}
				if part[0] == '*' && len(part) > 1 {
					params[part[1:]] = strings.Join(searchParts[index:], "/")
					break
				}
			}
			return n, params
		}
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}