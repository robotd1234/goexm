package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// 等待客户端发送请求信息
		// 如果客户端没有发送数据， 则goroutine就阻塞在这里
		fmt.Printf("服务器在等待客户端：%s 发送信息\n", conn.RemoteAddr().String())
		var buf [1024]byte

		// 接受不到/接收到如何处理
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Print(string(buf[:n]))

	}
}

func Server() {

	// 创建一个服务端监听的地址
	// net.Listen net.accept
	// 使用tcp协议，端口号：127.0.0.1:8088

	service := "127.0.0.1:8088"
	l, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("访问客户端信息： con = %v\n, 客户端 IP= %v\n", conn, conn.RemoteAddr().String())
		go handleConnection(conn)
	}

	// tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	// fmt.Println("tcpAddr: ", tcpAddr)
	// fmt.Printf("IP is :%s\n, Port is :%d\n", tcpAddr.IP, tcpAddr.Port)

	// myConn, err1 := net.DialTCP("tcp", nil, tcpAddr)
	// //checkError(err1)
	// fmt.Println("myConn :")
	// //typeof(myConn)

	// 	checkError(err)
	// 	fmt.Println("tcpAddr: ", tcpAddr)
	// 	typeof(tcpAddr)

	//    myConn, err1 := net.DialTCP("tcp", tcpAddr)

}

func main() {
	Server()
}
