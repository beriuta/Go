package main

import (
	"fmt"
	// "math/rand"
	"time"
	// "strings"
	// "strconv"
	// "sync"
)

// 如果要计算1-200的数的阶乘,并且把哥哥书的阶乘放入到map中,最后显示出来,要求使用goroutine
// 分析思路
// 使用goroutine来完成,效率高,但是会出现并发/并行安全问题
// 这就提出不同goroutine如何通信
// 代码实现
// 使用goroutine来完成(如果用goroutine并发会出现什么问题)
// 在运行某个程序时,如何知道是否存在资源竞争问题, 在编译改程序时,增加一个参数 -race即可
// var (
// 	myMap = make(map[int]int, 10)
// 	// 声明一个全局互斥锁
// 	// lock 是全局额互斥锁
// 	// sync : 包, 同步的意思
// 	// Mutex : 互斥的意思
// 	lock sync.Mutex
// )

// func test(n int) {
// 	// 计算数字之间的阶乘
// 	// 因为是并发执行,又要阶乘,所以不确定执行到哪里,只能将每次传入的值所对应的阶乘存储到map中
// 	res := 1
// 	for i := 1; i <= n; i++ {
// 		res += i
// 	}
// 	lock.Lock()  // 加锁
// 	myMap[n] = res  // concurrent map writer
// 	lock.Unlock()
// }

// func main() {

// 	// 这里开启多个协程
// 	for i := 1; i <= 200; i++ {
// 		go test(i)
// 	}
// 	// 休眠十秒钟
// 	time.Sleep(time.Second * 10)  // 这里要等多少秒,是未知的,没办法判断主线程的结束时间,所以需要管道Channel
// 	lock.Lock()  // 加锁
// 	for i, v := range myMap {
// 		fmt.Printf("map[%d]=%d\n", i, v)  // 当开启200个协程的时候,没有打印,因为主线程之行结束了,但是协程没有执行结束也被迫中断了
// 		// 打印出来的内容里面有负数,说明已经越界了
// 	}
// 	lock.Unlock()
// 	// 解决完之后空白问题,出现了并发问题,因为这200个协程在操作同一个map,go build -race channel.go 可以查看因为资源争夺的问题
// 		// 写必须要有保护,读可以直接读,但是写必须要加锁
// 	// 不同goroutine之间如何通讯
// 		// 1,全局变量加锁同步
// 		// 2. Channel
// 	// 使用全局变量加锁同步改进程序
// 		// 因为没有对全局变量m加锁,因此会出现资源争夺问题,代码会出现错误,提示concurrent map writer
// 		// 解决方案,加入互斥锁
// 		// 我们的数的阶乘很大,结果会越界,可以将求阶乘改成 sum += uin(64(i))
// 	// 当协程发现有锁并且没有释放锁,其余的协程会在队列中排队,加锁的包(sync)

// 	// 为什么需要Channel
// 		// 前面使用全局变量加锁同步来解决goroutine的通讯,并不完美
// 		// 主线程在等待所有goroutine全部完成的时间很难确定,所以暴力设置了10秒钟,仅仅是估算
// 		// 如果主线程休眠时间过长,会加长等待的时间,如果等待时间短了,可能会有一些goroutine在工作,而因为主线程退出而销毁
// 		// 通过全局变量加锁同步来实现通讯,也并不利用多个协程对全局变量的读写操作
// 		// 因为上述的不完美,所以引除了Channel管道
// 	// Channel介绍
// 		// Channel本质就是一个数据结构-队列(先进先出)[FIFO]
// 		// 线程安全,多goroutine访问时,不需要加锁,就是说Channel本身就是线程安全的
// 		// Channel是有类型的,一个string的Channel只能存放string类型的数据,如果要放多种类型,就把管道声明成interface,取的时候需要用类型断言
// 	// Channel基本使用
// 		// var 变量名 chan 数据类型
// 		// 比如 : var intChan chan int, var mapChan chan map[int]string, var perChan chan Person, var perChen chen *Person
// 		// channel 是引用类型
// 		// Channel必须初始化才能写入数据,即make()后才能使用
// 		// 管道是有类型的,比如intChan只能写入int类型
// }

// func main() {
// 	// 创建一个可以存放3个int类型的管道
// 	var intChan chan int
// 	intChan = make(chan int, 3)
// 	// 变量值是一个地址---intChan
// 	fmt.Printf("intChan=%v, intChan本身的地址=%p\n", intChan, &intChan)  // intChan=0xc00008c080, intChan本身的地址=0xc00009a018
// 	// 向管道写入值
// 	intChan<- 10
// 	num := 211
// 	intChan<- num
// 	// intChan<- 50
// 	intChan<- 60
// 	// 查看管道的长度和Cap(容量),cap不能增长
// 	fmt.Printf("Channel len=%v  cap = %v\n", len(intChan), cap(intChan))  // Channel len=3  cap = 3
// 	// 向管道添加数据的时候,最多只能添加声明管道时的数量,添加多了就会报错,也就是说,不成超过Cap的数量
// 	// 报的错误是 all goroutines are asleep - deadlock!
// 	// 从管道中读取数据
// 	// var num2 int
// 	num2 := <- intChan  // 显然是10,因为先进先出
// 	fmt.Println("num2=",num2)
// 	fmt.Printf("Channel len=%v  cap = %v\n", len(intChan), cap(intChan))  // Channel len=2  cap = 3

// 	// 在没有使用协程的情况下,如果管道数据全部取出之后再取,就会报错
// 	num3 := <- intChan
// 	num4 := <- intChan
// 	// num5 := <- intChan  // all goroutines are asleep - deadlock!
// 	// fmt.Println(num3, num4, num5)
// 	fmt.Println(num3, num4)
// 	// 其实管道中是流动的数据,一边放,一边取
// 	<- intChan  // 只是取值,不需要变量接收
// }

// func main() {
// 	// 创建一个mapChan,最多可以存放10个map[string]string的key-val
// 	var mapChan chan map[string]string
// 	mapChan = make(chan map[string]string, 10)  // 管道可以放10个map
// 	m1 := make(map[string]string, 20)  // 每个字典可以放20个键值对
// 	m1["city"] = "深圳"
// 	m1["city1"] = "天津"
// 	m2 := make(map[string]string, 20)
// 	m2["hero1"] = "哇哈哈"
// 	m2["hero2"] = "乳娃娃"
// 	mapChan <- m1
// 	mapChan <- m2
// 	fmt.Println(mapChan)
// }

// // 容易出错的题目

// type Cat struct {
// 	Name string
// 	Age int
// }

// func main() {
// 	// 定义一个可以存放任意数据类型的管道, 3个数据
// 	allChan := make(chan interface{}, 3)
// 	allChan<- 10
// 	allChan<- "jack"
// 	allChan<- Cat{"小花猫",5}
// 	// 如果想取第三个,那就只能把前两个推出去
// 	<- allChan
// 	<- allChan

// 	newCat := <- allChan // 从管道中取出来的是一个接口类型,虽然存入的是结构体类型
// 	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)  // newCat=main.Cat, newCat={小花猫 5}
// 	// fmt.Printf("newCat.Name=%v", newCat.Name)  // 这里会直接报错,因为接口没有Name这个字段
// 	// 所以如果想要取出Name字段,需要类型断言
// 	newCatTrue := newCat.(Cat)  // newCat.Name=小花猫
// 	fmt.Printf("newCat.Name=%v\n", newCatTrue.Name)
// }

// channel的遍历和关闭
// 使用内置函数close可以关闭Channel,当Channel关闭后,就不能再向Channel写数据了
// 但是仍然可以从该Channel读取数据
// Channel的遍历
// Channel支持for-range的方式进行遍历,请注意两个细节 (因为管道的性质是取一个,长度减一个,所以不能for循环直接取)
// 1. 在遍历时,如果Channel没有关闭,则会出现deadlock的错误
// 2. 在遍历时,如果Channel没有关闭,则会正常遍历数据,遍历完后,就会退出遍历, 可以确定主程序的时间

// func main() {

// 	intChan := make(chan int, 3)
// 	intChan<- 10
// 	intChan<- 20
// 	close(intChan)  // 内置函数close
// 	// 这时不能再写入内容到管道
// 	// intChan<- 100  // send on closed channel
// 	fmt.Println("ok~~~~~~~")
// 	// 但是取,能取到相应的值
// 	newInt := <- intChan
// 	fmt.Println("newInt=", newInt)  // newInt= 10
// }

// func main() {

// 	intChan2 := make(chan int, 100)
// 	for i := 0; i < 100; i++ {
// 		intChan2<- i  // 将100个数据放入管道
// 	}

// 	// 在遍历的时候,如果Channel没有关闭,则会出现deadlock的错误,死锁
// 	// 如果关闭管道,就会正常遍历
// 	close(intChan2)

// 	// 用for range遍历管道,不能使用普通的for循环,也不能用cap
// 	for v := range intChan2 {
// 		fmt.Println("v=",v)  // 管道没有下标的
// 	}
// }

// 开启一个writeData协程,从管道intChan中写入50个整数
// 开启一个readData协程,从管道中intChan读取writeData写入的数据
// 注意: writeData和readData操作的是同一个管道,因为是指针类型
// 主线程需要等待writeData和readData协程都完成工作才能退出

// func writeData(intChan chan int) {
// 	for i := 1; i < 51; i++ {
// 		// 放入数据
// 		intChan <- i
// 	}
// 	// 写完之后关闭,因为读的时候没有影响
// 	close(intChan)
// }

// func readData(intChan chan int, exitChan chan bool) {

// 	for {
// 		v, ok := <-intChan
// 		if !ok {
// 			break
// 		}
// 		fmt.Println("读到的数据:", v)
// 	}
// 	exitChan <- true // 任务完成,写入数据
// 	close(exitChan)  // 马上关闭
// }

// func main() {

// 	// 创建两个管道
// 	// 一个用于读写,一个用于退出
// 	intChan := make(chan int, 50)
// 	exitChan := make(chan bool, 1)

// 	go writeData(intChan)
// 	go readData(intChan, exitChan)
// 	// 主进程直接这样写会直接一闪而过,终端也没有打印
// 	// 所以要根据退出管道中有没有值来确定,协程的工作有没有做完数据
// 	for {
// 		_, ok := <-exitChan
// 		if !ok {
// 			break
// 		}
// 	}

// }

// 作业1
	// 创建一个Person结构体{Name, Age, Address}
	// 使用rand方法配合随机创建10个Person实例,并放入Channel中
	// 遍历Channel,将各个实例的内容展示出来



// type Person struct {
// 	Name string
// 	Age int
// 	Address string
// }


// func main() {
	
// 	// 定义一个管道
// 	personChan := make(chan Person, 10)

// 	// 加入时间粒子
// 	rand.Seed(time.Now().UnixNano())  // 微秒级别, 纳秒会更好
// 	for i := 1; i < 11; i++ {
// 		// 创建10个实例
// 		name1 := strings.Join([]string{"name", strconv.Itoa(1)}, "")
// 		address1 := strings.Join([]string{"address", strconv.Itoa(i)}, "")  // int转换成string类型需要strconv
// 		personChan<- Person{name1, i, address1,}  
// 	}
// 	close(personChan)

// 	for k := range personChan {
// 		fmt.Println("Person实例为 : ",k)
// 	}
// }


// 作业2
	// 启动一个协程, 将1-20000的数放入一个Channel中 numChan
	// 启动8个协程,从numChan中取出数字,然后计算,放入另一个Channel resChan
	// 最后8个协程同时完成工作后,再遍历resChan,显示的结果 res[1] = 1


func inData(numChan chan int) {
	for i := 0; i < 20000000; i++ {
		numChan<- i
	}
	close(numChan)
}

func outData(numChan chan int, resChan chan int, exitChan chan bool) {

	for k := range numChan {
		newNum := k + 9
		resChan<- newNum
	}
	exitChan<- true  // 一个协程做完,就写一个true
	close(exitChan)

}

func resdData(resChan chan int) {
	for {
		_, ok := <- resChan
		if !ok {
			break
		}
	}
	
}

// 使用协程的时间为3秒
// func main() {

// 	numChan := make(chan int, 20000000) 
// 	resChan := make(chan int, 20000000) 
// 	exitChan := make(chan bool, 8)
// 	start := time.Now().Unix()
// 	go inData(numChan)  // 写入数据
// 	for i := 0; i < 8; i++ {
// 		go outData(numChan, resChan, exitChan)  // 取出数据并且往一个管道写结果,每个goroutine执行完之后会到exitChan中放入一个True
// 	}
// 	go func() {
// 		for n := 0; n < 8; n++ {
// 			<- exitChan
// 		}

// 		end := time.Now().Unix()
// 		fmt.Println("使用协程耗费的时间:", end - start)
// 		close(resChan)
// 	}()  // 匿名函数要在后面跟()调用
// 	resdData(resChan)
	
	


// }


// 不使用协程的时间
func main() {
	numChan := make(chan int, 20000000) 
	resChan := make(chan int, 20000000) 
	exitChan := make(chan bool, 1)
	start := time.Now().Unix()
	inData(numChan)  // 写入数据
	outData(numChan, resChan, exitChan)  // 取出数据并且往一个管道写结果,每个goroutine执行完之后会到exitChan中放入一个True
	_, ok := <- exitChan
	if !ok {
		fmt.Println("结束")
	}
	close(resChan)
	end := time.Now().Unix()
	fmt.Println("使用协程耗费的时间:", end - start)
	resdData(resChan)
}
