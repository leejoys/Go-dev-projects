package main

import (
	"fmt"
)

var IndexOutOfBoundsError = fmt.Errorf("Ошибка доступа за границы структуры")

type linketList struct {
}

func (linketList List) find(index int) (Node, error) {
	if index < 0 || index >= linketList.size {
		return nil, IndexOutOfBoundsError
	}
	node := linketList.head
	for i := 1; i < index; i++ {
		node = node.next
	}
	return node
}

func (linkedList List) size() int {
	var size int = 0
	if linkedList.head != nil {
		var node Node = linkedList.head.next
		for size = 1; node.next != nil; size++ {
			node = node.next
		}
	}
	return size
}

func main() {

}
