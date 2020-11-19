package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {

	return len(name), strings.ToUpper(name)

}
func main() {
	// 2일차
	// something.SayHello()
	// something.saybye() 는 소문자로 시작하는 private 함수라서 반응하지 않음
	// ctrl+click하면 원하는 정보를 다 볼 수있다.
	// same with var name string = "nico"
	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}

func multiply(a, b int) int {
	return a * b

}
func repeatMe(words ...string) {
	fmt.Println(words)
}
