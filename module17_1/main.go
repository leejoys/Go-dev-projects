package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Количество горутин
const workers int64 = 10

// Шаг наращивания счетчика
const step int64 = 2

// Конечное значение счетчика
const endCounterValue int64 = 1000

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup

	// Не всегда вычисление этой переменной будет приводить к верному
	// результату в счётчике, но для правильных значений
	// и для удобства - можно
	var iterationCount int = int(endCounterValue / workers / step)

	increment := func() {
		defer wg.Done()
		for i := 0; i < iterationCount; i++ {
			atomic.AddInt64(&counter, step)
		}
	}

	for i := 1; i <= int(workers); i++ {
		wg.Add(1)
		go increment()
	}

	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
