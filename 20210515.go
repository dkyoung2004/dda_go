package main

type shape interface { // 함수 들의 집합
	area() float64
	perimeter() float64
}

type Rect struct { //변수들의 집함
	width  float64
	height float64
}

func (r Rect) area() float64 {

	return r.width * r.height

}

func (r Rect) perimeter() float64 {

	return (r.width * r.height) * 2
}

func main() {
	test := Rect{
		10., 5.,
	}

}
