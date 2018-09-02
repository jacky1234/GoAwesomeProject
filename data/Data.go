package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//string_test()
	//string_default()
	//string_other()
	//string_apple_tobytes()


}

func string_apple_tobytes() {
	string_convert()
	var bs []byte
	bs = append(bs, "abc"...)
	fmt.Println(bs)
}

func string_convert() {
	s := "hello world"
	pp("s: %x\n", &s)
	bs := []byte(s)
	s2 := string(bs)
	pp("string to []byte,bs: %x\n", &bs)
	pp("[]byte to string,s2: %x\n", &s2)
	rs := []rune(s)				//类型rune专门用来作为Unicode码点
	s3 := string(rs)
	pp("string to []rune,bs: %x\n", &rs)
	pp("[]rune to string,s2: %x\n", &s3)
}

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))

	fmt.Printf(format, *h)
}

func string_other() {
	s := `line\r\n
				hello world 
				line 3`
	println(s)

	x := "ab" +
		"cd" +
		"efg"
	println(x == "abcd")
	println(x > "abc")

	//允许以索引号访问字节数组，但不能获取元素地址
	println(x[1])
	//println(&x[1])

	//切片，内部依然指向原字节数组
	x1 := x[:3]
	x2 := x[1:4]
	x3 := x[1:]
	println(x1, x2, x3)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%d:[%c]\n", i, x[i])
	}
}

func string_default() {
	var s string
	println(s == "")
	//println(s == nil)			//无效操作，
}

func string_test() {
	s := "娱乐，的理发师地1233ldfsa"
	fmt.Printf("%s\n", s)
	fmt.Printf("% x len: %d\n", s, len(s))
}
