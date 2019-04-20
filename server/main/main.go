package main

import (
	"fmt"
	"net"
)

func process1(conn net.Conn) {
	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}

	err := processor.Process2()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	listen, err := net.Listen("tcp", "localhost:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("监听端口出错，退出！！！")
		return
	}

	fmt.Println("服务器在8889端口监听......")
	fmt.Println("等待客户端连接......")

	for {
		//获取连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
		}
		//一旦获取连接，处理函数进行处理
		go process1(conn)
	}
}
