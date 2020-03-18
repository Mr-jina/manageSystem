// 数据层

package models

import(
	"fmt"
)

// 创建Customer结构体,存储数据

type Customer struct {

	Id int
	Name string
	Gender string
	Age string
	Phone string
	Email string

}

// 使用工厂模式创建customer结构体实例
func NewCustomer(id int, name string, gender string, age string, phone string, email string) Customer {
	return Customer{
		Id: id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}

}

// 创建customerTwo对象，不带“Id”
func NewCustomerTwo(name string, gender string, age string, phone string, email string) Customer {
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}

}


// 返回用户的信息,格式化字符串
func (this *Customer) GetInfo() string {

	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
			this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
	
	return info
}
