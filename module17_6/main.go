package main

import (
	"fmt"
	"sync"
)

//Напишите код, в котором несколько горутин увеличивают значение целочисленного счётчика и
//синхронизируют свою работу через канал. Нужно предусмотреть возможность настройки количества
//используемых горутин и конечного значения счётчика, до которого его следует увеличивать.
//Попробуйте реализовать счётчик с элементами ООП (в виде структуры и методов структуры).
//Попробуйте реализовать динамическую проверку достижения счётчиком нужного значения.

const (
	workers int = 10
	max     int = 10000
)

//CountInt data structure
type CountInt struct {
	countVal *int
	countMax *int
	in       chan int
	out      chan int
	*sync.WaitGroup
}

//NewCount constructor
func NewCount(val, max int) *CountInt {
	c := &CountInt{
		&val,
		&max,
		make(chan int),
		make(chan int),
		&sync.WaitGroup{},
	}

	go func() {
		defer c.Done()
		c.out <- *c.countVal
		for add := range c.in {
			*c.countVal = *c.countVal + add
			c.out <- *c.countVal
		}
	}()

	return c
}

//Increase value
func (c CountInt) Increase(add int) {
	defer c.Done()
	for {
		outVal := <-c.out
		if *c.countMax <= outVal {
			c.in <- 0
			return
		}
		c.in <- add
	}
}

func main() {

	c := NewCount(0, max)
	for i := 0; i < workers; i++ {
		c.Add(1)
		go c.Increase(2)
	}

	c.Wait()
	fmt.Println(<-c.out)
	c.Add(1)
	close(c.in)
	c.Wait()
}
