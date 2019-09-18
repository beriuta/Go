package service
import (
	"go_code/CustomerAll/model"
)


// 该CustomerService完成对Customer的操作(增删改查)
type CustomerService struct {
	// 定义一个切片
	customers []model.Customer  // 这里需要引入model
	// 声明一个字段,表示当前切片含有多少个客户,该客户还可以作为新客户的id号
	customerNum int

}

// 写一个返回客户切片的方法
func (p *CustomerService) List() []model.Customer {
	return p.customers
}

// 写一个增加客户的方法*CustomerService一定要加*,不然就是值拷贝,只在一个数据中添加
func (p *CustomerService) Add(customer model.Customer) bool {

	p.customerNum ++
	customer.Id = p.customerNum
	p.customers = append(p.customers, customer)
	return true

}

// 根据id下标,删除对应的人员,如果是-1,就返回False
func (p *CustomerService) Delete(id int) bool {
	loop := true
	index := p.FindById(id)
	if index != -1 {
		// 删除该条数据
		p.customers = append(p.customers[:index], p.customers[index+1:]...)
	}else {
		loop = false
	}
	return loop
}




// 写一个查找此id的一个方法,如果找不到就返回-1
func (p *CustomerService) FindById(id int) int {
	// 默认返回的是-1
	index := -1
	for i := 0; i < len(p.customers); i++ {
		if p.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// 写一个工厂方法
func NewCustomerService() *CustomerService {
	// 返回的是客户的切片
	customerService := &CustomerService{}  // 实例化一个对象
	customerService.customerNum = 0    // 初始化一个客户id
	customer := model.NewCustomer("张三", "男", 20, "12934567654", "23453455@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

