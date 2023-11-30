package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilSlice(t *testing.T) {
	var s []int
	assert.Nil(t, s)
	assert.True(t, len(s) == 0)
}

func TestMakeSlice(t *testing.T) {
	a := make([]int, 0, 1)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(len(a))
}
