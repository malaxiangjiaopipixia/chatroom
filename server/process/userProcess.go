package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

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
		fmt.Println("son.Marshal(mes) faild! err=", err)
		return
	}

	//这里是使用分层模式(mvc),
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn ,data) faild! err =", err)
		return
	}
	return
}
