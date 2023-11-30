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
 * note: we can use select+default to implement non-blocking send/recv
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

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

/**
 * note: send & recv can be used in one select
 */
func Test_Send_Recv_At_The_Same_Time(t *testing.T) {
	messages := make(chan string)
	msg := "hi"
	select {
	case rmsg := <-messages:
		// 'messages' channel has no value, so this case will be blocked
		fmt.Println("received message", rmsg)
		t.Fail()
	case messages <- msg:
		// 'messages' channel has no buffer, so this case will be blocked
		fmt.Println("sent message", msg)
		t.Fail()
	default:
		fmt.Println("no message sent")
	}
}
