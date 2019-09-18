package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatRoom/common/message"
	"go_code/chatRoom/client/utils"
)

type SmsProcess struct {
}

// 发送群聊消息
func (s *SmsProcess) SendGroupMes(content string) (err error) {

	// 1. 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. 创建一个SmsMes实例对象
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUser.UserId
	smsMes.UserStatus = curUser.UserId

	// 3. 序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化消息出错:", err)
		return
	}
	mes.Data = string(data)
	// 4. 序列化最终消息
	reData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化最终消息数据错误:", err)
		return
	}
	// 5. 将mes发送到服务器
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	// 6. 发送
	err = tf.WriterPkg(reData)
	if err != nil {
		fmt.Println("发送消息包到服务端出错:", err)
		return
	}
	return
}
