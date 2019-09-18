package monster

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

// 给monster绑定方法Store,可以将一个monster变量(对象),序列化后保存到文件中
func (m *Monster) Store() bool {
	// 先序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("错误信息是:", err)
		return false
	}

	// 写入文件
	filePath := "/home/beriuta/Desktop/新建文本.txt"
	ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("写入文件出错:", err)
		return false
	}
	return true
}

// 反序列化
func (m *Monster) ReStore() bool {

	// 先到文件中读取字符串
	filePath := "/home/beriuta/Desktop/新建文本.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件出错:", err)
		return false
	}

	// 使用读取data[]byte, 对此反序列化
	err1 := json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("反序列化失败:", err1)
		return false
	}

	return true 
}