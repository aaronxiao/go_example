package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)


//runtime.MemStats这个结构体包含的字段比较多，但是大多都很有用：
//1、Alloc uint64 //golang语言框架堆空间分配的字节数
//2、TotalAlloc uint64 //从服务开始运行至今分配器为分配的堆空间总 和，只有增加，释放的时候不减少
//3、Sys uint64 //服务现在系统使用的内存
//4、Lookups uint64 //被runtime监视的指针数
//5、Mallocs uint64 //服务malloc heap objects的次数
//6、Frees uint64 //服务回收的heap objects的次数
//7、HeapAlloc uint64 //服务分配的堆内存字节数
//8、HeapSys uint64 //系统分配的作为运行栈的内存
//9、HeapIdle uint64 //申请但是未分配的堆内存或者回收了的堆内存（空闲）字节数
//10、HeapInuse uint64 //正在使用的堆内存字节数
//10、HeapReleased uint64 //返回给OS的堆内存，类似C/C++中的free。
//11、HeapObjects uint64 //堆内存块申请的量
//12、StackInuse uint64 //正在使用的栈字节数
//13、StackSys uint64 //系统分配的作为运行栈的内存
//14、MSpanInuse uint64 //用于测试用的结构体使用的字节数
//15、MSpanSys uint64 //系统为测试用的结构体分配的字节数
//16、MCacheInuse uint64 //mcache结构体申请的字节数(不会被视为垃圾回收)
//17、MCacheSys uint64 //操作系统申请的堆空间用于mcache的字节数
//18、BuckHashSys uint64 //用于剖析桶散列表的堆空间
//19、GCSys uint64 //垃圾回收标记元信息使用的内存
//20、OtherSys uint64 //golang系统架构占用的额外空间
//21、NextGC uint64 //垃圾回收器检视的内存大小
//22、LastGC uint64 // 垃圾回收器最后一次执行时间。
//23、PauseTotalNs uint64 // 垃圾回收或者其他信息收集导致服务暂停的次数。
//24、PauseNs [256]uint64 //一个循环队列，记录最近垃圾回收系统中断的时间
//25、PauseEnd [256]uint64 //一个循环队列，记录最近垃圾回收系统中断的时间开始点。
//26、NumForcedGC uint32 //服务调用runtime.GC()强制使用垃圾回收的次数。
//27、GCCPUFraction float64 //垃圾回收占用服务CPU工作的时间总和。如果有100个goroutine，垃圾回收的时间为1S,那么就占用了100S。
//28、BySize //内存分配器使用情况

type Monitor struct {
	Alloc 			uint64 	`json:"堆分配字节数"`
	TotalAlloc 		uint64 	`json:"堆分配字节总数"`
	Sys  			uint64 	`json:"获得系统总内存"`
	Mallocs 		uint64 	`json:"分配对象数"`
	Frees   		uint64 	`json:"释放对象数"`
	LiveObjects  	uint64 	`json:"存活对象数"`
	PauseTotalNs 	uint64 	`json:"应用开始总GC暂停数"`

	NumGC     		uint32 `json:"GC 循环完成数"`
	NumGoroutine 	int    `json:"goroutine数量"`					//goroutine数量
}

func NewMonitor() {
	var m Monitor
	var rtm runtime.MemStats
	//var interval = time.Duration(1) * time.Second
	//for {
		//<-time.After(interval)

		// Read full mem stats
		runtime.ReadMemStats(&rtm)

		// Number of goroutines
		m.NumGoroutine = runtime.NumGoroutine()

		// Misc memory stats
		m.Alloc 		= rtm.Alloc
		m.TotalAlloc 	= rtm.TotalAlloc
		m.Sys 			= rtm.Sys
		m.Mallocs 		= rtm.Mallocs
		m.Frees 		= rtm.Frees

		// Live objects = Mallocs - Frees
		m.LiveObjects = m.Mallocs - m.Frees

		// GC Stats
		m.PauseTotalNs 	= rtm.PauseTotalNs
		m.NumGC 		= rtm.NumGC

		// Just encode to json and print
		b, _ := json.Marshal(m)
		fmt.Println(string(b))
	//}
}
