package main
// import (
// 	// "fmt"
// 	// "math/rand"
// 	// "time"
// )

// func main(){
// 	// 运算符 /(除法) 和 %(取余)

// 	// 如果参与运算的都是整数,会把小数点去掉
// 	fmt.Println(10 / 4)  // 2
// 	var n1 float32 = 10 / 4  // 2
// 	fmt.Println(n1)  //不管数据类型是小数还是整数,仍然不会保留小数位
// 	// 如果需要保留小数部分,需要有小数参与计算
// 	var n2 float32 = 10.0 / 4
// 	fmt.Println(n2)  // 2.5

// 	// 取余  推导公式 : a % b = a-a / b*b
// 	var n3 float32 = 10 % 3 
// 	var n4 float32 = -10 % 3  // -10 - (-10) / 3 * 3 = -10 - (-9) = -1
// 	var n5 float32 = 10 % -3 
// 	var n6 float32 = -10 % -10 
// 	fmt.Println(n3)  // 1
// 	fmt.Println(n4)  // -1
// 	fmt.Println(n5)  // 1
// 	fmt.Println(n6)  // 0

// 	// ++ 和 --
// 	var i int = 10
// 	i++  // 等价于 i = i + 1
// 	fmt.Println(i)  // 11
// 	i-- // i = i - 1
// 	fmt.Println(i)  // 10
// 	// i-- 跟 i++ 只能独立使用,不能赋值给变量
// 	// a = i++  这是错的
// 	// i++ > 0  这也是错的 正确使用是先i++ 然后再 i > 0
// 	// i++ i--是可以的 但是不允许 ++i --i 

// 	// 练习
// 	// 假如还有97天放假,问,xx个星期零xx天
// 	// 定义一个变量保存华氏温度,华氏温度转换摄氏温度的公式为 5.0/9*(华氏温度-100)
// 	// 请求出华氏温度对应的摄氏温度
// 	var days int = 97
// 	var week int = days / 7
// 	var day int = days % 7
// 	fmt.Printf("%d个星期零%d天\n",week,day)
// 	var hua float32 = 134.2  // 华氏温度
// 	var num float32 = 5.0 / 9 * (hua - 100)
// 	fmt.Printf("%v 对应的摄氏温度=%v \n",hua,num)

// 	// 逻辑运算
// 	// 多个条件连起来判断 && 两边结果都为真的时候才会为真 || 一边结果8为真就为真 |逻辑非,整个结果的反向
// 	var age int = 40
// 	if age > 30 && age < 50{
// 		fmt.Println("ok1")
// 	}
// 	if age > 30 && age < 40{
// 		fmt.Println("ok2")
// 	}
// 	// 逻辑运算符号使用 ||
// 	if age < 30 || age > 40{
// 		fmt.Println("比30小的或者比40大的都会进这个")
// 	}
// 	if age <30 || age < 41{
// 		fmt.Println("比30小的或者比41小的")
// 	}
// }

// 声明一个函数(测试)
// func test() bool{
// 	fmt.Println("test....")
// 	return true
// }

// func main(){
// 	var i int = 10
// 	// if i > 9 && test(){
// 	// 	fmt.Println("ok.........")
// 	// }  // 先输出test......再输出ok.......... 

// 	// 短路与
// 	if i < 9 && test(){
// 		fmt.Println("ok.........")
// 	}  // 什么都不输出,因为i < 9这个判断就是错的,所以后面的就不会再判断了

// 	// 如果修改成||, 先判断前面是否为真,如果不为真,再判断后面的是否为真,如果为真,就为真
// 	if i < 9 || test(){
// 		fmt.Println("ok....hello.....")
// 	}  // 依然会先输出test........然后再输出 ok....hello.....
// }

// // 赋值运算符, =, +=, -=, /=, %=
// func main(){
// 	a := 9
// 	b := 2
// 	fmt.Printf("交换前的a=%d,b=%d\n", a, b)
// 	t := a
// 	a = b
// 	b = t
// 	fmt.Printf("交换后的a=%d,b=%d\n", a, b)
// 	// 组合赋值
// 	num2 := 30
// 	num2 += 7
// 	fmt.Println("num2=", num2)  // 37

// 	// 两个变量,mm 和 nn ,不允许用中间变量交换,打印最终结果
// 	mm := 30
// 	nn := 28
// 	mm,nn = nn,mm
// 	fmt.Printf("mm%d和nn%d\n", mm, nn)
// }


// // &和*的使用
// func main(){
// 	a := 100
// 	fmt.Println("a的地址:", &a)  // a的地址: 0xc000018120
// 	var ptr *int = &a  // 把a的地址交给ptr 用*取出对应的值
// 	fmt.Println("ptr指向的是", *ptr)  //ptr指向的是 100
// 	// golang不支持三元运算符, 如果要用只能是if else

// 	var nnn int
// 	var i int = 10
// 	var j int = 12
// 	// 传统三元运算, 如果i > j n=i 如果相反, n = j
// 	// n = i > j ? i : j

// 	// 在golang中的写法
// 	if i > j {
// 		nnn = i
// 	} else {
// 		nnn = j
// 	}
// 	fmt.Println("nnn=", nnn)  // 12
// }

// func main(){

// 	// 求两个数的最大值
// 	var n1 int = 10
// 	var n2 int = 33
// 	var max int
// 	if n1 < n2 {
// 		max = n2
// 		fmt.Println("最大值n2:", n2)
// 	} else {
// 		max = n1
// 		fmt.Println("最大值n1:", n1)
// 	}

// 	// 求出三个数的最大值
// 	// 先求出两个数的最大值,将求出的最大值跟第三个数比较,再求出最大值
// 	var n3 int = 2
// 	if max < n3 {
// 		max = n3
// 		fmt.Println("最大值max是n3", n3)
// 	}
// 	fmt.Println("最大值max是max", max)

// }

// func main(){
// 	// 要求可以从控制台接收用户信息[姓名,年龄,薪水, 是否通过测试]
// 	// 使用fmt.Scanin()获取  ==> 读取一行数据
// 	// 使用fmt.Scanf()获取

// 	// // 声明变量
// 	var name string
// 	var age int
// 	var sal float32
// 	var isPass bool
// 	// fmt.Println("请输入姓名")  
// 	// // 当程序执行到这里时,程序会等待用户输入,停止在这里,当用户输入之后才会继续执行以下代码
// 	// fmt.Scanln(&name)  // 把变量的地址传递到函数里面

// 	// fmt.Println("请输入年龄")
// 	// fmt.Scanln(&age)

// 	// fmt.Println("请输入薪水")
// 	// fmt.Scanln(&sal)

// 	// fmt.Println("是否通过测试")
// 	// fmt.Scanln(&isPass)

// 	// fmt.Printf("名字是 %v \n 年薪是 %v \n 年龄是 %v \n 是否通过考试 %v \n", name, sal, age, isPass)

// 	// 方式2:fmt.Scanf(), 可以按指定的格式输入
// 	fmt.Println("请输入你的姓名,年龄,薪水,是否通过考试,使用空格隔开")
// 	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
// 	fmt.Printf("名字是 %v \n 年薪是 %v \n 年龄是 %v \n 是否通过考试 %v \n", name, sal, age, isPass)
// }

// 进制的介绍
// ========对于整数,有四种表示方式
// 二进制: 0,1 满2进1
// 十进制: 0-9, 满10进1
// 八进制: 0-7, 满8进1, 以数字0开头表示
// 十六进制: 0-9及A-F,满16进1,以0x或0X开头表示   此处的A-F不区分大小写 如:0x21AF + 1 = 0X21B0

// // 在golang中不能直接使用二进制表示一个整数,它沿用了c的特点
// func main(){

// 	var i int = 5
// 	// 二进制转换
// 	fmt.Printf("%b\n", i)  // %b是输出二进制的  101

// 	// 八进制, 0-7 在go中以数字0开头表示
// 	var j int = 011  
// 	fmt.Println("j=", j)  // 9

// 	// 十六进制, 0-9以及A-F,在go中以0x或者0X开头
// 	var mm int = 0x11
// 	fmt.Println("mm=", mm)  // 17

// 	// 进制转换的介绍
// 	// 第一组(其他进制转十进制)
// 		// 二进制转十进制
// 		// 八进制转十进制
// 		// 十六进制转十进制
// 		// 规则: 从最低位开始(右边的),将每个位上的数提取出来,乘以2的(位数-1)次方
// 		//     然后求和

// 	// 案例: 请将二进制1011转成十进制的数
// 	// 1011 = 1 * 2的0次方 + 1 * 2的1次方 + 0 * 2的2次方 + 1 * 2的3次方
// 	//      = 1 + 2 + 0 + 8 = 11

// 	// 十进制转化成十进制
// 	// 134 = 4 * 10的0次方 + 3 * 10的1次方 + 1 *10的2次方
// 	//     = 4 + 30 + 100 = 134

// 	// 八进制转换成十进制
// 	// 0123 = 3 * 8的0次方 + 2 * 8的1次方 + 1 * 8的2次方 + 0 * 8的3次方
// 	//      = 3 + 16 + 64 + 0 = 83

// 	// 十六进制转十进制
// 	// 0x34A = 10 * 16的0次方 + 4 * 16的1次方 + 3 * 16的2次方 
// 	//       = 10 + 64 + 768 = 842

// 	// 第二组(十进制转其他进制)
	
// 	// 十进制转二进制
// 	// 规则: 将该数不断除以2,直到商为0为止,然后将每步得到的余数倒过来,就是对应的二进制

// 	// 请将56转成二进制
// 	//  56 / 2 = 28 余数0
// 	//  28 / 2 = 14 余数0
// 	//  14 / 2 = 7  余数0
// 	//  7 / 2 = 3   余数1
// 	//  3 / 2 = 1   余数1
// 	//   56 转换成二进制 : 111000

// 	// 十进制转八进制
// 	// 规则: 将该数不断除以8, 直到商为0为止,然后将每步的余数倒过来,就是对应的八进制

// 	// 156转化成八进制
// 	// 156 / 8 = 19  余数4
// 	// 19 / 8 = 2  余数3
// 	// 156转换成八进制 : 0234

// 	// 十进制转十六进制
// 	// 356转成十六进制
// 	// 356 / 16 = 22 余数4
// 	// 22 / 16 = 1  余数6
// 	// 356转换成十六进制 : 0x164

// 	// 第三组(二进制转其他进制)

// 	// 规则: 将二进制的每三位一组(0-7),从低位开始组合,转成对应的八进制数即可
// 	// 请将二进制11010101转换成八进制
// 	// 11    010   101   ====> 每三位分成一组
// 	// 101 = 5
// 	// 010 = 2
// 	// 11 = 3
// 	// 11010101二进制转八进制 = 0325

// 	// 二进制转十六进制
// 	// 规则: 将二进制数每四位一组(0-F)从低位开始组合,转成对应的十六进制即可
// 	// 将二进制11010101转成十六进制的
// 	// 1101     0101    =====> 每四位分成一组
// 	// 0101 = 5
// 	// 1101 = 13  13在十六进制中是 D
// 	// 11010101转换成十六进制 = 0xD5

// 	// 第四组(八进制,十六进制转二进制)

// 	// 将八进制数每1位,转化成对应的一个3位的二进制数即可
// 	// 0237 : 7 ===> 111
// 	//        3 ===> 011
// 	//        2 ===> 10
// 	// 0237转成二进制 10011111
	
// 	// 将十六进制转换成二进制
// 	// 规则: 将十六进制每1位,转成对应的一个4位的二进制数即可
// 	// 请将0x237转换成二进制
// 	// 7 ==> 0111
// 	// 3 ==> 0011
// 	// 2 ==> 10
// 	// 0xo237转换成二进制 : 1000110111
// }

// func main(){

// 	// 在计算机内部,运行各种运算时,都是以二进制的方式来进行运算的

// 	// 原码, 反码, 补码
// 	// - 对于有符号的而言:(在golang中分为有符号和无符号的数)
// 			// 1. 二进制的最高位是符号位:0表示正数,1表示负数
// 			//=========> 1 = [0000 0001]  -1 = [1000 0001]  首位的0和1来表示正负数
// 			// 2. 正数的原码,反码,补码都是一样的
// 			// 3. 负数的反码 = 它的原码符号位不变,其他位取反(0==>1, 1==>0)
// 			//  =========>  1的原码[0000 0001]  反码[0000 0001] 补码[0000 0001]  >正数
// 			//  =========>  -1的原码[1000 0001] 反码[1111 1110] 符号位不变,其余的取反   补码[1111 1111] 反码 + 1
// 			// 4. 负数的补码 = 它的反码 + 1
// 			// 5. 0的反码,补码都是0
// 			// 6. 在计算机运算的时候,都是以补码的方式来运算的

// 	// 位运算  只要是运算,就要转成补码
// 	// 右移运算 >> : 低位溢出符号位不变,并用符号位补溢出的高位
// 	// 左移运算 << : 符号位不变,低位补0
// 	var a int = 1 >> 2  // >> 位右移两位, 1的补码[0000 0001] 向右移两位 [0000 0000] = 0
// 	var b int = -1 >> 2  
// 	var c int = 1 << 2  // 1的补码[0000 0001] 向左移两位[0000 0100] = 4
// 	var d int = -1 << 2
// 	fmt.Println("a=", a)
// 	fmt.Println("b=", b)
// 	fmt.Println("c=", c)
// 	fmt.Println("d=", d)

// 	// 在golang中,下面的表达式运算结果是,计算机运算的时候都是以补码的方式来运算的

// 	fmt.Println(2&3)
// 	// 先写出2的补码[0000 0010]正数的原码,反码,补码都是一样的  3的补码[0000 0011]
// 	// & 都为1的时候才会为1否则为0 所以,2&3 [0000 0010]  = 2

// 	fmt.Println(2|3)
// 	// 写出2跟3的补码,然后对比,只要有1的就为1,否则为0
// 	// 2 | 3 [0000 0011] = 3

// 	fmt.Println(2^3)
// 	// ^ 亦或 ,两个都相等的时候它为0, 否则为1  两个不一样的为1
// 	// 2 ^ 3 [0000 0001] = 1

// 	fmt.Println(-2^2)
// 	// 先把两个数的补码写出来  
// 	// 2因为是正数,补码跟原码一样 [0000 0010]
// 	// -2是负数,补码 = 反码 + 1, 反码 = 原符号位不变,其余数取反
// 	// -2的原码 = [1000 0010] 反码 = [1111 1101] 补码 = [1111 1110]
// 	// -2 ^ 2 = [1111 1100] 此时得出来的是补码,还要翻译成原码才能得出结论
// 	// [1111 1100] ==> 推出反码 [1111 1011]  ==> 变成原码 [1000 0100] = -4

// 	fmt.Println(13&7)
// 	fmt.Println(6|4)

// }

// 程序流程控制  ==> 三大流程控制
// 一.顺序控制
	// 程序从上到下逐行执行,中间没有任何判断和跳转
	// golang中定义的变量采用前向引用,先定义,再引用

// 二.分支控制
	// 让程序有选择性的执行 分为三个分支

	// 单分支 
		// 当条件表达式为True时,就会执行{}的代码
		// {}必须要有,就算只有一行代码
		// if 条件表达式 {
			// 执行代码块
		// }

	// 例子
		// func main(){
		// 	// 用户输入大于18的年龄返回,你的年龄大于18,要对自己的行为负责
		// 	var age int 
		// 	fmt.Println("请输入年龄:")
		// 	fmt.Scanln(&age)
		// 	if age > 18 {
		// 		fmt.Println("你的年龄大于18岁,要对自己的行为负责")
		// 	}
		// }

		// Go的if还有一个强大的功能,就是条件判断语句里面允许声明一个变量,这个变量的作用域只能在该条件逻辑块内使用,其他地方就不起作用了
		// 	if age := 20; age > 18 {
		// 		fmt.Println("你的年龄大于18岁,要对自己的行为负责")
		// 	}

	// 双分支 
	// 当条件表达式成立,即执行代码块1,否则执行代码块2 {}必须要有
	// 	if age > 18 {
		// 		fmt.Println("你的年龄大于18岁,要对自己的行为负责")
		// 	} else {
			// fmt.Println("你的年龄小于18岁,不能进入")
		// }

	//例子
		// func main(){
		// 	// 判断一个年份是否是闰年,闰年的条件符号下面二者之一
		// 	// 年份能够被4整除 但不能被100整除  能被400整除
		// 	var year int = 2019
		// 	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0{
		// 		fmt.Println("这一年是闰年")
		// 	}
		// }

	// 多分支
	// if 条件表达式 {
			// 执行代码块1
		// } else if 条件表达式2 {
			// 执行代码块2
		// }...
		// else {   =========> 这个else是可有可没有的,不是必须的
			// 执行代码块n
		// }
	// 语法说明:
		// 1. 先判断条件表达式,是否成立,如果为真就执行代码块1
		// 2. 如果条件表达式1的结果为假,就执行条件表达式2
		// 3. 以此类推
		// 4. 如果以上代码全部都不成立,则执行最后else的代码块n
		// 5. 多分支最终只能有一个执行入口
	// 例子
		// 如果成绩为100分,奖励一辆BMW
		// 成绩为80-99, 奖励一台IP7
		// 成绩60-80, 奖励iPad
		// 其他什么也不奖励
		// func main(){
		
		// 	var fraction int 
		// 	fmt.Println("请输入成绩:")
		// 	fmt.Scanln(&fraction)
		// 	if fraction == 100 {
		// 		fmt.Println("奖励一台BMW")
		// 	} else if fraction > 80 && fraction <= 99 {
		// 		fmt.Println("奖励一台IP7")
		// 	} else if fraction > 60 && fraction <= 80 {
		// 		fmt.Println("奖励iPad")
		// 	} else {
		// 		fmt.Println("什么不奖励")
		// 	}
		// }
	
	// 嵌套分支
	// 在一个分支结构中又完整的嵌套了另一个完整的分支结构,里面的分支结构称为内层分支
	// 外面的分支结构称为外层分支
	// 基本语法:
		// if 条件表达式 {
			// if 条件表达式2 {
				// 代码块
			// } else {
				// 代码块
			// }
		// } else {
			// 代码块
		// } 
	// 一般建议嵌套最多三层

	// 例子
		// func main(){
		// 	// 参加百米运动会,如果用时8秒以内进入决赛,否则淘汰,并且根据性别进入男子队或者女子队,输入成绩和性别,系统自行判断
		// 	var s int 
		// 	var sex string
		// 	fmt.Println("请输入成绩跟性别,以空格分隔")
		// 	fmt.Scanln(&s, &sex)
		// 	if s <= 8 {
		// 		fmt.Println("进入决赛")
		// 		if sex == "男" {
		// 			fmt.Println("你进入决赛的男子队")
		// 		} else if sex == "女" {
		// 			fmt.Println("你进入决赛的女子队")
		// 		}
		// 	} else {
		// 		fmt.Println("你被淘汰了")
		// 	}
		// }
	
	//switch分支

	// 基本介绍: switch语句用于不同条件执行不同动作,每一个case分支都是唯一的,从上到下逐行测试,直到匹配为止
			// 匹配项后面也不需要再加break, 如果匹配不到就执行默认的分支

	// 基本语法
		// switch 表达式 {    ===> 如果switch匹配到第一个case的表达式的值就为真
		// case 表达式1,表达式2...:  ===>case是当....的时候,表达式是或者
			// 语句块1
		// case 表达式3,表达式4...:
			// 语句块2
		// 这里可以有多个case
		// default:
			// 语句块
		// }
	// 基本语法逻辑介绍
		// switch的执行流程,先执行switch表达式,得到值,然后和case的表达式进行比较,如果相等,就匹配到,然后执行对应case的语句块,然后退出switch控制
		// 如果switch的表达式的值没有和任何case的表达式匹配成功,则执行default语句块, 然后退出switch控制,执行下面的代码
		// golang的case后面的表达式,可以有多个,使用逗号间隔
		// golang中case语句块不需要写break,因为默认情况下,等程序执行完case语块后,就直接退出该switch控制

	// 案例
		// func main(){
		// 	// 接收一个字符,比如:a,b,c,d,e,f,g
		// 	// a表示星期一,b表示星期二,根据用户输入的显示相应的星期
		// 	var key byte
		// 	fmt.Println("请输入想要知道所代表含义的字母:")
		// 	fmt.Scanf("%c", &key)
		// 	switch key {
		// 		case 'a':  // "" : 双引号go中默认为字符串,  '': 单引号默认为byte类型
		// 			fmt.Println("星期一")
		// 		case 'b':
		// 			fmt.Println("星期二")
		// 		case 'c':
		// 			fmt.Println("星期三")
		// 		case 'd':
		// 			fmt.Println("星期四")
		// 		case 'e':
		// 			fmt.Println("星期五")
		// 		case 'f':
		// 			fmt.Println("星期六")
		// 		case 'g':
		// 			fmt.Println("星期天")
		// 		default:
		// 			fmt.Println("输入的字符没有任何含义")
				
		// 	}
		// }

	// switch细节讨论
		// case后是一个表达式,即:常变量,变量,一个有返回值的函数等都可以
		// case后的各个表达式的值的数据类型,必须和switch的表达式数据类型一致
		// case后面可以多个表达式,使用逗号间隔,比如,case表达式1,case表达式2....
		// case后面的表达式如果是常量值(字面量),则要求不能重复
		// case后面不需要带break,程序匹配发哦一个case后就会执行对应的代码块,然后退出switch,如果一个匹配不到,则执行default
		// default语句不是必须的
		// switch后面也可以不带表达式,类似一个if-else分支来使用

		// func main(){
		// 	// 把switch当做if-else来用
		// 	var age int = 10
		// 	switch {
		// 		case age == 10:
		// 			fmt.Println("年龄是十岁")
		// 		case age == 20:
		// 			fmt.Println("年龄是二十岁")
		// 	}
		// }
		
	// switch 穿透 fallthrough,如果在case语句块后增加fallthrough,则会继续执行下一个case
		// func main (){
		// 	// switch 穿透
		// 	var num int = 10
		// 	switch num {
		// 	case 10:
		// 		fmt.Println("ok1")
		// 		fallthrough  // 默认穿透一层
		// 	case 20:
		// 		fmt.Println("ok2")
		// 		fallthrough
		// 	case 30:
		// 		fmt.Println("ok4")
		
		// 	default:
		// 		fmt.Println("没有匹配项.....")
		
		// 	}
		// }

	// type_switch
		// func main(){
		// 	// type switch语句,用于type_switch来判断某个interface变量中实际指向的变量类型
		// 	var x interface{}
		// 	var y = 10.2
		// 	x = y
		// 	switch i := x.(type){
		// 	case nil:
		// 		fmt.Printf("x的类型~:%T\n", i)
		// 	case int:
		// 		fmt.Printf("x的类型~:%T\n", i)
		// 	case float64:
		// 		fmt.Printf("x的类型~:%T\n", i)
		// 	case func(int) float64:
		// 		fmt.Printf("x的类型~:%T\n", i)
		// 	case bool, string:
		// 		fmt.Printf("x的类型~:%T\n", i)
		// 	default:
		// 		fmt.Println("未知类型")
		// 	}
		// }

	// 练习题
		// func main(){

			// // 使用switch把小写类型的char类型转化为大写(键盘输入),只能转化a,b,c,d,e其他输出"other"
			// var char byte
			// fmt.Println("请输入想要转换的字母:") 
			// fmt.Scanf("%c", &char)
			// switch char {
			// 	case 'a':  // 也可以根据ASI码来加一个数就是大写
			// 		fmt.Println("A")
			// 	case 'b':
			// 		fmt.Println("B")
			// 	case 'c':
			// 		fmt.Println("C")
			// 	case 'd':
			// 		fmt.Println("D")
			// 	case 'e':
			// 		fmt.Println("E")
			// 	default:
			// 		fmt.Println("other")
			// }
		
			// 根据学生成绩,输出对应的情况,大于60输出"合格", 低于60, 输出"不合格"
			// var sak int 
			// fmt.Println("请输入考试成绩:")
			// fmt.Scanln(&sak)
			// switch {
			// 	case sak < 60 && sak >= 0:
			// 		fmt.Println("成绩不合格")
			// 	case sak > 60 && sak <= 100:
			// 		fmt.Println("成绩合格")
			// 	default:
			// 		fmt.Println("成绩不能小于0或者大于100")
		
			// }
		
			// // 根据用户指定月份,打印该月份所属的季节,3,4,5春季,6,7,8夏季,9,10,11秋季,12,1,2冬季
			// var sook int
			// fmt.Println("请输入月份:")
			// fmt.Scanln(&sook)
			// switch sook {
			// 	case 3,4,5:
			// 		fmt.Println("春季")
			// 	case 6,7,8:
			// 		fmt.Println("夏季")
			// 	case 9,10,11:
			// 		fmt.Println("秋季")
			// 	case 12,1,2:
			// 		fmt.Println("冬季")
			// 	default:
			// 		fmt.Println("没有此月份")
		
			// }
		
		// }

		// switch和if的比较
		// 如果判断的是具体的值,而且符合整数,浮点数,字符,字符串这几种类型,建议使用switch语句,简洁高效
		// 其他情况,对于区间判断和结果为bool类型的判断,使用if,if的使用返回更广

	// for循环
		// func main(){
		// 	// for循环练习,执行的流程是这样的
		// 	// i定义一个数字,然后判断i是不是小于5,然后打印语句,之后执行i++
		// 	for i := 1; i <= 5; i++ {
		// 		fmt.Println("hello world!", i)
		// 	}
		// }

	// for循环的基本语法
		// for循环变量初始化,循环条件,循环变量迭代{
			// 循环后操作(语句)
		// }
	// 上面语法说明
		// 循环四要素
			// 1. 循环变量初始化
			// 2. 循环条件
			// 3. 循环操作(语句),有人也叫循环体
			// 4. 循环变量的迭代
		// for循环执行顺序说明
			// 1. 执行变量初始化,比如i=1
			// 2. 执行循环条件,比如i<=5
			// 3. 如果循环条件为真,就执行循环操作,执行代码语句
			// 4. 执行循环遍历浪迭代, 比如i++
			// 5. 反复执行2,3,4步骤,知道条件为Flase,退出循环

		// 注意,for循环中的i跳出循环之后是不能再引用的,除非是在for循环外面定义了变量之后,可以在for循环结束之后继续打印i
	// for循环的注意事项和细节
		// 1. 循环条件是返回一个布尔值的表达式
		// 2. for循环的第二种使用方式
			// func main(){
			// 	// for循环的第二种方法
			// 	j := 1  // 循环变量初始化
			// 	for j <= 10 {  // 循环条件
			// 		fmt.Println("hello haha", j)
			// 		j++  // 循环迭代
			// 	}
			// }
		// 3. for循环的第三种方式
			// func main(){
			// 	// for循环的第三种方式,死循环,配合break使用
			// 	k := 1
			// 	for {  // 等价于 for ; ; {}
			// 		if k <= 10 {
			// 			fmt.Println("hhahahaha")
			// 		} else {
			// 			break  // 如果上面的判断不成立,直接退出
			// 		}
			// 		k++
			// 	}
			// }
		// 4. golang提供for_range的方式,可以方便遍历字符串和数组
			// func main(){
			// 	// 字符串遍历方式1
				// sss := "hello,world!"
				// for i := 0; i < len(sss); i++ {
				// 	fmt.Printf("%c\n", sss[i])
				// }
			// 	// 用for_range方式
			// 	for index, val := range sss {
			// 		fmt.Printf("index=%d, val=%c\n", index, val)
			// 	}
			// }
		// for循环细节
			// 1. 如果我们的字符串含有中文,那么传统的遍历自妇产方式就是错误,会出现乱码的情况
			// 传统的字符串遍历是按照字节来遍历的,而一个汉字在utf-8编码是对应的3个字节
			// 解决方法:
				// 需要将string类型转成[]rune切片
				// func main(){
				// 	// 字符串遍历中文
				// 	var sss string = "hello,world!北京"
				// 	ss2s := []rune(sss)  // 就是把字符串转换成了切片
				// 	for i := 0; i < len(sss); i++ {
				// 		fmt.Printf("%c \n", sss[i])
				// 	}
				// }
			// for_range是可以循环中文的,但是index坐标会每隔3跳一次

		// for 循环练习
			// func main(){
			// 	// for循环练习,打印1~100之间所有9的倍数的整数的个数及总和
			// 	max := 100
			// 	num := 0
			// 	sum := 0
			// 	for i := 1; i <= max; i++ {
			// 		if i % 9 == 0 {
			// 			num += 1
			// 			sum += i
			// 		}
			// 	}
			// 	fmt.Println("和是", sum)
			// 	fmt.Println("个数是", num)
			// }
			//////////////////////////////////////////////
			// func main(){
			// 	// 完成下面表达式,6是可变的
			// 	n := 6
			// 	for i := 0; i <= n; i++ {
			// 		fmt.Printf("%v + %v = %v\n", i, n-i, n)
			// 	}
			// }

	// while循环在golang中是没有的,可以通过for循环来实现使用效果
		// for {
		// 	if 循环条件表达式成立{
		// 		break // 跳出循环
		// 	}
		// 	循环操作语句
		// 	循环变量迭代
		// }
		
		// 循环说明
		//1. for循环是一个无限循环
		// 2. break语句就是跳出for循环
			// func main(){
			// 	// 使用for循环,打印10句hello world,while循环类型的
			// 	vin := 0
			// 	for {
			// 		if vin >= 10 {
			// 			break
			// 		}
			// 		fmt.Println("hello world!")
			// 		vin ++
			// 	}
			// }

			// func main(){
			// 	// 使用do....while实现输出打印十句"hello world!ok"
			// 	vint := 0
			// 	for {
			// 		fmt.Println("hello world!ok")
			// 		vint++
			// 		if vint >= 10 {
			// 			break
			// 		}
			// 	}
			// }

// 三.多重循环控制
// 将一个循环放在另一个循环体中,就形成了嵌套循环,在外面的for循环称为外层循环,在里面的循环称为内层循环(建议循环不要超过三层)
// 实际上,嵌套循环就是把内层循环当成外层循环的循环体,当只有内层循环的循环条件为False时,才会跳出内层循环,也就是说结束外层循环的当次循环,然后执行下次循环
// 设外层循环次数为m次,内层为n次,则内层循环实际上需要执行m*n=mn次
// 应用案例
// 统计3个班的成绩情况,每个班有5名同学,求各个班级的平均分和所有班级的平均分(学生的成绩从键盘输入)
// 先易后难,将一个复杂的问题分解成简单的问题
// 先死后活
// func main(){
	// 统计3个班成绩情况,每个班有5名同学
	// 求各个班的平均分,和所有班级的平均分
	// 先求出一个班的平均分
	// class_num := 3
	// student_num := 5
	// all_num := 0
	// for j := 1; j <= class_num; j++ {
	// 	sum := 0
	// 	fmt.Printf("这是%d班\n",j)
	// 	for i := 1; i <= student_num; i++ {
	// 		var sork int
	// 		fmt.Println("请输入学生成绩:")
	// 		fmt.Scanln(&sork)
	// 		sum += sork
	// 	}
	// 	fmt.Println("这个班级的平均成绩是:", sum / student_num)
	// 	all_num += sum
	// }
	// fmt.Printf("所有班的总成绩%d\n,所有班的平均分%d\n", all_num, all_num / (student_num * class_num))


	// 统计三个班的及格人数,每个班有5人
	// student_sum := 0
	// for j := 1; j <= class_num; j++ {
	// 	fmt.Printf("这是%d班\n",j)
	// 	for n := 1; n <= student_num; n++ {
	// 		var sork int
	// 		fmt.Printf("请输入第%d学生成绩:", n)
	// 		fmt.Scanln(&sork)
	// 		if sork >= 60 {
	// 			student_sum += 1
	// 		}
	// 	}
	// }
	// fmt.Printf("三个班的及格人数%d\n", student_sum)

	// for循环打印一个矩形
// 	for i := 1; i <= 5; i++ {
// 		for j := 1; j <= 5; j++ {
// 			fmt.Print("*")  // 打印的内容不换行
// 		}
// 		fmt.Println()
// 	}
// 	// for循环打印半个金字塔
// 	for m := 1; m <=5; m++ {
// 		// 循环第一层
// 		for jj := 1; jj <= m; jj++ {
// 			fmt.Print("*")
// 		} 
// 		fmt.Println()
// 	}
// 	// for循环打印一个金字塔
// 	for mm := 1; mm <= 5; mm++ {
// 		// 打印总层数 - 当前层数的空格
// 		for kk := 1; kk <= 5 - mm; kk++ {
// 			fmt.Print(" ")
// 		}
// 		// 打印的层数
// 		for nj := 1; nj <= 2 * mm - 1; nj++ {
// 			fmt.Print("*")
// 		} 
// 		fmt.Println()
// 	}
// 	// for循环打印一个空心金字塔
// 	// 每一层的第一个和最后一个是打印"*", 其他的是空的
// 	for mm := 1; mm <= 5; mm++ {
// 		// 打印总层数 - 当前层数的空格
// 		for kk := 1; kk <= 5 - mm; kk++ {
// 			fmt.Print(" ")
// 		}
// 		// 打印的层数
// 		for nj := 1; nj <= 2 * mm - 1; nj++ {
// 			// 每层的第一个跟最后一个打印"*"
// 			if nj == 1 || nj == 2 * mm - 1 || mm == 5{
// 				fmt.Print("*")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 		}
		
// 		fmt.Println()
// 	}

// 	// for循环打印九九乘法表
// 	for ss := 1; ss <= 9; ss++ {
// 		// 定义层数
// 		for sn := 1; sn <= ss; sn++ {
// 			// 乘法
// 			fmt.Printf("%d * %d = %d\t", sn, ss, sn * ss)
// 		}
// 		fmt.Println()
// 	}
// }

// // 判断一个整数属于什么范围,大于0,小于0,等于0
// func main(){
// 	// 用户终端输入
// 	var num int
// 	// fmt.Println("请输入数字:")
// 	// fmt.Printf(&num)
// 	// if num < 0 {
// 	// 	fmt.Println("这个数字小于0")
// 	// }
// 	// if num > 0 {
// 	// 	fmt.Println("这个数字大于0")
// 	// }
// 	// if num == 0 {
// 	// 	fmt.Println("这个数字等于0")
// 	// }
// 	// 判断一个整数是否是水仙花数,指一个三位数,其各个位数上的数立方和等于其本身
// 	// 例如:153 = 1*1*1 + 5*5*5 + 3*3*3
// 	fmt.Println("请输入数字,判断是否是水仙花数:")
// 	fmt.Scanln(&num)
// 	if len(num) == 3 {
// 		// 将每个数字都立方然后加起来
// 		var i int
// 		var m int
// 		var u int
// 		i = num[0] * num[0] * num[0]
// 		m = num[1] * num[1] * num[1]
// 		u = num[2] * num[2] * num[2]
// 		if i + m + u == num {
// 			fmt.Println("此数字是水仙花数")
// 		} else {
// 			fmt.Println("此数字不是水仙花数")
// 		}
// 	} else {
// 		fmt.Println("请输入3位数的数字")
// 	}
// }


// // 跳转控制语句-break
// // 随机生成1-100之间的一个数,直到生成99,查看一共用了几次就获取到了99这个数
// func main(){
// 	// go中随机生成的函数是rand.Intn()
// 	// rand.Intn(100) + 1  // 需要引入包,如果只是这样,随机生成的数就是[0 - 100),是不包含100的,所以要+1
// 	var num int = 0
// 	for {
// 		// 为了生成一个随机数,还需要给rand设置一个种子,引入time包
// 		// time.Now().Unix():会返回一个秒数,从1970:01:01的0时0分0秒到现在的秒数,时间戳时间
// 		rand.Seed(time.Now().UnixNano())  // 纳秒
// 		inn := rand.Intn(100) + 1
// 		fmt.Println(inn)
// 		num++
// 		if inn == 99 {
// 			break
// 		}
// 	}
// 	fmt.Printf("生成99时,用了%d次\n", num)
// }

// break跳转控制语句的注意事项跟细节说明
// break 语句出现在多层嵌套的语句块中,可以通过标签致命要终止的是那一层语句块
// 标签的基本使用
	// label1: {
	// 	label2: {
	// 		label3: {
	// 			if 条件 {
	// 				break label1 // 直接跳出整个循环
	// 				break label2  // 跳到label2的循环中
	// 				break label3 // 跳出当前循环  如果break后面没有标签,默认跳出最近的循环
	// 			}
	// 		}
	// 	}
	// }
// func main() {
// 	// 指定标签的形式来使用break
// 	label2:  // 设置一个标签
// 	for i := 0; i < 4; i++ {
// 		// label1:
// 		for j := 0; j < 10; j++ {
// 			if j == 2{
// 				break label2
// 			}
// 			fmt.Println("j = ", j)
// 		}
// 	}
// }

// func main() {
// 	// 实现登录验证,有三次机会,如果用户名为"张无忌",密码"888"提示登陆成功,否则提示还有几次机会
// 	num := 3
//  var name string
// 	var password int
// 	for i := 1; i <= 3; i++{
// 		fmt.Println("请输入用户名:")
// 		fmt.Scanln(&name)
// 		fmt.Println("请输入密码:")
// 		fmt.Scanln(&password)
// 		// fmt.Println(name,password)
// 		if name == "张无忌" && password == 888 {
// 			fmt.Println("登陆成功")
// 			break
// 		} else {
// 			num--
// 			fmt.Printf("登陆失败,还有%d机会\n", i)
// 		}
// 	}
// }

// func main() {
// 	// 100以内的数求和,求出 当和第一次大于20的当前数
// 	sum := 0
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 		if sum > 20 {
// 			fmt.Println(i)
// 			break
// 		}
// 	}
// }

// 跳转控制语句-continue
// continue语句用于结束本次循环,继续执行下一次循环
// continue语句出现在多层嵌套的循环语句体现中,可以通过标签指明要跳过的是哪一层循环,这个和break前面的标签的使用规则一样
// 基本语法
	// { ...
	// 	continue
	// } ...
// 案例分析continue
// func main() {
// 	for i := 1; i < 4; i++ {
// 		for j := 1; j < 10; j++ {
// 			if j == 3 {
// 				continue
// 			}
// 			fmt.Println(j)
// 		}
		
// 	}
// }

// 跳转控制语句-goto
// go语言的goto语句可以无条件地转移到程序中指定的行
// goto语句通常与条件语句配合使用,可用来实现条件转移,跳出循环体功能
// 在go程序设计中,一般不主张使用goto语句,以免造成程序流程的混乱,使理解和调试程序都产生困难
// 基本语法
	// goto label  //当程序执行到这一行时,会直接跳转到此label,下面的代码就不执行了
	// ...  // 不执行
	// label:statement  // 跳转到这一行

import (
	"fmt"
	"strings"
)


func main() {
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		fmt.Println(word)
		wordCount[word] = 1
	}
	fmt.Println(wordCount)
}