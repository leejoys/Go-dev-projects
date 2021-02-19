package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
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

//Positive Стадия фильтрации отрицательных чисел (не пропускать отрицательные числа).
//[+]
func (d DataSrc) Positive() {
	for {
		el := <-d.C1
		if el >= 0 {
			d.C2 <- el
		}
	}
}

//Trine Стадия фильтрации чисел, не кратных 3 (не пропускать такие числа), исключая также и 0 [+].
//[+]
func (d DataSrc) Trine() {
	for {
		if el := <-d.C2; el%3 == 0 && el != 0 {
			d.C3 <- el
		}
	}
}

//[]
//Стадия буферизации данных в кольцевом буфере с интерфейсом, соответствующим тому, который был дан в
//качестве задания в 19 модуле. [] В этой стадии предусмотреть опустошение буфера (и соответственно, передачу
//этих данных, если они есть, дальше) с определённым интервалом во времени. [] Значения размера буфера и этого
//интервала времени сделать настраиваемыми (как мы делали: через константы или глобальные переменные).

//ring bufer delay
var ringDelay int

//RingBuf make rBuf(size), work and exit
func RingBuf(ringSize int) {
	r := newrBuf(ringSize)

}

// rBuf data structure
type rBuf struct {
	dataRing *ring.Ring
}

//newrBuf make rBuf
func newrBuf(size int) *rBuf {

	return &rBuf{
		ring.New(size),
	}
}

//Pull() — добавление нового элемента в буфер, в случае
//добавления в заполненный буфер стирается самый «старый» элемент
func (r *rBuf) Pull(el int) {
	r.dataRing.Value = el
	r.dataRing = r.dataRing.Next()
}

//получение всех упорядоченных элементов буфера
//и последующая очистка буфера
func (r *rBuf) Get() []int {
	resultArr := make([]int, 0)
	size := r.dataRing.Len()

	for i := 0; i < size; i++ {
		if r.dataRing.Value == nil {
			r.dataRing = r.dataRing.Next()
			continue
		}
		resultArr = append(resultArr, r.dataRing.Value.(int))
		r.dataRing.Value = nil
		r.dataRing = r.dataRing.Next()

	}
	return resultArr
}

//[+]
//Написать источник данных для конвейера. Непосредственным источником данных должна быть консоль.

// ErrNoInts - error for data check
var ErrNoInts = fmt.Errorf("There is no ints to process")

// ErrWrongCmd - error for wrong command check
var ErrWrongCmd = fmt.Errorf("Unsupported command. You can use commands: input, quit.")

//DataSrc struct
type DataSrc struct {
	data *[]int
	C1   chan int
	C2   chan int
	C3   chan int
	C4   chan int
	*sync.WaitGroup
}

//NewDataSrc create DataSrc
func NewDataSrc() *DataSrc {
	d := make([]int, 0)
	return &DataSrc{&d,
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		&sync.WaitGroup{}}
}

// intInput input ints to DataSrc
func (d DataSrc) intInput(i int) {
	*d.data = append(*d.data, i)
}

// Process send data from DataSrc to process.
// return error if DataSrc is empty
func (d DataSrc) Process() error {

	if len(*d.data) == 0 {
		return ErrNoInts
	}
	fmt.Println("Processing...")
	for _, i := range *d.data {
		d.C1 <- i
	}
	ringSize = len(*d.data)
	*d.data = make([]int, 0)
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

			fmt.Println("Enter numbers separated by whitespaces:")
			in := bufio.NewScanner(os.Stdin)
			in.Scan()
			anything := in.Text()

			someSlice := strings.Fields(anything)

			for _, some := range someSlice {

				if num, err := strconv.ParseInt(some, 10, 0); err == nil {

					d.intInput(int(num))
				}
			}

			err := d.Process()
			if err != nil {
				fmt.Println(err)
				break
			}

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
		fmt.Printf("Int received: %d\n", <-c)
	}
}

func main() {
	d := NewDataSrc()
	d.Add(1)
	go d.ScanCmd()
	go d.Positive()
	go d.Trine()
	go receiver(d.C3)
	d.Wait()
}
