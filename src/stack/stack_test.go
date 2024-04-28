package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	stk "jy.org/csgo/src/stack"
)

func TestStack(t *testing.T) {
    t.Log("Start TestStack")
    s := stk.Stack{}
    s.Push(1)
    s.Push(2)
    s.Push(3)

    v := s.Pop()
    assert.Equal(t, 3, v)
    v = s.Pop()
    assert.Equal(t, 2, v)
    v = s.Pop()
    assert.Equal(t, 1, v)

    v = s.Pop()
    assert.Equal(t, nil, v)
}

