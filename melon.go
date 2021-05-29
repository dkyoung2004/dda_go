package main

import (
	_ "github.com/lib/pq"
)

var hanguel = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
var songIDList = []int{
	30287019,
}
var searchkeyWord = "ㄱㄴㅇ ㅇㅈ ㅁ ㅎ ㅂㅇㅂ"

type songinfo struct {
	id     int
	title  string
	artist string
	lyrics string
}

func main() {

}
