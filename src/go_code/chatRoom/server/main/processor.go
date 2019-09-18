package main

import (
	"fmt"
	"go_code/chatRoom/common/message"
	"go_code/chatRoom/server/processes"
	"go_code/chatRoom/server/utils"
	"io"
	"net"
)

//创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//根据客户端发送的种类不同,决定调用哪个函数来处理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {

	// 检测是否能接收到客户端发送的消息
	fmt.Println("从客户端接收的消息:", mes)

	switch mes.Type {
	case message.LoginMesType:
		// 处理登录的逻辑
		// 创建一个UserProcess 实例,因为一个用户一个实例,也有一个连接
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册的逻辑
		// 创建一个UserProcess 实例,因为一个用户一个实例,也有一个连接
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// 处理转发客户端发送消息的逻辑
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGroupServerMes(mes)

	default:
		fmt.Println("消息不存在,无法处理....")
	}
	return
}

// 处理和客户端的通讯
func (p *Processor) process() (err error) {

	// 完成指定用户的验证,用户id=100,密码pwd=123456,其他用户不能登录
	// 读取客户端发送过来的数据
	// 先将数据反序列化,然后做校验

	for {
		// 创建Transfer实例,完成读包的任务
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		// 这里读取数据包,封装一个函数用来专门读取接收的数据包readPkg(),传入那个conn连接,返回一个mes结构体
		ret, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了连接,服务端也退出")
				return err
			} else {
				fmt.Println("readPkg出错了,错误原因:", err)
				return err
			}

		}
		err = p.serverProcessMes(&ret)
		if err != nil {
			return err
		}

	}
}
