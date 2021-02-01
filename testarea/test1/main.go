package main

import (
	"fmt"
	"testarea/some"
)

func do(a, c int) int {
	return c
}

func main() {
	doSomething(func(number int) { fmt.Printf("number: %d", number) })
	fmt.Println(do(1, 2))

	iphone := some.Phone{IsSmartphone: true}

	fmt.Println(iphone.IsSmartphone)

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := arr[:]

	fmt.Printf("array address %p\n", &arr)
	fmt.Printf("array elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7], &arr[8], &arr[9])

	fmt.Printf("slice address              %p\n", &slice)
	fmt.Printf("slice elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4], &slice[5], &slice[6], &slice[7], &slice[8], &slice[9])
	slice = append(slice, 10)

	fmt.Printf("slice address after append %p\n", &slice)
	fmt.Printf("slice elems add: 0:%p | 1:%p | 2:%p | 3:%p | 4:%p | 5:%p | 6:%p | 7:%p | 8:%p | 9:%p | \n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4], &slice[5], &slice[6], &slice[7], &slice[8], &slice[9])

}

func doSomething(anonymousFunc func(int)) {
	anonymousFunc(100)
}
