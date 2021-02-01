package main

import (
	"fmt"
)

func hashint64(val int64) uint64 {
	return uint64(val) % 1000
}

func main() {
	fmt.Printf("хэш от 976576455: %d\n", hashint64(976576455))

}
