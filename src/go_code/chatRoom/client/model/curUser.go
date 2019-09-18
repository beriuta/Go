package model


import (
	"net"
	"go_code/chatRoom/common/message"
)


// 在客户端会有很多地方使用到CurUser
type CurUser struct {
	Conn net.Conn
	message.User
}