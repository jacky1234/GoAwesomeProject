package main

func main() {
	//method_point_1()

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