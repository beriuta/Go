package main
// 接口
// import (
// 	"fmt"
// )

	// type Usb interface {
	// 	// 定义一个接口
	// 	// 声明两个没有实现的方法
	// 	Start()
	// 	Stop()
	// }

	// type Computer struct {
	// 	// 计算机
	// }

	// type Phone struct {
	// 	// 定义一个手机结构体
	// }

	// func (p Phone) Start() {
	// 	fmt.Println("手机开始工作.......")
	// }

	// func (p Phone) Stop() {
	// 	fmt.Println("手机结束工作...")
	// }

	// func (c Computer) Working(usb Usb) {
	// 	usb.Start()
	// 	usb.Stop()
	// }

	// func main() {
	// 	// 创建结构体变量
	// 	computer := Computer{}
	// 	phone := Phone{}
	// 	computer.Working(phone)

	// }

// 接口的注意事项
	// 接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
	// 借口偶中所有的方法都没有方法体,即都没有实现的方法
	// 在Golang中,一个自定义类型需要将某个接口的所有方法都实现,我们说这个自定义类型实现了该接口
	// 一个自定义类型中只有实现某个接口,才能将该自定义类型的实例(变量)赋值给接口类型,如果接口有三个方法,而自定义实例只实现了两个方法,就会报错
	// 只要自定义数据类型,就可以实现接口,不仅仅是结构体类型
	// 一个自定义类型可以实现多个接口
	// Golang接口中不能有任何变量
	// 一个接口(比如A接口)可以继承多个别的接口(比如B,C接口),这时如果要实现A接口,也必须将B,C接口的方法全部实现
	// interface类型默认是一个指针(引用类型),如果没有对interface初始化就使用,那么会报错
	// 空接口interface{}没有任何方法,所以所有类型都实现了空接口
	// 如果是接口继承,那所继承的接口不能有相同的方法名
// 接口和继承
