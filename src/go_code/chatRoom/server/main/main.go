package main

import (
	"fmt"
	"go_code/chatRoom/server/model"
	"net"
	"time"
)

// 处理和客户端的通讯
func process(conn net.Conn) {

	// 延时关闭
	defer conn.Close()
	// 创建一个总控
	pro := &Processor{
		Conn: conn,
	}
	err := pro.process()
	if err != nil {
		fmt.Println("协程出了问题")
		return
	}

	fmt.Println("进入process : 处理客户端发送的信息----------")

}

// 这里编写一个函数,完成对UserDao的初始化任务
func initUserDao() {
	// 这里的pool,本身是redis文件中的全局变量,并且已经初始化成功,但是要注意,一定要先初始化initPool
	// 如果直接先初始化UserDao,这个pool就是空的
	// 所以这里需要注意初始化的顺序问题
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	//服务器一开始就初始化一个redis连接池,这些传入的值都可以读配置文件填入
	initPool(16, 0, "127.0.0.1:6379", 300*time.Second)
	initUserDao() // 这里注意初始化的顺序
	// 提示信息
	fmt.Println("服务器[新结构]在8889端口监听.......")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务器监听出错:", err)
		return
	}

	// 一旦监听成功,就等待客户端的连接服务
	for {
		fmt.Println("等待客户端的连接.......")
		conn, err := listen.Accept() // 客户端与服务端互通的连接      客户端  <------> 服务端
		if err != nil {
			fmt.Println("客户端连接错误:", err)
		}

		// 成功之后,则启动一个协程,和客户端保持数据的通讯
		go process(conn)

	}
}
