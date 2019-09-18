package main
import (
	"fmt"
	"go_code/utils"
)

// func main() {
// 	// 使用自定义函数
// 	var nn float64 = 1.2
// 	var mm float64 = 2.3
// 	var op byte = '+'
// 	result := utils.Cal(nn, mm, op)
// 	fmt.Println(result)
// }

func main() {
	// 通过Set方法添加值
	acc := model.NewAccount(123456,80,"123475")
	// acc.model.SetNum(6)
	// acc.model.SetBalance(80)
	// acc.model.SetPassword(123456)
	// fmt.Println("此时的acc=", acc)
	if acc != nil {
		fmt.Println("注册成功")
	} else {
		fmt.Println("注册失败")
	}
}