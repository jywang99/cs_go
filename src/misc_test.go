package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayAppend(t *testing.T) {
    arr := make([]int, 5)
    arr = append(arr, 1)
    fmt.Println(arr)
    assert.Equal(t, 6, len(arr))
}

