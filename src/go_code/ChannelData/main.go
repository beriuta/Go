package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// 管道可以声明为只读或者只写
// 	// 1. 在默认情况下,管道是双向
// 	// var chan1 chan int // 可读可写
// 	// 2. 声明为只写
// 	var chan2 chan<- int
// 	chan2 = make(chan int, 3)
// 	chan2<- 100
// 	// num := <-chan2  // 直接报错

// 	// 3. 声明为只读
// 	var chan3 <-chan int
// 	chan3 = make(chan int)
// 	// chan3<- 20   // 声明的只读模式,这里就不能写了
// 	num2 := <-chan3
// 	fmt.Println(num2)
// 	// 协程参数是可以限制的,只写,只读
// }

// Channel使用细节和注意事项
// 使用select可以解决从管道取数据的阻塞问题
// goroutine中使用recover,解决协程中出现的panic,导致程序崩溃问题
// 说明:
	// 如果我们起了一个协程,但是这个协程出现了panic,如果我们没有播火这个panic,就会造成整个程序崩溃,
	// 这时我们可以在goroutine中使用recover来捕获panic,进行处理,这样即使这个协程发生的问题,但是主线程仍然不受影响,可以继续执行
// func main() {
	
// 	// 使用select可以结局让从管道取数据的阻塞问题

// 	// 定义一个管道, 放入10个值
// 	intChan := make(chan int, 10)
// 	for i := 0; i < 10; i++ {
// 		intChan<- i
// 	}

// 	// 定义一个管道 5个数据string
// 	stringChan := make(chan string, 5)
// 	for i := 0; i < 5; i++ {
// 		stringChan<- "测试" + fmt.Sprintf("%d", i)
// 	}

// 	// 传统方法在遍历管道事,如果不关闭会阻塞导致deadlock
// 	// 在实际开发中,可能我们不能确定什么时候关闭该管道
// 	// 可以使用select方式可以解决
// 	for {
// 		select {
// 			case v := <-intChan:  // 如果intChan一直没有关闭,不会一直阻塞而变成deadlock
// 				fmt.Printf("从intChan读取的数据%d\n", v)  // 如果在管道中取不到值,也不会报错,而是默认到下一个case中
// 			case v := <-stringChan:	
// 				fmt.Printf("从stringChan读取的数据为%s\n", v)
// 			default:
// 				fmt.Printf("都取到了~~~不玩了,这里可以加入业务逻辑\n")
// 				return  // 这里要添加return,如果添加break.不会退出循环
// 		}
// 	}
// }

// 写一个函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)  // 睡一秒钟
		fmt.Println("hello, world!!!")
	}
}

func test() {

	// // 定义了一个map
	// var myMap map[int]string
	// // 这里应该先make一下,才能插入值
	// myMap[0] = "golang"

	// 这里可以使用defer + recover
	defer func() {
		// 捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test函数出错了,错误是:", err)
		}
	}()  // 这样如果这个函数出错了,只是提示一下,但是主函数不会停止,而是继续执行
	var myMap map[int]string
	// 这里应该先make一下,才能插入值
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()  // 即使这个协程发现了问题,主线程也不会有影响

	for i := 0; i < 20; i++ {
		fmt.Println("main()函数", i)
		time.Sleep(time.Second)  // test函数出了错误,导致主函数也停止了.在实际开发中,不应该这样 
		
	}
}