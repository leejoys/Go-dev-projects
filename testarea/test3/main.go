package main

import (
	"fmt"
)

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := arr[:]
	slice2 := arr[:]
	fmt.Println("slice ", slice)
	fmt.Println("slice2 ", slice2)
	fmt.Println("array ", arr)
	fmt.Printf("array address %p\n", &arr)
	fmt.Printf("slice address %p\n", &slice)
	fmt.Printf("slice2 address %p\n", &slice2)
	fmt.Printf("array elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7], &arr[8], &arr[9])
	fmt.Printf("slice elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4], &slice[5], &slice[6], &slice[7], &slice[8], &slice[9])
	fmt.Printf("slice elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &slice2[0], &slice2[1], &slice2[2], &slice2[3], &slice2[4], &slice2[5], &slice2[6], &slice2[7], &slice2[8], &slice2[9])

}
