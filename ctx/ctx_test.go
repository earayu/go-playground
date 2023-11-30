package ctx

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCancelCtx(t *testing.T) {
	baseCtx := context.Background()
	doneC, cancelFunc := context.WithCancel(baseCtx)

	cancelFunc()

	v, ok := <-doneC.Done()
	// note:
	// since we call cancelFunc, the doneC.Done() chan will be closed, so we recv Zero Value from a closed chan
	assert.Equal(t, struct{}{}, v)
	// At the same time, the second return value is false, which means the chan is closed
	assert.Equal(t, false, ok)
}
