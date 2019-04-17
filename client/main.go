package main

import (
	"fmt"
	"os"
)

var (
	//userid   int
	username string
	userpwd  string
)

func main() {
	var key int
	loop := true

	for loop {
		fmt.Println("-----------欢迎进入聊天室！！！-----------")
		fmt.Println("\t-----------1.登陆-----------")
		fmt.Println("\t-----------2.注册-----------")
		fmt.Println("\t-----------3.退出-----------")
		fmt.Println("请输入数字1~3：")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("进入登录！！！")
			loop = false
		case 2:
			fmt.Println("进入注册！！！")
			loop = false
		case 3:
			fmt.Println("退出程序！！！")
			os.Exit(0)
		default:
			fmt.Println("套你猴子")
		}

	}

	if key == 1 {
		fmt.Println("请输入用户名：")
		fmt.Scanf("%s\n", &username)
		fmt.Println("请输入密码:")
		fmt.Scanf("%s\n", &userpwd)

		err := login(username, userpwd)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("用户名密码正确，欢迎你：", username)
		}
	} else if key == 2 {

	}
}
