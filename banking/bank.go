package main

type Bank struct {
	Users []Account `json: "users`
}

func (bank *Bank) FindUser(u string) (int, error) {
	for k, v := range bank.Users {
		if v.Username == u {
			return k, nil
		}
	}
	return 0, NotFound
}

func (bank *Bank) AddUser(u string, b int) {
	bank.Users = append(bank.Users, CreateAccount(u, b))
}

func (bank *Bank) DeleteUser(u string) {
	for _, i := range bank.Users {
		if i.Username == u {

		}
	}
}
