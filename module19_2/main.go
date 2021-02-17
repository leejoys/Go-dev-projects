package main

import (
	"container/ring"
	"fmt"
	"sync"
)

// rBuf data structure
type rBuf struct {
	dataRing *ring.Ring
	sync.Mutex
}

//newrBuf make rBuf
func newrBuf(size int) *rBuf {

	return &rBuf{
		ring.New(size),
		sync.Mutex{},
	}
}

//Pull() — добавление нового элемента в буфер, в случае
//добавления в заполненный буфер стирается самый «старый» элемент
func (r *rBuf) Pull(el int) {
	r.Lock()
	r.dataRing.Value = el
	r.dataRing = r.dataRing.Next()
	r.Unlock()
}

//получение всех упорядоченных элементов буфера
//и последующая очистка буфера
func (r *rBuf) Get() []int {
	resultArr := make([]int, 0)
	r.Lock()
	size := r.dataRing.Len()

	for i := 0; i < size; i++ {
		if r.dataRing.Value == nil {
			r.dataRing = r.dataRing.Next()
			continue
		}
		resultArr = append(resultArr, r.dataRing.Value.(int))
		r.dataRing.Value = nil
		r.dataRing = r.dataRing.Next()

	}
	r.Unlock()
	return resultArr
}
func main() {
	var wg sync.WaitGroup
	r := newrBuf(5)

	wg.Add(2)
	go func() {
		r.Pull(1)
		r.Pull(2)
		wg.Done()
	}()

	go func() {
		r.Pull(3)
		r.Pull(4)
		wg.Done()
	}()

	// r.Pull(5)
	// r.Pull(6)
	wg.Wait()
	fmt.Println(r.Get())
	fmt.Println(r.Get())
}
