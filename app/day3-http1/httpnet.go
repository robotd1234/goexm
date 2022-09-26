// http/net 创建一个基本web server
package main

import (
	"fmt"
	"net/http"
)

// 新建handler对象由http.handler来处理
type indexHandler struct {
	content string
}

func handlefunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	// http handlefunc
	http.HandleFunc("/", handlefunc)

	// http handler
	http.Handle("/handle", &indexHandler{content: "Hello, world!"})
	server.ListenAndServe()

}
