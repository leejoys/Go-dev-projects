package main

import (
	"container/heap"
	"fmt"
)

// IntHeap - В основе нашей кучи будет обычный срез
//В куче будем хранить целые числа
type IntHeap []int

// Len - размер кучи
func (h IntHeap) Len() int {
	return len(h)
}

// Определяем логику определения приоритета элементов кучи
// в нашем случае приоритет у того элемента, который просто
// меньше по значению
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Здесь определяем логику перестановки элементов кучи - простое
// копирование в нашем случае
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push - Логика добавления нового элемента в кучу.
// Так как у нас срез в основе кучи -  то простое добавление в
// конец среза
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop - Логика извлечения - обычная для списка, основанного на срезе
// с изменением размера среза
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{20, 30, -5}
	heap.Init(h)
	heap.Push(h, 100)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(h))
	}
}
