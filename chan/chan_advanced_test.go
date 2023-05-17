package _chan

import (
	"fmt"
	"github.com/earayu/go-playground/global"
	"testing"
)

func Test_Recv_From_A_Closed_Chan(t *testing.T) {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	close(ch)
	// note: we can still recv from a closed chan
	v, ok := <-ch
	fmt.Println(v)
	fmt.Println(ok)
	v, ok = <-ch
	// note: after drained, we can call '<-' from a closed chan, but the value is zero value
	fmt.Println(v)
	fmt.Println(ok)
}

func Test_Send_To_A_Closed_Chan(t *testing.T) {
	defer global.HandlePanic()
	ch := make(chan int, 100)
	close(ch)
	ch <- 1
	// note: we can't send to a closed chan, otherwise panic
}
