package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main()  {
	_,err := os.Stat("abc.txt")

	if err != nil {
		fmt.Println(err)
	}

	f, err := filepath.Abs("abc.txt")
	if err == nil {
		fmt.Println(f)
	}
}
