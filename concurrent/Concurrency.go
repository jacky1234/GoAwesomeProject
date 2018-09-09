package main

import (
	"time"
)

/**
1. 并发:concurrency
2. 并行:parallesim	物理上同一时刻执行多个并发任务
 */
func main() {
	//goTest1()

	exit:=make(chan struct{})			//创建通道。近通知，数据并没有实际的意义

	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")

		close(exit)		//关闭通道，发出信号
	}()

	println("main ...")
	println("main exit.")
}

func goTest1() {
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second) //让goroutine在main逻辑之后执行
		println("go:", x, y)
	}(a, counter())
	//立即计算并复制参数
	a += 100
	println("main:", a, counter())
	time.Sleep(time.Second * 3)
}

var c int

func counter() int {
	c++
	return c
}
