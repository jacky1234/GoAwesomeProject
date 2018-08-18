package main

import "fmt"

type X int

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

func main() {
	//slice_test()
	//method_test()
	//method2_test()
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
