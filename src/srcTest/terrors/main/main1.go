package main

import (
	//"errors"
	"fmt"
	"github.com/pkg/errors"
)

// 全局的 错误号 类型，用于API调用之间传递
type MyErrorCode int

// 全局的 错误号 的具体定义
const (
	ErrorBookNotFoundCode MyErrorCode = iota + 1
	ErrorBookHasBeenBorrowedCode
)

// 内部的错误map，用来对应 错误号和错误信息
var errCodeMap = map[MyErrorCode]string{
	ErrorBookNotFoundCode:        "Book was not found",
	ErrorBookHasBeenBorrowedCode: "Book has been borrowed",
}

// Sentinel Error： 即全局定义的Static错误变量
// 注意，这里的全局error是没有保存堆栈信息的，所以需要在初始调用处使用 errors.Wrap
var (
	ErrorBookNotFound        = NewMyError(ErrorBookNotFoundCode)
	ErrorBookHasBeenBorrowed = NewMyError(ErrorBookHasBeenBorrowedCode)
)

func NewMyError(code MyErrorCode) *MyError {
	return &MyError{
		Code:    code,
		Message: errCodeMap[code],
	}
}

// error的具体实现
type MyError struct {
	// 对外使用 - 错误码
	Code MyErrorCode
	// 对外使用 - 错误信息
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func main() {
	books := []string{
		"Hamlet",
		"Jane Eyre",
		"War and Peace",
	}

	for _, bookName := range books {
		fmt.Printf("%s start\n===\n", bookName)

		err := borrowOne(bookName)
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
		fmt.Printf("===\n%s end\n\n", bookName)
	}
}

//借一本书
func borrowOne(bookName string) error {
	// Step1: 找书
	err := searchBook(bookName)

	// Step2: 处理
	// 特殊业务场景：如果发现书被借走了，下次再来就行了，不需要作为错误处理
	if err != nil {
		// 提取error这个interface底层的错误码，一般在API的返回前才提取
		// As - 获取错误的具体实现
		var myError = new(MyError)

		if errors.As(err, &myError) {
			fmt.Printf("error code is %d, message is %s\n", myError.Code, myError.Message)
		}

		// 特殊逻辑: 对应场景2，指定错误(ErrorBookHasBeenBorrowed)时，打印即可，不返回错误
		// Is - 判断错误是否为指定类型
		if errors.Is(err, ErrorBookHasBeenBorrowed) {
			fmt.Printf("book %s has been borrowed, I will come back later!\n", bookName)
			err = nil
		}
	}
	//return err
	return errors.WithMessage(err, "borrowOne")
}

func searchBook(bookName string) error {
	// 下面两个 error 都是不带堆栈信息的，所以初次调用得用Wrap方法
	// 如果已有堆栈信息，应调用WithMessage方法

	// 3 发现图书馆不存在这本书 - 认为是错误，需要打印详细的错误信息
	if len(bookName) > 10 {
		return errors.Wrapf(ErrorBookNotFound, "bookName is %s", bookName)
	} else if len(bookName) > 8 {
		// 2 发现书被借走了 - 打印一下被接走的提示即可，不认为是错误
		return errors.Wrapf(ErrorBookHasBeenBorrowed, "bookName is %s", bookName)
	}
	// 1 找到书 - 不需要任何处理
	return nil
}
