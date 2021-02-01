package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Создаем кольцо размером 5 элементов
	ring1 := ring.New(5)
	// ring2 := ring.New(3)

	// Получаем длину кольца
	ring1Len := ring1.Len()
	// ring2Len := ring2.Len()

	// Заполняем первое кольцо индексами
	for i := 0; i < ring1Len; i++ {
		ring1.Value = i
		ring1 = ring1.Next()
	}
	//создаем второе кольцо ссылкой на первое
	ring2 := ring1.Next()

	//проверяем содержимое
	for i := 0; i < ring1Len; i++ {
		fmt.Print(ring2.Value)
		fmt.Printf("\t")
		fmt.Println(ring1.Value)
		ring1 = ring1.Next()
		ring2 = ring2.Next()
	}

	// Заполняем второе кольцо единичками
	// for j := 0; j < ring2Len; j++ {
	// 	ring2.Value = 1
	// 	ring2 = ring2.Next()
	// }

	// Соединяем второе кольцо с первым вот таким образом
	ring1.Next().Link(ring2)
	// Так как мы соединили два кольца -
	// размер нового кольца равен сумме размеров исходных
	unionLen := ring1Len * 2

	// Заполняем первые 4 элемента пятерками
	for i := 0; i < 4; i++ {
		ring1.Value = 5
		ring1 = ring1.Next()
	}
	// Проверяем содержимое нового кольца
	for i := 0; i < unionLen; i++ {
		fmt.Println(ring2.Value)
		ring1 = ring1.Next()
	}

	// for i := 0; i < unionLen; i++ {
	// 	ring2.Value = i
	// 	fmt.Println(ring1.Value)
	// 	ring1 = ring1.Next()
	// 	ring2 = ring2.Next()
	// }

	// for i := 0; i < unionLen; i++ {
	// 	ring2.Value = i
	// 	fmt.Println(ring1.Value)
	// 	ring1 = ring1.Next()
	// 	ring2 = ring2.Next()
	// }
}
