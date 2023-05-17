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

/*
* 1. callTickerInGoRoutine() returns immediately
* 2. the main goroutine sleeps for 5 seconds
* 3. QUESTION: whether the ticker goroutine is still running before main goroutine end?
*
* ANSWER: yes, the ticker goroutine is still running

* note: 当一个 goroutine 不再被引用时，并不会立即结束。与 Java 中线程的生命周期不同，Go 中的 goroutine 生命周期只有在自身逻辑执行完毕之后才会结束。
 */
func Test_callTickerInGoRoutineAndReturnImmediately(t *testing.T) {
	callTickerInGoRoutine()
	time.Sleep(5 * time.Second)
}

func callTickerInGoRoutine() {
	go tickerLoop()
}

func tickerLoop() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}
