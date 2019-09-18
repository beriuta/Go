package main
// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )
// // 字符串中常用的函数
// 说明:字符串在我们程序开发中,使用的是非常多的,常用的函数有
	// 统计字符串的长度,按字节len(str)
	// 字符串遍历,同事处理有中文的问题 r := []rune(str)
	// 字符串转整数 : n, ner := strconv.Atoi("12")
	// 整数转字符串 : str = strconv.ltoa(12345)
	// 字符串 转 byte类型 : var bates = []byte["hello go"]
	// []byte转字符串 : str = string([]byte{97,98,99})
	// 10进制转2,8,16进制: str = strconv.Formatint(123,2)  // 将123转成二进制
	// 查找字符是否在字符串中: strings.Contains("seafood", "foo")  // true
	// 统计一个字符串有几个指定的字符: strings.Count("ceheese", "e") // 4
	// 不区分大小写的字符串比较: fmt.Println(strings.EqualFold("abc", "ABC")) // true
	// 返回字符在字符串第一次出现的index值,如果没有返回-1: strings.Index("NLT_abc", "abc") // 4坐标

	//字符串函数练习
		// func main() {
		// 	// // 统计字符串的长度,按字节len(str)
		// 	// str := "hello上海"  // golang编码统一是utf8(ascii的字符(子母和数字占一个字节),汉字占三个字节)
		// 	// fmt.Println("str的长度:", len(str))  // 11

		// 	// // 字符串的遍历,解决中文的问题
		// 	// str2 := "hello上海" 
		// 	// rr := []rune(str2)  // 想要展示中文,需要先用rune切片,然后再使用
		// 	// for i := 0; i < len(rr); i++ {
		// 	// 	fmt.Printf("每个字节:%c\n", rr[i])
		// 	// }

		// 	// 转数字
		// 	n, err := strconv.Atoi("123")
		// 	if err != nil {
		// 		fmt.Println("转换错误:", err)
		// 	} else {
		// 		fmt.Println("转换结果是:", n)
		// 	}

		// 	// 整数转字符串
		// 	str := strconv.Itoa(12345)
		// 	fmt.Printf("数字转字符串:%v, 类型是%T\n", str, str)

		// 	// 字符串转byte类型
		// 	var num = []byte("hello go")
		// 	fmt.Println("输出的byte:", num)  // [104 101 108 108 111 32 103 111]

		// 	// []byte转字符串
		// 	str = string([]byte{97, 98, 99})  // 注意是{}
		// 	fmt.Println("byte转字符串:", str)  // abc

		// 	// 10进制转2,8,16进制,
		// 	str = strconv.FormatInt(123,2)  // 将123转成3进制,返回的是string类型
		// 	fmt.Println("十进制转2进制:", str)  // 1111011
		// 	str = strconv.FormatInt(123,16)
		// 	fmt.Println("十进制转16进制:", str)  // 7b

		// 	// 查找的字符是否存在字符串中
		// 	bbb := strings.Contains("seafood", "foo")  // 返回的是bool值
		// 	fmt.Println("是否在此字符串中:", bbb)   // true

		// 	// 计算一个字符串有几个
		// 	ccc := strings.Count("ceheese", "e")
		// 	fmt.Printf("此字符串有%d个e\n", ccc)

		// 	// 比较这俩字符串是否相同,如果用"=="来区分,就会对比到字母的大小写,如果不区分大小写
		// 	nn := strings.EqualFold("abc", "ABC")
		// 	fmt.Println("不区分大小写的对比:", nn)  // true
		// 	fmt.Println("abc" == "ABC")  // false

		// 	// 返回某个指定的字符在这个字符串中第一次出现的索引位置,如果没有就返回-1
		// 	dd := strings.Index("MLT_abc", "ab")
		// 	fmt.Println(dd)  // 4

		// 	// 返回字符串中字符最后出现的一次的坐标位置,没有返回-1
		// 	zd := strings.LastIndex("go golang", "go")  // 3 到g
		// 	fmt.Println(zd)

		// 	// 将指定的字符串替换成别的信息
		// 	str1 := "go golang go go go"
		// 	sp := strings.Replace(str1, "go", "go语言", 2)  // 最后一个数字可以指定替换几个,如果全部替换就是-1
		// 	fmt.Println(sp)  // str1本身并没有变化

		// 	// 按照指定的某个字符,为分割标识,将一个字符拆分成字符串数组
		// 	meu := "we,rt,yu,df,we,tt"
		// 	sqq := strings.Split(meu,",") 
		// 	fmt.Println(sqq)  // [we rt yu df we tt]

		// 	// 将字符串的字母进行大小写的转换
		// 	as := "AANEeds"
		// 	sqqq := strings.ToLower(as)
		// 	fmt.Println("转换成小写", sqqq)
		// 	sqa := strings.ToUpper(as)
		// 	fmt.Println("转换成大写", sqa)

		// 	// 将字符串左右两边的空格去掉
		// 	ndn := "  swd dlsk dlke sddsf "
		// 	saw := strings.TrimSpace(ndn)
		// 	fmt.Println(saw)  // swd dlsk dlke sddsf

		// 	// 将字符串左右两边的指定的字符去掉
		// 	ndnw := "!eow!"
		// 	aaad := strings.Trim(ndnw, "!")
		// 	fmt.Println(aaad)  // eow

		// 	// 将字符串左边指定的字符去掉
		// 	aaad = strings.TrimLeft(ndnw, "!")
		// 	fmt.Println(aaad)  // eow!
		// 	aaad = strings.TrimRight(ndnw, "!")
		// 	fmt.Println(aaad)  // !eow

		// 	// 判断字符串中是否以指定字符开头的
		// 	df := "Wfjf,dfrr,Ddd,Wgfg"
		// 	sai := strings.HasPrefix(df, "W")
		// 	fmt.Println(sai)  // true

		// 	// 判断字符串中是否以指定字符结束的
		// 	jiu := strings.HasSuffix(df, "W")
		// 	fmt.Println(jiu)  // false

		// }