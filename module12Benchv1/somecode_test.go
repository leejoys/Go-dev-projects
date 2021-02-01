package main

import (
	"testing"
)

func BenchmarkSample(b *testing.B) {

	b.StartTimer()
	b.Run("small arrays bubbles", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}

	})

	b.Run("middle arrays bubbles", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays bubbles", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays sort pkg", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.StartTimer()
			sortSort(ar)
			b.StopTimer()
		}
	})

	b.StopTimer()
}
