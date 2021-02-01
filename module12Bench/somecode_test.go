package main

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.ReportAllocs()
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big bad case", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateDownSlice(100000)
			b.ReportAllocs()
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big good case", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateUpSlice(100000)
			b.ReportAllocs()
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big sort pkg", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.StartTimer()
			sortSort(ar)
			b.StopTimer()
		}
	})

}

func BenchmarkSelectionSort(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.ReportAllocs()
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big good case", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateUpSlice(100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big bad case", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateDownSlice(100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkSelectionSortMinMax(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100)
			b.StartTimer()
			selectionSortMinMax(ar)
			b.StopTimer()
		}
	})

	b.Run("middle", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSortMinMax(ar)
			b.StopTimer()
		}
	})

	b.Run("big", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.ReportAllocs()
			b.StartTimer()
			selectionSortMinMax(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkInsertionSort(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.ReportAllocs()
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkMergeSort(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100)
			b.StartTimer()
			ar = mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			ar = mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.ReportAllocs()
			b.StartTimer()
			ar = mergeSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkQuickSort(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 100000)
			b.ReportAllocs()
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("too big 1k max", func(b *testing.B) {
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000, 200000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("too big 50 max", func(b *testing.B) {
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(50, 100000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
}
