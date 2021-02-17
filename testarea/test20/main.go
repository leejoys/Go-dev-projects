// Вариант кода с бесконечным приращением и остановкой по проверке значения
// (почему-то не всегда работает)
// ==========================================================================

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // Шаг наращивания счётчика
// const step int64 = 1

// // Конечное значение счетчика
// const endCounterValue int64 = 1000

// func main() {

// 	var counter int64 = 0
// 	var wg sync.WaitGroup
// 	var c = sync.NewCond(&sync.Mutex{})
// 	increment := func() {
// 		for {
// 			c.L.Lock()
// 			counter += step
// 			c.Wait()
// 			c.L.Unlock()
// 		}
// 	}

// 	var iterationCount int = int(endCounterValue / step)
// 	for i := 1; i <= iterationCount; i++ {
// 		go increment()
// 	}

// 	wg.Add(1)
// 	go func() {

// 		for {
// 			c.L.Lock()
// 			if counter < endCounterValue {
// 				c.Signal()
// 			} else if counter >= endCounterValue {
// 				wg.Done()
// 				return
// 			}
// 			c.L.Unlock()
// 		}
// 	}()

// 	// Ожидаем поступления сигнала
// 	wg.Wait()
// 	// Печатаем результат, надеясь, что будет 1000
// 	fmt.Println(counter)
// }

package main

import (
	"fmt"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 1000

func main() {

	var counter int64 = 0
	var done = make(chan bool, 0)
	var ready = make(chan bool, 0)
	var doWork = make(chan bool, 0)

	increment := func(done, ready, doWork chan bool) {
		ready <- true
		for {
			select {
			case <-done:
				return
			case <-doWork:
				counter += step
				ready <- true
			}
		}
	}

	go increment(done, ready, doWork)

	for {
		<-ready

		if counter < endCounterValue {
			doWork <- true
		} else {
			done <- true
			break
		}

	}

	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
