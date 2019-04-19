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

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(bytes) faild!", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

	//先从mes中去除mes.data，反序列化成loginmes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data) faild! err=", err)
		return
	}

	var resMes message.Message
	var loginResMes message.LoginResMes
	//mes.Data
	if loginMes.UserName == "yang" && loginMes.UserPwd == "tian" {
		//合法
		fmt.Println("合法用户")
		loginResMes.Code = 200
	} else {
		//不合法
		fmt.Println("不合法用户")
		loginResMes.Code = 500
		loginResMes.Error = "用户名或密码错误，请重新输入......"
	}

	fmt.Println("loginResMes=", loginResMes)
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) faild! err=", err)
		return
	}

	resMes.Type = message.LoginResMesType
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(mes) faild! err=", err)
		return
	}

	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("writePkg(conn ,data) faild! err =", err)
		return
	}
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
		err = serverPeocessMes(conn, &mes)
		if err != nil {
			return
		}
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
