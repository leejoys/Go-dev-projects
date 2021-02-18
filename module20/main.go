package main

import (
	"fmt"
	"sync"
)

//Напишите код, реализующий пайплайн, работающий с целыми числами и состоящий из следующих стадий:

//[]
//Стадия фильтрации отрицательных чисел (не пропускать отрицательные числа).
//[]
//Стадия фильтрации чисел, не кратных 3 (не пропускать такие числа), исключая также и 0 [].
//[]
//Стадия буферизации данных в кольцевом буфере с интерфейсом, соответствующим тому, который был дан в
//качестве задания в 19 модуле. [] В этой стадии предусмотреть опустошение буфера (и соответственно, передачу
//этих данных, если они есть, дальше) с определённым интервалом во времени. [] Значения размера буфера и этого
//интервала времени сделать настраиваемыми (как мы делали: через константы или глобальные переменные).
//[]
//Написать источник данных для конвейера. Непосредственным источником данных должна быть консоль.
//[]
//Также написать код потребителя данных конвейера. Данные от конвейера можно направить снова в консоль
//построчно, сопроводив их каким-нибудь поясняющим текстом, например: «Получены данные …».
//[]
//При написании источника данных подумайте о фильтрации нечисловых данных, которые можно ввести через
//консоль. Как и где их фильтровать, решайте сами.
//=========================================================================================================

//[+]
//Positive Стадия фильтрации отрицательных чисел (не пропускать отрицательные числа).
func Positive(c <-chan int) <-chan int {
	ch := make(chan int)
	if el := <-c; el >= 0 {
		ch <- el
	}
	return ch
}

//[+]
//Trine Стадия фильтрации чисел, не кратных 3 (не пропускать такие числа), исключая также и 0 [].
func Trine(c <-chan int) <-chan int {
	ch := make(chan int)
	if el := <-c; el%3 != 0 || el == 0 {
		ch <- el
	}
	return ch
}

//[]
//Стадия буферизации данных в кольцевом буфере с интерфейсом, соответствующим тому, который был дан в
//качестве задания в 19 модуле. [] В этой стадии предусмотреть опустошение буфера (и соответственно, передачу
//этих данных, если они есть, дальше) с определённым интервалом во времени. [] Значения размера буфера и этого
//интервала времени сделать настраиваемыми (как мы делали: через константы или глобальные переменные).

//[+]
//Написать источник данных для конвейера. Непосредственным источником данных должна быть консоль.

//ring bufer delay
var ringDelay int

//ring bufer size
var ringSize int

// ErrNoData - error for process metod
var ErrNoData = fmt.Errorf("There is no data to process")

// ErrNotInt - error for int check
var ErrNotInt = fmt.Errorf("Input value not int")

// ErrWrongCmd - error for wrong command check
var ErrWrongCmd = fmt.Errorf("Unsupported command. You can use commands: input, process, quit.")

//DataSrc struct
type DataSrc struct {
	data []int
	C    chan int
	*sync.WaitGroup
}

//NewDataSrc create DataSrc
func NewDataSrc() *DataSrc {
	return &DataSrc{make([]int, 0), make(chan int), &sync.WaitGroup{}}
}

// intInput input ints to DataSrc
func (d DataSrc) intInput(i int) {

	d.data = append(d.data, i)
}

// Process send data from DataSrc to process.
// return error if DataSrc is empty
func (d DataSrc) Process() error {

	if len(d.data) == 0 {
		return ErrNoData
	}
	for i := range d.data {
		d.C <- i
	}
	ringSize = len(d.data)
	d.data = make([]int, 0)
	return nil
}

//ScanCmd scan cmd
func (d DataSrc) ScanCmd() {

	for {
		var cmd string
		fmt.Println("Enter command:")
		fmt.Scanln(&cmd)

		switch cmd {

		//[-]
		//При написании источника данных подумайте о фильтрации нечисловых данных,
		//которые можно ввести через консоль. Как и где их фильтровать, решайте сами.
		case "input":
			var anything interface{}
			fmt.Println("Enter int number:")
			fmt.Scanln(&anything)
			num, ok := anything.(int)
			if !ok {
				fmt.Println(ErrNotInt)
				break
			}
			d.intInput(num)

		case "process":
			err := d.Process()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Processing...")

		case "quit":
			d.Done()
			return

		default:
			fmt.Println(ErrWrongCmd)
		}
	}
}

//[+]
//Также написать код потребителя данных конвейера. Данные от конвейера можно направить снова в консоль
//построчно, сопроводив их каким-нибудь поясняющим текстом, например: «Получены данные …».

func receiver(c <-chan int) {
	for {
		fmt.Printf("Int received: %d", <-c)
	}
}

func main() {
	d := NewDataSrc()
	d.Add(1)
	go d.ScanCmd()
	go receiver(d.C)
	d.Wait()
}
