package main

import (
	"fmt"
	"math"
)

func hashstr(val string) (valIntUint64 uint64) {
	valBytes := []rune(val)
	fmt.Println(valBytes)
	for i := 0; i < len(valBytes)-1; i++ {
		valIntUint64 = valIntUint64 + uint64(math.Pow(float64(valBytes[i])/10, float64(valBytes[i+1])/20))
	}
	return valIntUint64 % 1000
}

func main() {

	fmt.Printf("хэш от Golang: %d\n", hashstr("Golang"))
	fmt.Printf("хэш от Gangol: %d\n", hashstr("Gangol"))
	fmt.Printf("хэш от Gonlag: %d\n", hashstr("Gonlag"))
	fmt.Printf("хэш от Goglan: %d\n", hashstr("Goglan"))
	fmt.Printf("хэш от Gongla: %d\n", hashstr("Gongla"))
	fmt.Printf("хэш от Ganolg: %d\n", hashstr("Ganolg"))

	fmt.Printf("хэш от Иван: %d\n", hashstr("Иаав"))
	fmt.Printf("хэш от Ивна: %d\n", hashstr("Иваа"))
}
