package global

import "fmt"

func HandlePanic() {
	r := recover()
	fmt.Println(r)
}
