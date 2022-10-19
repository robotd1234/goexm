package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	// 读取系统参数信息
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1] // 从参数中读取主机信息

	// 建立连接
	conn, err := net.Dial("tcp", service)
	validateError(err)

	// 对服务器请求write()方法
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	validateError(err)

	// 答应响应数据
	res, err := fullyRead(conn)
	validateError(err)
	fmt.Println(string(res))
	os.Exit(0)
}

// 连接出错，打印错误信息并退出程序
func validateError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(0)
		return
	}

}

// 接受连接的对象 读取响应数据
func fullyRead(conn net.Conn) ([]byte, error) {
	// 关闭conn, 延迟执行
	defer conn.Close()

	// 读取流数据到buffer， 并返回结果
	res := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		res.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return res.Bytes(), nil
}
