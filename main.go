// 패키지를 꼭 써줘야함
// main은 컴파일을 위해서 필요한 것 (필수)
package main

import (
	"fmt"
	"strings"
)

// type을 꼭 명시 해줘야함
func multiply(a, b int) int {
	return a * b
}

// 표준 라이브러리가 매우 거대
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	//함수가 마친 뒤에 실행함 -> defer
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	//return length,uppercase
	return
}

// 여러인자 받을때
func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func superAdd2(numbers ...int) int {
	total := 0
	//index ignore
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func canIDrink(age int) bool {
	//조건을 체크하기 전에 변수 만듬
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	switch age {
	case 10:
		return false
	case 20:
		return true
	}
	return false
}

func main() {
	// //var name string = "DH"
	// //축약시킨 코드와 위와 똑같다. Go가 알아서 type을 찾아준다.
	// //축약형은 func안에서만 가능하고 변수에만 적용가능
	// name := "DH"

	// name = "YDH"
	// fmt.Println(name)

	// somthing.SayHello()

	// fmt.Println(multiply(2, 2))

	// totalLength, up := lenAndUpper2("ydh")
	// fmt.Println(totalLength, up)

	// // totalLength, upperName := lenAndUpper("ydh")
	// // fmt.Println(totalLength, upperName)

	// //하나의 인자값은 무시 -> _
	// // totalLength2, _ := lenAndUpper("dh")
	// // fmt.Println(totalLength2)

	// repeatMe("nico", "lynn", "dal", "marl", "flynn")

	// superAdd(1, 2, 3, 4, 5, 6)

	// result := superAdd2(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)

	fmt.Println(canIDrink2(20))
}
