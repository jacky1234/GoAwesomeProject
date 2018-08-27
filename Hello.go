package main

import (
	"fmt"
	"unsafe"
	"math"
	"strconv"
	"errors"
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

type Printer interface {
	//接口
	Print()
}

func (u user) Print() {
	sprint := fmt.Sprint("%+v\n", u)
	fmt.Print(sprint)
}

func forSliceRange(s []int) { //切片做函数参数
	s[0] = -1
	for i, v := range s {
		fmt.Printf("s[%d]=%d, ", i, v)
		//s[0]=-1, s[1]=1, s[2]=2, s[3]=3, s[4]=4, s[5]=5, s[6]=6, s[7]=7, s[8]=8, s[9]=9,
	}
	fmt.Println("\n")
}

func main() {
	//slice_test()
	//slice_copy()
	//slice_method()

	//method_test()
	//method2_test()
	//method_multiple_return()
	//interface_test1()
	//const_test1()
	//enum_iota_test1()
	//enum_iota_test2()
	//math_test()
	//number_convert()
	//enum_iota_test3()
	//puzzle()
	//for_range()
	//for_test()
	//switch_test()
	//defer_test()
	//delay_toConfirm1()

	//var a = struct {
	//	x int
	//}{100}
	//
	//println(a)

	var a *int = test()
	println(a, *a)
}

func slice_copy() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:]
	//{8, 9}
	s2 := data[:5]
	//{0, 1, 2, 3, 4}
	copy(s2, s1)
	// dst:s2, src:s1
	fmt.Println(s1)
	fmt.Println(s2)
	//for i, n := range data {
	//	fmt.Printf("key = %d,value = %d\n", i, n)
	//}
	fmt.Println(s2)
	//[8 9 2 3 4]
	fmt.Println(data)
	//[8 9 2 3 4 5 6 7 8 9]
}

func slice_method() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	forSliceRange(slice)
}

func base_compare() {
	//? max_float 表示的范围不是很大吗？
	var a float32 = 1.1234567899
	var b float32 = 1.12345678
	var c float32 = 1.123456781
	println(a, b, c)
	println(a == b, a == c)
	fmt.Printf("%v %v, %v\n", a, b, c)
}

func number_convert() {
	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)
	println(a, b, c)
	println("0b" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(a, 8))
	println("0x" + strconv.FormatInt(a, 16))
}

func delay_toConfirm1() {
	a := 1 << 2
	a |= 1 << 6
	fmt.Printf("0b%d", a)
}

func
enum_iota_test3() {
	println(Sunday)
	println(Monday)
	println(Tuesday)
}

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//func defer_test() {
//	defer println("dispose...")
//	c, e := div(10, 0)
//	fmt.Println(10 / 0)
//	fmt.Println(c, e)
//}

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

func
switch_test() {
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

func
for_test() {
	for i := 0; i < 4; i++ {
		println(i)
	}
}

func
for_range() {
	x := []int{100, 101, 102}
	for i, n := range x {
		println(i, ":", n)
	}
}

func
math_test() {
	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)

	fmt.Printf("0b%b, %#o, %#x\n", a, a, a)
	fmt.Println(math.MinInt8, math.MaxInt8)

}

func
enum_iota_test2() {
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
	x := make([]int, 0) //切片
	c := cap(x)
	for i := 0; i < 25; i++ {
		x = append(x, i, i*10)
		if n := cap(x); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}

	fmt.Println(x)

	//var s1 []int
	//s2 := []int{}
	s5 := []int{1, 2, 3}
	s5 = append(s5, 4);
	fmt.Println(s5)
	fmt.Println(s5[2])

	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[2:5]
	fmt.Println(s1)

	array[1] = 100
	fmt.Println(array)
}
