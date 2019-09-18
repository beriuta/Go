/*
	TCP网络编程
		实际应用: QQ, 迅雷, 百度网盘客户端, 新浪网站, 京东商城, 淘宝...
	网络编程的两种:
		1. TCP socket编程, 是网络编程的主流,之所以叫TCP socket变成,是因为底层是基于TCP/IP协议,比如QQ聊天
		2. B/S结构的HTTP编程, 我们使用的浏览器去访问服务器时,使用的就是HTTP协议,而HTTP协议的底层依旧是用TCP socket实现的,比如,京东商城
	端口
		只要是做服务程序,都必须监听一个端口
		该端口就是其他程序和该服务通讯的通道
		一台电脑有65535个端口
		一旦一个端口被某个程序占用了,那其他程序就不能再此端口监听
		客户端的端口是随意的,根据当时的端口场景
	端口分类:
		0号是保留端口
		1-1024是固定端口,又叫有名端口,即被某些程序固定使用,一般程序员不使用
			22:ssh远程登录协议
			23:teinet使用
			21:ftp使用
			25:smtp服务使用
			80:iis使用
			7:echo服务
		1025-65535是动态端口
	TCP socket编程的快速入门
		服务端的处理流程
			1. 监听端口
			2.接收客户端的TCP链接,简历客户端和服务端的链接
			3. 创建goroutine,处理该链接的请求(通常客户端会通过链接发送请求包)

		客户端的处理流程
			1, 建立服务端的链接
			2, 发送请求数据,接收服务器端返回的数据
		3, 关闭链接

*/
package TCP

