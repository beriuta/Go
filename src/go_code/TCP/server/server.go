/*
	服务端
*/

package main


import (
	"fmt"
	"net"  // 做网络socket开发时,net包含有我们需要的方法和函数
)

func process(conn net.Conn) {
	// 循环的接收客户端发送的数据
	defer conn.Close() // 关闭conn

	for {

		// 创建一个新的切片
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待%s的数据\n", conn.RemoteAddr().String())

		// 1, 等待客户端通过conn发送信息
		// 2, 如果客户端没有wrtie[发送]内容,那么协程就会一直在阻塞
		n, err := conn.Read(buf)  
		if err != nil {
			fmt.Println("客户端已退出....")
			return
		}
		// 将客户端发送的数据显示在终端
		fmt.Print(string(buf[:n]))  // n很重要,只显示实际从conn中读取的数据
	}
}


func main() {

	fmt.Println("服务器开始监听了......")
	// tcp: 表示使用网络协议是tcp
	// 0.0.0.0:8088 : 表示本地监听8088端口
	listen, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()  // 延时关闭

	// 循环等待客户端的连接
	for {
		// 等待客户端连接
		fmt.Println("等待客户端连接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("一个连接出了问题:", err)  // 这里不需要return,因为可能有的连接能连上,有的连接不能连上
		} else {
			fmt.Printf("连上的接口是什么:%v, 连接的客户端地址:%v\n", conn, conn.RemoteAddr().String())
		}

		// 这里接收用户的数据,但是不能阻塞,不然别的客户端发不过来数据,所以要使用协程
		go process(conn)
	}


	// fmt.Println("listen=", listen)  // listen= &{0xc00009c000}

}