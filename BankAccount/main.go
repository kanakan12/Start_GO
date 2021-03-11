package main

import (
	"fmt"

	"github.com/JINHYUN/BankAccount/accounts"
)

func main() {
	// 다음과 같이 작성하는 경우, Owner와 Balance 모두 접근가능
	// account := banking.Account{Owner: "JINHYUN", Balance: 10000}

	account := accounts.NewAccount("JINHYUN")
	fmt.Println(account)

	// owner, balance는 private로 설정되어 있기 때문에 해당과 같이 이름을 변경할 수 없음
	// account.owner = "KAKA"
	// account.balance = 10000

	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	// GO에서 error를 핸들하는 방법
	if err != nil {
		// Println()을 실행하고 프로그램을 종료
		//log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("KAKA")
	fmt.Println(account.Balance(), account.Owner())

	fmt.Println(account)
}
