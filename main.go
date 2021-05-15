package main

import "fmt"

//사용되지 않는 모듈은 에러를 냅니다.
// include<stdio.h>
func addAll(inp ...int) (total int) {
	total = 0
	for _, val := range inp {
		total += val

	}
	return total
}
func main() {
	//fmt에서 불러오는 함수다.
	//Println이라는게
	//그래서 어떤 모듈에서 export되는 imprt가 가능한 함수들은 무조건 대문자 시작
	// var 변수명 변수타입
	// 이름 := 값
	//fmt.Println(addAll(1, 2, 3, 3, 4, 5, 6, 7, 8, 9))

	age := 20

	if manAge := age - 1; manAge >= 19 {
		fmt.Println("You can buy an alchol")
	} else {
		fmt.Println("You cannot buy an alchol")
	}

	switch manAge := age - 1; {
	case manAge <= 13:
		fmt.Println("어린이 요금")
	case manAge < 65:
		fmt.Println("일반 요금")
	case manAge >= 65:
		fmt.Println("노약자 요금")
	}

	arr := []int(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(arr)

	append(arr, 11)
}
