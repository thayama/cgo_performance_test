package cgotest

import "testing"

func BenchmarkCallGoFromC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallCFunction()
	}
}

func BenchmarkCallGoFromC2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallCFunction2()
	}
}

func BenchmarkCallGoFromC3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallCFunction3()
	}
}

func BenchmarkCallGoFromCNop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallCFunctionNop()
	}
}

func BenchmarkCallGoFromCNop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallCFunctionNop2()
	}
}
