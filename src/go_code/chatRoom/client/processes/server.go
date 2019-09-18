package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatRoom/client/utils"
	"go_code/chatRoom/common/message"
	"net"
	"os"
)

// 显示登录成功之后的界面
func ShowMenu() {
	fmt.Println("-----------恭喜xxx登录成功----------")
	fmt.Println("-----------1.显示在线用户列表----------")
	fmt.Println("-----------2.发送消息----------")
	fmt.Println("-----------3.信息列表----------")
	fmt.Println("-----------4.退出系统----------")
	fmt.Println("请选择1-4:")
	var loginkey int
	var content string
	fmt.Scanln(&loginkey)

	// 因为会经常使用到SmsProcess实例,因此我们将其定义在switch外部
	smsProcess := &SmsProcess{}
	switch loginkey {
	case 1:
		// 显示在线用户列表
		// fmt.Println("你选择了查看显示在线用户")
		outputOnlineUser()
	case 2:
		// 发送消息
		fmt.Println("请输入你想对大家说的话:")
		fmt.Scanln(&content)
		fmt.Println("你输入的内容是", content)
		smsProcess.SendGroupMes(content)

	case 3:
		// 信息列表
		fmt.Println("你选择了信息列表")
	case 4:
		// 退出系统
		fmt.Println("你选择了退出系统....")
		os.Exit(0)
	default:
		fmt.Println("输入选项不正确,请重新输入")
	}

}

func ProcessServerMes(conn net.Conn) {
	// 创建一个Transfer实例,不停的读取消息
	tp := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在读取服务器发送的消息")
		mes, err := tp.ReadPkg()
		if err != nil {
			fmt.Println("客户端读取服务端数据失败", err)
			return
		}
		// 如果读取到数据
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 1. 取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 2. 把这个用户的信息,状态都保存到客户map中 [int]User
			updateUserStatus(&notifyUserStatusMes)
			fmt.Println("有人上线了")
		case message.SmsMesType: // 有人群发消息了
			fmt.Println("收到群发消息")
			outputGroupMes(&mes)

		default:
			fmt.Println("服务器端返回了一个未知的消息类型")
		}
	}
}
