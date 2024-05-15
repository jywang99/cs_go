package heap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ct "jy.org/csgo/src/custtypes"
	hp "jy.org/csgo/src/heap"
)

func TestHeap(t *testing.T) {
    h := hp.NewMinHeap(5)
    h.Insert(ct.CompInt{Data: 10})
    h.Insert(ct.CompInt{Data: 7})
    h.Insert(ct.CompInt{Data: 8})
    h.Insert(ct.CompInt{Data: 4})
    h.Insert(ct.CompInt{Data: 1})

    err := h.Insert(&ct.CompInt{Data: 100})
    assert.NotNil(t, err)

    assert.Equal(t, 1, (*h.GetMin()).(ct.CompInt).Data)
    assert.Equal(t, 1, (*h.ExtractMin()).(ct.CompInt).Data)
    assert.Equal(t, 4, (*h.GetMin()).(ct.CompInt).Data)
    assert.Equal(t, 4, (*h.ExtractMin()).(ct.CompInt).Data)
    assert.Equal(t, 7,(*h.ExtractMin()).(ct.CompInt).Data) 
    assert.Equal(t, 8, (*h.ExtractMin()).(ct.CompInt).Data)
    assert.Equal(t, 10, (*h.ExtractMin()).(ct.CompInt).Data)

    assert.Nil(t, h.ExtractMin())
}

