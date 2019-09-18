package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatRoom/common/message"
	"go_code/chatRoom/server/model"
	"go_code/chatRoom/server/utils"
	"net"
)

// 创建一个user连接
type UserProcess struct {
	Conn net.Conn
	// 增加一个字段,表示该Conn是哪个用户的
	UserId int
}

// 编写通知所有在线用户的方法
// userId 这个用户要通知其他在线用户,我上线了
func (u *UserProcess) NotifyOtherOnlineUser(userId int) {

	// 遍历 onlineUser ,然后一个个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUser {
		if id == userId {
			// 过滤掉自己
			continue
		}
		// 开始通知[单独写一个方法], 这里的up是每一个在线的用户,给每一个在线的用户组装NotifyUserStatusMes结构物
		up.NotifyMeOnline(userId)
	}

}

func (u *UserProcess) NotifyMeOnline(userId int) {
	// 通知每个人我上线了
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// 将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyUserStatusMes序列化出错:", err)
		return
	}
	mes.Data = string(data)

	reData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes序列化出错:", err)
		return
	}

	// 创建Transfer实例
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WriterPkg(reData)
	if err != nil {
		fmt.Println("通知别人我上线的时候错误:", err)
		return
	}
}

// 编写一个函数serverProcessLogin,专门处理登录请求,变成UserProcess的方法
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 从Message中取出data,并直接反序列化成loginMes
	var loginMes message.LoginMes

	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("登录信息反序列化错误:", err)
		return
	}

	//  声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginReMesType

	// 声明一个 LoginResMes  服务端返回的登录结果
	var loginResMes message.LoginReMes

	// 需要到redis数据库中完成验证
	// 1. 使用model.MyUserDao,到redis验证
	user, err := model.MyUserDao.LoginCheck(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 // 该用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403 // 密码错误
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505 // 服务器内部错误
			loginResMes.Error = "服务器内部错误..."
		}

	} else {
		loginResMes.Code = 200
		// 用户登录成功,就把该登陆成功的用户放入userMgr中
		// 将登录成功的userId赋给u
		u.UserId = loginMes.UserId
		userMgr.AddOnlineUser(u)

		// 通知其他在线的用户,我上线了
		u.NotifyOtherOnlineUser(loginMes.UserId)

		// 将当前在线的用户的id,放入loginResMes.UserIds中,方便客户端拿到数据展示当前在线人数
		// 遍历userMgr.onlineUser
		for id, _ := range userMgr.onlineUser {
			// 将在线的id循环添加到登录返回的信息中,并赋值给登录的返回信息
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println("登录成功")
	}

	fmt.Println("user是什么:", user)

	// 将这个loginResMes序列化
	resData, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("服务端返回客户端登录序列化失败loginResMes", err)
		return
	}

	// 将resData赋值给resMes
	resMes.Data = string(resData)
	fmt.Println("服务端将什么数据返回给客户端了", resData)

	// 将resMes序列化
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("服务端返回客户端登录序列化失败resMes", err)
		return
	}
	// 发送data代码逻辑封装到WriterPkg函数中
	// 因为使用分层模式(MVC),所以需要先创建一个Transfer实例,然后读取
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WriterPkg(data)

	return

}

// 编写一个ServerProcessRegister,专门处理注册请求
func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	// 从Message中取出data,并直接反序列化成register
	var register message.RegisterMes

	err = json.Unmarshal([]byte(mes.Data), &register)
	if err != nil {
		fmt.Println("服务端注册信息反序列化错误:", err)
		return
	}

	//  声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterReMestype

	// 声明一个 RegisterReMes  服务端返回的注册结果
	var registerReMes message.RegisterReMes
	// 将接收到客户端的数据反序列化之后,传入处理注册的函数中将用户信息添加得到redis中
	err = model.MyUserDao.RegisterCheck(register.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			// 说明此用户已被注册
			registerReMes.Code = 505
			registerReMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerReMes.Code = 506
			registerReMes.Error = "注册发生未知错误"
		}
	} else {
		registerReMes.Code = 200
	}

	// 将回复的结构体反序列化,然后再反序列化整个resMes
	resData, err := json.Marshal(registerReMes)
	if err != nil {
		fmt.Println("服务端返回客户端登录序列化失败loginResMes", err)
		return
	}
	resMes.Data = string(resData)
	// 将resMes序列化
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("服务端返回客户端注册序列化失败resMes", err)
		return
	}
	// 发送data代码逻辑封装到WriterPkg函数中
	// 因为使用分层模式(MVC),所以需要先创建一个Transfer实例,然后读取
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WriterPkg(data)

	return

}
