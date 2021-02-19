package main

import (
	"fmt"
	"strconv"
)

func main() {
	var anything [...]string
	fmt.Println("Enter numbers separated by whitespaces:")
	fmt.Scanln(&anything)
	fmt.Println(anything)
	//someSlice := strings.Fields(anything)
	//fmt.Println(someSlice)
	for _, some := range anything {
		fmt.Println(some)
		if num, err := strconv.ParseInt(some, 10, 0); err == nil {
			fmt.Println(num)
			fmt.Println(int(num))
		}
	}

}
