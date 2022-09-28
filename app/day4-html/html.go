package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func webPageFunc(req http.ResponseWriter, res *http.Request) {
	t, err := template.ParseFiles("./tmpl_exmpl.html")
	if err != nil {
		fmt.Println("template parse error:", err)
		return
	}
	name := "我爱Go语言"
	t.Execute(req, name)
}

func main() {
	http.HandleFunc("/", webPageFunc)
	http.ListenAndServe(":8080", nil)
}
