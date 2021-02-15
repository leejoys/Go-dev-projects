package main

import (
	"fmt"
	"sync"
	"time"
)

//Напишите код, в котором имеются два канала сообщений из целых чисел, так,
//чтобы приём сообщений из них никогда не приводил к блокировке
//и чтобы вероятность приёма сообщения из первого канала была выше в 2 раза, чем из второго.
//*Если хотите, можете написать код, который бы демонстрировал это соотношение.

func main() {

	var sumOfOne int
	var sumOfTwo int

	oneChan, twoChan := make(chan int, 0), make(chan int, 0)
	close(oneChan)
	close(twoChan)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Scanln()
		wg.Done()
	}()

	go func() {
		for {
			<-time.Tick(time.Second / 4)
			fmt.Println(float64(sumOfTwo) / float64(sumOfOne))
		}
	}()

	go func() {
		for {

			select {
			case <-oneChan:
				sumOfOne++

			case <-twoChan:
				sumOfTwo++

			case <-twoChan:
				sumOfTwo++
			}
		}
	}()
	wg.Wait()
}
