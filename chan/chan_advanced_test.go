package _chan

import (
	"fmt"
	"github.com/earayu/go-playground/global"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Two_Select_Stmts_Wait_For_One_Channel(t *testing.T) {
	doneC := make(chan struct{})
	c1 := make(chan int)
	c2 := make(chan int)

	doneChannelCount := 0

	waitGoRoutine := func() {
		for {
			select {
			case <-doneC:
				fmt.Println("recv doneC")
				doneChannelCount++
				return
			case <-c1:
			case <-c2:
			}
		}
	}
	go waitGoRoutine()
	go waitGoRoutine()

	close(doneC)
	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, doneChannelCount, 2)
}

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
