package main

import "fmt"

func main() {
	var d data
	//var t tester = d
	var t tester = &d
	t.test()
	t.string()
}

//接口通常以er结尾
type tester interface {
	test()
	string() string
}

type data struct {

}

func (*data) test(){
	fmt.Println("test....")
}

func (*data) string() string{
	return "123"
}