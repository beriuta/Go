package main

import (
	"fmt"
	// "bytes"
	// "log"
	// "runtime"
	// "time"
	"strconv"
	"strings"
//    "runtime"
//     "os"
)
// // num := 10  // 全局变量不能使用:=赋值, 全局变量是允许声明但是不使用的
// func main() {
//     var goos string = runtime.GOOS
//     fmt.Printf("The operating system is: %s\n", goos)  // The operating system is: linux
//     path := os.Getenv("PATH")
//     fmt.Printf("Path is %s\n", path)  // Path is /home/beriuta/.pyenv/plugins/pyenv-virtualenv/shims:/home/beriuta/.pyenv/shims:/home/linuxbrew/.linuxbrew/bin:/home/linuxbrew/.linuxbrew/sbin:/home/beriuta/.pyenv/bin:/sbin:/home/beriuta/.local/bin:/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games:/usr/sbin:/opt/go/bin
//     fmt.Printf("Path is %s\n", num)  // Path is /home/beriuta/.pyenv/plugins/pyenv-virtualenv/shims:/home/beriuta/.pyenv/shims:/home/linuxbrew/.linuxbrew/bin:/home/linuxbrew/.linuxbrew/sbin:/home/beriuta/.pyenv/bin:/sbin:/home/beriuta/.local/bin:/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games:/usr/sbin:/opt/go/bin
// }


// func main() {

// 	// 复数使用 re+imI 来表示，其中 re 代表实数部分，im 代表虚数部分，I 代表根号负 1

// 	// var c1 complex64 = 5 + 10i
// 	// fmt.Printf("The value is: %v", c1)  // The value is: (5+10i)
// 	// var re, im float32
// 	// c := complex(re, im)  // 将float32转换成complex64复数
// 	// fmt.Printf("The value is: %v", c)  // The value is: (0+0i)
// 	// fmt.Printf("实数为: %v", real(c))  // 实数为: 0
// 	// fmt.Printf("虚数是: %v", imag(c))  // 虚数是: 0

// 	// var ch int = '\u0041'
// 	// var ch2 int = '\u03B2'
// 	// var ch3 int = '\U00101234'
// 	// fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer  数字类型  65 - 946 - 1053236
// 	// fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character  字符类型  A - β - r
// 	// fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes  字节类型  41 - 3B2 - 101234
// 	// fmt.Printf("%U - %U - %U\n", ch, ch2, ch3) // UTF-8 code point  字码类型  U+0041 - U+03B2 - U+101234

// 	// unicode 包
// 	// 判断是否为字母：unicode.IsLetter(ch)
// 	// 判断是否为数字：unicode.IsDigit(ch)
// 	// 判断是否为空白符号：unicode.IsSpace(ch)

// 	// // strings 包
// 	// s := "fdss"
// 	// // HasPrefix 判断 s 是否以 prefix 开头
// 	// ret := strings.HasPrefix(s, "prefix")
// 	// fmt.Print(ret)

// 	// // HasSuffix 判断 s 是否以 prefix 结尾
// 	// ret1 := strings.HasSuffix(s, "prefix")
// 	// fmt.Print(ret1)

// 	// // Contains 判断字符串 s 是否包含 substr
// 	// ret2 := strings.Contains(s, "substr")
// 	// fmt.Print(ret2)

// 	// // 测试Map函数跟strings.IndexFunc()
// 	// rett := "ss;lfk;sdkfsf"
// 	// res= strings.Map(rett)
// 	// rr := strings.IndexFunc()
// 	// fmt.Println(res)

// 	// // strconv 包
// 	// var orig string = "666"
//     // var an int
//     // var newS string

//     // fmt.Printf("The size of ints is: %d\n", strconv.IntSize)      

//     // an, _ = strconv.Atoi(orig)  // 将字符串转换为int类型 666
//     // fmt.Printf("The integer is: %d\n", an) 
//     // an = an + 5
//     // newS = strconv.Itoa(an)  // 将int类型数字转换成字符串
// 	// fmt.Printf("The new string is: %s\n", newS)  // 671
	
// 	// // time 包
// 	// t := time.Now()  // 实例化一个时间对象
//     // fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
//     // fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
//     // // 21.12.2011
//     // t = time.Now().UTC()
//     // fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
//     // fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
// 	// // calculating times:
//     // week := 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
//     // week_from_now := t.Add(week)
//     // fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
//     // // formatting times:
//     // fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
//     // fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
//     // fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
//     // s := t.Format("20060102")
// 	// fmt.Println(t, "=>", s)
	
// 	// // switch 分支
// 	// k := 6
// 	// switch k {
// 	// 	case 4: fmt.Println("was <= 4"); fallthrough;
// 	// 	case 5: fmt.Println("was <= 5"); fallthrough;
// 	// 	case 6: fmt.Println("was <= 6"); fallthrough;
// 	// 	case 7: fmt.Println("was <= 7"); fallthrough;
// 	// 	case 8: fmt.Println("was <= 8"); fallthrough;
// 	// 	default: fmt.Println("default case")
// 	// }

// 	// // 结果
// 	// // was <= 6
// 	// // was <= 7
// 	// // was <= 8
// 	// // default case


// 	// // for 循环
// 	// s := ""
// 	// for ; s != "aaaaa"; {
// 	// 	fmt.Println("Value of s:", s)
// 	// 	s = s + "a"
// 	// }
// 	// for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
// 	// 	s = i+1, j+1, s + "a" {
// 	// 	fmt.Println("Value of i, j, s:", i, j, s)
// 	// }

// 	// label

// 	// LABEL1:
// 	// for i := 0; i <= 5; i++ {
// 	// 	for j := 0; j <= 5; j++ {
// 	// 		if j == 4 {
// 	// 			continue LABEL1
// 	// 		}
// 	// 		fmt.Printf("i is: %d, and j is: %d\n", i, j)
// 	// 	}
// 	// }

// 	// x := min(1, 3, 2, 0)
//     // fmt.Printf("The minimum is: %d\n", x)
//     // arr := []int{7,9,3,5,1}
//     // x = min(arr...)
// 	// fmt.Printf("The minimum in the array arr is: %d\n", x)
// 	// fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
//     // fmt.Printf("%d is odd: is %t\n", 17, odd(17))
//     // // 17 is odd: is true
// 	// fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	

// 	// addBmp := MakeAddSuffix(".bmp")
// 	// addJpeg := MakeAddSuffix(".jpeg")

// 	// d1 := addBmp("file") // returns: file.bmp
// 	// d2 := addJpeg("file") // returns: file.jpeg
// 	// fmt.Println(d1)
// 	// fmt.Println(d2)

// 	// // 检测执行到哪个函数了,用log文件
// 	// where := func() {
// 	// 	_, file, line, _ := runtime.Caller(1)
// 	// 	log.Printf("%s:%d", file, line)
// 	// }
// 	// where()
// 	// // some code
// 	// fmt.Println("完成")
// 	// fmt.Println("完成")
// 	// fmt.Println("完成")
// 	// fmt.Println("完成")
// 	// fmt.Println("完成")
// 	// fmt.Println("完成")
// 	// where()
// 	// // some more code
// 	// where()

// 	// log.SetFlags(log.Llongfile)
// 	// log.Print("")
// 	// fmt.Println("完成")

// 	// start := time.Now()  // 开始时间 
// 	// longCalculation()
// 	// end := time.Now()
// 	// delta := end.Sub(start)  // 结束时间的差值,计算花费了多少时间
// 	// fmt.Printf("longCalculation took this amount of time: %s\n", delta)

// 	// var result uint64 = 0
//     // start := time.Now()
//     // for i := 0; i < LIM; i++ {
//     //     result = fibonacci(i)
//     //     fmt.Printf("fibonacci(%d) is: %d\n", i, result)
//     // }
//     // end := time.Now()
//     // delta := end.Sub(start)
// 	// fmt.Printf("longCalculation took this amount of time: %s\n", delta)
// 	// // fmt.Println(fibs)
	
// 	// new()跟直接var arr2 = [5]int new是指针类型的数据,arr2是值拷贝
// 	// var arr1 = new([5]int)
// 	// fmt.Println(arr1)  // &[0 0 0 0 0] 直接new()是一个指针类型
// 	// fmt.Println(*arr1)  // [0 0 0 0 0] 取值用*指针变量名
// 	// var splice []int
// 	// fmt.Println(splice)  // []
// 	// arr1 := []int{1, 2, 3, 4, 6, 7, 8, 9}
// 	// splice = arr1[:len(arr1)]
// 	// fmt.Println(splice)  // [1 2 3 4 6 7 8 9]
// 	// splice1 := make([]int, 20)
// 	// fmt.Println(splice1)  // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// 	// fmt.Println(cap(splice1))  // 20
// 	// splice2 := make([]int, 3, 10)
// 	// fmt.Println(splice2)  // [0 0 0]
// 	// // new 跟 make的区别
// 	// // new(T) : 返回一个指向类型为 T，值为 0 的地址的指针, 一般适用于数组跟结构体 它相当于 &T{}
// 	// // make(T) : 返回类型为T的初始值, 它只适用于3种内建的引用类型：切片、map 和 channel
// 	// var p *[]int = new([]int) 
// 	// fmt.Println("p=",p)  // p= &[]
// 	// fmt.Println("p=",*p)  // p= []
// 	// fmt.Println("p的内容=",&p)  // p的内容= 0xc00009a020
// 	// p1 := new([]int)
// 	// fmt.Println(*p1)  // []
// 	// p2 := make([]int, 0)
// 	// fmt.Println(p2)  // []
// 	// 综上所述,new是一个指针类型,本身存储的是切片的地址,make是切片已经被初始化,但是指向了一个空的/数组
// 	// s := make([]byte, 5)
// 	// fmt.Println(s)  // [0 0 0 0 0]
// 	// s1 := s[2:4]
// 	// fmt.Println(cap(s1))  // 3
// 	// fmt.Println(len(s1))  // 2
// 	// s1[1] = 55
// 	// fmt.Println(s)  // [0 0 0 55 0]
// 	// fmt.Println(s1)  // [0 55]
// 	// s[4] = 22
// 	// fmt.Println(s)  // [0 0 0 55 22]
// 	// fmt.Println(s1)  // [0 55]
// 	// // 切片是引用类型,直接能操作原数组

// 	// // bytes包
// 	// var buffer bytes.Buffer
// 	// for {
// 	// 	if s, ok := getNextString(); ok { //method getNextString() not shown here
// 	// 		buffer.WriteString(s)  // 追加在后面
// 	// 	} else {
// 	// 		break
// 	// 	}
// 	// }
// 	// fmt.Print(buffer.String(), "\n")
// 	// // 这种实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候。

// 	// for-range结构
// 	// splice := make([]int, 10)
// 	// for k, v := range splice {
// 	// 	splice[k] = v + 3 
// 	// }
// 	// fmt.Println(splice)
// 	// session := []string{"Spring", "Summer", "Autumn", "Winter"}
// 	// for key := range session {
// 	// 	fmt.Println(key)  // 0 1 2 3 可以直接不写值,只取index
// 	// }
// 	// for _, value := range session {
// 	// 	fmt.Println(value)  // 取值
// 	// }
// 	// splice := [...]int{2, 200, 32, 432, 32, 333}
// 	// ret := minAttr(splice)
// 	// fmt.Println(ret)

// 	// slice1 := make([]int, 0, 10)
//     // // load the slice, cap(slice1) is 10:
//     // for i := 0; i < cap(slice1); i++ {
//     //     slice1 = slice1[0:i+1]
//     //     slice1[i] = i
//     //     fmt.Printf("The length of slice is %d\n", len(slice1))
//     //     fmt.Printf("The length of slice is %v\n", slice1)
// 	// }
	
// 	// 切片的复制(copy)和追加(append)
// 	selFrom := []int{1,2,3}
// 	selTo := make([]int, 10)
// 	n := copy(selTo, selFrom)
// 	fmt.Println(n)  // 3
// 	sel3 := []int{1,2,3}
// 	sel3 = append(sel3, 4, 5, 6)


// }
// const LIM = 41
// var fibs [LIM]uint64

// func fibonacci(n int) (res uint64) {
//     // memoization: check if fibonacci(n) is already known in array:
//     if fibs[n] != 0 {
// 		fmt.Println("执行我了...........", fibs[n])
// 		res = fibs[n]
//         return
//     }
//     if n <= 1 {
//         res = 1
//     } else {
//         res = fibonacci(n-1) + fibonacci(n-2)
//     }
//     fibs[n] = res
//     return
// }

// var where = log.Print
// func func1() {
// where()
// 	// ... some code
// where()
// 	// ... some code
// where()
// 	}

// func MakeAddSuffix(suffix string) func(string) string {
//     return func(name string) string {
//         if !strings.HasSuffix(name, suffix) {
//             return name + suffix
//         }
//         return name
//     }
// }
// func minAttr(a [6]int) int {
// 	minNum := 0
// 	for index:= range a {
// 		if index == 0 {
// 			minNum = a[index]
// 		} else if a[index] < minNum {
// 			minNum = a[index]
// 		}
// 	}
// 	return minNum
// }

func main() {
	m := "2333"
	m = strings.Join([]string{m, strconv.Itoa(1)}, "+")
	fmt.Println(m)
}
