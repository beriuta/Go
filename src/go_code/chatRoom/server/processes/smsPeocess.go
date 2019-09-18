package processes

import (
	"go_code/chatRoom/common/message"
	"go_code/chatRoom/server/utils"
	"net"
	"fmt"
	"encoding/json"
)

type SmsProcess struct {
}

// 转发消息
func (s *SmsProcess) SendGroupServerMes(mes *message.Message) {

	// 遍历服务器端的onlineUsers map[int]*UserProcess
	// 将消息转发到相应的客户端
	// 取出mes的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("服务器发序列化转发消息错误:", err)
		return
	}
	data, err := json.Marshal(mes)
	for id, up := range userMgr.onlineUser {
		if id == smsMes.UserId {
			continue
		}
		s.SendMesToEachOnline(data, up.Conn)
	}
}

func (s *SmsProcess) SendMesToEachOnline(info []byte, conn net.Conn) {

	// 创建一个Transfer实例,发送data
	tf := &utils.Transfer{
		Conn:conn,
	}
	err := tf.WriterPkg(info)
	if err != nil {
		fmt.Println("转发消息失败:", err)
		return
	}
}
