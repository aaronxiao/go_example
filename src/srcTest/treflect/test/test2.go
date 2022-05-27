package main

import (
	"fmt"
	"reflect"
)

//结构体与放射

type User struct {
	Id 		int		`json:"id" db:"Id"`
	Name 	string	`json:"name" db:"NickName"`
	Age     int		`json:"age" db:"Age"`
}

func (u User) Hello(name string) {
	fmt.Println("hello", name)
}

func Poni(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("类型", t)
	fmt.Println("字符串类型", t.Name() )
	
	v := reflect.ValueOf(o)
	fmt.Println(v)

	//t = t.Elem()

	for i:= 0; i < t.NumField(); i++ {
		//获取每个字段
		f := t.Field(i)
		fmt.Printf("%s :%v \n", f.Name , f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		//val := v.Elem().Field(i).Interface()
		fmt.Println("val", val)
	}

	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}


// 匿名字段
type Boy struct {
	User
	Addr string
}


//修改结构体值
func SetValue(o interface{})  {
	v := reflect.ValueOf(o)
	//获得指针指向的元素
	v = v.Elem()
	//取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String{
		f.SetString("kitty")
	}
}

//测试修改结构体值
func TestSetValue()  {
	u := User{1,"aaron", 21}
	SetValue(&u)
	fmt.Println(u)
}

//测试调用结构体方法
func TestCallStructFunc()  {
	u := User{1, "kitty", 20}
	v := reflect.ValueOf(u)
	//获取方法
	m := v.MethodByName("Hello")
	//构建一些参数
	args := []reflect.Value{ reflect.ValueOf("6666")}
	//var args []reflect.Value
	// 没参数的情况下：var args2 []reflect.Value
	// 调用方法，需要传入方法的参数
	m.Call(args)
}

//获取结构体字段tag
func TestStructTag()  {
	u := User{1, "kitty", 20}
	v := reflect.ValueOf(&u)

	//类型
	t := v.Type()
	//获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}


func main()  {
	u := User{1,"sz", 20}
	Poni(u)
	//Poni(&u)

	fmt.Println("================================")
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t,  t.Name())
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))

	fmt.Println("================================")
	TestSetValue()

	fmt.Println("================================")
	TestCallStructFunc()

	fmt.Println("================================")
	TestStructTag()
}