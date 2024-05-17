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

func TestHuffman(t *testing.T) {
    rs, ht, err := greedy.NewHuffmanEncode("AAAAAABBCCDDEEFFFFF")
    assert.Nil(t, err)
    assert.Equal(t, 46, len(rs)) // since order of traversing a map is not guaranteed, specific codes can be different
    assert.NotNil(t, ht)
}

