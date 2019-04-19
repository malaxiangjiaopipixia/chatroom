package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
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

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}

	//time.Sleep(20 * time.Second)
	//处理服务器返回消息
	mes, err = readPkg(conn)
	if err != nil {
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	//fmt.Println("loginResMes=", loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功！！！")
	} else if loginResMes.Code == 500 {
		fmt.Println("登陆失败！！！")
		return errors.New(loginResMes.Error)
	}

	return
}
