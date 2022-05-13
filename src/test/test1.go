package main
import (
	"flag"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("USAGE: ./GameServer -c ../conf/server.ini --sid 8777  -d")
	flag.PrintDefaults()

	var opts int64 = 0

	for k := 0; k < 1 ;k++ {
		go func() {
			for i := 0; i < 50; i++ {
				// 注意第一个参数必须是地址
				new := atomic.AddInt64(&opts, 3) //加操作
				fmt.Println("opts: ", new)
				//atomic.AddInt64(&opts, -1) 减操作
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	fmt.Println("opts: ", atomic.LoadInt64(&opts))
}