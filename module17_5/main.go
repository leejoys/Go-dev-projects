package main

//Напишите программу, которая делает следующее: одна горутина по порядку отсылает числа от 1 до 100 в канал,
//вторая горутина их принимает в правильном порядке и печатает на экран (в консоль).

var intChan chan int

func sender(c chan<- int) {

}

func receiver(c <-chan int) {

}

func main() {
	intChan := make(chan int, 0)

	go sender(intChan)
	go receiver(intChan)
}
