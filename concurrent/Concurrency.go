package main

import (
	"time"
	"sync"
	"github.com/golang/go/src/pkg/math"
	"runtime"
)

/**
1. 并发:concurrency
2. 并行:parallesim	物理上同一时刻执行多个并发任务
 */
func main() {
	//goTest1()
	//go_exit()
	//wait_group()
	//wait_group2()

	//GOMAXPROCS仅返回当前设置值，不做任何修改
	n := runtime.GOMAXPROCS(0)
	println((1+math.MaxUint32)*math.MaxUint32/2)
	println(runtime.NumCPU())
	test(n)
	//test2(n)
}

//测试目标函数
func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}

	println(x)
}

func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

//并发执行
func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}

	wg.Wait()
}

func wait_group2() {
	var wg sync.WaitGroup
	//建议在goroutine外增加计数，防止Add未执行，Wait已经退出
	wg.Add(1)
	go func() {
		wg.Wait() //等待归零，接触阻塞
		println("wait exit!")
	}()
	go func() {
		time.Sleep(time.Second)
		println("done")
		wg.Done()
	}()
	wg.Wait()
	//等待归零，接触阻塞
	println("main exit.")
}

func wait_group() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) //累加计数,实现了原子操作

		go func(id int) {
			defer wg.Done() //计数减一

			time.Sleep(time.Second)
			println("goroutine", id, "done.")
		}(i)
	}
	println("main...")
	wg.Wait()
	//阻塞，直到计数归零
	println("main exit.")
}

func go_exit() {
	exit := make(chan struct{})
	//创建通道。仅通知，数据并没有实际的意义
	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")

		close(exit) //关闭通道，发出信号
	}()
	println("main ...")
	<-exit
	//如通道关闭，立即解除阻塞
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
