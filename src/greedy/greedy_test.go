package greedy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jy.org/csgo/src/greedy"
)

func TestSuffixArray(t *testing.T) {
    sarr := greedy.NewSuffixArray("banana")
    assert.Equal(t, 2, sarr.FindFirstOccr("nan"))
    assert.Equal(t, 0, sarr.FindFirstOccr("ban"))
    assert.Equal(t, 2, sarr.FindFirstOccr("nana"))
    assert.Equal(t, 1, sarr.FindFirstOccr("anana"))
    assert.Equal(t, -1, sarr.FindFirstOccr("asdf"))
}
