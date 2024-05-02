package heap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	hp "jy.org/csgo/src/heap"
)

func TestHeap(t *testing.T) {
    h := hp.NewMinHeap(5)
    h.Insert(10)
    h.Insert(7)
    h.Insert(8)
    h.Insert(4)
    h.Insert(1)

    err := h.Insert(100)
    assert.NotNil(t, err)

    assert.Equal(t, 1, *h.GetMin())
    assert.Equal(t, 1, *h.ExtractMin())
    assert.Equal(t, 4, *h.GetMin())
    assert.Equal(t, 4, *h.ExtractMin())
    assert.Equal(t, 7, *h.ExtractMin())
    assert.Equal(t, 8, *h.ExtractMin())
    assert.Equal(t, 10, *h.ExtractMin())

    var iptr *int = nil
    assert.Equal(t, iptr, h.ExtractMin())
}
