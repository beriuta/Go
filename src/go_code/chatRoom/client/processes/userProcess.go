package processes

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatRoom/client/utils"
	"go_code/chatRoom/common/message"
	"net"
)

type UserProcess struct {
}

// 一个函数,实现登录功能

func (u *UserProcess) Login(userId int, userPwd string) (err error) {

	// 开始定协议
	fmt.Printf("userId = %d\n, userPwd = %s\n", userId, userPwd)
	// 连接服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("客户端连接服务端出错~~~~~~:", err)
		return
	}

	defer conn.Close() // 延时关闭

	// 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType // 发送的消息类型

	// 创建一个LoginMes 结构体, 将用户名跟密码填入LoginMes实例对象
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 将loginMes序列化
	data, err := json.Marshal(loginMes) // 返回的data是一个[]byte切片
	if err != nil {
		fmt.Println("将loginMes序列化出错:", err)
		return
	}
	// 如果序列化成功,将data赋值给mes中的Data
	mes.Data = string(data)
	// 将mes序列化
	dataMes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("将mes序列化出错:", err)
		return
	}
	// 这里的dataMes就是要发送的消息,是byte切片
	// 获取到dataMes的长度转成一个表示长度的byte切片
	// 先把dataMes的长度发送到服务器中
	var pkgLen uint32                           // 无符号的,能表示比较大的数字
	var buf [4]byte                             // 4*8=32,一个字节8位
	pkgLen = uint32(len(dataMes))               // 将int类型的数字强转成uint32
	binary.BigEndian.PutUint32(buf[:4], pkgLen) // PutUint32将一个uint32的数据类型转换成一个byte切片

	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送长度出错", err)
		return
	}
	// fmt.Println("发送长度的n = ", n)
	fmt.Printf("login中客户端发送的消息长度为:%d,发送的内容是%v\n", len(dataMes), string(dataMes))

	// 发送消息本身
	_, err = conn.Write(dataMes)
	if err != nil {
		fmt.Println("发送内容出错", err)
		return
	}
	// 创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	// 处理从服务器中返回的消息
	mesRet, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("客户端解析服务端的login数据出错:", err)
		return
	}
	// 将mes的data部门反序列化成 LoginResMes
	var reLoginMes message.LoginReMes
	err = json.Unmarshal([]byte(mesRet.Data), &reLoginMes)
	fmt.Println("我是处理服务端返回的消息中的登录信息------------", reLoginMes)
	if err != nil {
		fmt.Println("客户端解析服务端的login数据出错---->", err)
		return
	} else if reLoginMes.Code == 500 {
		fmt.Println(reLoginMes.Error)
	} else if reLoginMes.Code == 200 {

		// 初始化curUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline
		// 将服务端发送的数据中的在线列表展示出来
		for _, v := range reLoginMes.UserIds {
			// 每个用户不显示自己的id
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			// 在这里完成客户端onlineClient的初始化s
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineClient[v] = user
		}
		fmt.Println("登录成功")
		fmt.Print("\n\n")
		// 这里客户端启动一个协程
		// 该协程保持和服务器端的通讯,如果服务器有数据推送给客户端,则接收并显示客户端的终端中
		go ProcessServerMes(conn)
		// 显示登陆成功后的菜单
		for {
			ShowMenu()
		}
	}

	return nil
}

func (u *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	fmt.Printf("userId = %d\n, userPwd = %s\n", userId, userPwd)
	// 连接服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("客户端连接服务端出错~~~~~~:", err)
		return
	}

	defer conn.Close() // 延时关闭

	// 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType // 发送的消息类型

	// 创建一个LoginMes 结构体, 将用户名跟密码填入LoginMes实例对象
	var registerMes message.RegisterMes

	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	// 将registerMes序列化
	data, err := json.Marshal(registerMes) // 返回的data是一个[]byte切片
	if err != nil {
		fmt.Println("将registerMes序列化出错:", err)
		return
	}
	// 如果序列化成功,将data赋值给mes中的Data
	mes.Data = string(data)
	// 将mes序列化
	dataMes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("将mes序列化出错:", err)
		return
	}

	// 创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WriterPkg(dataMes)
	if err != nil {
		fmt.Println("WeiterPkg包包出错了")
		return
	}

	// 处理从服务器中返回的消息
	mesRet, err := tf.ReadPkg() // mesRet就是RegisterReMes
	if err != nil {
		fmt.Println("客户端解析服务端的注册数据出错:", err)
		return
	}
	// 将mes的data部门反序列化成 RegisterReMes
	var registerReMes message.RegisterReMes
	err = json.Unmarshal([]byte(mesRet.Data), &registerReMes)
	fmt.Println("我是处理服务端返回的消息中的登录信息------------", registerReMes)
	if err != nil {
		fmt.Println("客户端解析服务端的注册数据出错---->", err)
		return
	} else if registerReMes.Code == 400 { // 已经被注册了
		fmt.Println(registerReMes.Error)
	} else if registerReMes.Code == 200 { // 注册成功
		fmt.Println("注册成功,可以重新登录")
	}
	return

}
