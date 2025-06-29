package main

import "fmt"

type BankAccount struct {
	Balance float64
}

func (acc *BankAccount) Deposit(amount float64) {
	acc.Balance += amount
}

func (acc *BankAccount) Withdraw(amount float64) (float64, bool) {
	if amount > acc.Balance {
		return acc.Balance, false
	}
	acc.Balance -= amount
	return acc.Balance, true
}

func main() {
	account := BankAccount{Balance: 1000}

	account.Deposit(500)
	fmt.Println("Баланс после пополнения:", account.Balance)

	// Снятие
	if newBalance, ok := account.Withdraw(500); ok {
		fmt.Println("Снятие успешно. Новый баланс:", newBalance)
	} else {
		fmt.Println("Ошибка: недостаточно средств")
	}
}
