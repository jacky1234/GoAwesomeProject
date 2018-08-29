package main

import (
	"fmt"
)

func main() {
	//method_point_1()
	//method_point_2()
	//add_method_for_object()
	//unFixed_param()
	//anonymous_method()
	//closure_test()
	//class_inherit_test()

	base := Base{}
	t1 := &TypeOne{1,base}
	t2 := &TypeTwo{2,base}
	t3 := &TypeThree{3,base}
	//https://blog.csdn.net/wangshubo1989/article/details/73287204
}

type Base struct {

}

type Baser interface {
	Get() float32
}

type TypeOne struct {
	value float32
	Base
}

type TypeTwo struct {
	value float32
	Base
}

type TypeThree struct {
	value float32
	Base
}

func (t *TypeOne) Get() float32 {
	return t.value
}

func (t *TypeTwo) Get() float32 {
	return t.value * t.value
}

func (t *TypeThree) Get() float32 {
	return t.value + t.value
}

func class_inherit_test() {
	mark := Student{Human{"Jack", 12, "18516606250"}, "nantong university"}
	sam := Employee{Human{"Alice", 30, "983094012"}, "Xiaoxin Company"}
	mark.SayHi()
	sam.SayHi()
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//重写
func (e *Employee) SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

func closure_test() {
	f := closure(123)
	f()
}

func closure(x int) func() {
	return func() {
		fmt.Printf("the value is %v\n", x)
	}
}

func anonymous_method() {
	//直接执行
	func(s string) {
		println(s)
	}("hello world")

	//赋值给变量
	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 2))
	println(add(10, 20))

	//作为参数
	as_param(func() {
		println("hello world")
	})

	//作为返回值
	result := asfunc()
	println(result(10, 29))
}

func asfunc() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func as_param(f func()) {
	println("before func todo")
	f()
	println("after func todo")
}

func unFixed_param() {
	print("1", 2, "3", 4, "asdf")
}

func print(a ...interface{}) {
	for _, v := range a {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
}

func add_method_for_object() {
	p := person{name: "jack"}
	fmt.Println(p.toString())
	p.modifyName()
	fmt.Println(p.toString())
	p.modifyName2()
	fmt.Println(p.toString())
	(&p).modifyName() //如果我们没有这么强制使用指针进行调用，Go的编译器自动会帮我们取指针
	fmt.Println(p.toString())
}

//值接收者
func (p person) modifyName() {
	p.name = "里斯"
}

//指针接收者,指针具有指向原有值的特性
func (p *person) modifyName2() {
	p.name = "里斯"
}

type person struct {
	name string
}

//现在我们说，类型person有了一个toString方法，现在我们看下如何使用它
func (p person) toString() string {
	return "the person name is " + p.name
}

func method_point_2() {
	a := 0x100
	p := &a
	fmt.Printf("pointer : %p, target: %v, value:%v\n", &p, p, *p)
	test_param(p)
}

func test_param(x *int) {
	fmt.Printf("pointer : %p, target: %v, value:%v\n", &x, x, *x)
}

func method_point_1() {
	var a *int = test()
	println(a, *a)
}

//////////////method begin////////////
/**
1. 不支持overload
2. 函数属于第一类对象
 */

//1.  go build -gcflags "-l -m"  禁用函数内联，输出优化信息
//2.  go tool objdump -s "main\.main" test		反汇编确认?
//3. .test
func test() *int {
	a := 0x110
	return &a
}

func methodBegin(n int) {

}

///////////////////////////
