package main

import (
	"fmt"
	"time"
)

func main() {
	// 앞에 go를 적으줌으로써, 단계처리 방식에서 병렬처리 방식으로 변경
	// go가 없을 경우, 20초 작업
	// go가 있을 경우, 10초 작업 : 동시에 진행하기 때문에
	// go는 main 메서드가 동작하는 경우에만 동작을 진행, main 메서드는 goroutine을 기다리지 않음
	// 아래 두 메서드에 모두 go를 적을 시, 프로그램이 종료됨
	//go sexyCount("A")
	//go sexyCount("B")
	// 위 두 메서드에 모두 go를 적었으며 해당 문구를 추가할 경우, main 메서드는 5초동안 살아있고 그 후 프로그램이 종료 됨
	//time.Sleep(time.Second * 5)

	c := make(chan string)
	people := [2]string{"A", "B"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// blocking operation
	// channel로 부터 메세지를 받을 때는 main 메서드는 기다리고 있음
	fmt.Println("Waiting for message")
	fmt.Println("Received this message:", <-c)
	fmt.Println("Received this message:", <-c)
	// goroutine에서 만든 channel을 2개이기 때문에 deadlock 발생
	// 컴파일시 알 순 없지만 실행시 문제가 있다는 것을 알 수 있음
	//fmt.Println(<-c)

	c1 := make(chan string)
	people1 := [5]string{"A", "B", "C", "D", "E"}
	for _, person1 := range people1 {
		go isSexy(person1, c1)
	}
	for i := 0; i < len(people1); i++ {
		fmt.Println("Waiting for", i)
		fmt.Println(<-c1)
	}
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	// channel로 메세지를 보내고 있음
	c <- person + " is sexy"
}

// Goroutine이란 기본적으로 다른 함수와 동싱에 실행시키는 함수
// Channel은 파이프와 같은 것
// 이러한 파이프를 통해 메세지를 보낼 수도, 받을 수도 있음
