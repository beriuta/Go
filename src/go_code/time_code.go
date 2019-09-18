package main
// import (
// 	"fmt"
// 	"time"
// 	// "strconv"
// )


// func test03() {
// 	str := ""
// 	for i := 0; i < 100000; i++ {
// 		str += strconv.Itoa(i)
// 		// fmt.Println(i)
// 	}
// }






// // 关于日期跟时间的函数  time包
// func main() {
// 	// // time.time时间类型
// 	// now := time.Now()
// 	// fmt.Printf("now=%v, now类型%T\n", now, now)  // now=2019-07-10 16:08:32.640807807 +0800 CST m=+0.000031846, now类型time.Time

// 	// // 通过这个now,能获取到年月日
// 	// fmt.Printf("年=%v\n", now.Year())
// 	// fmt.Printf("月=%v\n", now.Month())
// 	// fmt.Printf("日=%v\n", now.Day())
// 	// fmt.Printf("时=%v\n", now.Hour())
// 	// fmt.Printf("分=%v\n", now.Minute())
// 	// fmt.Printf("秒=%v\n", now.Second())

// 	// // 格式化日期时间
// 	// // 第一种方式就按照上面的格式输出
// 	// // 第二种方式,用Format来输出,里面的数字不能改
// 	// fmt.Printf(now.Format("2006/01/02 15:04:05"))
// 	// fmt.Println()
// 	// fmt.Printf(now.Format("2006-01-02"))
// 	// fmt.Println()
// 	// fmt.Printf(now.Format("15:04:05"))
// 	// fmt.Println()
// 	// fmt.Printf(now.Format("01"))  // 返回的就是月份
// 	// fmt.Println()

// 	// // 时间常量
// 	// // 每隔1秒钟打印一个数字,输出100个就结束
// 	// // time.Second 秒
// 	// // time.Millisecond  毫秒
// 	// for i := 0; i <= 100; i++ {
// 	// 	fmt.Println(i)
// 	// 	time.Sleep(time.Millisecond * 100)
// 	// }

// 	// // 获取当前unix时间戳,和unixnano时间戳纳秒
// 	// fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
// 	// // unix时间戳=1562747788 unixnano时间戳=1562747788490285633

// 	// 统计某一段函数执行的时间
// 	// 编写一个函数,执行之前先执行一遍时间
// 	// 执行之后再执行一遍时间
// 	// 两个时间相减就是函数执行的时间
// 	s_time := time.Now().Unix()
// 	// test03()
// 	e_tiem := time.Now().Unix()
// 	fmt.Printf("耗时%v秒\n", e_tiem-s_time)  // 耗时6秒

// }

