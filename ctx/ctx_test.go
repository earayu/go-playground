package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func waitAndExec(d time.Duration, f func()) {
	time.Sleep(d)
	f()
}

func TestCancelCtx(t *testing.T) {
	baseCtx := context.Background()
	cancelCtx, _ := context.WithCancel(baseCtx)
	//d := cancelCtx.Done()
	//fmt.Println(<-d)
	//waitAndExec(1*time.Second, func() {
	//	cancel()
	//})
	d := cancelCtx.Done()
	fmt.Println(<-d)
}
