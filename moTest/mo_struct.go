package moTest

import (
	"fmt"
)

type student struct {
	Id      int
	name    string
	address string
	cls     string
	age     int
}

type addr struct {
	City     string
	Province string
}
type user struct {
	Name    string
	age     int
	address addr
}
type Class struct {
}

func StudentInfo() {
	//实例化-- var 结构体实例名 结构体名称
	var st student
	fmt.Println("结构体实例化默认值 student:", st)
	st.Id = 1000
	st.name = "沟浪"
	st.address = "沟浪_address"
	st.cls = "沟浪_cls"
	st.age = 10
	fmt.Println("结构体实例化赋值 student:", st)

	//匿名结构体-- var 结构体实例名 struct{}
	var person struct {
		Name string
		Age  int
	}
	fmt.Println("匿名结构体声名:", person)
	person.Name = "匿名结构体名称"
	person.Age = 20
	fmt.Println("匿名结构体赋值:", person)

	//指针类型结构体
	var pStu = new(student)
	fmt.Println("指针类型结构体实例化默认值:", pStu)
	pStu.Id = 1001
	pStu.name = "指针类型结构名称"
	pStu.address = "指针类型结构地址"
	pStu.cls = "指针类型结构_cls"
	pStu.age = 12
	fmt.Println("指针类型结构体赋值：", pStu)

	//取结构体地址实例化
	var pdStu = &student{}
	fmt.Println("取结构体地址定义默认值：", pdStu)
	pdStu.Id = 1001
	pdStu.name = "取结构体地址_name"
	pdStu.address = "取结构体地址_address"
	pdStu.cls = "取结构体地址_cls"
	pdStu.age = 12
	fmt.Println("取结构体地址定义赋值：", pdStu)

	//键值对赋值结构体
	var kvStu = student{Id: 10003, name: "键值对赋值_name", address: "键值对赋值_address", cls: "键值对赋值_cls", age: 12}
	fmt.Printf("键值对赋值=%#v\n", kvStu)

	//取结构体地址键值对实例化
	var pdKvStu = &student{Id: 1004, name: "取结构体地址键值对实例化_name"}
	fmt.Printf("取结构体地址键值对实例化:%v \n", pdKvStu)

	//构造函数
	newStu := NewStudent(1005, "构造函数_name", "构造函数_address", "构造函数_cls", 15)
	fmt.Printf("构造函数%v \n", newStu)

	//方法和接受者
	methodStu := &student{Id: 10004, name: "method_name", address: "method_address", cls: "method_cls", age: 16}
	methodStu.Dream()

	//嵌套结构体
	pUser := user{Name: "嵌套结构体_name", age: 11, address: addr{City: "嵌套结构体_city", Province: "嵌套结构体_Province"}}
	fmt.Printf("嵌套结构体 %#v \n", pUser)

	m := make(map[string]student)
	stus := []student{
		{Id: 1, name: "张三", address: "shanghai", cls: "三年级", age: 10},
		{Id: 2, name: "张五", address: "shanghai Pudong", cls: "三年级", age: 9},
		{Id: 3, name: "张六", address: "shanghai qingpu", cls: "三年级", age: 11},
	}
	for _, stu := range stus {
		m[stu.name] = stu
	}
	fmt.Println("m:", m)
	for k, v := range m {
		fmt.Println(k, "-->", v)
	}
}

//构造函数
func NewStudent(Id int, Name string, Address string, Cls string, Age int) *student {
	return &student{Id: Id, name: Name, address: Address, cls: Cls, age: Age}
}

//方法和接受者
func (s student) Dream() {
	fmt.Printf("沟浪接受者：%v \n", s.name)
}
