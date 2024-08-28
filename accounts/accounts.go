package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance uint
}

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount deposit
func (a *Account) Deposit(amount uint) {
	a.balance += amount
}

func (a *Account) Withdraw(amount uint) error {
	if a.balance < amount {
		return errors.New("돈이 부족함")
	}
	a.balance -= amount
	return nil
}

func (a Account) Balance() uint {
	return a.balance
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

// 해당 struct를 호출하면 자동 호출되는 함수
func (a Account) String() string {
	//fmt.Println(a) 내부에서 호출하면 재귀호출됨

	return fmt.Sprint(a.Owner(), "예약 함수인가\n ?? : ", a.Balance())
}
