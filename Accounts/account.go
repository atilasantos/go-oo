package accounts

import (
	"fmt"

	customer "github.com/atilasantos/go-oo/Customers"
)

type Account struct {
	Customer      customer.Customer
	AgencyNumber  int16
	AccountNumber int32
	balance       float32
}

func (account *Account) ToString() string {
	message := "Name: " + account.Customer.CustomerName + "\n" +
		"Agency: " + fmt.Sprintf("%d", account.AgencyNumber) + "\n" +
		"Account Number: " + fmt.Sprintf("%d", account.AccountNumber) + "\n" +
		"Balance " + fmt.Sprintf("%f", account.GetBalance())
	return message
}

func (account *Account) Withdraw(amount float32) {
	if amount <= account.GetBalance() && amount > 0 {
		account.balance -= amount
		fmt.Println("\nWithdraw made with success\n")
	} else {
		fmt.Println("\nNot enough balance\n")
	}

}

func (account *Account) Deposit(amount float32) (string, bool) {
	if amount > 0 {
		account.balance += amount
		return "Deposit made with success.", true
	}
	return "Deposit failed due to value less than zero.", false
}

func (account *Account) TransferTo(receiverAccount *Account, amount float32) bool {
	if account.GetBalance() >= amount && amount > 0 {
		account.Withdraw(amount)
		receiverAccount.Deposit(amount)
		return true
	} else {
		fmt.Println("Balance not enough..")
		return false
	}
}

func (account *Account) GetBalance() float32 {
	return account.balance
}

func (account *Account) SetBalance(amount float32) {
	if amount > 0 {
		account.balance = amount
	} else {
		fmt.Println("Value less than zero.")
	}
}
