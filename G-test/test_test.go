package main

import "testing"

func BenchmarkFibonachi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacchi(10)
	}
}

func BenchmarkFibonachi2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacchi(10)
	}
}
