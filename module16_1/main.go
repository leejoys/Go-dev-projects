package main

//Напишите программу, которая запустит n горутин, каждая из которых будет один раз в секунду
//выводить свой идентификатор в консоль. Идентификатором считается порядковый номер горутины.
//Число n задаётся пользователем путём ввода числа в консоль.
//Программа должна выполняться до тех пор, пока пользователь не нажмёт клавишу Enter.

import (
	"fmt"
	"time"
)

func reporter(i int) {
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("Goroutine %d reporting\n", i)
	}
}

func main() {
	var n int
	fmt.Println("Enter number of goroutines:")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		go reporter(i)
	}
	fmt.Scanln()
}
