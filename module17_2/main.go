package main

import (
	"fmt"
	"sync"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 1000

func main() {

	var counter int64 = 0
	var c = sync.NewCond(&sync.RWMutex{})
	increment := func() {
		c.Lock()
		counter += step
		c.Unlock()
	}

	var iterationCount int = int(endCounterValue / step)
	for i := 1; i <= iterationCount; i++ {
		go increment()
	}

	go func(){
		for{
			if counter==endCounterValue{
				c.Signal()
			}
		}
	}

	// Ожидаем поступления сигнала
	c.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
