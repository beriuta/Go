// 排序和查找,二分,冒泡
// 排序
	// 将一组数据,依指定的顺序进行排列的过程
	// 1,内部排序
		// 指将需要处理的所有数据都加载到内部储存器中进行排序
		// 包括(交换式排序法,选择式排序法和插入式排序法)
	// 2,外部排序法
		// 数据量过大,无法全部加载到内存,需要借助外部存储进行排序,包括(合并排序法和直接合并排序法)
	// 交换式排序法
		// 交换式排序属于内部排序法,是运用数据值比较后,依判断规则对数据位置进行交换,以达到排序的目的
		// 分为两种
			// 1. 冒泡排序法(Bubble sort)	
			// 2. 快速排序法(Quick sort)
		// 冒泡排序
			// 通过对待排序序列从后向前(从下标较大的元素开始),依次比较相邻元素的排序码,若发现逆序则交换
			// 使排序码较小的元素主键从后部向前移(从下标较大的单元一想下标较小的单元)
			// 注意:
				// 因为排序的过程中,个元素不断接近自己的位置,如果一趟比较下来没有进行过交换,就说明序列有序,因此要在排序过程中设置
				// 一个标志flag判断元素是否进行过交换,从而减少不必要的比较
// 案例
// 有五个无序的数字 24, 69, 80, 57, 13,使用冒泡排序法将其排成一个从小到大的有序数列
// 冒泡的规则,思想最重要,把思想转换成代码
package main
// import "fmt"

	// func test(list *[5]int) {
	// 	fmt.Println("原始数据", (*list))

	// 	for j := 0; j < len(*list) - 1; j++ {

	// 		for i := 0; i < len(*list) - j - 1; i++ {

	// 			if (*list)[i] > (*list)[i + 1] {
	// 				(*list)[i], (*list)[i + 1] = (*list)[i + 1], (*list)[i]
	// 			}
	// 		}
	// 	}
	// 	fmt.Println("第一次排序:", (*list))
	// }


	// func main() {
	// 	// 先将最大的数排到最后面
	// 	// 把第二大的数放到倒数第二位
	// 	// 多重循环
	// 	list := [5]int{24, 69, 80, 57, 13}
	// 	// 将数组传到一个函数,完成排序
	// 	test(&list)  // 将地址传进去
	// }

// 查找
	// 顺序查找
	// 二分查找(必须是有序的)
// 练习
	// 数列,白眉鹰王,金毛狮王,紫衫龙王,青翼蝠王
		// func main() {

		// 	listStr := [4]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼蝠王"}
		// 	var heroName = ""
		// 	fmt.Println("请输入要查找的人名:")
		// 	fmt.Scanln(&heroName)

		// 	// 第一种方法(顺序查找)
		// 	// for i := 0; i < len(listStr); i++ {
		// 	// 	if heroName == listStr[i] {
		// 	// 		fmt.Printf("找到了%v,在第%v索引\n", listStr[i], i)
		// 	// 		break
		// 	// 	} else if i == len(listStr) - 1 {
		// 	// 		fmt.Printf("输入的是无效值\n")
		// 	// 	}
		// 	// }

		// 	// 第二种方法(顺序查找)
		// 	index := -1
		// 	for i := 0; i < len(listStr); i++ {
		// 		if heroName == listStr[i] {
		// 			index = i
		// 			break
		// 		}
		// 	}
		// 	if index != -1 {
		// 		fmt.Printf("找到了%v,在第%v索引\n", heroName, index)
		// 	} else {
		// 		fmt.Printf("输入的是无效值\n")
		// 	}

		// }

// 二分查找的思路分析
// arr是一个数组,是从小到大的有序数组
// 如果查找的值为findValue
// 先找中间的下标,middle = (leftindex + rightindex) / 2
// 如果middle的值 > findValue 就将 [middle:rightindex]这一段值传进去查找
// 如果middle的值 < findValue 就将 [middle:leftindex]这一段值传进去查找
// 如果middle的值==findValue
// 退出的条件为 leftindex > rightindex 就为退出值,找不到

	// func test(arr *[30]int, leftindex int, rightindex int, findValue int) {

	// 	// 将退出条件写上
	// 	// 当leftindex > rightindex时,就会退出
	// 	if leftindex > rightindex {
	// 		fmt.Println("没有找到")
	// 		return
	// 	}

	// 	// 将传入的列表取出中间值
	// 	middle := (leftindex + rightindex) / 2
	// 	// fmt.Println(middle)
	// 	// fmt.Println((*arr)[middle])

	// 	if (*arr)[middle] == findValue {

	// 		fmt.Printf("此数字在下标为%v的位置\n", middle)

	// 	} else if (*arr)[middle] > findValue {
	// 			// arr因为是在自己的函数中,不需要传递地址
	// 			test(arr, leftindex, middle - 1, findValue)

	// 	} else if (*arr)[middle] < findValue {

	// 			test(arr, middle + 1, rightindex, findValue)
	// 	}
	// }

	// func main() {
	// 	// 定义一个需要查找的数
	// 	findValue := 8
	// 	// 定义一个数组
	// 	arr := [30]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30}
	// 	test(&arr, 0, len(arr) - 1, findValue)
	// }

// 二维数组 
	// func main() {
	// 	// 定义一个二维数组
	// 	var arr [4][6]int
	// 	fmt.Println(arr)  // [[0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0]]
	// 	// 把二行3列的数变为8, 三行二列的数变为3,三行四列的数变为6,四行三列的数变为9
	// 	arr[1][2] = 8
	// 	arr[2][1] = 3
	// 	arr[2][3] = 6
	// 	arr[3][2] = 9
	// 	fmt.Println(arr)  // [[0 0 0 0 0 0] [0 0 8 0 0 0] [0 3 0 6 0 0] [0 0 9 0 0 0]]
	// 	// 竖向打印
	// 	for i := 0; i < 4; i++ {
	// 		// 此时打印的时候只是把每个数组打印出来了
	// 		// fmt.Println(arr[i])
	// 		// [0 0 0 0 0 0]
	// 		// [0 0 8 0 0 0]
	// 		// [0 3 0 6 0 0]
	// 		// [0 0 9 0 0 0]
	// 		// 所以需要再遍历一遍,把每个元素打印出来
	// 		for j := 0; j < 6; j++ {
	// 			// fmt.Printf("%v ",arr[i][j])
	// 			fmt.Print(arr[i][j], " ")
	// 		}
	// 		// 每打印完一次就换行
	// 		fmt.Println()
	// 	}
	// }

// 二维数组的内存
// 1. 先声明再赋值
	// 语法:  var 数组名 [行大小][列大小]数据类型
	// 比如: var arr [2][3]int 再赋值
	// 在内存中,是存储了两个指针,打印出来是两个地址相差24字节,一个int占8个字节
	// 分别有两个三位数的数组
	// 如果改成 var arr [2][2]int, 就分别有两个二位数的数组, 两个地址就相差16
		// fmt.Printf("行指针%p", &arr[0])
		// fmt.Printf("列指针%p", &arr[1])
// 2. 直接初始化(各种写法)
	// var arr [2][3]int = [2][3]int{{1,2,3},{5,2,3}}
	// var arr = [2][3]int{{1,2,3},{5,2,3}}
	// arr := [2][3]int{{1,2,3},{5,2,3}}
// 3. 二维数组的遍历
	// func main() {
	// 	// // 二维数组的遍历
	// 	// arr := [2][3]int{{1,2,3},{3,4,67}}

	// 	// // for 循环遍历
	// 	// for i := 0; i < len(arr); i++ {
	// 	// 	// for循环内层数字
	// 	// 	for j := 0; j < len(arr[i]); j++ {
	// 	// 		fmt.Printf("%v\t\n", arr[i][j])
	// 	// 	}
	// 	// }
	// 	// // for_range遍历数组
	// 	// for _, v := range arr {
	// 	// 	for _, vj := range v {
	// 	// 		fmt.Println("每个元素为:", vj)
	// 	// 	}
	// 	// 	fmt.Println("v=", v)
	// 	// }
	// 	// 定义一个二维数组,用于保存三个班,每个班五名同学成绩
	// 	// 并求出每个班级平均分,以及所有班级的平均分
	// 	var scores [3][5]float64
	// 	var num float64
	// 	for i:= 0; i < len(scores); i++ {
	// 		// 循环输入成绩
	// 		var all float64
	// 		for j := 0; j < len(scores[i]); j++ {
	// 			fmt.Printf("请输入第%d班第%d人的成绩\n", i + 1, j + 1)
	// 			fmt.Scanln(&scores[i][j])
	// 			all += scores[i][j] 
	// 		}
	// 		num += all
	// 		fmt.Printf("%d班的总成绩:%v\n", i + 1, all)
	// 		fmt.Printf("%d班的平均成绩:%v\n", i + 1, all / float64(len(scores[i])))
	// 	} 
	// 	fmt.Printf("所有班的总成绩:%v\n", num)
	// 	fmt.Printf("所有班的平均成绩:%v\n", num / float64(len(scores)))
	// }

// map是key:value数据结构,又称之为字典或者关联数组,类似其他编程语言的集合
	// map基本语法
		// var data(变量名) map[keytype(key的数据类型)]valuetype(value的数据类型)
		// key可以放什么数据类型
			// golang中map的key可以是很多种类型,比如bool,数字,string,指针,Channel,还可以是只包含前面几个类型的,接口,结构体,数组
			// 通常为int, string 99%
			// 注意: slice,map还有function不可以,因为这几个没法用来判断
		// value可以是什么类型
			//跟key的数据类型基本一样
			// 通常为: 数字(整数,浮点数), string, map, struct 
		// map声明举例
			// var a map[string]string
			// var a map[string]int
			// var a map[int]string
			// var a map[string]map[string]string
			// 注意: 声明是不会分配内存的,初始化需要make,分配内存后才能赋值和使用
				// func main() {
				// 	// 声明一个map
				// 	var a map[string]string
				// 	fmt.Println(a)  // map[]
				// 	// a["n01"] = "b"  // 此时没有声明make前就赋值,会产生报错
				// 	// fmt.Println(a)  // assignment to entry in nil map
				// 	// make的作用就是给map分配数据空间
				// 	a = make(map[string]string, 10)  // 最大可以放10对key:value
				// 	a["no1"] = "深圳"
				// 	a["no2"] = "上海"
				// 	a["no3"] = "先"  // key不能重复,如果是相同的值会覆盖值
				// 	fmt.Println(a)  // map是没有顺序的,无序的
				// 	// make内置函数的说明
				// }
// map的使用方式
// 1. var cities map[string]string   cities = make(map[string]string, 10)
// 此时的map = nil
// 2. var cities = make(map[string]string)  直接make
// 3. var cities = map[string]string = map[string]string{"no2":"哇哈哈", "n03": "笑嘻嘻"}
// cities["no5"] = "深圳"
	// func main() {
	// 	// // 第二种方式
	// 	// cities := make(map[string]string)
	// 	// cities["no1"] ="深圳"
	// 	// cities["no2"] ="上海"
	// 	// fmt.Println("此map打印出来的格式", cities)
	// 	// 第三种方式
	// 	cities := map[string]string{"haha":"222", "waw":"43333",}  // 最后不能省略逗号
	// 	fmt.Println("第三种方式", cities)
	// }

	// func main() {
	// 	// 存放3个学生的信息,姓名:小花,性别:女
	// 	studentData := make(map[string]map[string]string)
		
	// 	for i := 0; i < 3; i++ {
	// 		studentData[string(i)] = make(map[string]string, 3)
	// 		var name string
	// 		var sex string
	// 		fmt.Println("请输入学生姓名:")
	// 		fmt.Scanln(&name)
	// 		fmt.Println("请输入学生性别:")
	// 		fmt.Scanln(&sex)
	// 		studentData[string(i)]["name"] = name
	// 		studentData[string(i)]["sex"] = sex
	// 	}
	// 	fmt.Println(studentData)
	// }
// map增删改查(crud)
	// 增加和更新,有key就更新,没有key就添加
		// 变量["key"] = "日日日"
	// 删除
		// delete(变量名, "key"),delete是一个内置函数如果key存在就删除,不存在也不报错
	// 删除细节说明
		// 如果要删除整个map所有的值,go中没有方法一次性删除,只能遍历map逐个删除
		// 或者 变量名 = make(...) make一个新的,让原来的成为垃圾,被gc回收  (推荐用这种)
// 演示map的查找
	// func main() {
	// 	cities := map[string]string{"jsjs":"ddd", "sna":"kdjds"}
	// 	val, ok := cities["se"]
	// 	if ok {
	// 		fmt.Println("寻找的键有值,值是:", val)
	// 	} else {
	// 		fmt.Println("寻找的键没有值")
	// 	}
	// }
// map遍历,只能用for-range遍历,不能用for循环
	// func main() {
	// 		// 存放3个学生的信息,姓名:小花,性别:女
	// 		studentData := make(map[string]map[string]string)
			
	// 		for i := 0; i < 3; i++ {
	// 			studentData[string(i)] = make(map[string]string, 3)
	// 			var name string
	// 			var sex string
	// 			fmt.Println("请输入学生姓名:")
	// 			fmt.Scanln(&name)
	// 			fmt.Println("请输入学生性别:")
	// 			fmt.Scanln(&sex)
	// 			studentData[string(i)]["name"] = name
	// 			studentData[string(i)]["sex"] = sex
	// 		}
	// 		fmt.Println(studentData)

	// 		for k, v := range studentData {
	// 			fmt.Println(k, v)
	// 			for k2, v2 := range v {
	// 				fmt.Printf("\t k2%v, v2%v", k2, v2)
	// 			}
	// 		}
	// }
// map 拿长度len()
// map切片
	// 切片的数据类型如果是map,那么map的个数是动态变化的
		// func main() {

		// 	// 声明一个map切片
		// 	var monsters []map[string]string
		// 	monsters = make([]map[string]string, 2) // 切片本身先要make一下,此切片目前能放入两个map
		// 	// 增加一个妖怪信息

		// 	if monsters[0] == nil {
		// 		monsters[0] = map[string]string{"name":"牛魔王", "age":"43333",}
		// 	}

		// 	if monsters[1] == nil {
		// 		monsters[1] = map[string]string{"name":"玉兔精", "age":"933",}
		// 	}
		// 	// 此时如果想要再添加,需要用到切片的append
		// 	// 定义一个新的map类型
		// 	newMonsters := map[string]string{"name":"铁扇公主", "age":"888",}
		// 	// 将新的添加到旧的
		// 	monsters = append(monsters, newMonsters)
		// 	fmt.Println(monsters)  // [map[age:43333 name:牛魔王] map[age:933 name:玉兔精] map[age:888 name:铁扇公主]]
		// }
// map排序
	// func main() {
	// 	// 定义一个map
	// 	map1 := make(map[int]int, 10)
		
	// 	map1[8] = 23	
	// 	map1[1] = 999
	// 	map1[10] = 122
	// 	map1[4] = 432

	// 	fmt.Println(map1)  // map[1:999 4:432 8:23 10:122]
	// }
// map使用细节
// map是引用类型,遵守引用类型的传递机制, 一个函数接收map之后如果修改,原来的map也会一起修改
// map的容量达到后,再加入新的元素也不会发生panic,也就是说map能动态增长键值对
	// func modify(map1 map[int]int) {
	// 	map1[10] = 999
	// }

	// func main() {
	// 	// 展示修改了map之后会不会原来的map也会发生变化
	// 	map1 := map[int]int{1:23,4:56,78:43}
	// 	modify(map1)
	// 	fmt.Println(map1)  // map[1:23 4:56 10:999 78:43]
	// }
// 在实际开发中,map的value其实更多的是用struct结构体,可以自定义不同的数据类型,使存入的数据类型更丰富

	// type Stu struct {
	// 	Name string
	// 	Age int
	// 	Address string
	// }

	// func main() {
	// 	// 定义一个学生map,每个学生key为学号,value为结构体,里面有姓名,年龄,地址
	// 	students := make(map[string]Stu, 10)
	// 	// 创建两个学生
	// 	stu1 := Stu{"tom", 29, "上海"}
	// 	stu2 := Stu{"marry", 19, "深圳"}
	// 	students["032334"] = stu1
	// 	students["032335"] = stu2
	// 	fmt.Println("最后的map内容:", students)  //最后的map内容: map[032334:{tom 29 上海} 032335:{marry 19 深圳}]
	// }

// map练习


	// func modifyUser(user map[string]map[string]string, name string) {
	// 	value, ok := user[name]
	// 	if ok {
	// 		// 说明用户名重复,修改密码
	// 		value["password"] = "88888"
	// 	} else {
	// 		// 说明不重复,新建用户
	// 		user[name] = make(map[string]string)
	// 		user[name]["nikeName"] = name
	// 		user[name]["password"] = "222222222"
	// 	}
	// 	fmt.Println("程序执行完毕")
	// }



	// func main() {
	// 	// 使用map[string]map[string]string类型
	// 	// key表示用户名,如果某个用户名存在,就把密码修改为88888
	// 	// 如果不存在,就增加这个用户的名称跟密码
	// 	// 编写一个函数 modifyUser(user map[string]map[string]string, name string)
	// 	user := make(map[string]map[string]string, 10)
	// 	user["韩蕾"] = make(map[string]string)
	// 	user["韩蕾"]["nikeName"] = "韩蕾"
	// 	user["韩蕾"]["password"] = "029183848"
	// 	fmt.Println(user)
	// 	name := "呵呵"
	// 	modifyUser(user, name)
	// 	fmt.Println(user)
	// 	// map[韩蕾:map[nikeName:韩蕾 password:029183848]]
	// 	// 程序执行完毕
	// 	// map[韩蕾:map[nikeName:韩蕾 password:88888]]
	// }

// map的遍历
	// type Stu struct {
	// 		Name string
	// 		Age int
	// 		Address string
	// 	}

	// func main() {
	// 	// 定义一个map
	// 	student := make(map[string]Stu, 10)
	// 	// 插入信息
	// 	stu1 := Stu{"tome", 18, "深圳"}
	// 	stu2 := Stu{"mary", 16, "上海"}
	// 	student["stu1"] = stu1
	// 	student["stu2"] = stu2

	// 	for k, v := range student {
	// 		fmt.Printf("学生的编号:%v\n", k)
	// 		fmt.Printf("学生的姓名:%v\n", v.Name)  // v.Name是把结构体拿出来
	// 		fmt.Printf("学生的年龄:%v\n", v.Age)
	// 		fmt.Printf("学生的住址:%v\n", v.Address)  
	// 	}

	// }