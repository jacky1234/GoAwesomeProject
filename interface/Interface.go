package main

import (
	"fmt"
	"github.com/golang/go/src/pkg/sort"
)

type ByAge []Person //自定义
func main() {
	//interface_test1()
	//puzzle1()
	//sort_test()
	interface_switch()
}

func interface_switch() {
	var x interface{} = func(x int) string {
		return fmt.Sprintf("d: %d", x)
	}
	//y := 199
	//x = &y
	switch v := x.(type) {
	case nil:
		println("nil")
	case *int:
		println(*v)
	case func(int) string:
		println(v(100))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		println("unKnown")
	}
}

func sort_test() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	for k, v := range people {
		fmt.Println(k, v)
	}
	//fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

//Sort 函数的形参是一个 interface
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type Person struct {
	Name string
	Age  int
}

func puzzle1() {
	var d iData = 15
	var x interface{} = d
	if n, ok := x.(fmt.Stringer); ok { //转换为更具体的接口类型
		fmt.Println(n, ok)
	}
	if d2, ok := x.(data); ok {
		fmt.Println(d2)
	}
	e := x.(error)
	fmt.Println(e)
}

type iData int

func (i iData) string() string {
	return fmt.Sprintf("data:%d\n", i)
}

type stringer interface {
	string() string
}

func interface_test1() {
	var d data
	//var t tester = d
	var t tester = &d
	t.test()
	t.string()
}

//接口通常以er结尾
type tester interface {
	test()
	stringer //嵌入其他接口
}

type data struct {
}

func (*data) test() {
	fmt.Println("test....")
}

func (*data) string() string {
	return "123"
}
