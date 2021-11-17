package main

import (
	"fmt"

	accounts "github.com/atilasantos/go-oo/Accounts"
	customers "github.com/atilasantos/go-oo/Customers"
)

func main() {
	accountJoselito := accounts.Account{Customer: customers.Customer{CustomerName: "John", Cpf: "192.333.123.09", Job: "Developer"}, AgencyNumber: 12345, AccountNumber: 33245}
	accountJoselito.Deposit(1000.)
	fmt.Println(accountJoselito.ToString())
	fmt.Println()
	accountMaria := accounts.Account{Customer: customers.Customer{CustomerName: "Maria", Cpf: "391.223.019.11", Job: "Ui/Ux designer"}, AgencyNumber: 4321, AccountNumber: 44532}
	accountMaria.Deposit(750.)
	fmt.Println(accountMaria.ToString())

	fmt.Println("Balance after transfer:")
	accountJoselito.TransferTo(&accountMaria, 500.)

	fmt.Println("Balance John: ", accountJoselito.GetBalance())
	fmt.Println("Balance Maria: ", accountMaria.GetBalance())
	fmt.Println()
	fmt.Println()

	fmt.Println(accountJoselito)
	fmt.Println(accountMaria)

}
