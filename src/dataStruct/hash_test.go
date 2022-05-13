package dataStruct

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)


func TestHash_1W(t *testing.T)  {
	rand.Seed(time.Now().UnixNano() )

	hash := make(map[int]int)

	now := time.Now()
	for i := 0;i < 10000 ;i++  {
		hash[i] = i
	}

	finish := time.Since(now)
	fmt.Println("插入10000条数据耗时：", finish)

	var ok = false
	for i := 0 ;i < 1000000 ;i++ {
		nr := rand.Intn(10000)
		_, ok = hash[nr]
	}
	finish1 := time.Since(now)
	if ok {
		fmt.Println("找到查找键值耗时: ", finish1 - finish)
	}else{
		fmt.Println("未找到查找键值耗时: ", finish - finish)
	}
}


func TestHash_10W(t *testing.T)  {
	rand.Seed(time.Now().UnixNano() )

	hash := make(map[int]int)

	now := time.Now()
	for i := 0 ;i < 100000 ;i++  {
		hash[i] = i
	}

	finish := time.Since(now)
	fmt.Println("插入100000条数据耗时：", finish)

	var ok = false
	for i := 0 ;i < 1000000 ;i++ {
		nr := rand.Intn(10000)
		_, ok = hash[nr]
	}

	finish1 := time.Since(now)
	if ok {
		fmt.Println("找到查找键值耗时: ", finish1 - finish)
	}else{
		fmt.Println("未找到查找键值耗时: ", finish1 - finish)
	}
}



func TestHash_100W(t *testing.T)  {
	rand.Seed(time.Now().UnixNano() )

	hash := make(map[int]int)

	now := time.Now()
	for i := 0 ;i < 1000000 ;i++  {
		hash[i] = i
	}

	finish := time.Since(now)
	fmt.Println("插入1000000条数据耗时：", finish)

	var ok = false
	for i := 0 ;i < 1000000 ;i++ {
		nr := rand.Intn(10000)
		_, ok = hash[nr]
	}
	finish1 := time.Since(now)
	if ok {
		fmt.Println("找到查找键值耗时: ", finish1 - finish)
	}else{
		fmt.Println("未找到查找键值耗时: ", finish1 - finish)
	}
}

