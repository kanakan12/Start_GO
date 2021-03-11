package accounts

import (
	"errors"
	"fmt"
)

// Account struct
// struct의 타입명의 시작이 소문자일 경우 : private
//                     대문자일 경우 : public
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can`t withdraw")

// NewAccount creates Account
// func를 통해 object를 만드는 방법
func NewAccount(owner string) *Account {
	// account 초기화
	account := Account{owner: owner, balance: 0}
	// 새로운 object를 return하고 싶을 때, 새로운 object를 만들지 않고 복사본 자체를 return
	// 프로그램 속도 저하 방지 목적
	return &account
}

// (a Account) : 해당과 같은 기능을 receiver라 함, 변수 a와 타입인 Account로 이루워져 있음
// receiver는 struct의 첫 글자를 따서 소문자로 지어야 함
// Method라 볼 수 있음, func 뒤에 struct를 가짐
// a는 main.go에서 account의 복사본을 만들어서 사용하기 때문에 a의 값은 amount에 따라 증가하나, account의 값은 증가하지 않음
// GO에서는 보안을 위해서 복사본을 만들어서 사용하기 때문
/* func (a Account) Deposit(amount int) {
 	fmt.Println("Gonna deposit", amount)
 	a.balance += amount
}*/

// Deposit x amount on your account
// 다음과 같이 *를 사용해서 a를 복사본이 아닌, main.go의 account를 직접 사용하는것 가능
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	// nil은 null, none과 같은 의미
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// GO에서 자동으로 호출해주는 메서드
// 따라서 main.go에서 account를 호출할 경우, 형식을 바꿀 수 있음
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "`s account\nHas: ", a.Balance())
}
