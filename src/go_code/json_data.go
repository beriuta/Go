package main
import (
	"fmt"
	"encoding/json"
)

// 定义一个结构体
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string
	Sal float64
	Skill string
}

func testStruct() {
	// 演示
	monster := Monster{
		Name:"牛魔王" ,
		Age:100 , 
		Birthday:"1902-10-2",
		Sal:5000.0,
		Skill:"牛魔拳",
	}
	// 将monster序列化, 使用Marshal包序列化,成功返回一个byte类型,失败返回error信息, 接收的是一个接口
	data, err := json.Marshal(&monster)  //  直接传递一个指针类型
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println(string(data))  // {"Name":"牛魔王","Age":100,"Birthday":"1902-10-2","Sal":5000,"Skill":"牛魔拳"}
	// 因为结构体key是大写,所以要把它变成小写
	fmt.Println(string(data))  // {"name":"牛魔王","age":100,"Birthday":"1902-10-2","Sal":5000,"Skill":"牛魔拳"}




}

func testSlice() {
	// 使用切片序列化
	var slice []map[string]interface{}  // 定义这个数组
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "深圳"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "7"
	m2["address"] = "上海"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)  //  直接传递一个指针类型
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println(string(data))
}

func unmarshalStruct() {
	// 反序列化成struct
	str := "{\"Name\":\"牛魔王\",\"Age\":100,\"Birthd/ay\":\"1902-10-2\",\"Sal\":5000,\"Skill\":\"牛魔拳\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(monster)

}

func unmarshalMap() {
	// 反序列化成map
	str := "{\"address\":\"上海\",\"age\":\"7\",\"name\":\"tom\"}"
	// 定义一个map
	var a map[string]interface{}
	// 反序列化的map不需要再make一下,因为在反序列化的时候已经make了
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(a)  // map[address:上海 age:7 name:tom]
}

func unmarshalSlice(){
	str := "[{\"address\":\"深圳\",\"age\":\"7\",\"name\":\"jack\"}," +
	"{\"address\":\"上海\",\"age\":\"7\",\"name\":\"tom\"}]"
	// 定义一个slice
	var slice []map[string]interface{}
	// 反序列化的map不需要再make一下,因为在反序列化的时候已经make了
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(slice)  // [map[address:深圳 age:7 name:jack] map[address:上海 age:7 name:tom]]
}


func main()  {

	// json 一般web编程应用
	// 服务器,比如go数组,json序列化编程json字符串,传到浏览器中
	// {"name":"tome","age":18, "address":["北京","上海"]}
	// 任何数据类型,都能转换成json数据
	// 将结构体,map,切片的序列化
	testStruct()
	// 切片序列化
	testSlice()  // [{"address":"深圳","age":"7","name":"jack"},{"address":"上海","age":"7","name":"tom"}]

	// json的反序列化, 反序列化的类型要跟序列化的类型一样
	// 将json字符串反序列化结构体,map和切片
	unmarshalStruct()
	// 反序列化map
	unmarshalMap()
	// 反序列化切片
	unmarshalSlice()





}

// 一个被测试的函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}