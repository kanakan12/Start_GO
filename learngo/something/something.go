package something

import "fmt"

// private func
// 대문자로 시작하지 않았기 때문에
func sayHello() {
	fmt.Println("Bye")
}

// export된 func
// 대문자로 시작했기 때문에
func SayHello() {
	fmt.Println("Hello")
}
