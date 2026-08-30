package main

type Account struct {
	Username string `json: "username"`
	Balance  int    `json: "balance"`
}

func CreateAccount(u string, b int) Account {
	return Account{
		Username: u,
		Balance:  b,
	}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.Balance-amount < 0 {
		return NotEnoughMoney
	}
	a.Balance -= amount
	return nil
}

func (a *Account) Transfer(amount int, recepient string, bank Bank) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	r, err := bank.FindUser(recepient)
	if err != nil {
		return err
	}
	bank.Users[r].Withdraw(amount)

	return nil
}
