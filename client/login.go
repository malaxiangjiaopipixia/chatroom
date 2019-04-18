package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(username string, userpwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial faild!", err)
		return
	}

	defer conn.Close()

	//准备向服务器发送消息
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserName = username
	loginMes.UserPwd = userpwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal faild!", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Message json.Marshal faild!", err)
		return
	}

	//将message的长度转换成一个可以表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(bytes) faild!", err)
		return
	}

	fmt.Printf("客户端发送数据长度=%d 内容=%s", pkgLen, string(data))

	return nil
}
