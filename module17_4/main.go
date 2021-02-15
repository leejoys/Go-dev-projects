package main

import (
	"fmt"
	"time"
)

//Напишите код, в котором имеются два канала сообщений из целых чисел так,
//чтобы приём сообщений всегда приводил к блокировке.
//Приёмом сообщений из обоих каналов будет заниматься главная горутина. Сделайте так, чтобы во время такого
//«бесконечного ожидания» сообщений выполнялась фоновая работа в виде вывода текущего времени в консоль.

func main() {

	evenChan, oddChan := make(chan int, 0), make(chan int, 0)

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second / 2)
			if i%2 != 0 {
				oddChan <- i
			} else {
				evenChan <- i
			}
		}
		close(evenChan)
		close(oddChan)
	}()

	for {

		select {
		case _, notclosed := <-evenChan:
			if !notclosed {
				return
			}
			fmt.Printf("Even!\n")

		case _, notclosed := <-oddChan:
			if !notclosed {
				return
			}
			fmt.Printf("Odd!\n")

		case <-time.Tick(time.Second / 10):
			fmt.Println(time.Now().Format(time.StampMilli))
		}
	}
}
