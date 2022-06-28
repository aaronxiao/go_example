package iterator

import "fmt"

//用于使用相同方式送代不同类型集合或者隐藏集合类型的具体实现

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}


type Numbers struct {
	start, end int
}
func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}
func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}



type NumbersIterator struct {
	numbers *Numbers
	next    int
}
func (i *NumbersIterator) First() {
	i.next = i.numbers.start
}
func (i *NumbersIterator) IsDone() bool {
	return i.next > i.numbers.end
}
func (i *NumbersIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}


func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		a := c.(int)
		fmt.Printf("%#v %d\n", c, a)
	}
}
