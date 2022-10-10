package main

import (
	"net/http"
)

type indexHandler struct {
}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))

}

func hi(w http.ResponseWriter, r *http.Request) {
	s := "hello web"
	w.Write([]byte(s))
}

func main() {

	mux := http.NewServeMux()
	// mux.HandleFunc("/", hi)
	// http.ListenAndServe(":8090", mux)

	mux.Handle("/handler1", &indexHandler{})
	mux.HandleFunc("/handler2", hi)
	http.ListenAndServe(":8090", mux)

}
