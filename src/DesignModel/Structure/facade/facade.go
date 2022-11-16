package facade

import "fmt"

//API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。

//facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

//API is facade interface of facade package
type API interface {
	Test() string
}

//facade implement
type apiImpl struct {
	a AModuleAPI			//小写 隐藏细节 简化访问
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}


//AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}


//BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}