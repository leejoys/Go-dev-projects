package main

import (
	"bufio"
	"fmt"
	"os"
)

// RuneRing - кольцевой массив
type RuneRing struct {
	size  int
	data  []*rune
	Start int
	End   int
}

// NewRuneRing - конструктор кольца
func NewRuneRing(size int, start int) (*RuneRing, error) {
	if start >= size {
		return nil, fmt.Errorf("Стартовая позиция не соответствует размеру кольцевого массива")
	}
	return &RuneRing{size, make([]*rune, size, size), start, start}, nil
}

// Size - получение размера кольца
func (r RuneRing) Size() int {
	return r.size
}

// IsEmpty - пусто ли кольцо
func (r *RuneRing) IsEmpty() bool {
	return r.Start == r.End
}

// IsFull - достигнут ди конец
func (r *RuneRing) IsFull() bool {
	return (r.End < r.Start && r.Start-r.End == 1) || (r.Start == 0 && r.End == r.size-1)
}

// Read - чтение элемента
func (r *RuneRing) Read() (rune, error) {
	if !r.IsEmpty() {
		el := r.data[r.Start]
		r.Start++
		for el == nil && r.Start < r.End {
			el = r.data[r.Start]
			r.Start++
		}
		if el == nil {
			return 0, fmt.Errorf("Нет новых данных в буфере")
		}
		return *el, nil
	}
	return 0, fmt.Errorf("Нет новых данных в буфере")
}

// Write - добавление одного элемента
func (r *RuneRing) Write(v rune) error {
	if r.IsEmpty() || !r.IsFull() {
		r.data[r.End] = &v
		r.End++
		return nil
	}
	return fmt.Errorf("Буфер переполнен")
}

// // Remove - удаление элемента
// func (r *RuneRing) Remove(index int) error {
// 	if index < 0 || index > r.size {
// 		return fmt.Errorf("Неверный индекс списка")
// 	}
// 	r.data[index] = nil
// 	return nil
// }

// // Print - печать содержимого кольца
// func (r RuneRing) Print() {
// 	if r.Start < r.End {
// 		for _, el := range r.data[r.Start:r.End] {
// 			fmt.Printf("%d\t", *el)
// 		}
// 		fmt.Printf("\n")
// 	} else if r.Start > r.End {
// 		tempData := append(r.data[r.Start:], r.data[:r.End]...)
// 		for _, el := range tempData {
// 			fmt.Printf("%d\t", *el)
// 		}
// 		fmt.Printf("\n")
// 	} else {
// 		fmt.Println("Список пуст!")
// 	}
// }

func main() {

	fmt.Println("Программа инновационного шифрования")
	fmt.Println("Введите строку не более 126 символов:")

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	inString := in.Text()

	var outString string
	runeRing, err := NewRuneRing(127, 0)
	if err != nil {
		fmt.Println(err)
	}

	for _, char := range inString {
		err := runeRing.Write(char)

		if err != nil {
			fmt.Println(err)
			break
		}
	}

	for range inString {
		char, err := runeRing.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		outString = outString + string(char)
	}

	fmt.Printf("Результат: %s\n", outString)
}
