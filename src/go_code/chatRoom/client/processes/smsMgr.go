package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatRoom/common/message"
)

func outputGroupMes(mes *message.Message) (err error) {

	// 反序列化mes.Data
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化转发过来的消息失败:", err)
		return
	}

	// 显示信息
	info, err := fmt.Scanf("用户id:%d\t对大家说:%s\t", smsMes.UserId, smsMes.Content)
	if err != nil {
		fmt.Println("显示信息错误:", err)
		return
	}
	fmt.Println(info)
	fmt.Println()
	return
}
