package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8080"
	body := `{"email": "test@homuchen.com", "name": "homuchen"}`
	ct := "application/json"
	resp, err := http.Post(url, ct, bytes.NewReader([]byte(body)))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

