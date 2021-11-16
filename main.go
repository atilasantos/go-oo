package main

import (
	"fmt"

	accounts "github.com/atilasantos/go-oo"
)

func main() {
	accountJoselito := accounts.Account{Name: "Joselito", AgencyNumber: 12345, AccountNumber: 33245, Balance: 785.12}
	fmt.Println(accountJoselito.ToString())
	fmt.Println()
	accountMaria := accounts.Account{Name: "Maria", AgencyNumber: 4321, AccountNumber: 44532, Balance: 200.}
	fmt.Println(accountMaria.ToString())

	fmt.Println("Balance after transfer:")
	accountJoselito.TransferTo(&accountMaria, 500.)

	fmt.Println("Balance Joselito: ", accountJoselito.Balance)
	fmt.Println("Balance Maria: ", accountMaria.Balance)

}
