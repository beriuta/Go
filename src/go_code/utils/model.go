package model
import "fmt"

// // 为了让其他包的文件可以使用这个函数,需要把函数名首字母大写  该函数可导出
// func Cal(n1 float64, n2 float64, operator byte) float64 {
// 	// 定义一个变量承接结果
// 	var res float64
// 	switch operator {
// 		case '+':
// 			res = n1 + n2
// 		case '-':
// 			res = n1 - n2
// 		case '*':
// 			res = n1 * n2
// 		case '/':
// 			res = n1 / n2
// 		default:
// 			fmt.Println("操作符号错误")	
// 	}
// 	return res
// }

type account struct {
	num int  // 6-10之间
	balance int  // 余额>20
	password string // 密码必须6位
}

// 构建一个工厂函数
func NewAccount(num int, balance int, password string) *account {

	num_len := len(string(num))
	num_pwd := len(password)
	fmt.Println("密码个数=", num_pwd)
	if num_len > 10 || num < 6 {

		fmt.Println("输入的号码不在6-10范围之内,请重新输入")
		return nil
	}
	if balance < 20 {

		fmt.Println("余额小于20,请重新输入")
		return nil
	}
	if num_pwd != 6 {
		
		fmt.Println("密码只能输入6位,请重新输入")
		return nil
	}
	return &account{
		num:num,
		balance:balance,
		password:password,
	}
}

// func (account *account) SetNum(num int) {
// 	if num > 6 || num < 10 {
// 		account.num = num
// 	} else {
// 		fmt.Println("输入的号码不在6-10范围之内,请重新输入")
// 	}
// }

// func (account *account) SetBalance(balance int) {
// 	if balance > 20 {
// 		account.balance = balance
// 	} else {
// 		fmt.Println("余额小于20,请重新输入")
// 	}
// }

// func (account *account) SetPassword(password int) {
// 	num := len(string(password))
// 	if num == 20 {
// 		account.password = password
// 	} else {
// 		fmt.Println("密码只能输入6位,请重新输入")
// 	}
// }