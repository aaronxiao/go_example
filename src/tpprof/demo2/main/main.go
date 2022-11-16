package main

import (
	"Practice/tpprof/demo2/data"
	"log"
	"net/http"
	_ "net/http/pprof"

)




func main() {
	go func() {			//这是一个有问题的程序 内存会一直增加知道撑爆
		for {
			log.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
