package main

import "fmt"

func main() {
	// 2일차
	// something.SayHello()
	// something.saybye() 는 소문자로 시작하는 private 함수라서 반응하지 않음
	// ctrl+click하면 원하는 정보를 다 볼 수있다.
	name := "nico"
	// same with var name string = "nico"
	name = "lynn"
	fmt.Println(name)
	fmt.Println(multiply(2, 2))
}

func multiply(a, b int) int {
	return a * b

}
