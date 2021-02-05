package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//interface implementation test
var _ Account = BankClient{}

// ErrNotEnough - error for withdrawal method
var ErrNotEnough = fmt.Errorf("Not enough money to withdraw")

// ErrWrongCmd - error for wrong command check
var ErrWrongCmd = fmt.Errorf("Unsupported command. You can use commands: balance, deposit, withdrawal, quit.")

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
	a.Lock()
	a.balance += amount
	a.Unlock()
}

// Withdrawal withdraws given amount from clients account.
// return error if clients balance less the withdrawal amount
func (a *Account) Withdrawal(amount int) error {
	defer a.Unlock()
	a.Lock()
	if a.balance < amount {
		return ErrNotEnough
	}
	a.balance -= amount
	return nil
}

// Balance returns clients balance
func (a *Account) Balance() int {
	a.RLock()
	b := a.balance
	a.RUnlock()
	return b
}

func scanCommand(a *Account, wg *sync.WaitGroup) {

	for {
		var cmd string
		var ammount int = 0
		fmt.Println("Enter command:")
		fmt.Scanln(&cmd)

		switch cmd {
		case "balance":
			fmt.Printf("Balance: %d\n", a.Balance())

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

//с промежутком от 0.5 секунд до 1 секунды зачисляет на счёт клиента случайную сумму от 1 до 10
func credit(a *Account) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
		a.Deposit(rand.Intn(9) + 1)
	}
}

//с промежутком 0.5 секунд до 1 секунды снимаeт с клиента случайную сумму от 1 до 5
//Если снятие невозможно, в консоль выводится сообщение об ошибке, и приложение продолжает работу.
func debet(a *Account) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
		err := a.Withdrawal(rand.Intn(4) + 1)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	a := NewAccount()
	var wg sync.WaitGroup
	wg.Add(1)

	go scanCommand(a, &wg)

	for i := 0; i <= 10; i++ {
		go credit(a)
	}
	for i := 0; i <= 5; i++ {
		go debet(a)
	}

	wg.Wait()
}
