package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ErrWrongListIndex -  ошибка при передачи несуществующего в списке индекса
var ErrWrongListIndex = fmt.Errorf("Неверный индекс списка")

// StringNode - описание типа Узел списка
type StringNode struct {
	Value string
	Next  *StringNode
}

// NewStringNode - создание нового узла списка
func NewStringNode(value string) *StringNode {
	return &StringNode{value, nil}
}

// StringList - описание типа Список рун
type StringList struct {
	size int
	Head *StringNode
}

// NewStringList - создание нового списка рун
func NewStringList() *StringList {
	return &StringList{0, nil}
}

// Size - получение размера списка
func (l StringList) Size() int {
	return l.size
}

// Get - получение произвольного элемента списка
func (l StringList) Get(index int) (*StringNode, error) {
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
func (l *StringList) Add(el string) {
	newNode := NewStringNode(el)
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.size++
}

//Remove - удаление элемента из проивольной позиции
func (l *StringList) Remove(index int) error {
	if index < 0 || index >= l.size {
		return ErrWrongListIndex
	}

	if index == 0 {
		node := l.Head
		l.Head = node.Next
		l.size--
		return nil
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
// func (l *StringList) Set(el string, index int) error {
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
// func (l *StringList) Insert(el string, index int) error {
// 	if index < 0 || index >= l.size {
// 		return ErrWrongListIndex
// 	}
// 	newNode := NewStringNode(el)
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
// func (l StringList) Print() {
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
func (l *StringList) Queue(el string) error {

	if l.size == 0 {
		l.Add(el)
		return nil
	}

	newNode := NewStringNode(el)
	node, err := l.Get(l.size - 1)
	if err != nil {
		return err
	}

	node.Next = newNode
	l.size++
	return nil
}

// QueueHead - добавление элемента в очередь с головы
func (l *StringList) QueueHead(el string) {
	l.Add(el)
	return
}

// Dequeue - извлечение элемента из очереди
func (l *StringList) Dequeue() (*StringNode, error) {
	if l.size == 0 {
		return nil, ErrQueueEmpty
	}

	node := l.Head

	l.Head = node.Next
	l.size--

	return node, nil

}

// DequeueEnd - извлечение элемента из очереди c хвоста
func (l *StringList) DequeueEnd() (*StringNode, error) {
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

	fmt.Println("Программа обработки заказов")
	fmt.Println("Введите номера заказов через пробел и нажмите ВВОД:")

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	inString := strings.Fields(in.Text())

	var outString string

	queue := NewStringList()
	for _, char := range inString {
		err := queue.Queue(char)

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("Введите порядковый номер заказа в очереди для отмены: ")
	var cancelOrder int
	fmt.Scanf("%d", &cancelOrder)

	queue.Remove(cancelOrder - 1)

	size := queue.Size()
	for i := 0; i < size; i++ {
		char, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
		}
		outString = outString + " " + string(char.Value)
	}

	fmt.Printf("Номера заказов для обработки: %s\n", outString)

}
