package main

import "testing"

func BenchmarkEchoStr(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		EchoStr()
	}
}
