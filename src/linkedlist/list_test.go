package linkedlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ll "jy.org/csgo/src/linkedlist"
)

func TestSinglyLinkedList(t *testing.T) {
    head := ll.NewSinglyLinkedList(1)
    n := head
    for _, i := range([]int{2, 3, 4}) {
        n = n.Append(i)
    }
    two := head.Next
    two.Append(5)

    assert.Equal(t, head.Data, 1)
    assert.Equal(t, two.Data, 2)
    five := two.Next
    assert.Equal(t, five.Data, 5)
    three := five.Next
    assert.Equal(t, three.Data, 3)
    four := three.Next
    assert.Equal(t, four.Data, 4)
}

func TestDoublyLinkedList(t *testing.T) {
    head := ll.NewDoublyLinkedList(1)
    two := head.Append(3)
    two.Prepend(2)

    assert.Equal(t, head.Data, 1)
    two = head.Next()
    assert.Equal(t, two.Data, 2)
    three := two.Next()
    assert.Equal(t, three.Data, 3)

    two = three.Prev()
    assert.Equal(t, two.Data, 2)
    head = two.Prev()
    assert.Equal(t, head.Data, 1)
}
