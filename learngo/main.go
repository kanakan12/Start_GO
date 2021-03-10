// package main
// function main is undeclared in the main package
// 특정 function 필요
// 컴파일러는 main package와 main func을 먼저 찾고 실행
package main

import (
	"fmt" // formatting을 위한 package
	"strings"

	"github.com/JINHYUN/learngo/something"
)

// struct는 object와 비슷하며, map보다는 유연함
// map에서의 value는 특정 타입 하나로 고정되어 있기 때문에
// GO는 Class가 존재하지 않으며, object도 존재하지 않음
type person struct {
	name    string
	age     int
	favFood []string
}

func canIDrink(age int) bool {
	// variable experssion
	// ; 사용 이전에는 변수를 선언할 수 있으며, ; 이후에는 해당 변수를 조건식으로 사용 가능
	// koreanAge는 if 내에서만 사용할 수 있는 지역변수로서 사용
	if koreanAge := age + 2; koreanAge < 21 {
		return false
	}

	// switch 또한 변수 선언이 가능하며, 해당 변수를 조건으로 사용 가능
	switch KoreanAge := age + 2; KoreanAge {
	case 10:
		return false
	case 11:
		return false
	}
	return true
}

func superAdd(numbers ...int) int {
	// range는 index를 사용하기 때문에 결과가 0 1 2 3 4 5로 출력
	// range는 for에서만 사용 가능하며 index와 number를 받음
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
}

// naked return
// return의 값을 이미 선언했기 때문에 메서드 마지막의 return에 추가적으로 작성할 필요 없음
func lenAndUpper1(name string) (lenght int, uppercase string) {
	// defer는 메서드가 끝나고 나서 코드를 실행
	defer fmt.Println("I`m done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// 타입앞에 ...을 이용해서 여러개의 변수를 받을 수 있음
func repeatMe(words ...string) {
	fmt.Println(words)
}

// GO는 하나의 func에서 여러개의 리턴을 가질 수 있음
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// a, b의 타입이 같다면
// func multiply(a, b int) int 와 같이 사용할 수 있음
func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello World!")
	something.SayHello()
	// something.sayBye()는 사용할 수 없음

	// 타입이 없는 상수(constant)
	const name = "nico"
	fmt.Println(name)

	// var name1 string = "kaka"
	// fmt.Println(name1)
	// var 변수는 변경 가능
	// name1 = "nana"
	// fmt.Println(name1)

	// 축약형은 func안에서만 사용 가능하며, 변수에만 적용가능
	name1 := "jinhyun"
	// var name1 string = "kaka"와 같은 의미
	fmt.Println(name1)

	fmt.Println(multiply(2, 2))

	// lenAndUpper 메서드는 2개의 리턴 값을 가지기 때문에 2개의 변수를 생성해야 함
	totalLength, upperName := lenAndUpper("JINHYUN")
	fmt.Println(totalLength, upperName)

	// _는 무시된 변수
	totalLength1, _ := lenAndUpper("JINHYUN")
	fmt.Println(totalLength1)

	repeatMe("A", "B", "C", "D", "E")

	totalLenght2, up := lenAndUpper1("JINHYUN")
	fmt.Println(totalLenght2, up)

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(18))

	// 메모리에 저장된 같은 object를 복사하지 않고, 참조로 확인할 수 있음
	a := 2
	b := a
	// c는 a의 메모리 주소에 연결되어 있음
	c := &a
	a = 5
	fmt.Println(a, b, c)
	// &는 주소와 같은 의미
	fmt.Println(&a, &b, &c)
	// *를 사용해서 메모리 주소의 값을 볼 수 있음
	fmt.Println(*c)

	// c는 a의 메모리 주소에 연결되어 있기 때문에 c의 값을 변경하면 a의 값이 변경됨
	// c는 a를 살펴보는 pointer이기 때문에 값이 변경
	*c = 20
	fmt.Println(a)

	// Array의 범위가 설정되어 있는 case
	names := [5]string{"A", "B", "C"}
	names[3] = "D"
	names[4] = "E"
	fmt.Println(names)

	// Go에서는 slice라는 형태의 배열을 사용
	// 배열의 크기를 지정하지 않음
	names1 := []string{"A", "B", "C"}
	fmt.Println(names1)

	// slice에 item을 추가하기 위해서는 append() 메서드 사용
	// append() 메서드는 값을 수정하지 않음
	names1 = append(names1, "D")
	fmt.Println(names1)

	// map은 key와 value로 이루워져있음
	//           key   value
	KAKA := map[string]string{"name": "KAKA", "age": "16"}
	fmt.Println(KAKA)

	// map도 for문으로 사용 가능
	for key, value := range KAKA {
		fmt.Println(key, value)
	}

	// struct 입력 방법에는 두 가지가 존재
	// 1. struct의 순서대로 기입
	favFood := []string{"A", "B"}
	NANA := person{"NANA", 20, favFood}
	fmt.Println(NANA)

	// 2. struct의 type을 같이 기입
	BABA := person{name: "BABA", age: 20, favFood: favFood}
	fmt.Println(BABA)
}
