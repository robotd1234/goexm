package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func Client() {
	// Client 建立conn
	// net.Dial 返回的Conn, 再处理Conn读写数据

	service := "127.0.0.1:8088"
	conn, err := net.Dial("tcp", service)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端连接成功，服务端地址:%s", conn.RemoteAddr().String())

	// 连接后向服务器发送请求数据
	// 客户端获取取终端输入数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入【终端】

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.Trim(line, "\r\n")

		if line == "exit" {
			fmt.Println("退出客户端")
			break
		}

		// 将line 发送给服务器端
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	Client()
}
