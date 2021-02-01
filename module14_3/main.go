package main

import (
	"fmt"
	"math"
)

type hashmap struct {
	data [1000]string
}

//NewHashmap create hashmap{[1000]string{}}
func NewHashmap() hashmap {
	return hashmap{[1000]string{}}
}

func (h *hashmap) hashstr(val string) (valIntUint64 uint64) {
	valBytes := []byte(val)
	for i := 0; i < len(valBytes)-1; i++ {
		valIntUint64 = valIntUint64 + uint64(math.Pow(float64(valBytes[i])/10, float64(valBytes[i+1])/10))
	}
	return valIntUint64 % 1000
}

func (h *hashmap) Set(key, val string) {
	i := h.hashstr(key)
	h.data[i] = val
}

func (h *hashmap) Get(key string) (value string, ok bool) {
	i := h.hashstr(key)
	if h.data[i] == "" {
		return h.data[i], false
	}
	return h.data[i], true
}

func (h *hashmap) Delete(key string) {
	i := h.hashstr(key)
	h.data[i] = ""
}

func main() {
	someMap := NewHashmap()

	someMap.Set("white", "white")
	someMap.Set("black", "black")
	someMap.Set("red", "red")
	someMap.Set("brown", "brown")
	someMap.Set("pink", "pink")
	someMap.Set("yellow", "yellow")
	someMap.Set("green", "green")
	someMap.Set("blue", "blue")

	el, ok := someMap.Get("white")
	if ok {
		fmt.Printf("элемент white: %s\n", el)
	}

	el, ok = someMap.Get("black")
	if ok {
		fmt.Printf("элемент black: %s\n", el)
	}

	el, ok = someMap.Get("red")
	if ok {
		fmt.Printf("элемент red: %s\n", el)
	}

	el, ok = someMap.Get("brown")
	if ok {
		fmt.Printf("элемент brown: %s\n", el)
	}

	el, ok = someMap.Get("pink")
	if ok {
		fmt.Printf("элемент pink: %s\n", el)
	}

	el, ok = someMap.Get("yellow")
	if ok {
		fmt.Printf("элемент yellow: %s\n", el)
	}

	el, ok = someMap.Get("green")
	if ok {
		fmt.Printf("элемент green: %s\n", el)
	}

	el, ok = someMap.Get("blue")
	if ok {
		fmt.Printf("элемент blue: %s\n", el)
	}

	someMap.Delete("yellow")
	someMap.Delete("brown")

	el, ok = someMap.Get("yellow")
	if ok {
		fmt.Printf("элемент yellow: %s\n", el)
	}

	el, ok = someMap.Get("brown")
	if ok {
		fmt.Printf("элемент brown: %s\n", el)
	}
}
