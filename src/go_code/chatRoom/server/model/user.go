package model


// 定义一个用户的结构体
type User struct {
	UserId int `json:"userid"`  // 用户信息的json字符串的key,要跟结构体的tag字段一致
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`

}