package main

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
