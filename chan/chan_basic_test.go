package _chan

import (
	"fmt"
	"testing"
)

func Test_Send_Recv(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
}

/**
 * bingo: we can use select+default to implement non-blocking send/recv
 */
func Test_NonBlock(t *testing.T) {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	// note: send & recv can be used in one select
	case msg := <-messages:
		fmt.Println("received message", msg)
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
