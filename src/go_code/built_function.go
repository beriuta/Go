package main
// import (
// 	"fmt"
// 	_"time"
// 	_"math/rand"
// 	_"errors"
// )

// // 内置函数
// // Golang设计者为了编程方便,提供了一些函数,这些函数可以直接使用,我们称为Go的内置函数
// 	// len()用来求长度,比如string,array,slice,map,channel
// 	// new()用来分配内存,主要用来分配值类型,比如 int,float32, struct,返回的是指针
// 	// make()用来分配内存,主要用来分配引用类型,比如channel(管道),map,slice

// func main() {

// 	// 常规的赋值方式
// 	num1 := 100
// 	fmt.Printf("num1的值%v, num1的类型%T,num1的地址%v\n", num1, num1, &num1)
// 	// num1的值100, num1的类型int,num1的地址0xc000018120

// 	// 用new()方法来分配内存
// 	num2 := new(int)
// 	*num2 = 100  // 直接赋值 赋值之后打印的  num2的值0xc000018138, num2的类型*int,num2的地址0xc00000e030,num2指向的值100
// 	fmt.Printf("num2的值%v, num2的类型%T,num2的地址%v,num2指向的值%v\n", num2, num2, &num2, *num2)
// 	// num2的值0xc000018138, num2的类型*int,num2的地址0xc00000e030,num2指向的值0
// }

// 错误处理机制
// Go语言追求简洁优雅,不支持try...catch...finally这种操作
// Go中引入的处理方式为:defer,panic,recover
// 在程序运行中,Go可以抛出一个panic的异常,然后在defer中通过recover捕获这个异常,然后正常处理
		// func test() {
		// 	// 引出错误机制
		// 	// 使用defer + recover来捕获处理异常
		// 	defer func() {
		// 		// err := recover()  // recover是内置函数,可以捕获异常
		// 		if err := recover(); err != nil {
		// 			fmt.Println("异常错误是:", err)
		// 			fmt.Println("发送邮件给2819345643@163.com")
		// 		}
		// 	}()  // 匿名函数直接加()执行
		// 	num1 := 100
		// 	num2 := 0
		// 	res := num1 / num2
		// 	fmt.Println("计算结果:", res)

		// }

		// func main() {
			
		// 	fmt.Println("开始出现错误")
		// 	test()
		// 	fmt.Println("错误触发")
		// }
// 错误处理机制------自定义错误
// errors.New() 和 panic内置函数
	// 1. errors.New("错误说明")会返回一个error类型的值,表示一个错误
	// 2. panic内置函数,接收一个空接口interface{}类型的值(也就是说任何值)作为参数,可以接收error类型的变量,
	// 输出错误信息,并退出程序
		// func readConf(name string) (err error) {
		// 	// 此函数读取配置文件init.conf的信息
		// 	// 如果文件名传入不正确,就返回一个自定义错误
		// 	if name == "config.ini" {
		// 		// 读取
		// 		return nil
		// 	} else {
		// 		// 返回一个自定义错误
		// 		return errors.New("读取文件错误.....")
		// 	}
		// }

		// func test02() {
		// 	err := readConf("config2.ini")
		// 	if err != nil {
		// 		// 如果读取文件发送错误,就输出这个错误,并终止程序
		// 		panic(err)
		// 	}
		// 	fmt.Println("程序继续执行.........")
		// }

		// func main() {

		// 	test02()
		// 	fmt.Println("mian程序会继续执行......")

		// }

// 数组,切片
	// 数组可以存放多个同一类型数据,数组也是一种数据类型,在Go中,数组是值类型
// 6只鸡,3,5,1,3.4,2,50kg,求总重量以及平均值
// 小数点精确到两位,打印的时候 %.2f
	// func main() {
	// 	// 利用数组的方式求出
	// 	var hens [6]float64
	// 	hens[0] = 3.0
	// 	hens[1] = 5.0
	// 	hens[2] = 1.0
	// 	hens[3] = 3.4
	// 	hens[4] = 2.0
	// 	hens[5] = 50.0

	// 	// 遍历求出总体重
	// 	totalWeight := 0.0
	// 	for i := 0; i < len(hens); i++ {
	// 		totalWeight += hens[i]
	// 	}
	// 	fmt.Println("总体重:", totalWeight)

	// 	// 求平均值 注意:如果除以一个数字,他是可以的,因为数字是一个常量值,而赋值给一个变量,变量是有数据类型的
	// 	avgWeight := fmt.Sprintf("%.2f",totalWeight / float64(len(hens)))  // 取两位小数
	// 	fmt.Println("平均体重:", avgWeight)
	// }

// 数组定义和内存布局
// 数组的定义
	// var 数组名 [数组大小]数据类型
	// var a [20] int
	// 赋值a[0]=1...
	// 数组的地址可以通过数组名获取,数组的地址就是数组的首地址,就是第一个元素的地址
	// 数组第二个元素的地址就是第一个元素地址基础上加第二个元素的字节数 int占8个字节
		// func main() {
		// 	// 获取数组的地址
		// 	var intArrt [3]int
		// 	fmt.Printf("数组的地址%p,数组的值%v, intArrt[0]的地址%p\n", &intArrt, intArrt, &intArrt[0])  // 数组的地址0xc000014340,数组的值[0 0 0]
		// 	// 数组的地址0xc000014340,数组的值[0 0 0], intArrt[0]的地址0xc000014340
		// }
	// 循环输入5个成绩,并保存到数组中
		// func main() {

		// 	// 从终端输入成绩
		// 	var list [5]int
		// 	for i := 0; i < len(list); i++ {
		// 		fmt.Println("输入成绩:")
		// 		fmt.Scanln(&list[i])
		// 	}
		// 	fmt.Println("所有的成绩为:", list)
		// }

	// 四种初始化数组的方式
		// func main() {

		// 	// 直接定义初始化
		// 	var numArr01 [3]int = [3]int{1, 2, 3}
		// 	fmt.Println("numArr01===>", numArr01)

		// 	// 直接使用类型推导
		// 	numArr02 := [3]int{5, 6, 7}
		// 	fmt.Println("numArr02===>", numArr02)

		// 	// 不写长度,固定写...
		// 	var numArr03 = [...]int{8, 9, 710, 23, 34, 56}
		// 	fmt.Println("numArr03===>", numArr03)

		// 	// 不按照默认数据给值,指定下标来赋值
		// 	var numArr04 = [...]int{3: 800, 0: 900, 2: 710, 1: 230}
		// 	fmt.Println("numArr04===>", numArr04)
		// }
// 数组的遍历
// 方式1,常规遍历
// 方式2, for-range结构遍历
// 结构语法:
	// for index, value := range array01 {
	// 	...
	// }
// 结构说明
	// 1. 第一个返回值index是数组的下标
	// 2. 第二个value是在下标位置的值
	// 3. 他们都是仅在for循环内部可用的局部变量
	// 4. 遍历数组元素时,如果不想使用下标index,可以直接把下标标为_
	// 5. index和value的名称不是固定的,可自行指定
		// func main() {

		// 	// 演示for -range遍历数组
		// 	heroes := [...]string{"松江", "吴勇", "巴拉拉"}
		// 	for k, v := range heroes {
		// 		fmt.Printf("输出的下标%v, 输出的值为%v\n", k, v)
		// 	}
		// }
// 数组使用注意事项和细节
// 1. 数组是多个相同类型数据的组合,一个数组一旦定义了,长度是固定的,不能动态变化
// 2. var arr []int 这时arr就是一个slice切片
// 3. 数组中的元素可以是任何数据类型,保罗类型和引用类型,但是不能混用
// 4. 数组创建后,如果没有赋值,有默认值为 int默认值0, 字符串默认值"",bool默认为False
// 5. 使用数组的步骤,1.声明数组并开辟空间,2.给数组哥哥元素赋值, 3.使用数组
// 6. 数组下标都是0开始的
// 7. 数组下标必须指定范围内使用,否则报panic,数组越界
// 8. Go的数组属于值类型,在默认情况下是值传递,因此会进行值拷贝,数组间不会相互影响
// 9. 如想在其他函数中去修改原来的数组,需要吧地址穿进去,引用传递(指针方式),当数据庞大时,而又想修改此时的数组,就用指针把地址传递进去修改

// 数组应用案例
// 1. 创建一个byte类型的26个元素的数组,分别放置'A'-'Z',使用for循环访问所有元素并打印出来,提示:字符数据运算"A"+1 = "B"
	// func main() {
	// 	var num [26]byte
	// 	for i := 0; i < len(num); i++ {
	// 		num[i] = 'A' + byte(i)
	// 	}
	// 	fmt.Println(num)
	// 	for i := 0; i < len(num); i++ {
	// 		fmt.Printf("%c\n", num[i])
	// 	}
	// }
// 2. 请求出一个数组的最大值,并得到对应的下标
	// func main() {
	// 	num := [...]int{2,56,998,32}
	// 	maxIndex := 0
	// 	maxValue := num[0]
	// 	for i := 1; i < len(num); i++ {
	// 		if num[i] > maxValue {
	// 			maxValue = num[i]
	// 			maxIndex = i
	// 		}
	// 	}
	// 	fmt.Printf("最大的数:%v,它的下标为:%v\n", maxValue, maxIndex)
	// }
// 3. 请求出一个数组的和以及它的平均值 for-range
	// func main() {
	// 	num := [...]int{2,56,998,32}
	// 	allNum := 0
	// 	for _, v := range num {
	// 		allNum += v
	// 	}
	// 	avgNum := float64(allNum) / float64(len(num))
	// 	fmt.Printf("平均值:%v,总和:%v\n", avgNum, allNum)
	// }
// 4. 随机生成一个数组,将他反转打印
// 逻辑分析:
	// 1. 先用rand.Intn()随机生成四个数字,循环他们的一半长度
	// 2. 要定义一个时间随机,不然会一直打印一个数字
	// 3. 然后定义一个临时变量,存放交换的时候的数字
	// 4. 将最后一个数字交给临时变量,将第一个数字赋值给最后一个数字的位置,将临时变量的最后数字赋值给第一个数字的位置

		// func main() {
		// 	var num [4]int
		// 	lenNum := len(num) 
		// 	rand.Seed(time.Now().UnixNano())
		// 	for i := 0; i < lenNum; i++ {
		// 		num[i] = rand.Intn(100)
		// 	}
		// 	fmt.Println("交换前数组是:", num)

		// 	// 开始交换
		// 	temp := 0
		// 	for i := 0; i < lenNum / 2; i++ {
		// 		// 将最后一个值给temp
		// 		temp = num[lenNum - 1 - i]
		// 		// 将第一个赋值给最后一个位置
		// 		num[lenNum - 1 - i] = num[i]
		// 		// 再将temp的值赋值给第一个位置
		// 		num[i] = temp
		// 	}
		// 	fmt.Println("交换后数组是:", num)
		// }

// 切片的基本介绍
// 切片可以简单的理解成是一个动态的数组,用于不知道长度的情况
// 切片是数组的一个引用,因此切片是引用类型,在进行传递时,遵守引用传递的机制
// 切片的使用和数组类似,遍历切片.访问切片的元素和求切片长度len(slice)都一样
// 切片的长度是可以变化的,因此切片是一个可以动态变化数组
// 切片定义的基本语法:
	// var 变量名 []类型 : var a []int

	// 切片的演示
		// func main() {

		// 	// 定义一个基本的数组
		// 	intArr := [...]int{1,2,56,87,34}
		// 	// 声明一个切片, 切片的名字必须叫slice
		// 	slice := intArr[1:3]
		// 	fmt.Println("slice的元素是", slice)
		// 	fmt.Println("slice的个数是", len(slice))  // 2
		// 	fmt.Println("slice的容量是", cap(slice))  // 4   切片的容量是可以动态变化的

		// }
// 切片在内存中的形式
// 切片底层的数据结构可以理解长一个结构体struct
// 输出切片和切片引用的地址
// 切片的第一个部分就是指向的一个最开始引用的那个元素的地址, 第二部分是存放这个切片的长度, 第三部分可以用来存放这个切片的容量

// 定义切片有两种方式
	// 1. 定义一个数组,在数组的基础上切片
	// 2. 用make内置函数来创建切片
		// 基本语法:
		// var 切片名 []type = make([]int,len,[cap]), 创建好的内容全部都是0
		// 说明:type就是数据类型 len:长度 cap:指定切片容量,可选, 如果定义Cap,Cap要>=len
		// 通过make方式创建切片可以指定切片的大小和容量
		// 如果没有给切片的哥哥元素赋值,那么就会使用默认值(int,float==>0, string==>"", bool==>False)
		// 通过make方式创建的切片对应的数组是由make底层维护的,对外不可见,即只能通过slice去访问各个元素
			// func main() {
			// 	// 用make创建切片
			// 	var slice []float64 = make([]float64, 5, 10)  // slice 默认指向第一个元素的地址
			// 	fmt.Println("创建好的切片内容是:", slice)  // [0 0 0 0 0]
			// }
	// 3. 第三种方式,定义一个切片,直接就指定具体数组,使用原理类似make的方式
		// func main() {
		// 	var slice []string = []string{"tom", "jack", "mary"}
		// 	fmt.Println("切片的元素:", slice)  // [tom jack mary]
		// 	fmt.Println("切片的个数:", len(slice))  // 3
		// 	fmt.Println("切片的元素:", cap(slice))  // 3
		// } 
// 切片的方式1跟方式2的区别
	// 方式1是直接引用一个数组,数组是事先存在的,程序员是可见的
	// 方式2 make创建切片,也会创建一个数组,是由切片在底层进行维护,程序员是看不见的
	// 指针是第一个元素地址
// 遍历切片
	// func main() {
	// 	// 使用常规的for循环遍历切片
	// 	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	// 	slice := arr[1:4] // 20,30,40
	// 	for i := 0; i < len(slice); i++ {
	// 		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	// 	}

	// 	// 使用for-range方式遍历切片
	// 	for k, v := range slice {
	// 		fmt.Printf("slice[%v]=%v\n", k, v)
	// 	}
	// }

// 切片注意事项和细节说明
	// 切片初始化时 var slice = arr[start_index:end_index]
	// 说明:从arr数组下标为start_index,取到下标为endindex的元素,不含endindex的元素
	// 切片初始化时仍然不能越界,范围在[0-len(arr)]之间,但是可以动态增长
		// var slice = arr[0:end]可以简写 var slice = arr[:end]
		// var slice = arr[start:len(arr)] 可以简写 var slice = arr[strat:]
		// var slice = arr[0:len(arr)] 可以简写 var slice = arr[:]
	// cap是一个内置函数,用于统计切片的容量,即最大可以存放多少个元素
	// 切片定义完后,还不能使用,因为本身是一个空的,需要让其引用一个数组或者make一个空间供切片来使用
	// 切片可以继续切片
	// arr,slice,和slice2如果其中一个值改变了,都会一起改变,因为他们指向的都是同一块内存空间

		// // slice从底层来说,其实就是一个结构体(struct结构体)
		// type slice struct{
		// 	prt *[2]int  // 指针
		// 	len int  // 长度
		// 	cap int  // 容量
		// }
// 切片的append函数
// 将元素追加到数组的后面,扩容,追加的数据类型要一致
	// func main() {
	// 	var slice3 []int = []int{299,333,44,55,453}
	// 	// 通过append直接给slice3追加具体的元素
	// 	// slice3 = append(slice3, 400,500,600)  // 要有个变量接收追加之后的数组
	// 	// 如果是别的变量接收,本身slice3这个数组是不改变的,一般是本身接收
	// 	slice4 := append(slice3, 400, 500, 600) 
	// 	fmt.Println("slice3=", slice3)       //  slice3= [299 333 44 55 453]
	// 	fmt.Println("slice4=", slice4)      // slice4= [299 333 44 55 453 400 500 600]
	// 	// 将一个切片追加到slice3切片中
	// 	slice3 = append(slice3, slice3...)  // 后面必须是切片,不能是数组
	// 	fmt.Println("slice3=", slice3)  // slice3= [299 333 44 55 453 299 333 44 55 453]
	// }
// 切片append操作的底层原理分析
// 1. 切片append操作的本质就是对数组扩容
// 2. go底层会创建一下新的数组newArr(数组的大小就是安装扩容后大小)
// 3. 将slice原来包含的元素拷贝到新的数组newArr
// 4. slice重新引用newArr
// 5. 注意newArr是在底层来维护的,程序员不可见

// 切片类型的拷贝(注意:不是数组,数组不能执行)
	// func main() {
	// 	var a []int = []int{1,2,3,4,5}
	// 	// 再定义一个切片
	// 	var slice = make([]int, 30)
	// 	fmt.Println("slice=",slice)  // slice= [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	// 	copy(slice, a)
	// 	fmt.Println("slice=", slice)  // slice= [1 2 3 4 5 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	// }
	// 按上面的代码,slice和a的数据空间是独立的,不相互依赖的,也就是说如果修改了某一个,另一个并不会跟着修改
	// 如果拷贝的那个切片长度小于被拷贝的长度,那么只拷贝相应的个数的值,多余的忽略
	// 切片是引用类型,所以在传递时,遵守引用传递机制,如果拷贝的数组,当某个值修改都修改


// string和slice
	// 1. string是byte类型,所以可以用切片
	// 2. string是不可变的,如果要通过下标来修改字符串会报错
	// 3. 如果需要修改字符串,需要先将字符串>>转化成byte类型,或者[]rune>>>修改>>>重新转成string
		// func main() {
		// 	str := "hello@world!"
		// 	slice := str[:]
		// 	fmt.Println("str=", str)
		// 	fmt.Println("slice=", slice)

		// 	// arr1 := []byte(str)
		// 	// arr1[0] = 'z'
		// 	// str = string(arr1)
		// 	// fmt.Println("str=", str)  // str= zello@world!

		// 	// 如果是中文,会报错,一个汉字是三个字节,会报错
		// 	// 需要将string转成[]rune即可,因为rune是按照字符处理的,兼容汉字

		// 	arr1 := []rune(str)  // 用[]rune
		// 	arr1[0] = '深'
		// 	str = string(arr1)
		// 	fmt.Println("str=", str)  // str= 深ello@world!

		// }

// 编写一个函数fbn(n int)
// 可以接收一个 n int
// 能够将斐波那契数列放到切片中
// 前两位是1,之后没加一位就是前两位的和
// 1. 声明一个函数,接收一个整数,返回一个切片,数据类型需要存放最大数量[]unit64
// 2. for循环存放斐波那契数列
	// func feBo(n int) ([]uint64) {
	// 	// 声明一个切片数据类型
	// 	fbnSlice := make([]uint64, n)
	// 	// 第一个数和第二个数是等于1
	// 	fbnSlice[0] = 1
	// 	fbnSlice[1] = 1
	// 	for i := 2; i < n; i++ {
	// 		fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
	// 	}
	// 	return fbnSlice
	// }

	// func main() {
	// 	fb := feBo(10)
	// 	fmt.Println("斐波那契数列=", fb)  // 斐波那契数列= [1 1 2 3 5 8 13 21 34 55]
	// }
	
