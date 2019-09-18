/*
	客户端
*/
package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "192.168.0.102:8088")
	if err != nil {
		fmt.Println("客户端来连接错误:", err)
		return
	}
	// fmt.Printf("看一下客户端连接的内容:%v, 连接的客户端地址:%v\n", conn, conn.RemoteAddr().String())
	// 功能一: 客户端可以发送单行数据,然后退出
	reader := bufio.NewReader(os.Stdin)  // os.Stdin : 代表标准输入[终端]

	// 从终端读取一行用户输入,并准备发送到服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("读取客户端输入错误:%v\n", err)

	}

	// 再将line发送到服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Printf("发送数据给客户端时出错:%v\n", err)

	}
	fmt.Printf("客户端向服务端发送了%d字节数据,并退出", n)
}