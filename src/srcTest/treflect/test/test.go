package main

import (
	"fmt"
	"reflect"
)

func reflect_type(a interface{} )  {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)

	k := t.Kind()		//获得具体的类型
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("k is float64 \n")
	case reflect.String:
		fmt.Printf("k is string \n")

	}
}

func reflect_value(a interface{})  {
	v := reflect.ValueOf(a)

	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k )
	switch k {
	case reflect.Float64:
		fmt.Println(v.Float() )
	}
}

//反射修改值信息
func reflect_set_value(a interface{})  {
	v := reflect.ValueOf(a)
	//v = v.Elem()				//加和不加会有区别
	k := v.Kind()
	fmt.Println(v.Type())
	fmt.Println("k kind", k)
	switch k {
	case reflect.Float64:
		v.SetFloat(2.0)
		println("a is", v.Float() )
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(3.0)
		println("case", v.Elem().Float() )
		fmt.Println(v.Pointer() )
	}
}


func main() {
	var x float64 = 1.0
	reflect_type(x)
	fmt.Println("=========================")
	reflect_value(x)
	fmt.Println("=========================")
	reflect_set_value(&x)		//不能填x   因为是一个副本  对一个副本修改值显然没意义 运行会报错
}
