package model

import (
	"fmt"
)


// 声明一个Customer结构体,表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Emaile string

}


// 一个方法,返回的是一个客户的信息,字符串
func(this Customer) GetInfo() string {
	Info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", 
	this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Emaile)
	return Info
}



// 编写一个工厂模式,返回一个Customer的实例
func NewCustomer(name string, gender string, age int,
	phone string, emaile string) Customer {
	return Customer{
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Emaile:emaile,
	}

}

