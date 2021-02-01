package main

import (
	"bufio"
	"fmt"
	"os"
)

// ErrEmptyStack Ошибка при извлечении данных из пустого стека
var ErrEmptyStack error = fmt.Errorf("%s", "Стек пуст")
var inString string
var outString string

// RuneStack - структура стека
type RuneStack struct {
	data []rune
}

// NewRuneStack - конструктор стека
func NewRuneStack() RuneStack {
	return RuneStack{[]rune{}}
}

// Size - метод возвращает размер стека
func (stack RuneStack) Size() int {
	return len(stack.data)
}

// Push - метод добавляет элемент в стек
func (stack *RuneStack) Push(el rune) {
	stack.data = append(stack.data, el)
}

// Pop - метод возвращает элемент из стека или ошибку, если он пуст
func (stack *RuneStack) Pop() (rune, error) {
	size := stack.Size()
	if size == 0 {
		return 0, ErrEmptyStack
	}
	el := stack.data[size-1]
	stack.data = stack.data[:size-1]
	return el, nil
}

func main() {

	fmt.Println("Программа реверса строк")
	fmt.Println("Введите строку и нажмите ВВОД:")

	in := bufio.NewReader(os.Stdin)
	inString, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("%q", inString))

	stack := NewRuneStack()
	for _, char := range inString {
		stack.Push(char)
	}
	for range inString {
		char, _ := stack.Pop()
		outString = outString + string(char)
	}

	fmt.Printf("Результат: %s\n", outString)
	fmt.Println("нажмите ВВОД для выхода")
	fmt.Scanln(&inString)
}
