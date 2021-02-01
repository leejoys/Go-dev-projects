package main

import (
	"fmt"
)

// ErrWrongListIndex - ошибка при чтении элемента списка
var ErrWrongListIndex = fmt.Errorf("Неверный индекс элемента")

// IntQNode - описание типа Узел списка
type IntQNode struct {
	Key  string
	Next *IntQNode
}

// NewIntQNode - создание нового узла списка
func NewIntQNode(key string) *IntQNode {
	return &IntQNode{key, nil}
}

// IntQueue - описание типа очередь
type IntQueue struct {
	size int
	Head *IntQNode
}

// NewIntQueue - создание новой очереди
func NewIntQueue() *IntQueue {
	return &IntQueue{0, nil}
}

// Size - получение размера очереди
func (l IntQueue) Size() int {
	return l.size
}

// Get - получение произвольного элемента очереди
func (l IntQueue) Get(index int) (*IntQNode, error) {
	if index < 0 || index >= l.Size() {
		return nil, ErrWrongListIndex
	}
	node := l.Head
	for i := 1; i <= index; i++ {
		node = node.Next
	}
	return node, nil
}

// Add - добавление нового элемента в начало очереди
func (l *IntQueue) Add(el string) {
	newNode := NewIntQNode(el)
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.size++
}

// //Remove - удаление элемента из проивольной позиции
// func (l *IntQueue) Remove(index int) error {
// 	if index < 0 || index >= l.size {
// 		return ErrWrongListIndex
// 	}

// 	if index == 0 {
// 		node := l.Head
// 		l.Head = node.Next
// 		l.size--
// 		return nil
// 	}

// 	node, err := l.Get(index - 1)
// 	if err != nil {
// 		return err
// 	}

// 	node.Next = node.Next.Next
// 	l.size--
// 	return nil
// }

// // QueueHead - добавление элемента в очередь с головы
// func (l *IntQueue) QueueHead(el string) {
// 	l.Add(el)
// 	return
// }

// // DequeueEnd - извлечение элемента из очереди c хвоста
// func (l *IntQueue) DequeueEnd() (*IntQNode, error) {
// 	if l.size == 0 {
// 		return nil, ErrQueueEmpty
// 	}

// 	node, err := l.Get(l.size - 1)
// 	if err != nil {
// 		return nil, err
// 	}
// 	l.Remove(l.size - 1)
// 	return node, nil
// }

// Set - обновление произвольного элемента списка
// func (l *IntQueue) Set(el string, index int) error {
// 	if index < 0 || index >= l.Size() {
// 		return ErrWrongListIndex
// 	}
// 	node, err := l.Get(index)
// 	if err != nil {
// 		return err
// 	}
// 	node.Key = el
// 	return nil
// }

// Insert - вставка нового элемента в произвольную позицию
// func (l *IntQueue) Insert(el string, index int) error {
// 	if index < 0 || index >= l.size {
// 		return ErrWrongListIndex
// 	}
// 	newNode := NewIntQNode(el)
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
// func (l IntQueue) Print() {
// 	node := l.Head
// 	if node != nil {
// 		for node != nil {
// 			fmt.Printf("%s\t", string(node.Key))
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
func (l *IntQueue) Queue(el string) error {

	if l.size == 0 {
		l.Add(el)
		return nil
	}

	newNode := NewIntQNode(el)
	node, err := l.Get(l.size - 1)
	if err != nil {
		return err
	}

	node.Next = newNode
	l.size++
	return nil
}

// Dequeue - извлечение элемента из очереди
func (l *IntQueue) Dequeue() (*IntQNode, error) {
	if l.size == 0 {
		return nil, ErrQueueEmpty
	}

	node := l.Head

	l.Head = node.Next
	l.size--

	return node, nil

}
