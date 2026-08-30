package main

import (
	"fmt"
	"net/http"
	"restapi/banking"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	bank := banking.Bank{
		Users: []banking.Account{
			banking.CreateAccount("Masha", 10000),
			banking.CreateAccount("Jirinovsky", 200000),
			banking.CreateAccount("Putin", 1234567),
		},
	}

	fmt.Println(bank)

}
