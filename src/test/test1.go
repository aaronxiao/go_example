package main

import "fmt"

func test()(int, error)  {
	return 0, nil
}

// 连接复制
func ConnCopy(ch chan error) {
	_, err := test()
	ch <- err
}

func main() {
	ch := make(chan error, 1)
	go ConnCopy(ch)
	//go ConnCopy(ch)
	fmt.Println( <-ch)
	fmt.Println("1111111111111")


}