package main

import (
	"fmt"
	"errors"
	"unsafe"
)

type X int

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

type Printer interface {		//接口
	Print()
}

func (u user) Print() {
	sprint := fmt.Sprint("%+v\n", u)
	fmt.Print(sprint)
}

func main() {
	//slice_test()
	//method_test()
	//method2_test()
	//method_multiple_return()
	//interface_test1()
	//const_test1()
	//enum_iota_test1()
	//enum_iota_test2()
	//puzzle()
	//for_range()
	//for_test()
	//switch_test()
	//defer_test()
}

func defer_test() {
	defer println("dispose...")
	c, e := div(10, 0)
	fmt.Println(10 / 0)
	fmt.Println(c, e)
}

func method_multiple_return() {
	a, b := 10, 0
	c, err := div(a, b)
	fmt.Println(c, err)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func switch_test() {
	x := 100
	switch {
	case x > 10:
		println("x")
	case x < 0:
		println("-x")
	case x == 0:
		println("0")
	}
}

func for_test() {
	for i := 0; i < 4; i++ {
		println(i)
	}
}

func for_range() {
	x := []int{100, 101, 102}
	for i, n := range x {
		println(i, ":", n)
	}
}

func puzzle() {
	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)
}

func enum_iota_test2() {
	test_color(red)
	test_color(1<<8 - 1)
}

type color byte

const (
	black color = iota
	red
	blue
)

func test_color(c color) {
	println(c)
}

func enum_iota_test1() {
	const (
		B  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)
	println(B)
	println(KB)
	println(MB)
	println(GB)
}

func const_test1() {
	//import "unsafe"
	const (
		ptrSize = unsafe.Sizeof(uintptr(0))
		strSize = len("hello world")
	)
	println(ptrSize)
	println(strSize)
}

func interface_test1() {
	var u user
	u.name = "jack"
	u.age = 30
	var p Printer = u
	p.Print()
}

func method2_test() {
	var m manager
	m.name = "Tom"
	m.age = 29
	println(m.ToString())
}

func (u user) ToString() string {
	return fmt.Sprint("%+v", u)
}

func method_test() {
	var x X
	x.inc()
	println(x)
}
func (x *X) inc() {
	*x++
}

func slice_test() {
	x := make([]int, 0, 5) //切片

	for i := 0; i < 8; i++ {
		x = append(x, i)
	}

	fmt.Println(x)
}
