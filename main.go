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

type person struct {
	name    string
	age     int
	favFood []string
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

	//fmt.Println(canIDrink2(20))
	// a := 2
	// b := &a
	// a = 5

	//메모리 주소는 &표시
	// fmt.Println(&a, &b)
	// fmt.Println(&a, b)

	//살펴보는거
	// fmt.Println(*b)

	//d를 이용해서 c바꾸기
	//d는 c를 주소에 있는 것을 살펴보는 Pointer
	// c := 2
	// d := &c
	// *d = 20
	// fmt.Println(c)

	//0,1,2,3,4
	// names := [5]string{"dh", "dh1", "dh2"}
	// names[3] = "dh3"
	// names[4] = "dh4"
	// //names[5] = "dh5" 4까지라 오류남
	// fmt.Println(names)

	// //length없이 사용
	// names2 := []string{"dh", "dh1", "dh2"}

	// //추가해주는 것
	// names2 = append(names2, "dh3")
	// fmt.Println(names2)

	// key Type, value Type
	// nico := map[string]string{"name": "nico", "age": "12"}
	// fmt.Println(nico)

	// for key, value := range nico {
	// 	fmt.Println(key, value)
	// }

	favFood := []string{"kimchi", "ramen"}

	//dh := person{"dh", 24, favFood} 이렇게 써도되지만 아래가 더 가독성이 좋다
	dh := person{name: "dh", age: 24, favFood: favFood}
	fmt.Println(dh)

	//특정값만 출력
	fmt.Println(dh.name)
	fmt.Println(dh.age)
}
