package time

import (
	"fmt"
	"testing"
	"time"
)

// typical usage of time.Ticker
func Test_ticker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-done:
			fmt.Println("done")
			return
		}
	}
}
