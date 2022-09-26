// 练习http服务端和客户端请求
package main

import (
	"net/http"
)

type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}
func hellofunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myhandler),
		refer:   "",
	}

	http.HandleFunc("/hello", hellofunc)
	http.ListenAndServe(":8080", referer)

}
