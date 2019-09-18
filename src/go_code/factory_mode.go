/*
	工厂模式:
	Golang的结构体没有构造函数,通常可以使用工厂模式来解决这个问题
	   
*/
package main
// import (
// 	"fmt"
// )


	// type Account struct {
	// 	Username string
	// 	Password string
	// 	Money float64
	// }

	// // 存款
	// func (account *Account) Deposite(money float64, pwd string) {

	// 	// 校验密码是否正确
	// 	if pwd != account.Password {
	// 		fmt.Println("密码不正确,请重新输入")
	// 		return
	// 	}

	// 	// 校验输入的钱数是否符合实际情况
	// 	if money <= 0 {
	// 		fmt.Println("输入金额不正确,请重新输入")
	// 		return
	// 	}

	// 	// 如果上述都校验成功,则将该账户的金额加上现存入的金额
	// 	account.Money += money
	// 	fmt.Println("存款成功")
	// 	fmt.Printf("存款现有%v元\n", account.Money)
	// }	

	// func main() {
	// 	num1 := Account{
	// 		Username:"韩蕾", Password:"qaz123wsx",
	// 		Money:23456,
	// 	}
	// 	num1.Deposite(3838, "qaz123wsx")

	// }

// 面向对象的三大特征
// Golang仍然有面向对象变成的继承,封装和多态的特性,只是实现的方式和传统OOP语言不同
	// 封装就是把抽象出的字段和对字段的操作封装在一期,数据被保护在内部,程序的其它包只有通过被授权的操作(方法),才能对字段进行操作
	// 实现的方式 就是将结构体,字段(属性)的首字母小写
// 课堂练习
// 创建程序,在Model包中定义Account结构体, 在main函数中体会Golang的封装性
// 要求:Account结构体要求具有字段: 账号(长度在6-10之间), 余额(必须>20), 密码(必须是6位)
	// 通过Setxxx方法给Account的字段赋值

