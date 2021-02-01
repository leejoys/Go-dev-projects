package main

// Importing fmt
import (
	"fmt"
	"runtime"
)

// Calling main
func main() {
	fmt.Print(runtime.GOMAXPROCS(0))
}
