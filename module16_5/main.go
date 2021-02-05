package main

import (
	"fmt"
	"sync"
)

// ErrNotEnough - error for withdrawal method
var ErrNotEnough = fmt.Errorf("Not enough money to withdraw")

// ErrWrongCmd - error for wrong command check
var ErrWrongCmd = fmt.Errorf("Unsupported command. You can use commands: balance, deposit, withdrawal, quit.")

// BankClient interface
type BankClient interface {
	Deposit(amount int)
	Withdrawal(amount int) error
	Balance() int
}

//Account struct
type Account struct {
	balance int
	*sync.RWMutex
}

//NewAccount create account
func NewAccount() *Account {
	return &Account{0, &sync.RWMutex{}}
}

// Deposit deposits given amount to clients account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdrawal withdraws given amount from clients account.
// return error if clients balance less the withdrawal amount
func (a *Account) Withdrawal(amount int) error {
	if a.balance < amount {
		return ErrNotEnough
	}
	a.balance -= amount
	return nil
}

// Balance returns clients balance
func (a *Account) Balance() int {
	return a.balance
}

func scanCommand(a *Account, wg *sync.WaitGroup) {

	for {
		var cmd string
		var ammount int = 0
		fmt.Println("Enter command:")
		fmt.Scanln(&cmd)

		switch cmd {
		case "balance":
			fmt.Printf("Balance: %d", a.Balance())

		case "deposit":
			fmt.Println("Enter ammount:")
			fmt.Scanln(&ammount)
			a.Deposit(ammount)

		case "withdrawal":
			fmt.Println("Enter ammount:")
			fmt.Scanln(&ammount)
			a.Withdrawal(ammount)

		case "quit":
			wg.Done()

		default:
			fmt.Println(ErrWrongCmd)
		}
	}
}

//запускает 10 горутин, каждая из которых с промежутком
//от 0.5 секунд до 1 секунды зачисляет на счёт клиента случайную сумму от 1 до 10.
func debet() {

}

//запускается 5 горутин, которые с промежутком 0.5 секунд до 1 секунды снимают с клиента случайную сумму от 1 до 5.
//Если снятие невозможно, в консоль выводится сообщение об ошибке, и приложение продолжает работу.
func credit() {

}

func main() {
	a := NewAccount()
	var wg sync.WaitGroup
	wg.Add(1)

	go scanCommand(a, wg)
	go debet()
	go credit()

	wg.Wait()
}
