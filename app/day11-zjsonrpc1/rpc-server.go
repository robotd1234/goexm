// 实现json-rpc 服务端

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 服务端编写接口以及实现
func init() {
	fmt.Println("JSON编码RPC")
}

type ArgsLanguage struct {
	Java, Go string
}

type Programmer string

func (m *Programmer) GetSkill(al *ArgsLanguage, skill *string) error {
	*skill = "Skill:" + al.Java + ", Skill2" + al.Go
	return nil
}

func main() {

	// 主函数注册实例化对象
	str := new(Programmer)
	fmt.Printf("this is the Programmer: %v\n", str)

	// 注册服务
	rpc.Register(str)

	// net包开启TCP服务端
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8085")
	if err != nil {
		fmt.Println("ResolveTCPAddr err=", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("tcp listen err=", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
