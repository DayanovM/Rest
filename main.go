package main

import (
	"fmt"
	"net/http"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	bank := Bank{
		Users: []Account{
			CreateAccount("Masha", 10000),
			CreateAccount("Jirinovsky", 200000),
			CreateAccount("Putin", 1234567),
		},
	}

	fmt.Println(bank)

}
