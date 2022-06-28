package simpleFactory

import (
	"fmt"
	"testing"
)

//TestType1 test get hiapi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Type1 test fail")
	}

	if hi, ok := api.(*hiAPI); ok{
		*hi = 2
		fmt.Println( hi.Hello("Tom"), "value: ", *hi )
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Type2 test fail")
	}

	if _,ok := api.(*hiAPI); ok{
		t.Fatal( "Type2 test fail" )
	}
}
