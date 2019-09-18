package cal

import (
	// "fmt"
	"reflect"
	"testing"
)

// // 一个被测试的函数
// func addUpper(n int) int {
// 	res := 0
// 	for i := 1; i <= n; i++ {
// 		res += i
// 	}
// 	return res
// }

// func main() {
// 	// 单元测试的传统方式测试, 在main函数中执行被测试的函数看结果是否正确
// 	res := addUpper(10)
// 	if res != 55 {
// 		fmt.Println("函数正确")
// 	} else {
// 		fmt.Println("函数正确")
// 	}
// 	// 传统测试方法不方便,我们需要修改main函数,如果项目正在运行,就会停止项目中
// 	// 不利于管理,如果有两个函数需要在main函数中注销
// 	// 单元测试-基本介绍
// 		// go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试,testing框架和其他语言中的测试框架类似,可以基于这个框架
// 		// 写针对响应函数的测试用例,也可以基于该框架写相应的压力测试用例,通过单元测试,可以解决如下问题:
// 			// 1, 确保每个函数是可运行的,并且运行结果正确
// 			// 2, 确保写出来的代码性能是好的
// 			// 3, 单元测试能即使的发现程序的实际或实现的逻辑错误,使问题及早暴露,便于问题的定位解决,而性能测试的重点在于发现程序设计上的一些问题,让程序能够在高并发的情况下还能保持稳定
// 测试文件名必须以_test.go为结尾,
// 测试用例函数必须以Testxxxx开头
// Testxxxx(t *testing.T)的形参类型必须是*testing.T
// 一个测试文件可以有多个测试函数
// 运行测试用例的命令:
// go test ==> 运行正确无日志,运行错误,会输出日志
// go test -v  ===> 运行正确或是错误都会输出日志
// 当出现错误时,可以使用t.Fatalf来格式化输出错误日志信息
// t.Logf可以输出相应的日志
// 测试用例函数并没有放在main函数中,也执行了
// pass表示全部成功, fall表示某一个失败
// 测试单个文件,一定要带上被测试的原文件
// go test -v 测试文件名
// 测试单个方法
//go test -v -test.run TestSddUppe(方法名)

// 单元测试总和案例要求
// 编写一个Monster结构体,字段Name,Age,Skill
// 给Monster绑定方法Store,可以将一个Monster变量(对象),序列化后保存到文件中
// 给Monster绑定方法ReStore,可以将一个序列化的Monster,从文件中读取,并反序列化为Monster对象
// 变成测试用例文件 store_test.go,编写测试用例函数TestStore和TestRestore进行测试

// }

// func TestAddUpper(t *testing.T) {
// 	// 此函数的前缀是Test后面紧跟要测试的函数名,要测试的函数名需要大写
// 	// 调用
// 	res := addUpper(10)
// 	if res != 55 {
// 		t.Fatalf("addUpper(10) 执行错误, 期望值=%v, 实际值=%v\n", 55, res)
// 	}

// 	// 如果正确,输出日志
// 	t.Logf("addUpper(10), 执行正确......")
// }
// // 单元测试的时候的终端命令是 go test -v

// func TestReflectFunc(t *testing.T) {
// 	call1 := func(v1, v2 int) {
// 		t.Log(v1, v2)
// 	}
// 	call2 := func(v1,v2 int, s string) {
// 		t.Log(v1, v2, s)
// 	}
// 	var (
// 		function reflect.Value
// 		inValue []reflect.Value
// 		n int
// 	)
// 	// 桥连接
// 	bridge := func(call interface{}, args...interface{}) {
// 		n = len(args)
// 		inValue = make([]reflect.Value, n)
// 		for i := 0; i < n; i++ {
// 			inValue[i] = reflect.ValueOf(args[i])
// 		}
// 		function = reflect.ValueOf(call)
// 		function.Call(inValue)
// 	}
// 	bridge(call1, 1, 2)
// 	bridge(call2, 1, 2, "test2")
// }
type user struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model)                             // 获取类型 *user
	t.Log("reflect.TypePf", st.Kind().String())            // ptr
	st = st.Elem()                                         // 取出结构体对应的对象,而不是指针
	t.Log("refelct.TypeOf.Elem", st.Kind().String())       // 这时是struct
	elem = reflect.New(st)                                 // New返回一个Value类型值,该值持有一个执行类型为type的新申请的零值的指针
	t.Log("refelct.New(st)", elem.Kind().String())         // ptr
	t.Log("refelct.New.Elem", elem.Elem().Kind().String()) // struct
	// model 就是创建的user结构体实例
	model = elem.Interface().(*user)                   // 类型断言,model是*user指向和elem一样
	elem = elem.Elem()                                 // 取到elem指向的值
	elem.FieldByName("UserId").SetString("1282828282") // 赋值
	elem.FieldByName("Name").SetString("哇哈哈")          // 赋值
	t.Log("model model.Name=", model, model.Name)

}
