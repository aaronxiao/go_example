package main

import (
	"fmt"
	"time"
)



func main() {
	d := time.Duration(time.Second*2)
	t := time.NewTimer(d)
	defer t.Stop()
	for {
		select {
		case <- t.C:
			fmt.Println("timeout...")
			return
		default:
			fmt.Println("Not timeout...")
			time.Sleep(time.Second * 1)
			t.Reset(time.Second*2)

		}
	}
}