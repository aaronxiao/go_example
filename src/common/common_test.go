package common

import (
	"testing"
)

func TestGetCurPwd(t *testing.T) {
	GetCurPwd()
}


func TestConsistent(t *testing.T) {
	hostArray := []string{"127.0.0.1:30000","127.0.0.1:30001"}

	hashConsistent := NewConsistent()
	//采用一致性hash算法，添加节点
	for _, v := range hostArray {
		hashConsistent.Add(v)
	}
	resStr, err :=  hashConsistent.Get("123456")
	if err != nil{
		t.Errorf("Get err: %v", err)
		return
	}
	t.Logf("Get key:123456 resSrt: %s", resStr)

	hashConsistent.add("127.0.0.1:30002")

	resStr, err =  hashConsistent.Get("123456")
	if err != nil{
		t.Errorf("Get err: %v", err)
		return
	}
	t.Logf("After Add Port:127.0.0.1:30002 Get key:123456 resSrt: %s", resStr)


	hashConsistent.Remove("127.0.0.1:30001")
	resStr, err =  hashConsistent.Get("123456")
	if err != nil{
		t.Errorf("Get err: %v", err)
		return
	}
	t.Logf("After Remove Port:127.0.0.1:30001 Get key:123456 resSrt: %s", resStr)
}