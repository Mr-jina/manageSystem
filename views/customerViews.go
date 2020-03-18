// 显示层

package main

import(
	"fmt"
	"../services"
	"../models"
)

// 创建customerView结构体对象

type customerViews struct {

	// 定义必要字段
	key string // 接收用户输入...

	exit bool // 退出标记
	 
	customcerServices *services.CustomerServices // customerServices对象

}

// 显示所有的顾客信息
func (this *customerViews) find() {

	// 首先，获取到当前所有的顾客信息(在切片中)
	customers := this.customcerServices.List()

	// 显示
	fmt.Println("--------------------顾客列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	// for-rang 遍历 list切片
	for _, v := range customers {
			
		fmt.Println(v.GetInfo())
		
	}	

	fmt.Println("------------------顾客列表完成----------------")
}

// 添加顾客信息
func (this *customerViews) add() {

	// 添加
	fmt.Println("--------------------添加顾客-----------------")
	
	fmt.Println("姓名: ")
	name := ""   // 1
	fmt.Scanln(&name)

	fmt.Println("性别: ")
	gender := "" // 2
	fmt.Scanln(&gender)
	
	fmt.Println("年龄: ")
	age := ""    // 3
	fmt.Scanln(&age)

	fmt.Println("电话: ")
	phone := ""  // 4
	fmt.Scanln(&phone)
	
	fmt.Println("邮箱: ")
	email := ""  // 5
	fmt.Scanln(&email)

	// 构建一个新的Customer实例
	// 注意: id，没有让用户输入，id是"唯一"的，需要系统分配
	customer := models.NewCustomerTwo(name,gender,age,phone,email)

	// 调用customcerServices里的Add()把 顾客信息添加 到 customers 切片里
	flag := this.customcerServices.Add(customer)

	if flag {
		fmt.Println("------------------添加完成----------------")
	}else {
		fmt.Println("------------------添加失败----------------")
	}

}

// 删除顾客信息
func (this *customerViews) del() {

	id := 0 //初始化为0
	fmt.Println("请输入id号: ")
	fmt.Scanln(&id)
	this.customcerServices.Del(id)

}

// 更新顾客信息
func (this *customerViews) updated() {

	fmt.Println("请输入要更新的顾客id: ")
	id := 0
	fmt.Scanln(&id)

	existed := this.isExistId(id)
	// 判断该 “顾客id” 是否存在
	if existed { // existed 为 true，则证明有该“顾客的id”

		name := ""
		gender := ""
		age := ""
		phone := ""
		email := ""
		
		fmt.Println("请输入姓名: ")
		fmt.Scanln(&name)
		fmt.Println("请输入性别: ")
		fmt.Scanln(&gender)
		fmt.Println("请输入年龄: ")
		fmt.Scanln(&age)
		fmt.Println("请输入手机号码: ")
		fmt.Scanln(&phone)
		fmt.Println("请输入邮箱: ")
		fmt.Scanln(&email)
		
		this.update(id,name,gender,age,phone,email)

	}else {

		return // 没有该顾客id,结束方法

	}

}

// 更新顾客信息的附加方法
func (this *customerViews) update(id int, name, gender, age, phone, email string) {

	// 把顾客“id号”传入方法,和相关信息
	this.customcerServices.Update(id,name,gender,age,phone,email)

}

// 把 "顾客id号” 传入,判断是否存在这个顾客id
func (this *customerViews) isExistId(id int) bool {

	isExist := this.customcerServices.IsExistId(id)

	return isExist

}



// 退出程序
func (this *customerViews) quit() {

	for {
		
		fmt.Println("您是否要退出？请输入y/n: ")
		
		fmt.Scanln(&this.key)

		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
		
			this.exit = false

		}else {

			fmt.Println("您未输入y/n,请重新输入！")

		}

		if this.key == "Y" || this.key == "y" {
			
			break // 确定是y/Y , 退出死循环

		}
		
	}
	

}


//主菜单 -- 调用具体功能方法

func (this *customerViews)mainMenu() {


	for {

		fmt.Println("\n-----------客户信息管理软件----------")
		fmt.Println("------------1 添加客户--------------")
		fmt.Println("------------2 修改客户--------------")
		fmt.Println("------------3 删除客户--------------")
		fmt.Println("------------4 客户列表--------------")
		fmt.Println("------------5 退    出--------------")
		fmt.Println("           请选择(1 - 5):           ") 

		//  1.fmt.Scanln(&变量名)  
		//	2.fmt.Scanf("%s\n",&变量名) 
		//  注意: &变量名: 获取变量的内存地址，可以修改数据(忘了这个&符号！！！)
		fmt.Scanln(&this.key)

		switch this.key {
			case "1":
				fmt.Println("添 加 客 户")
				this.add()
				
			case "2":
				fmt.Println("修 改 客 户")
				this.updated()
			
			case "3":
				fmt.Println("删 除 客 户")
				this.del()
				
			case "4":
				fmt.Println("客 户 列 表")
				this.find()

			case "5":
				this.quit()

			default :
				fmt.Println("您输入的选项不存在,请重新输入!")
		}

		// 判断exit是否为false
		if !this.exit {
			break
		}

	}

}

func main() {
	
	// 创建customerViews结构体对象,并赋值
	customerViews := customerViews{

		key : "",
		exit : true,

	}

	// 这里完成对customerViews结构体的CustomerServices字段进行初始化
	customerViews.customcerServices = services.NewCustomerServices()

	// customerViews结构体对象调用“主菜单方法”
	customerViews.mainMenu()

}