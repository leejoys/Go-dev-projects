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
		for i := 1; i <= 100; i++ {
			time.Sleep(time.Second / 4)
			if i%2 != 0 {
				oddChan <- i
			} else {
				evenChan <- i
			}
		}
		close(evenChan)
		close(oddChan)
	}()

	//for {

	select {
	case message, notclosed := <-evenChan:
		if !notclosed {
			break
		} else {
			fmt.Printf("%d %v is even!\n", message, notclosed)
		}
	case message, notclosed := <-oddChan:
		if !notclosed {
			break
		} else {
			fmt.Printf("%d %v is odd!\n", message, notclosed)
		}
	case <-time.Tick(time.Second / 10):
		fmt.Println(time.Now().Format(time.StampMilli))
	}
	//}
}
