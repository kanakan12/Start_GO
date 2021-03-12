package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// 다음과 같은 방식으로 results를 선언하면 초기화가 되지 않았기 때문에 값을 넣을 수 없음
	// 1. 초기화 하기 위해서는 선언문 뒤에 {}를 입력하는 방법
	// 2. make 메서드를 이용하는 방법
	// var results map[string]string
	// results := map[string]string{}
	results := make(map[string]string)

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	// SatusCode는 0부터 100, 200, 300까지는 리다이렉션
	// 429 Too many request
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.Status)
		return errRequestFailed
	}
	return nil
}
