package main

import (
	"fmt"
)

// ErrWrongTreeKey -  ошибка при передаче несуществующего в дереве значения
var ErrWrongTreeKey = fmt.Errorf("Неверное значение узла")

// IntNode - описание типа Узел дерева
type IntNode struct {
	Key    int
	Parent *IntNode
	Left   *IntNode
	Right  *IntNode
}

// NewIntNode - создание нового узла дерева
func NewIntNode(key int) *IntNode {
	return &IntNode{key, nil, nil, nil}
}

// IntTree - описание типа Дерево целых чисел
type IntTree struct {
	height int
	Root   *IntNode
}

// NewIntTree - создание нового дерева целых чисел
func NewIntTree() *IntTree {
	return &IntTree{0, nil}
}

// Height - получение размера дерева
func (t IntTree) Height() int {
	return t.height
}

//Print - печать всего дерева с корнем root
func (t IntTree) Print(root *IntNode) {
	if root != nil {
		t.Print(root.Left)
		fmt.Printf("%d ", root.Key)
		t.Print(root.Right)
	}
}

// Get - поиск указанного элемента дерева
func (t IntTree) Get(key int) *IntNode {
	node := t.Root
	for node != nil && key != node.Key {
		if key > node.Key {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return node
}

//Min - поиск минимального элемента дерева с корнем root
func (t IntTree) Min(root *IntNode) *IntNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

//Max - поиск максимального элемента дерева с корнем root
func (t IntTree) Max(root *IntNode) *IntNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

//Successor - поиск последующего для root элемента
func (t IntTree) Successor(node *IntNode) *IntNode {
	if node.Right != nil {
		return t.Min(node.Right)
	}
	other := node.Parent
	for other != nil && node == other.Right {
		node = other
		other = other.Parent
	}
	return other
}

//Predecessor - поиск предыдущего для root элемента
func (t IntTree) Predecessor(node *IntNode) *IntNode {
	if node.Left != nil {
		return t.Max(node.Left)
	}
	other := node.Parent
	for other != nil && node == other.Left {
		node = other
		other = other.Parent
	}
	return other
}

// Insert - вставка нового элемента в дерево, если он уникальный
func (t *IntTree) Insert(key int) {
	var other *IntNode
	node := t.Root
	newNode := NewIntNode(key)

	for node != nil {
		other = node
		if newNode.Key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	newNode.Parent = other
	if other == nil {
		t.Root = newNode
	} else if newNode.Key < other.Key {
		other.Left = newNode
	} else {
		other.Right = newNode
	}
}

//Transplant - процедура перемещения поддеревьев
func (t *IntTree) Transplant(node, other *IntNode) {
	if node.Parent == nil {
		t.Root = other
	} else if node == node.Parent.Left {
		node.Parent.Left = other
	} else {
		node.Parent.Right = other
	}
	if other != nil {
		other.Parent = node.Parent
	}
}

// Delete - удаление элемента из дерева
func (t *IntTree) Delete(key int) {
	node := t.Get(key)
	var other *IntNode
	if node.Left == nil {
		t.Transplant(node, node.Right)
	} else if node.Right == nil {
		t.Transplant(node, node.Left)
	} else {
		other = t.Min(node.Right)
		if other.Parent != node {
			t.Transplant(other, other.Right)
			other.Right = node.Right
			other.Right.Parent = other
		}
		t.Transplant(node, other)
		other.Left = node.Left
		other.Left.Parent = other
	}
}

// someTree := NewIntTree()
// someTree.Print(nil)
// someTree.Insert(1)
// someTree.Insert(9)
// someTree.Insert(2)
// someTree.Insert(6)
// someTree.Insert(4)
// someTree.Insert(3)
// someTree.Insert(5)
// someTree.Insert(8)
// someTree.Insert(7)
// someTree.Print(someTree.Root)
// println()
// someTree.Print(someTree.Get(1))
// println()
// someTree.Print(someTree.Get(2))
// println()
// someTree.Print(someTree.Get(3))
// println()
// someTree.Print(someTree.Get(4))
// println()
// someTree.Print(someTree.Get(5))
// println()
// someTree.Print(someTree.Get(6))
// println()
// someTree.Print(someTree.Get(7))
// println()
// someTree.Print(someTree.Get(8))
// println()
// someTree.Print(someTree.Get(9))
// println()
// someTree.Delete(9)
// someTree.Print(someTree.Root)
// println()
// someTree.Delete(2)
// someTree.Print(someTree.Root)
// println()
// someTree.Delete(1)
// someTree.Print(someTree.Root)
// println()
