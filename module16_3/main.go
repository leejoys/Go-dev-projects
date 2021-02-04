package main

import (
	"fmt"
	"sync"
	"time"
)

func printValue(i *int, m *sync.Mutex) {
	for {
		time.Sleep(time.Second * 1 / 3)
		m.Lock()
		fmt.Printf("Value is %d \n", *i)
		m.Unlock()
	}
}

func addValue(i *int, m *sync.Mutex) {
	for {

		time.Sleep(time.Second * 1)
		m.Lock()
		*i++
		m.Unlock()
	}
}

func main() {
	var someValue int
	var n int
	var mutex = &sync.Mutex{}
	fmt.Println("Enter duration in seconds")
	fmt.Scanln(&n)

	go addValue(&someValue, mutex)
	go printValue(&someValue, mutex)

	time.Sleep(time.Second * time.Duration(n))
}
