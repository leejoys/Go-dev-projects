package main

import (
	"bufio"
	"fmt"
	"os"
)

// ErrWrongListIndex -  ошибка при передачи несуществующего в списке индекса
var ErrWrongListIndex = fmt.Errorf("Неверный индекс списка")

// RuneNode - описание типа Узел списка
type RuneNode struct {
	Value rune
	Next  *RuneNode
}

// NewRuneNode - создание нового узла списка
func NewRuneNode(value rune) *RuneNode {
	return &RuneNode{value, nil}
}

// RuneList - описание типа Список рун
type RuneList struct {
	size int
	Head *RuneNode
}

// NewRuneList - создание нового списка рун
func NewRuneList() *RuneList {
	return &RuneList{0, nil}
}

// Size - получение размера списка
func (l RuneList) Size() int {
	return l.size
}

// Get - получение произвольного элемента списка
func (l RuneList) Get(index int) (*RuneNode, error) {
	if index < 0 || index >= l.Size() {
		return nil, ErrWrongListIndex
	}
	node := l.Head
	for i := 1; i <= index; i++ {
		node = node.Next
	}
	return node, nil
}

// Add - добавление нового элемента в начало списка
func (l *RuneList) Add(el rune) {
	newNode := NewRuneNode(el)
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.size++
}

//Remove - удаление элемента из проивольной позиции
func (l *RuneList) Remove(index int) error {
	if index < 0 || index >= l.size {
		return ErrWrongListIndex
	}
	node, err := l.Get(index - 1)
	if err != nil {
		return err
	}
	node.Next = node.Next.Next
	l.size--
	return nil
}

// Set - обновление произвольного элемента списка
// func (l *RuneList) Set(el rune, index int) error {
// 	if index < 0 || index >= l.Size() {
// 		return ErrWrongListIndex
// 	}
// 	node, err := l.Get(index)
// 	if err != nil {
// 		return err
// 	}
// 	node.Value = el
// 	return nil
// }

// Insert - вставка нового элемента в произвольную позицию
// func (l *RuneList) Insert(el rune, index int) error {
// 	if index < 0 || index >= l.size {
// 		return ErrWrongListIndex
// 	}
// 	newNode := NewRuneNode(el)
// 	if index == 0 {
// 		l.Add(el)
// 		return nil
// 	}
// 	node, err := l.Get(index - 1)
// 	if err != nil {
// 		return err
// 	}
// 	newNode.Next = node.Next
// 	node.Next = newNode
// 	l.size++
// 	return nil
// }

// Print - печать списка
// func (l RuneList) Print() {
// 	node := l.Head
// 	if node != nil {
// 		for node != nil {
// 			fmt.Printf("%s\t", string(node.Value))
// 			node = node.Next
// 		}
// 		fmt.Printf("\n")
// 		return
// 	}
// 	fmt.Println("Список пуст!")
// }

// ErrQueueEmpty - ошибка при чтении из пустой очереди
var ErrQueueEmpty = fmt.Errorf("Очередь пуста")

// Queue - добавление элемента в очередь
func (l *RuneList) Queue(el rune) error {

	if l.size == 0 {
		l.Add(el)
		return nil
	}

	newNode := NewRuneNode(el)
	node, err := l.Get(l.size - 1)
	if err != nil {
		return err
	}

	node.Next = newNode
	l.size++
	return nil
}

// QueueHead - добавление элемента в очередь с головы
func (l *RuneList) QueueHead(el rune) {
	l.Add(el)
	return
}

// Dequeue - извлечение элемента из очереди
func (l *RuneList) Dequeue() (*RuneNode, error) {
	if l.size == 0 {
		return nil, ErrQueueEmpty
	}

	node := l.Head

	l.Head = node.Next
	l.size--

	return node, nil

}

// DequeueEnd - извлечение элемента из очереди c хвоста
func (l *RuneList) DequeueEnd() (*RuneNode, error) {
	if l.size == 0 {
		return nil, ErrQueueEmpty
	}

	node, err := l.Get(l.size - 1)
	if err != nil {
		return nil, err
	}
	l.Remove(l.size - 1)
	return node, nil
}

func main() {

	fmt.Println("Программа буфера ввода/вывода")
	fmt.Println("Введите строку и нажмите ВВОД:")

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	inString := in.Text()

	var outString string

	queue := NewRuneList()
	for _, char := range inString {
		err := queue.Queue(char)

		if err != nil {
			fmt.Println(err)
		}
	}

	for range inString {
		char, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
		}
		outString = outString + string(char.Value)
	}

	fmt.Printf("Результат: %s\n", outString)

	for _, char := range inString {
		queue.QueueHead(char)
	}

	outString = ""
	for range inString {
		char, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
		}
		outString = outString + string(char.Value)
	}

	fmt.Printf("Наоборот: %s\n", outString)

	for _, char := range inString {
		queue.QueueHead(char)
	}

	outString = ""
	for range inString {
		char, err := queue.DequeueEnd()
		if err != nil {
			fmt.Println(err)
		}
		outString = outString + string(char.Value)
	}

	fmt.Printf("В обратную сторону: %s\n", outString)

}
