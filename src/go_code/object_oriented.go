/*
	面向对象编程(OOP)
	不是传统意义的面向对象
	是结构体
	结构体(变量)(结构体对象):
		Golang也支持面向对象编程(oop),但是和传统的面向对象编程的区别,并不是纯粹的面向对象语言,所以我们说Golang支持面向对象变成特性
		Golang没有类(class)的概念,Go语言的结构体(struct)和其他变成语言的类(class)有同等低位,可以理解为Golang是基于struct来实现OOP特性的
		Golang面向对象编程非常简洁,去掉传统OOP的继承,方法重载,构造函数和析构函数,隐藏的self等等
		Golang仍然有面向对象编程的继承,封装和多态的特性,只是实现的方式和其他OOP语言不一样,比如继承,golang的继承是通过匿名字段来实现的
		Golang面向对象(OOP)很优雅,OOP本身就是语言类型系统的一部分,通过接口(interface)关联,耦合性低,也非常灵活,所以在Golang中面向接口编程是非常重要的特征
*/
package main
// import "fmt"
// import "encoding/json"  // 引入json包
	// // 定义一个struct结构体,是管理猫的种类
	// type Cat struct {
	// 	Name string
	// 	Age int
	// 	Color string
	// }

	// func main() {
	// 	// 实例化这个猫的结构体
	// 	var cat1 Cat

	// 	// 给这个猫定义内容
	// 	cat1.Name = "小花"
	// 	cat1.Age = 18
	// 	cat1.Color = "花色"
	// }

// 结构体是自定义的数据类型,代表一类事物
// 结构体变量(实例)是具体的,实际的,代表一个具体变量

// 结构体变量(实例)在内存的布局,是值类型

// 基本结构体

	// type 结构体名字 struct {
	// 	field1 type  字段名称1
	// 	field2 type  字段名称2
	// }
// 字段的结构体是一个组成部分,一般是基本数据类型,数组,也可以是引用型,

// 结构体注意事项和细节说明
// 1. 字段声明语法同变量,示例:字段名 字段类型
// 2. 字段类型可以为基本类型,数组 或引用类型
// 3. 在床架一个结构体变量后,如果没有给字段赋值,都对应一个零值(默认值)
//    如果结构体是map,slice,指针,nil 需要先声明空间才能使用,如果没有make,就是没有分配空间
// 4. 不同结构题变量的字段是独立的,互不影响,一个结构体变量字段的更改不影响另一个

// 字段的声明
	// type Person struct {
	// 	Name string
	// 	Age int
	// }

	// func main() {
	// 	// 方式二
	// 	p2 := Person{"tome", 15}
	// 	fmt.Println("打印出来的方式二", p2)  // 打印出来的方式二 {tome 15}
	// 	fmt.Printf("此人的名字是%v\n", p2.Name)  // 此人的名字是tome
	// 	fmt.Printf("此人的年龄是%v\n", p2.Age)  // 此人的年龄是15

	// 	// 利用类型推导赋值一个类型指针
	// 	p3 := Person{"mary", 19}
	// 	fmt.Println("此时的p3:", p3)  // 此时的p3: &{mary 19}
	// 	var p4 *Person = new(Person)  // 这样等同于类型推导
	// 	fmt.Println("此时的p4:", p4)  // 此时的p4: &{ 0}
	// 	// 因为p3,p4是一个指针,因此标准给字段赋值方式
	// 	(*p3).Name = "哇哈哈" // p3.Name = "哇哈哈"
	// 	(*p3).Age = 20  // p3.Age = 20   这种方式也可以,推荐使用不带括号的
	// 	(*p4).Name = "乳娃娃"
	// 	(*p4).Age = 18
	// 	fmt.Println("赋值之后的p3:", p3)  // 赋值之后的p3: &{哇哈哈 20}
	// 	fmt.Println("赋值之后的p4:", p4)  // 赋值之后的p4: &{乳娃娃 18}

	// 	// 第四种方式
	// 	var n5 *Person = &Person{}
	// 	n5.Name = "scott"
	// 	n5.Age = 209
	// 	fmt.Println("赋值之后的n5:", n5)

	// }

// 结构体的所有与字段在内存中是连续的

	// // 定义两个结构体
	// type Point struct {
	// 	x int
	// 	y int
	// }
	// type Rect struct {
	// 	leftUp, rightDown Point
	// }
	// func main() {

	// 	// r1里面有四个int, 在内存中是连续分布的
	// 	r1 := Rect{Point{1,2}, Point{3,4}}
	// 	// 打印地址
	// 	fmt.Printf(
	// 		"r1.lsefUp.x的地址:%p\nr1.leftUp.y的地址%p\nr1.rightDown.x的地址%p\nr1.rightDown.y的地址%p\n",
	// 		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	// 		// r1.lsefUp.x的地址:0xc000014340
	// 		// r1.leftUp.y的地址0xc000014348
	// 		// r1.rightDown.x的地址0xc000014350
	// 		// r1.rightDown.y的地址0xc000014358
	// 	// r2有两个Point,两个*Point类型本身的地址是连续的,但是不代表他们里面的内容地址是连续的
	// 	fmt.Printf("r2.leftUp的地址:%p\nr2.rightDown的地址:%p", &r2.leftUp, &r2.rightDown)

	// }

// 结构体的注意事项和使用细节
	// 结构体是用户单独定义的类型,和其他类型进行转换需要有完全相同的字段(名字,个数和类型)
		// type A struct {
		// 	Num int
		// }
		// type B struct {
		// 	Num int
		// }
		// func main() {

		// 	var a A
		// 	var b B
		// 	// a = b  // 这样直接等于是不行的,会报错
		// 	a = A(b)  // 需要强转型一下
		// 	fmt.Println("转换成功")
		// 	fmt.Println(a,b)  // {0} {0}
		// }
	// 结构体进行type重新定义之后(相当于取了别名),Golang认为是新的数据类型,但是,相互间可以强传
		// type Student struct {
		// 	Name string
		// 	Age int
		// }
		// type Stu Student

		// func main() {
		// 	var stu1 Student
		// 	var stu2 Stu
		// 	// stu1 = stu2  // 这是错误的
		// 	stu1 = Student(stu2)
		// 	fmt.Println("转型之后", stu1)  // 转型之后 { 0}
		// }
// 结构体的注意事项和使用细节
	// struct的每个字段上,可以写一个tag,该tag可以通过反射机制获取,常见的使用场景就是序列号和反序列化
	// 代码说明:将struct变量进行json处理
	// 问题:json处理后得到字段也是首字母大写,如果将json字符串返回给其他程序该如何处理
	// 用tag解决

		// type Monster struct {
		// 	Name string `json:"name"`
		// 	Age int `json:"age"`
		// 	Skill string `json:"skill"`
		// }

		// func main() {

		// 	monster := Monster{"牛魔王", 500, "吹牛~~"}

		// 	// 将monster 变量序列化为json字符串
		// 	jsonMonster, err := json.Marshal(monster)  // 这个函数中使用了反射
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	// fmt.Println("json数据", string(jsonMonster))  // json数据 {"Name":"牛魔王","Age":500,"Skill":"吹牛~~"}
		// 	// 但是打印的是大写的,所以需要在结构体上添加如果是json数据,就使用后面的字段名
		// 	// 这时打印的就是小写
		// 	fmt.Println("json数据", string(jsonMonster))  // json数据 {"name":"牛魔王","age":500,"skill":"吹牛~~"}
		// }
// 方法的基本介绍
// 在某些情况下,我们需要声明(定义)方法,比如Person结构体,除了一些字段外,还会有一些方法,比如学习,跑步之类的,这时就要用到方法才能完成
// golang中的方法是作用在指定的数据类型(和指定的数据类型绑定),因此自定义类型都可以有方法,而不仅仅是struct
	// type Person struct {
	// 	Name string
	// }
	// // 定义一个A结构体的方法,方法名为test,只能被A来调用
	// func (p Person) test() {
	// 	fmt.Println("test函数:", p.Name)
	// }
	// func main() {

	// 	var person Person
	// 	person.Name = "tom"
	// 	person.test()  // 调用方法
	// }
// 总结说明: 
	// test方法和Person类型绑定
	// test方法只能通过Person类型的变量来调用
	// person表示,哪个Person变量调用,person就是哪个,就跟函数传参一样
	// person这个值不是固定的,相当于是形参
// 方法快速入门
	// 1,给Person结构体添加speak方法,输出xxx是一个好人
	// 2,给Person结构体添加compute方法,可以计算从1+...1000的结果
	// 3,给Person结构体添加compute2方法,该方法接收一个数n,计算从1+...n
	// 4,给Person结构体添加getSum方法,可以计算两个数的和,并返回结果


		// type Person struct {
		// Name string
		// Age int
		// }

		// func (person Person) speak() {
		// fmt.Println("xxx是一个好人")
		// }

		// func (person Person) compute() {
		// sum := 0
		// for i := 1; i < 1000; i++ {
		// 	sum += i
		// }
		// fmt.Println("1000的和:", sum)
		// }

		// func (person Person) compute2(n int) {
		// // 接收n
		// sum2 := 0
		// for i := 1; i < n; i++ {
		// 	sum2 += i
		// }
		// fmt.Println("sum2的和:", sum2)
		// }

		// func (person Person) getSum(x, y int) int {

		// numSum := x + y
		// return numSum
		// }

		// func main() {
		// p_student := Person{"娃娃", 10}
		// p_student.speak()
		// p_student.compute()
		// p_student.compute2(10)
		// ret := p_student.getSum(55,89)
		// fmt.Println("他俩的和为:", ret)

		// }
// 方法的调用和传参机制
	// type Circle struct {
	// 	radius float64
	// }

	// func (cricle Circle) area() float64 {
	// 	return 3.14 * cricle.radius * cricle.radius
	// }
	// 这个方法传递的是结构体的地址
	// func (cricle *Circle) area2() float64 {
	// 	return 3.14 * cricle.radius * cricle.radius
	// }

	// func main() {

	// 	// 输入这个结构体的半径
	// 	cll := Circle{3.5}
	// 	res := cll.area()
	// 	fmt.Println("此圆的面积是:", res)  // 此圆的面积是: 38.465
	// }

// 方法的声明(定义)
	// func (receiver type) methodName (参数列表) (返回值列表){ 方法体 return 返回值}
	// 参数列表:表示方法输入
	// receiver type : 表示这个方法和type这个类型进行绑定,或者说该方法作用于type类型
	// receiver type: 可以是结构体,也可以是其他自定义类型
	// receiver: 表示type类型的一个变量(实例), 比如 Person结构体的一个变量(实例)
	// 参数列表: 表示方法输入
	// 返回值列表: 表示返回值,可以是多个,一个可以不加()
	// 方法主体: 逻辑代码块
	// return : 语句不是必须的

// 方法注意事项和细节讨论
	// 1, 结构体类型是值类型,在方法调用中,遵守值类型的传递机制,是值拷贝传递方式
	// 2, 如果程序员希望在方法中,改变结构体变量的值,可以通过结构体指针的方式来处理
	// 3, Golang中的方法作用在指定的数据类型上(即: 和制定的数据类型绑定)
	// 4, 方法的访问范围控制的规则,和函数一样,方法名首字母小写,只能在本包访问,方法首字母大写,可以在本包和其他包访问
	// 5, 如果一个变量实现了String()这个方法,那么fmt.Prinln()默认会调用这个变量的string()进行输出
// 方法练习
	// 编写结构体,不需要传参,打印一个10*8的矩形,在main函数中调用这个方法
		// type Rectangle struct {
		// 	width float64
		// 	long float64
		// }
		// func (rectangle *Rectangle) test() {
		// 	// 打印一个矩形
		// 	for i := 1; i < 11; i++ {
		// 		for j := 0; j < 8; j++ {
		// 			fmt.Print("*")
		// 		}
		// 		fmt.Print("\n")
		// 	}
			
		// }

		// func main() {
		// 	re := Rectangle{1,3}
		// 	re.test()
		// }

		// type Student struct {
		// 	Name string  `json:"name"`
		// 	Gender string `json:"gender"`
		// 	Age int `json:"age"`
		// 	Id int `json:"id"`
		// 	Score float64 `json:"score"`
		// }

		// func (student *Student) say() string {
		// 	students, err := json.Marshal(student)
		// 	if err != nil {
		// 		return "序列化出错"
		// 	}
		// 	return string(students)
		// }

		// func main() {
		// 	stu1 := Student{"tom", "男", 18, 00213, 97}
		// 	respons := stu1.say()
		// 	fmt.Println("返回的字符串为:", respons)  // 返回的字符串为: {"name":"tom","gender":"男","age":18,"id":139,"score":97}
		// }

// 面向对象编程继承
// 如果一个struct嵌套了另一个匿名结构体,那么这个结构体可以直接访问匿名结构体的字段和方法,从而实现了继承特性
	// type A struct {
	// 	Name string
	// 	Age int
	// }
	// type B struct {
	// 	a A  // 有名结构体  如果访问有名结构体的字段,必须要指定结构体的命名
	// 	A  // 匿名结构体
	// }







		