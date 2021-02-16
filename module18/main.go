package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(num int) {
			for i := 0; i < 10; i++ {
				fmt.Println(num)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}
