package main

//사용되지 않는 모듈은 에러를 냅니다.
import "fmt" // include<stdio.h>

func main() {
	//fmt에서 불러오는 함수다.
	//Println이라는게
	//그래서 어떤 모듈에서 export되는 imprt가 가능한 함수들은 무조건 대문자 시작
	fmt.Println("Hello, World")

}
