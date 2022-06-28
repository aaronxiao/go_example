package main

import (
	"fmt"
	"time"
)

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	//a := Animal{
	//	Name: "Gopher",
	//	Age:  2,
	//}
	//fmt.Println(a)
	//// Output: Gopher (2)

	frameTicker :=  time.NewTicker(time.Duration(2000) * time.Millisecond)

	//frameTicker1 := time.NewTicker(time.Duration(5000) * time.Millisecond)
	//frameTicker2 := time.NewTicker(time.Duration(5000) * time.Millisecond)

	frameTicker1 := frameTicker
	frameTicker2 := frameTicker

	fmt.Println(frameTicker1)
	fmt.Println(frameTicker2)

	var pushDone = make(chan struct{}, 1)

	go func() {
		for {
			select {
			case <-pushDone:
				fmt.Println("groutine 1 done!")
				return
			case <-frameTicker1.C:
				fmt.Println("frameTicker1 done!")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-pushDone:
				fmt.Println("groutine 2 done!")
				return
			case <-frameTicker2.C:
				fmt.Println("frameTicker2 done!")
			}
		}
	}()

	//time.Sleep(1 * time.Second)
	//close(pushDone)
	frameTicker.Stop()
	frameTicker.Reset(time.Duration(4000) * time.Millisecond)


	time.Sleep(10 * time.Second)

	//select {
	//
	//}

	//rand.Seed(time.Now().Unix() )
	//
	//fmt.Println(rand.Perm(5) )
}