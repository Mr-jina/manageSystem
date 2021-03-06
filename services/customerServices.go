// 控制层(业务层)

package services

import(
	"fmt"
	"../models"
)

type CustomerServices struct {

	// 创建Customer切片

	customers []models.Customer

	// 新顾客的编号

	customeNum int
}

// 工厂模式一
func NewCustomerServices() *CustomerServices {

	// 知识点:  var p *Person = &Person{}  返回结构体指针！调用Person结构体里的属性时,直接使用p.name,go内部已经优化了指针格式
	
	// 为了能够看到有顾客在切片中，我们初始化一个顾客信息
	CustomerServices := &CustomerServices{} 

	// 预加入一个 “顾客对象数据”
	// CustomerServices.customeNum = 1 // 顾客编号为1
	// customer1 := models.NewCustomer(1,"Mr-jin","female","22","10086","1040037245@qq.com") // 初始化一个customer对象
	// 将 customer1 存入 customers切片 里
	// CustomerServices.customers = append(CustomerServices.customers,customer1)

	return CustomerServices
}

// 需求: 增、删、查、改

// (不带id，让id“自增”)
// 添加customer顾客信息 到 customers切片 里
func (this *CustomerServices) Add(customer models.Customer)bool {

	this.customeNum++  // 新用户Id“自增方式”

	customer.Id = this.customeNum  // 定义customer对象“id自增”

	this.customers = append(this.customers,customer)  // 将 customer 存入 customers切片 里

	return true // 返回true，代表添加成功

}

// 删除 customers切片 里的 指定customer顾客信息
func (this *CustomerServices) Del(id int) {

	index := this.FindById(id) // 获取返回的id号

	// 如果 index == -1
	if index == -1 {

		// 则表示该顾客切片里没有该顾客的id，也就是没有该顾客的信息
		fmt.Println("该顾客切片里没有该顾客的id,请重新输入id！")
		return // 结束这个方法的所有操作！
		
		}else {
			
			// 如果有包含这个顾客的id，则删除数据
			// 将删除点前后的元素连接起来
			this.customers = append(this.customers[:index], this.customers[index+1:]...)

			fmt.Println("删除该顾客数据成功！")

	}

}

// 通过顾客id去查找对应的切片里的顾客下标，并返回该下标
func (this *CustomerServices) FindById(id int) int {

	index := -1 // 预定义一个标识符，默认为“-1”

	//遍历顾客切片
	for i := 0 ; i < len(this.customers) ; i++ {

		// 如果 顾客的Id 等于 我们传进来的id
		if this.customers[i].Id == id {
		
			//则返回该顾客的下标
			index = i

		}

	}

	return index
}

// 更新顾客信息
func (this *CustomerServices) Update(id int, name, gender, age, phone, email string) {

	haveId := this.FindById(id)

	if haveId == -1 {  // 如果haveId等于-1,则证明没有该id号的顾客信息

		return
		
	}else {
				
		// 更新该顾客的信息
		this.customers[haveId].Name = name
		this.customers[haveId].Gender = gender
		this.customers[haveId].Age = age
		this.customers[haveId].Phone = phone
		this.customers[haveId].Email = email
		
	}

}

// 判断是否存在该顾客id，返回一个bool值
func (this *CustomerServices) IsExistId(id int) bool {

	// 把 "顾客id号” 传入
	haveId := this.FindById(id)

	if haveId != -1 {  // 如果haveId等于-1,则证明没有该id号的顾客信息

		return true
		
	}
	return false
}


// customer的List方法
// 返回顾客切片
func (this *CustomerServices) List() []models.Customer {

	return this.customers
}

