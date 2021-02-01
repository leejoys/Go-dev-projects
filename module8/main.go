package main

import (
	"fmt"
	"module8/electronic"
)

func main() {
	apple := electronic.ConsrustorApple("iPhone X", "iOS 11")
	android := electronic.ConsrustorAndroid("Meizu", "X8", "Android 8")
	radio := electronic.ConsrustorRadio("General Electric", "26928GE2", 20)

	printCharacteristics(apple)
	printCharacteristics(android)
	printCharacteristics(radio)
}

func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Бренд: %s Модель: %s Тип: %s", p.Brand(), p.Model(), p.Type())

	switch p.(type) {
	case electronic.StationPhone:
		fmt.Printf(" Кнопок %d\n", p.(electronic.StationPhone).ButtonsCount())
	case electronic.Smartphone:
		fmt.Printf(" Операционная система %s\n", p.(electronic.Smartphone).OS())

	}
}
