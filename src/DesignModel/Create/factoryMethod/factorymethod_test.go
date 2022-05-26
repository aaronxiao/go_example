package factoryMethod

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	//op1,_ := op.(*OperatorBase)	//这是不成立的   因为OperatorBase没有实现全部的Operator接口方法
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
	)

	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) != 4 {
		t.Fatal("error with factory method pattern")
	}

	factory = MinusOperatorFactory{}
	if compute(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
