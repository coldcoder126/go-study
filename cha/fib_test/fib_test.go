package fib_test

import (
	"go_study/cha"
	"testing"
)

func benFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		cha.Fib(n)
	}
}

func BenchmarkFib2(b *testing.B) {
	benFib(b, 10)
}
func BenchmarkFib3(b *testing.B) {
	benFib(b, 20)
}
