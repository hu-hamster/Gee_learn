package gee

import (
	"fmt"
	"net/http"
)

// HandleFunc定义了gee使用的请求处理程序
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine实现了ServeHTTP接口（处理器）
type Engine struct {
	router map[string]HandleFunc
}

// Engine构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (e *Engine) addRoute(mothod string, pattern string, handle HandleFunc) {
	key := mothod + "-" + pattern
	e.router[key] = handle
}

// 定义GET请求方法
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

// 定义POST请求方法
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

// 定义了启动http服务器的方法
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND：%s\n", req.URL)
	}
}
