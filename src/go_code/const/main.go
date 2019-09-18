package main

import (
	_"fmt"
)

/*
	常量介绍
		常量使用const修改
		常量在定义的时候,必须初始化
		常量不能修改
		常量只能修饰bool,数据类型(int,float系列), string类型
		语法:const identifier(变量名称) [type] = value
		举例说明,看看下面的写法是否正确
			const name = "tom"  // 对,常量本身有类型推导
			const tax float64 = 0.8  // 对
			const a int  // 错误,必须给值
			const b = 9 / 3  // 对的
			num := 3
			const c = 9 / num  // 错误,不能直接放变量
			const c = getVal()  // 错误, 编译器编译的时候,常量必须是确定的
	常量的写法
		const (
			a = iota  // iota表示给a赋值为0,b在a的基础上+1, c在b的基础上+1
			b
			c
		)
	Golang中没有常量名必须大写的规范
	仍然可以通过首字母的大小写来控制常量的访问范围
	面试题
		const (  // iota只在这个const起效
			a = iota
			b = iota
			c
			d,e = iota,iota  // 如果写在一行,值是相同的
		)
		fmt.Println(a,b,c,d,e)  // 0 1 2 3 3 
*/

func main() {
	
}