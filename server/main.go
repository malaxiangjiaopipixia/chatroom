package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

}

func main() {

	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("监听端口出错，退出！！！")
		return
	}

	for {
		//获取连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
		}

		//一旦获取连接，处理函数进行处理
		go process(conn)
	}
}
