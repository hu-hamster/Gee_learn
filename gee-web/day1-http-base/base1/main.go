package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
* 功能：实现两个路由
 */

func main() {
	http.HandleFunc("/", indexHander)
	http.HandleFunc("/hello", helloHandle)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
