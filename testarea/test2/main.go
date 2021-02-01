package main

import (
	"fmt"
)

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := arr[:]

	fmt.Println("slice ", slice)
	fmt.Println("array ", arr)
	fmt.Printf("array address %p\n", &arr)
	fmt.Printf("array elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7], &arr[8], &arr[9])
	fmt.Printf("slice elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4], &slice[5], &slice[6], &slice[7], &slice[8], &slice[9])

	arr = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println("slice ", slice)
	fmt.Printf("array address %p\n", &arr)
	fmt.Printf("array elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7], &arr[8], &arr[9])
	fmt.Printf("slice elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4], &slice[5], &slice[6], &slice[7], &slice[8], &slice[9])

	arr[0] = 9
	arr[1] = 8
	arr[2] = 7
	arr[3] = 6
	arr[4] = 5
	arr[5] = 4
	arr[6] = 3
	arr[7] = 2
	arr[8] = 1
	arr[9] = 0
	fmt.Println("slice ", slice)
	fmt.Printf("slice elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4], &slice[5], &slice[6], &slice[7], &slice[8], &slice[9])

}
