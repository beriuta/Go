package message

// 定义消息类型

const (
	LoginMesType            = "loginMes"            // 登录的消息类型
	LoginReMesType          = "loginReMes"          // 登录返回的消息类型
	RegisterMesType         = "RegisterMes"         // 注册的消息类型
	RegisterReMestype       = "RegisterReMes"       // 注册返回的消息类型
	NotifyUserStatusMesType = "NotifyUserStatusMes" // 服务器主动推送消息
	SmsMesType              = "SmsMes"
)

// 定义一些客户的状态,在线,离线,发呆,隐身
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"data"` // 消息的内容
}

type LoginMes struct {
	UserId   int    `json:"userid"`   // 用户的id
	UserPwd  string `json:"userpwd"`  // 用户的密码
	UserName string `json:"username"` // 用户名
}

type LoginReMes struct {
	// 返回给客户端的消息
	Code    int    `json:"code"`    // 返回状态码 500 表示用户未注册 200 表示登陆成功
	UserIds []int  `json:"userids"` // 增加字段,保存用户id的切片
	Error   string `json:"error"`   // 返回错误信息

}

type RegisterMes struct {
	// 注册的一些字段
	User User `json:"User"` // 类型就是User的结构体
}

type RegisterReMes struct {
	// 返回给客户端的消息
	Code  int    `json:"code"`  // 返回状态码 200表示注册成功, 400是已经注册了
	Error string `json:"error"` // 返回错误信息

}

// 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userid"` // 用户id
	Status int `json:"status"` // 用户状态
}

// 增加一个SmsMes发送消息
type SmsMes struct {
	Content string `json:"content"`
	User           // 匿名结构体
}
