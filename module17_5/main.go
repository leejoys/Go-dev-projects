package main

import (
	"fmt"
	"sync"
)

//Напишите программу, которая делает следующее: одна горутина по порядку отсылает числа от 1 до 100 в канал,
//вторая горутина их принимает в правильном порядке и печатает на экран (в консоль).

var intChan chan int

func sender(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

func main() {
	var wg sync.WaitGroup
	intChan := make(chan int, 0)
	wg.Add(1)

	go sender(intChan)

	go func(c <-chan int) {
		for i := range c {
			fmt.Println(i)
		}
		wg.Done()
	}(intChan)

	wg.Wait()
}
