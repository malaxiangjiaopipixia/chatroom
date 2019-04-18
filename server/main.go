package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("等待读取数据...")
	time.Sleep(time.Second * 2)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	pkgLen := binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen] err=", err)
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal faild err=", err)
		return
	}

	return

}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

	//先从mes中去除mes.data，反序列化成loginmes
	var loginResMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	//mes.Data

	return

}

//处理消息类型函数
func serverPeocessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆函数
		err = serverProcessLogin(conn, mes)
		if err != nil {
			fmt.Println("loginProcessMes(conn, mes) faild err=", err)

		}
		return
	case message.RegisterMesType:
		//处理注册函数

	default:
		//未匹配
		fmt.Println("消息类型不存在，无法处理！！！")

	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()

	//循环的读取客户端信息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了连接，我也退出")
				return
			} else {
				fmt.Println("readPkg(conn) err=", err)
				return
			}
		}

		fmt.Println("mes=", mes)
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
