package main

import (
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()

	//循环的读取客户端信息
	for {
		buf := make([]byte, 8096)
		fmt.Println("等待读取数据...")
		time.Sleep(time.Second * 2)
		n, err := conn.Read(buf[:4])
		if err != nil || n != 4 {
			fmt.Println("conn.Read err=", err)
			return
		}

		fmt.Println("读到的buf=", buf[:4])
		fmt.Println("等待客户端连接......")
		return
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
		go process(conn)
	}
}
