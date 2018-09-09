package main

import "testing"

/**
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
