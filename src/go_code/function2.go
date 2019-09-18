// 此文档是练习Go语言函数跟包
// 函数定义
	// func 函数名 (形参列表) (返回值列表) {
		// 执行语句...
		// return 返回值列表
	// } 
// 形参列表,表示函数的输入
// 函数中的语句,表示为了实现某一功能代码块
// 函数可以有返回值,也可以没有
package main
// import (
// 	"fmt"
// 	// "strings"
// )

// 将计算的功能放到一个函数中,然后 在使用的时候直接调用即可
	// func acl(n1 float64, n2 float64, operator byte) float64 {
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


// func main() {
// 	// 使用自定义函数
// 	var nn float64 = 1.2
// 	var mm float64 = 2.3
// 	var op byte = '+'
// 	result := acl(nn, mm, op)
// 	fmt.Println(result)
// }

// 包的介绍
	// 在实际的而开发中,我们往往需要在不同文件中,去调用其他文件的定义的函数,比如,main.go中,去使用utils.go文件中的函数
	// 每个文件夹就是一个包,文件夹是相应的go文件,在每个go文件中写函数
	// go的每一个文件都属于一个包,也就是说go是以包的形式来管理文件和项目目录结构的

// 包的三大作用
// 1. 区分相同名字的函数,变量等标识
// 2. 当程序文件很多时,可以很好的管理项目
// 3. 控制函数.变量等访问范围,即作用域

// 打包基本语法
// package util

// 引入包的基本语法
// import "包的路径"

// 包的快速入门
// 说明: 在包内相互调用函数,我们将funcCal定义到文件utils.go中,将utils.go放到一个包中,当其他文件需要使用到utils.go的方法时,可以import该包,就可以使用了

// 包的注意事项
	// 1. 在给一个文件打包时,该包对应一个文件夹,文件包名通常跟文件夹名一致,一般为小写字母
	// 2. 当一个文件要使用其他包函数或者变量的时候,需要先引入对应的包
	// 3. package指令在文件第一行,然后是import指令
	// 4. 在import包时,路径从$GOPATH的src下开始,不用带src,编译器或自动从src下开始引用
	// 5. 如果是需要被别的包引用的函数以及变量,名字首字母都要大写,如果不大写就说明是私有变量或函数 
	// 6. 如果包名过长,可以令取名字 比如"go_code/chapter06/function/utils" 如果取别名  util "go_code/chapter06/function/utils"
	//    如果取了别名之后,之前的名字就不能使用了
	// 7. 在同一包下,不能有相同的函数名,否则报重复定义错误, 如果相同的文件夹下面不同的文件,也不允许定义相同的函数名或者变量名
	// 8. 如果要编译成可执行文件,需要将这个包声明为main,即package main这个就是一个语法规范,如果你是写一个库,包名可以自定义

// 编译一个项目的方式
// 1. 要到项目所在的文件中也就是$GOPATH中去编译
// 2. 编译时需要编译main包所在的文件夹
// 3. 项目的目录结构最好按照规范来组织
// 4. 编译后生成一个有默认名的可执行文件,在$GOPATH目录下,可以指定名字和目录,比如:放在bin目录下
// 5. 编译指令: go build -o bin 项目名.exe go_code/main

// 函数-调用机制
// ===调用过程

// // 编写一个函数
// func test(n1 int) {

// 	n1 = n1 + 1
// 	fmt.Println("test() n1=", n1)  // 11
// }

// func getSum(n1 int, n2 int) int {
	
// 	return n1 + n2
// }

// func getSumAndSub(n1 int, n2 int) (int, int) {

// 	sum := n1 + n2
// 	sub := n1 - n2
// 	return sum, sub
// }

// func main() {

// 	n1 := 10
// 	test(n1)  // 函数执行完毕之后,系统会把开辟的空间回收
// 	fmt.Println("main() n1=", n1)  // 10  这两个n1虽然一样,但是就像两个班级里面有同名的同学,虽然名字一样,但是所代表的人是不一样的

// 	sum := getSum(10, 20)
// 	fmt.Println("他俩的和是~~~:", sum)

// 	sum1, sub1 := getSumAndSub(30, 20) 
// 	fmt.Println("他俩的和:", sum1)
// 	fmt.Println("他俩的差:", sub1)

// 	// 希望忽略某个值,可以使用_占位
// 	_, res := getSumAndSub(444, 44444)
// 	fmt.Println("他们的差:", res)

// }

//对上层代码的解释
	// 调用一个函数时,会给该函数分配一个新的空间,编译器会通过自身的处理,range这个新空间和其他的栈空间区分开来
	// 在每个函数对应的栈中,数据空间是空间的,不会混淆
	// 当一个函数调用完毕(执行完毕)后,程序会销毁这个函数对应的栈空间

// golang函数调用的机制的底层分析
	// 栈区,堆区,代码区
	// 栈区:基本的数据类型,一般说分配到栈区,因为在go中编译器存在一个逃逸分析
	// 堆区:引用数据类型
	// 代码区:代码本身存在这里

// 函数---递归调用
	// 基本介绍
		// 一个函数在函数体内又调用了本身,称为递归函数
	// 递归调用的重要原则
		// 执行一个函数时,就创建一个新的受保护的独立空间(新函数栈)
		// 函数的局部变量是独立的,不会互相影响
		// 递归必须向提出递归的条件逼近,否则就是无线递归,死龟了
		// 当一个函数执行完毕,或者遇到return,就会返回,遵守谁调用,就将结果返回给谁,集合石当函数执行完毕或者返回时,该函数本身也会被系统销毁
			// func tete1(n int) {
			// 	if n > 2 {
			// 		n--
			// 		tete1(n)
			// 	} 
			// 	fmt.Println("n= ", n)
			// }

			// func main() {

			// 	tete1(4)
			// }
	// 递归练习
		// 使用递归的方式,求出斐波那契数1,1,2,3,5,8,13...第一个数是1,第二个数是1,第三个数是前面两个数的和,依次类推
		// 给你一个整数n,求出它的值是多少
			// func fbn(n int) int {
			// 	if (n == 1 || n == 2) {
			// 		return 1
			// 	} else {
			// 		return fbn(n - 1) + fbn(n - 2)
			// 	}
			// }

			// func num(n int) int {
			// 	if n == 1 {
			// 		return 3
			// 	} else {
			// 		return 2 * num(n - 1) + 1
			// 	}
			// }
		// 猴子吃桃
		// 有一堆桃子,猴子第一天吃了其中一半,并再多吃一个,以后每天猴子都吃其中一半,然后再多吃一个
		// 当到第十天的时候,想再吃(还没吃),发现只有一个桃子,求最初共有多少个桃子
		// 思路分析:
			// 已知第十天的桃子是1个
			// 第九天的桃子 = (第十天的桃子 + 1) * 2
			// 第八天的桃子 = (第九天的桃子 + 1) * 2
			// 规律 : 第n天的桃子的数量 peach(n) = (peach(n + 1) + 1) * 2
				// func peach(mm int) int {
				// 	if mm > 10 || mm < 1 {
				// 		fmt.Println("输入数字超出范围")
				// 		return 0 // 表示不正确
				// 	}
				// 	if mm == 10 {
				// 		return 1
				// 	} else {
				// 		return (peach(mm + 1) + 1) * 2
				// 	}
				// }

				// func main() {
				// 	// 第一天的桃子
				// 	res := peach(1)
				// 	fmt.Println("第一天桃子的数量:", res)
				// }
	// 函数注意事项和细节讨论
		// 1. 函数的形参列表可以是多个,返回值列表也可以是多个
		// 2. 形参列表和返回值列表的数据类型可以是值类型和引用类型
		// 3. 函数的命名遵循表示符号命名规范,首字母不能是数字,首字母大写该函数可以被本包文件和其他包文件使用,类似public,首字母小写
		//    只能被本包文件使用,其他包文件不能使用,类似private
		// 4. 函数中的变量是局部的,函数外不生效
		// 5. 基本数据类型和数组默认都是值传递的,即进行值拷贝,在函数内修改,不会影响到原来的值
		// 6. 如果希望函数内的变量能修改函数外的变量,可以传入变量的地址&,函数内以指针的方式操作变量
		// 7. Go函数不支持传统重载,他有自己定义的函数重载
		// 8. 函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量, 通过该变量可以实现对函数的调用
		// 9. 函数既然是一种数据类型,因此在Go中,函数可以作为形参,并且调用
		// 10.为了简化数据类型定义,Go支持自定义数据类型
		// 11.Go支持对函数返回值命名
		// 12.使用下划线_标识符,忽略返回值
		// 13.在Go中支持可变参数
			// 说明:
				// args是slice切片,通过args[index]可以访问各个值
				// 如果一个函数的形参列表中有可变参数,可变参数要放在形参列表的后面
		// 代码展示
			// func test(n1 int) {
			// 	// 这里的函数定义的n1只能这个函数的局部中使用,全局中不能使用
			// 	n1 = n1 + 20
			// 	fmt.Println("test函数中的n1:", n1)
			// 	fmt.Printf("test函数中的n1地址%v\n", &n1)  // 0xc00008c018
			// }

			// // 如果真的想把全局变量中的n1也改变,那么可以传入指针
			// // 函数中的指针数据类型*int  *这是取指针指向某个值的地址 &这是取指针的地址
			// func test1(n1 *int) {
			// 	*n1 = *n1 + 10
			// 	fmt.Println("test1 n1=" ,*n1)
			// }

			// func main() {
			// 	// 定义一个num
			// 	num := 10
			// 	test(num)  //30
			// 	test1(&num)  // 20
			// 	fmt.Println("全局变量中的num的地址", &num)  // 20  地址0xc000018120
			// }

			// func getSum(n1 int, n2 int) int {
			// 	return n1 + n2
			// }

			// // 函数可以作为形参,并且被调用
			// func myFunc(funcvar func(int, int) int, num1 int, num2 int) int{
			// 	return funcvar(num1, num2)

			// }
			// func main() {
			// 	res := myFunc(getSum, 50, 60)
			// 	fmt.Println("res是:", res)
			// }

			// // 自定义数据类型案例
			// func main() {
			// 	type myInt int  // 定义一个自己的数据类型
			// 	var num1 myInt
			// 	num1 = 10
			// 	var num2 int
			// 	num2 = num1  // 如果这样写会报错,因为num2是int类型,num1是myInt类型,会报类型错误 cannot use num1 (type myInt) as type int in assignment
			// 	// num2 = int(num1)
			// 	fmt.Println("num1=", num1, "num2=", num2)  // num1= 10 num2= 10

			// }

			// // 自定义数据类型(全局定义)

			// // func myFunc(funcvar func(int, int) int, num1 int, num2 int) int{
			// // 	return funcvar(num1, num2)

			// // }  // 原先的函数是这样写的

			// func getSum(n1 int, n2 int) int {
			// 	return n1 + n2
			// }

			// // 当我们定义了一个全局数据类型
			// type myFuncClass func(int, int) int

			// func myFunc(funcvar myFuncClass, num1 int, num2 int) int{
			// 	return funcvar(num1, num2)

			// }

			// func main() {
			// 	res := myFunc(getSum, 500, 600)
			// 	fmt.Println("用自己定义的全局变量函数类型来输出的内容:", res)  // 1100
			// }

			// // 支持对返回值命名
			// func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
			// 	sub = n1 - n2
			// 	sum = n1 + n2 // 这里需要注意,这时不需要 := 因为在函数返回值命名的时候,相当于是声明了两个数据类型以及变量
			// 	return  // 直接return就可以了
			// }

			// func main() {
			// 	a, b := getSumAndSub(32,45)  
			// 	fmt.Println("和是:", a, "差是:", b)  // 和是: 77 差是: -13
			// }

			// // go支持可变参数
			// func sum(n1 int, args...int) int {
			// 	// 编写一个函数,可以求出1-多个int值
			// 	sem := n1
			// 	for i := 0; i < len(args); i++ {
			// 		sem += args[i]  // i表示索引
			// 	}
			// 	return sem
			// }

			// func main() {
			// 	res := sum(10, 0, 94, 23, 12, 45, 45, 67)
			// 	fmt.Println("他们的和:", res)  // 他们的和: 296
			// }

			// // 函数练习
			// func swap(n1,n2 *int) {
			// 	*n1,*n2 = *n2,*n1
			// }

			// func main() {
			// 	n1 := 20
			// 	n2 := 30
			// 	swap(&n1,&n2)
			// 	fmt.Println(n1, n2)

			// }

		// // init函数
		// // 基本介绍:
		// 	// 每一个源文件都可以包含一个init函数,该函数会在main函数执行前,被Go运行框架调用,也就是说init会在mian函数前被调用
		// func init() {
		// 	fmt.Println("init函数.......")  // 先执行init函数,通常可以在init函数中完成初始化工作
		// }

		// func main() {
		// 	fmt.Println("mian函数..........")
		// }

		// // 函数执行的顺序
		// // 全局变量 > init函数 > main函数
		// var age = test()

		// func test() int {
		// 	fmt.Println("执行了test函数.....")
		// 	return 80
		// }

		// func init() {
		// 	fmt.Println("执行了init函数......")
		// }

		// func main() {
		// 	fmt.Println("执行了main函数.......")
		// 	fmt.Println("age=", age)
		// }

// 匿名函数
// 介绍:
	//Go支持匿名函数,如果我们某个函数只是希望使用一次,可以考虑使用匿名函数,匿名函数也可以实现多次调用
// 匿名函数使用方式
	// 1. 在定义匿名函数时就直接调用
	// 2. 将匿名函数付给一个变量(函数变量),再通过该变量来调用匿名函数
	// 3. 如果将匿名函数赋值给一个全局变量,那么这个匿名函数,就成为一个全局匿名函数,可以在程序有效

// func main() {
// 	// 求两个数的和,匿名函数来写
// 	res := func (n1,n2 int) int {
// 		return n1 + n2
// 	}(10, 20)
// 	fmt.Println("res = ", res)  // 30

// 	// 将匿名函数赋值给一个变量
	
// }

// 闭包
// 累加器
// func AddUpper() func (int) int {
// 	var n int = 10
// 	return func (x int) int {
// 		n = n + x
// 		return n
// 	}
// }

// // 对上面的讲解
// 	// AddUpper是一个函数,返回的数据类型是func(int) int
// 	// 闭包返回的是一个函数,但是这个匿名函数引用了函数外的变量
// 	// 因此这个匿名函数就和n形成了一个整体,构成了闭包,相当于类中的一个静态属性
// 	// 当我们反复的调用函数时,n只初始化一次,因此每调用一次就会累加
// 	// 要搞清楚闭包的关键,就是要分析出返回的函数它使用到哪些变量,因为函数和他引用到的变量构成了闭包

// func main() {
// 	// 调用前面的代码
// 	f := AddUpper()
// 	fmt.Println(f(20))  // 30
// 	fmt.Println(f(20))  // 50
// }

// 练习
// 编写一个函数makeSuffix(suffix, string) 可以接收一个文件后缀名(比如.jpg)
// 并返回一个闭包
// 调用闭包,可以传入一个文件名,如果该文件名没有指定的后缀(比如.jpg),则返回文件名.jpg
// 如果已经有.jpg后缀,则返回源文件名
// 要求使用闭包的方式完成
// strings.HasSuffix来判断是否有后缀名

// func makeSuffix(suffix string) func(string) string {
// 	// suffx是传进来的文件名后缀
// 	// pathName是文件名
// 	return func (pathName string) string {
// 		if !strings.HasSuffix(pathName,suffix) {
// 			return pathName + suffix
// 		} else {
// 			return pathName
// 		}
// 	}
// }

// // 代码说明
// // 返回的匿名函数和makeSuffix(suffix string)的Suffix变量组成一个闭包,因为返回的函数引用到Suffix这个变量
// // 我们体会一下闭包的好处,如果使用传统的方法,也可以轻松实现这个功能,但是传统方法需要每次都传入后缀名,而闭包因为可以保留上次引用的某个值,所以我们传入一次就可以反复使用


// func main() {
// 	// 测试文件名的闭包
// 	name := makeSuffix(".jpg")
// 	// 打印传入文件之后的名字
// 	fmt.Println(name("花儿朵朵"))
// 	fmt.Println(name("花儿画画.jpg"))
// }

// 函数中的defer
// 在函数中,程序员需要创建资源(比如,数据库连接,文件句柄,锁等),为了在函数执行完毕后,及时的释放资源,Gode的设计者提供defer(延时机制)
// 在压入栈的时候是原始数据,外面对n1,跟n2的操作没有影响,所以从栈取出的时候还是原始数据
// // defer一般用到读取文件的时候,关闭文件句柄,关闭数据库连接,等整个函数执行结束完毕之后才会执行
// func sum(n1, n2 int) int {

// 	// 当程序执行到defer时,会将defer压入到独立的栈中(defer栈)
// 	// 当函数执行完毕后,再从defer栈中按照,先进后出的顺序执行代码
// 	defer fmt.Println("this is n1 :", n1)
// 	defer fmt.Println("this is n2 :", n2)
// 	n1++
// 	n2++
// 	res := n1 + n2
// 	fmt.Println("this is res:", res)
// 	return res
// }

// func main() {
// 	res := sum(30,49)
// 	fmt.Println("this is all:",res)
// }

// 函数参数传递方式
// 值传递,基本类型传递(int,float,bool,string),数组,结构体struct
// 引用传递(传地址),指针,slice切片,map,管道chan,interface  ===>效率高

// 变量的作用域
// 函数内部声明/定义叫局部变量,作用域仅限于函数内部
// 函数外面声明的变量,作用域在整个包都是有效的,如果首字母大写,在整个程序都是有效的
// 如果变量在代码块,比如if for中,那么这个变量的作用域就在该代码块中
// var a = "tom"
// // Name := "eee"  // 赋值语句不能在函数体外,所以这句错误
// func main() {
// 	//测试
	
// 	fmt.Println("没有定义数据类型的全局变量", a)
// }

