package main

import (
	"container/ring"
	"fmt"
)

// rBuf data structure
type rBuf struct {
	dataRing *ring.Ring
}

//newrBuf make rBuf
func newrBuf(size int) *rBuf {

	return &rBuf{
		ring.New(size),
	}
}

//Pull() — добавление нового элемента в буфер, в случае
//добавления в заполненный буфер стирается самый «старый» элемент
func (r *rBuf) Pull(el int) {
	r.dataRing.Value = el
	r.dataRing = r.dataRing.Next()
}

//получение всех упорядоченных элементов буфера
//и последующая очистка буфера
func (r *rBuf) Get() []int {
	resultArr := make([]int, 0)
	size := r.dataRing.Len()

	for i := 0; i < size; i++ {
		if r.dataRing.Value == nil {
			r.dataRing = r.dataRing.Next()
			continue
		}
		resultArr = append(resultArr, r.dataRing.Value.(int))
		r.dataRing = r.dataRing.Next()

	}

	return resultArr
}
func main() {
	r := newrBuf(5)
	r.Pull(1)
	r.Pull(2)
	r.Pull(3)
	// r.Pull(4)
	// r.Pull(5)
	// r.Pull(6)
	fmt.Println(r.Get())
}
