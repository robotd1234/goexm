package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func helloTmpl(w http.ResponseWriter, r *http.Request) {

	// 解析文本模板
	t, err := template.ParseFiles("exm_tmpl.html")
	if err != nil {
		fmt.Println(err)
	}
	// 渲染文本模版
	user := UserInfo{
		Name:   "Jack",
		Gender: "male",
		Age:    18,
	}

	t.Execute(w, user)

}

func main() {

	http.HandleFunc("/", helloTmpl)
	http.ListenAndServe(":8086", nil)

}
