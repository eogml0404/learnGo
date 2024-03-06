package somthing

import "fmt"

//export 하고싶으면 함수 이름이 대문자로 시작
func sayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}
