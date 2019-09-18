package main

import(
	"fmt"
	"go_code/CustomerAll/service"
	"go_code/CustomerAll/model"
)


type customerView struct {
	// 此结构体定义必要字段
	key string  // 接收用户输入...
	loop bool  // 表示是否循环的显示主菜单
	// 添加一个字段
	customerService *service.CustomerService
}


// 显示所有的客户信息
func (p *customerView) list() {
	// 获取当前所有用户的信息
	customerList := p.customerService.List()
	fmt.Println("\n---------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for customer := 0; customer < len(customerList); customer++ {
		fmt.Println(customerList[customer].GetInfo())
	}
	fmt.Println("-------------客户信息完成-----------------")
}


// 添加用户方法
func (p *customerView) add() {
	// 得到用户的输入信息
	fmt.Println("\n---------------添加客户-------------------")
	fmt.Println("请输入用户名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入性别:")
	gander := ""
	fmt.Scanln(&gander)
	fmt.Println("请输入年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入手机号码:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("请输入邮箱:")
	emaile := ""
	fmt.Scanln(&emaile)
	customer := model.NewCustomer(name, gander, age, phone,emaile)
	if p.customerService.Add(customer) {
		fmt.Println("-----------------添加成功--------------------")
	} else {
		fmt.Println("-----------------添加失败--------------------")
	}
}


// 删除用户
func (p *customerView) delete() {
	fmt.Println("\n---------------删除客户-------------------")
	id := 0
	fmt.Println("请输入客户的id:")
	fmt.Scanln(&id)
	choice := ""
	fmt.Println("是否确定删除,请输入y或者n:")
	fmt.Scanln(&choice)
	if choice == "y" || choice == "n" {	
		fmt.Println("这里进来了")
		if choice == "n" {
			return
		}
		loopBool := p.customerService.Delete(id)
		if loopBool {
				fmt.Println("--------------删除用户成功---------------")
				return
			
		} else {
			fmt.Println("--------------删除用户失败,该id不存在---------------")
		}
	} else {
		return
	}
}

// 退出系统
func (p *customerView) exit() {
	exit := ""
	for {
		fmt.Println("是否确定退出系统,y/n:")
		fmt.Scanln(&exit)
		if exit == "y" || exit == "n" {
			if exit == "y" {
				p.loop = false
				break
			}else{
				break
			}
		} else {
			fmt.Println("请输入正确选择")
		}
	}
}

// 显示主菜单
func (p *customerView) mainMenu() {

	for {
		fmt.Println("\n-----------------客户管理系统------------------")
		fmt.Println("		1.添加客户")
		fmt.Println("		2.修改客户")
		fmt.Println("		3.删除客户")
		fmt.Println("		4.客户列表")
		fmt.Println("		5.退出")
		fmt.Print("------->请选择(1-5):")
		fmt.Scanln(&p.key)

		switch p.key {
		case "1":
			p.add()
		case "2":
			fmt.Println("修\t改\t客\t户")
		case "3":
			p.delete()
		case "4":
			p.list()
		case "5":
			p.exit()
		default :
			fmt.Println("你的输入有误,请重新输入")
		}

		if !p.loop {
			// 如果loop为假,那就跳出循环
			break
		}
	}
	fmt.Println("你已退出客户管理系统~~~")
}

func main() {
	// 实例化一个结构对象
	customerView := customerView{
		key:"",
		loop:true,
		
	}
	customerView.customerService = service.NewCustomerService()  // 此函数什么都不传参,有默认值
	customerView.mainMenu()
}