package main

import (
	"chatroom/common/message"

	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverPeocessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆函数
		userPro := &process.UserProcess{
			Conn: this.Conn,
		}
		err = userPro.ServerProcessLogin(mes)
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

func (this *Processor) Process2() (err error) {

	//循环的读取客户端信息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了连接，我也退出")
				return err
			} else {
				fmt.Println("readPkg(conn) err=", err)
				return err
			}
		}

		fmt.Println("mes=", mes)

		err = this.serverPeocessMes(&mes)
		if err != nil {
			return err
		}
	}
}
