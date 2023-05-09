package _select

import (
	"fmt"
	"testing"
	"time"
)

// this is a infinite loop
// the first case in the select statement will create a channel for every loop
// and since the default branch does nothing, it will enter next loop immediately.
// and the channel created in the previous loop will be garbage collected.
// for the second loop, a new channel will be created again.
func _Test_Forloop_And_Timeout(t *testing.T) {
	for {
		select {
		case <-time.After(1 * time.Second):
			t.Log("timeout")
		default:
			// do nothing
			fmt.Println("looping")
		}
	}
}

// this is the right version of the above test
// the timeout channel is created outside the loop
func Test_Forloop_And_Timeout2(t *testing.T) {
	timeout1 := time.After(1 * time.Second)
	for {
		select {
		case <-timeout1:
			t.Log("timeout")
			return
		default:
			// do nothing
			fmt.Println("looping")
		}
	}
}

// this is a interesting idiom
// the first timeout channel is created outside the loop, it controls the timeout of the whole loop
// the second timeout channel is created inside the loop, it controls the timeout of each loop
// the testcase will timeout after 10 seconds, and print "looping" every 500 milliseconds
func Test_Forloop_And_Timeout3(t *testing.T) {
	timeout1 := time.After(10 * time.Second)
	for {
		select {
		case <-timeout1:
			t.Fatalf("timeout")
			return
		case <-time.After(500 * time.Millisecond):
			// do nothing
			fmt.Println("looping")
		}
	}
}
