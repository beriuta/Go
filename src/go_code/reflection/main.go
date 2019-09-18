package main

import(
	"fmt"
	// "reflect"
)
/*
	反射的使用
		1. 在结构体的标签的使用
		2. 函数的适配器,使用反射,也称为桥连接,bridge
	反射的基本介绍
		1. 反射可以在运行时动态获取变量的各种信息,比如变量的类型(type), 类别(kind)
		2. 如果是结构体变量,还可以获取到结构体本身的信息(包括结构体的字段,方法)
		3. 通过反射,可以修改变量的值,可以调用关联的方法
		4. 使用反射,需要import("reflect")
	反射重要的函数和概念
		1. reflect.TypeOf(变量名), 获取变量的类型,返回reflect.Type类型
		2. reflect.ValueOf(变量名), 获取变量的值,返回reflect.Value是一个结构体类型,从这个结构体里面可以取到很多变量的信息
		3. 变量,interface{}, 和 reflect.Value是可以相互转换的,这点在实际开发中会经常用到
		反射的转换
		定义一个测试函数,接收的是一个接口
		func test(b interface{}) {
			1. 如何将interface{}转换成reflect.Value
			rVal := reflect.ValueOf(b)
			2. 如何将reflect.Value转换成interface{}
			iVal := rVal.Interface()
			3. 如何将interface{}转换成源来的变量类型,需要用到类型断言
			v := iVal.(Stu)
		}
	反射的快速入门
		编写一个案例,演示对(基本数据类型, interface{}, reflect.Value)进行反射的基本操作
		反射的本质是在运行时,所以当reflect.Value转成interface之后,即便类型已经是结构体了,他也不能直接取值,需要再继续用类型断言
		因为,反射是在运行的时候确定的,而在编译的时候没办法判断,所以编译会不通过
		如果类型断言想要变得灵活,可以用swtich的形式
	反射注意事项和细节说明
		reflect.ValueOf.Kind, 获取变量的类型,返回的是一个常量
		reflect.TypeOf.Kind,  获取的值是一样的
		kind获取的是类别,type是类型,kind的范畴更大,如果是结构体,kind获取的是struct,type获取的某个结构体的类型Student类型
		Type是类型,Kind是类别,type和kind可能是相同的,也可能是不同的
		var num int = 10 num的Type是int,Kind也是int
		var stu Student stu的Type是 包名.Student, KInd是struct

		通过反射可以让变量在interface{}和Reflect.Value之间相互转换,
			变量 <---->interface{}<----->reflect.Value

		使用反射的方式来获取变量的值(并返回对应的类型),要求数据类型匹配,比如x是int,那么就应该使用reflect.Value(x).INt()而不能使用其他,否则会报panic

		通过反射可以修改变量,注意当使用SetXxx方法来设置需要通过对应的指针类型来完成,这样才能改变传入的变量的值,同时需要使用到reflect.Value.Elem()方法
	
*/

// func reflectTest(b interface{}) {
// 	// 此函数专门用来反射
// 	// 通过反射获取传入的变量的Type, kind 值
// 	// 1. 先获取reflect.Type
// 	rTyp := reflect.TypeOf(b)
// 	fmt.Println("rTyp=", rTyp)  // rTyp=int  r开头的是reflect类型的
// 	// 如果真的要看这个实际的类型
// 	fmt.Printf("rTyp的值是%v, rTyp的类型是%T\n", rTyp, rTyp)  // rTyp的值是int, rTyp的类型是*reflect.rtype

// 	// 2. 获取到reflect.Value
// 	rVal := reflect.ValueOf(b)
// 	fmt.Println("rVal=", rVal)  // rVal=100  不是真正原先的100,它本身的类型是reflect.Value
// 	// 如果想要拿出里面的值rVal.Int()
// 	n2 := rVal.Int() + 2
// 	fmt.Println("他们的和是:", n2)  // 他们的和是: 102
// 	// 将rVal转成interface{}
// 	iV := rVal.Interface()
// 	// 将 interface{}通过类型断言转成需要的类型
// 	num3 := iV.(int)
// 	fmt.Println("num3 = ", num3)  // num3 =  100
// }
// func main() {
// 	// 先定义一个int
// 	num := 100
// 	reflectTest(num)
// }
// func reflect01(b interface{}) {
// 	// 获取到reflect.Value
// 	rVal := reflect.ValueOf(b) 
// 	fmt.Printf("rVal的值:%v, rVal的类型是:%v\n", rVal, rVal.Kind())  // rVal的值:0xc00008e010, rVal的类型是:ptr
// 	// 取出rVal这个指针所代指的值,然后再用SetInt
// 	rVal.Elem().SetInt(20)
// 	// rVal.SetInt(20)  // 只有rVal是int类型的时候,才能直接用SetInt这个方法
// }
// func main() {
// 	// 通过反射,修改 num int 的值
// 	// 修改student的值
// 	num := 10
// 	reflect01(&num)  // 反射如果传入的是地址类型,必须要Elem
// 	fmt.Println("num=", num)  // num= 20
// }

// 一个变量 var v float64 = 1.2, 使用反射得到reflect.Value,
// 然后获取对应的Type,Kind类型,并将reflect.Value转换成interface{},再将interface{}转换成float64
// func reflect01(n interface{}) {
// 	rVal := reflect.ValueOf(n)
// 	// 先转成值类型
// 	fmt.Printf("type=%v, kind=%v\n",rVal.Elem().Type(), rVal.Elem().Kind())
// 	ret := rVal.Elem().Interface()
// 	ret2 := ret.(float64)
// 	fmt.Println("ret2=", ret2)
// 	// 注意,如果是传进来的是值类型,不能直接Setxxx,需要转变成reflect.Value的时候传入值的地址
// 	// 必须要用Elem().Setxxx()的方式转换值
// }
// func main() {
// 	var num float64 = 1.2
// 	reflect01(&num)
// }

// // 使用反射来遍历结构体的字段,调用结构体的方法,并获取结构体标签的值,用到的方法Method跟Call

// // 定义一个结构体,并且定义三个结构体的方法
// type Monster struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// 	Score float32
// 	Sex string
// }

// func (m Monster) Print() {
// 	fmt.Println("------start---------")
// 	fmt.Println(m)
// 	fmt.Println("------end-------")
// }

// func (m Monster) GetSum(n1, n2 int) int {
// 	return n1 +n2
// }

// // 接收四个值,给Monster赋值
// func (m Monster) Set(name, sex string, age int, score float32) {
// 	m.Name = name
// 	m.Age = age
// 	m.Score = score
// 	m.Sex = sex
// }

// func TestStruct(a interface{}) {

// 	// 获取reflect.Type类型
// 	typ := reflect.TypeOf(a)
// 	// 获取reflect.Value类型
// 	val := reflect.ValueOf(a)
// 	// 获取a对应的类别
// 	kd := val.Kind()
// 	if kd != reflect.Struct{
// 		// 如果确定不是结构体,直接返回
// 		fmt.Println("expect struct")
// 		return
// 	}
// 	// 获取该结构体有几个字段,返回的是数字
// 	num := val.NumField()
// 	// fmt.Println("num=", num)  // num= 4 
// 	// 遍历结构体的所有字段\
// 	for i := 0; i < num; i++ {
// 		// 取结构体的字段是val.Field, 返回的是value类型
// 		fmt.Printf("Field %d: 值=%v\n", i, val.Field(i))
// 		// 获取结构体的标签,也就是json使用的名字,需要通过reflect.Type来获取tag
// 		// 用reflect.Type.Field,返回的数据类型structField类型,这是一个结构体,
// 		// 这个结构体里面有Tag这个字段,这个字段对应的数据类型是一个结构体StructTag,进入这个结构体就会发现有个Get方法,这个方法返回的才是string类型的数据
// 		tagVal := typ.Field(i).Tag.Get("json")
// 		if tagVal != "" {
// 			// 如果能获取到json的标签,就打印出来
// 			fmt.Printf("Field=%d, tag=%v\n", i, tagVal)
// 		}
// 	}

// 	// 查看这个结构有多少方法,用reflect.Value来取
// 	numOfMethod := val.NumMethod()
// 	fmt.Printf("此结构体的方法有%d个\n", numOfMethod)  // 此结构体的方法有3个
// 	// 用反射调用结构体中的方法,reflect.Value类型来调用, 注意方法得到排序是按照函数名的ASCII码进行排序的,不是按写的方法顺序排
// 	val.Method(1).Call(nil)  // 这是调用的按函数名的ASCII码排序的第二个方法, Call方法接收的是reflect.Value切片,返回的还是reflect.Value类型的切片
// 	// 调用结构体的第一个方法
// 	// 先声明一个reflect.Value的切片
// 	var params []reflect.Value
// 	params = append(params, reflect.ValueOf(10))  // 将一个常量变成一个reflect.Value类型,放到切片中
// 	params = append(params, reflect.ValueOf(20))  // 再在切片放一个值
// 	res := val.Method(0).Call(params)  // 第一个方法里面需要传值的,传入的是[]reflect.Value类型
// 	// res也是一个[]reflect.Value
// 	fmt.Println("res的值:", res[0].Int())  // 将结果返回,返回的结果是一个切片,从切片中获取值,再转成相应的数据类型,如果类型不确定,还需要加一个类型断言

// }

// func main() {
// 	// 实例化一个Monster对象
// 	neuf := Monster{
// 		Name: "牛大人",
// 		Age: 1000,
// 		Score: 10000.12,
// 	}
// 	TestStruct(neuf)
// }
// 使用反射的方式来获取结构体的tag标签,遍历字段的值,修改字段值,调用结构体方法(要求:通过传递地址的方式完成)
// type Student struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// 	Address string `json:"address"`
// }

// func reflectTest(r interface{}) {

// 	// 获取到值
// 	val := reflect.ValueOf(r)
// 	// 获取到类型
// 	kd := val.Elem().Kind()
// 	// 获取到类型
// 	typ := reflect.TypeOf(r)
// 	fmt.Println(kd)
// 	// 判断是否是结构体
// 	if kd != reflect.Struct {
// 		fmt.Println("该变量不是结构体")
// 		return
// 	}
// 	// 获取结构体字段的个数
// 	numFields := val.Elem().NumField()
// 	for i := 0; i < numFields; i++ {
// 		fmt.Printf("此结构体的第%d个字段是%v\n", i, val.Elem().Field(i))
// 		// reflect.Value是没有Tag这个方法的,所以还要用reflect.Type才行
// 		// fmt.Println("json字段", val.Elem().Field(0).Tag.Get("json"))
// 		fmt.Println("json字段", typ.Elem().Field(i).Tag.Get("json"))
		
		
// 	}
// 	// // 先转成interface,然后用类型断言转成结构体,然后再改值
// 	// newVal := val.Elem().Interface()
// 	// ret := newVal.(Student)
// 	// ret.Age = 20
// 	// fmt.Println(newVal)
// 	// fmt.Println(ret)
// 	// 修改值
// 	val.Elem().Field(0).SetString("白香精")

	
// }

// func main() {

// 	student := Student{
// 		Name: "韩蕾",
// 		Age: 10,
// 		Address: "新华路新安街道小胡同",
// 	}
// 	reflectTest(&student)
// 	fmt.Println("原始的结构体", student)
// }

// 适配器
// 定义了两个函数test1和test2,定义一个适配器函数用作统一处理接口
func main() {
	fmt.Println("wahaha")
}