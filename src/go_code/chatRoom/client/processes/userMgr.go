package processes

import (
	"go_code/chatRoom/client/model"
	"fmt"
	"go_code/chatRoom/common/message"
)

// 定义客户端维持的map
var onlineClient = make(map[int]*message.User, 10)
var curUser model.CurUser  // 用户登录成功之后,对CurUser进行初始化

// 在客户端显示当前在线的用户
func outputOnlineUser() {

	// 遍历一遍onlineClie
	fmt.Println("当前在线用户列表:")
	for k, _ := range onlineClient {
		fmt.Println("用户id:\t", k)
	}
}

// 编写一个方法,处理返回的notifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineClient[notifyUserStatusMes.UserId]
	fmt.Println("user是什么:", user)
	if !ok {
		// 创建一个User
		user = &message.User{
			UserId:     notifyUserStatusMes.UserId,
		}
	}
	// 如果有,只更新状态
	user.UserStatus = notifyUserStatusMes.Status
	onlineClient[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
