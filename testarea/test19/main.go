package main

import (
	"fmt"
	"sync"
)

func main() {
	// Небольшая обертка над множеством чисел, которая направляет их
	// в канал
	// и возвращает его наружному коду
	init := func(done <-chan int, wg *sync.WaitGroup, integers ...int) <-chan int {
		output := make(chan int)
		go func() {
			defer close(output)
			for _, i := range integers {
				select {
				case <-done:
					wg.Done()
					return
				case output <- i:
				}
			}
		}()
		return output
	}
	// Потокобезопасная версия стадии конвейера, осуществляющая
	// выполнения арифметического произведения данных на заданное
	// число
	multiply := func(done <-chan int, input <-chan int, multiplier int, wg *sync.WaitGroup) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range input {
				select {
				case <-done:
					wg.Done()
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}
	// Потокобезопасная версия стадии конвейера, осуществляющая
	// выполнения арифметического сложения данных на заданное число
	add := func(done <-chan int, input <-chan int, additive int, wg *sync.WaitGroup) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range input {

				select {
				case <-done:
					wg.Done()
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}
	// Этот канал используется для централизованной остановки
	// конвейера
	// Так как ни одна стадия не должна знать, когда конвейер будет
	// остановлен,
	// это не задача стадии, как вы понимаете,
	// Это лучше делать вот так - централизованно
	done := make(chan int)
	//defer close(done)
	var wg sync.WaitGroup
	wg.Add(4)
	intStream := init(done, &wg, 1, 2, 3, 4)
	// Одно из преимуществ использования каналов видно здесь -
	// каналы итерируемы, благодаря этому мы можем комбинировать
	// пайплайн как хотим,
	// а сама целостность пайплайна от этого не меняется

	pipeline := multiply(done, add(done, multiply(done, intStream, 2, &wg), 1, &wg), 2, &wg)
	for v := range pipeline {
		fmt.Println(v)
	}

	close(done)

	wg.Wait()
}
