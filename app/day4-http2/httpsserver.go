//https 服务器练习

package main

import (
	"log"
	"net/http"
)

type httpHandler struct {
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got Connection: %s", r.Proto)
	w.Write([]byte("Hello,this is http/2 message"))
}

func main() {

	//启动服务器
	srv := &http.Server{Addr: ":8080", Handler: &httpHandler{}}
	log.Printf("Serving on https://0.0.0.0:8080")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
	http.ListenAndServe(":8080", nil)
}
