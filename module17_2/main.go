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
	var c = sync.NewCond(&sync.Mutex{})
	increment := func() {
		c.L.Lock()
		counter += step
		c.L.Unlock()
	}

	var iterationCount int = int(endCounterValue / step)
	for i := 1; i <= iterationCount; i++ {
		go increment()
	}

	go func() {
		for {
			c.L.Lock()
			if counter >= endCounterValue {
				c.Signal()
			}
			c.L.Unlock()
		}
	}()

	// Ожидаем поступления сигнала
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
