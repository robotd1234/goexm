package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Send struct {
	Java, Go string
}

func main() {
	fmt.Println("client start ...")
	// Dial
	client, err := jsonrpc.Dial("tcp", "localhost:8085")
	if err != nil {
		log.Fatal("Dial error: ", err)
	}

	send := Send{"Java", "Go"}
	var receive string
	err = client.Call("Programmer.GetSkill", send, &receive)
	if err != nil {
		fmt.Println("Call error: ", err)
	}
	fmt.Println("receive:", receive)
}
