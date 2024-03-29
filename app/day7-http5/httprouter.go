package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Write([]byte("<html><head><title>Index</title></head><body><h1>Index</h1></body></html>"))
}

func main() {
	router := httprouter.New()

	// Create index controller
	router.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		w.Write([]byte("default get"))
		fmt.Println("I've been used default get")
	})

	// Create specific path matching request
	router.GET("/name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("default name"))
		fmt.Println("I've been used name")
	})

	// Create
	router.GET("/name/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name" + p.ByName("name")))
		fmt.Println("I've been used")
	})
	log.Fatal(http.ListenAndServe(":8083", router))

}
