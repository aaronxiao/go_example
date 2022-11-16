package main

import (
	"runtime/pprof"
	_ "net/http/pprof"
	"sync"
)

func allocate() {
	_ = make([]byte, 1<<20)
}

//GODEBUG=schedtrace=1000 ./main
//GODEBUG=scheddetail=1,schedtrace=1000 ./main
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				allocate()
				counter++
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}