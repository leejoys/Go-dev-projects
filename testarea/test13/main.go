package main

import (
	"bufio"
	"fmt"
	"os"
)

var i int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	for scanner.Scan() {

		msg := scanner.Text()
		fmt.Println(msg)
	}

}
