/*
	先安装第三方开源的库,操作redis_API
	之前的操作是手动连接,但是频繁的连接关闭消耗资源,所以需要连接池
	通过Golang对redis操作
		初始化一定数量的连接,方法哦连接池中
		当Go需要操作redis时,直接从redis连接池取出连接
		可以节省临时获取redis连接的时间,从而提高效率
	创建连接池的步骤
		1. 初始化连接池 init
		2. 拿到连接, pool.Get
		var pool *redis.Pool{
			MaxIdle:8,  // 最大空闲连接数
			MaxActive:0,  // 表示和数据库最大的连接数,0表示没有限制
			IdleTimeout:100,  // 最大空闲时间,在100秒没有用了,就会自动放到连接池中
			Dial:func()(redis.Conn,error) {  // 初始化连接代码,连接哪个id的redis
				return redis.Dial("tcp", "localhost:6379")
			},
		}
		conn := pool.Get()  // 从连接池中获取一个连接
		pool.Close()  // 关闭连接池,一旦关闭,就不能再从连接池中连接了
*/
package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"  // 引入redis包
)

// 定义一个redis连接池
var pool *redis.Pool


// 初始化的时候创建连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:8,  // 最大空闲连接数
		MaxActive:0,  // 表示和数据库最大的连接数,0表示没有限制
		IdleTimeout:100,  // 最大空闲时间,在100秒没有用了,就会自动放到连接池中
		Dial:func()(redis.Conn,error) {  // 初始化连接代码,连接哪个id的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}


// func main() {

// 	// 通过go向redis写入数据
// 	// 1. 连接redis
// 	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
// 	if err != nil {
// 		fmt.Println("连接redis出错:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	// 通过go向redis写数据,插入string类型
// 	_, err = conn.Do("Set", "family", "love")  // (操作类型,key,value)
// 	if err != nil {
// 		fmt.Println("插入数据失败:", err)
// 		return
// 	}
// 	// 通过go取出redis存入的值
// 	// 因为Get返回的是接口类型的,所以需要转换类型
// 	ret, err := redis.String(conn.Do("Get", "family"))  // (操作类型,key)
// 	if err != nil {
// 		fmt.Println("获取数据失败:", err)
// 		return
// 	}
// 	fmt.Println("获取的数据是:", ret)
// 	// 通过go对redis操作hash数据类型
// 	_, err = conn.Do("HMSet", "字典2", "name","make", "age", 18, "adress", "深圳世界之窗")  // 单个字段赋值,HMSet,是多个字段赋值
// 	if err != nil {
// 		fmt.Println("存入hash数据失败:", err)
// 		return
// 	}
// 	ret2, err := redis.Strings(conn.Do("HMGet", "字典2", "age", "name", "adress"))  // 这里的是Strings,需要添加一个s,返回的是一个切片
// 	fmt.Println("获取的字典数据是:", ret2)  // 这是一个切片
// 	for i, v := range ret2 {
// 		fmt.Printf("r[%d] = %s\n", i, v)
// 	}

// 	fmt.Println("连接成功:", conn)  // 这是conn的套接字,因为本身是结构体 连接成功: &{{0 0} 0 <nil> 0xc00009e020 0 0xc0000a01e0 0 0xc0000a20c0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}

// }

func main() {
	conn := pool.Get()  // 从连接池中获取一个连接
	defer conn.Close()  // 关闭连接池,一旦关闭,就不能再从连接池中连接了

	_, err := conn.Do("Set", "tommao", "汤姆猫~~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	ret, err := redis.String(conn.Do("Get", "tommao"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("打印ret===>", ret)
}