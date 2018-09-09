package main

import (
	"testing"
	"time"
	"os"
)

/**
单元测试示例
性能测试示例
 */
func add(x, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

//性能测试函数以 Benchmark 为名称前缀
func BenchmarkAdd(b *testing.B) {
	println("B.N =", b.N)
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}

func BenchmarkMinus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minus(1, 2)
	}
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 4 {
		t.FailNow()
	}
}

//Parallel 可有效利用多核并行优势，缩短测试时间
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	if os.Args[len(os.Args) - 1)] == "b"{
		t.Parallel()
	}

	time.Sleep(time.Second * 2)
}
