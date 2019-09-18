package main

import (
	"fmt"
	"go_code/chatRoom/client/processes"
)

// 定义两个全局变量
// 表示用户的id
var userId int

// 表示用户的密码
var userPwd string

// 表示用户的名称
var userName string

func main() {

	// 定义两个变量
	// 接收用户的输入
	var key int
	// 是否继续显示
	// loop := true

	for true {
		fmt.Println("-------欢迎登录聊天系统(新功能)---------")
		fmt.Println("\t1.登录聊天室")
		fmt.Println("\t2.注册用户")
		fmt.Println("\t3.退出系统")
		fmt.Println("\t请选择(1-3):")

		fmt.Scanf("%d", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			// 登录聊天室的功能
			fmt.Println("请输入id号:")
			fmt.Scanln(&userId) // 如果使用scanf,里面的字符串后面要加一个/n
			fmt.Println("请输入密码:")
			fmt.Scanln(&userPwd)
			// loop = false
			// 创建一个UserProcess的实例
			up := &processes.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入id号:")
			fmt.Scanln(&userId) // 如果使用scanf,里面的字符串后面要加一个/n
			fmt.Println("请输入密码:")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户昵称:")
			fmt.Scanln(&userName)
			// 创建一个UserProcess的实例,完成注册的请求
			up := &processes.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			// loop = false
		default:
			fmt.Println("你输入有误,请重新输入")
		}
	}

}
