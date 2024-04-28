package linkedlist_test

import (
	"fmt"
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
    two = head.Next
    assert.Equal(t, two.Data, 2)
    three := two.Next
    assert.Equal(t, three.Data, 3)

    two = three.Prev
    assert.Equal(t, two.Data, 2)
    head = two.Prev
    assert.Equal(t, head.Data, 1)
}

func TestDllPop(t *testing.T) {
    one := ll.NewDoublyLinkedList(1)
    two := one.Append(2)
    three := two.Append(3)
    four := three.Append(4)

    var nilptr *ll.DoublyLinkedList = nil
    assert.Equal(t, nilptr, one.Prev)
    assert.Equal(t, 2, two.Pop())
    assert.Equal(t, three, one.Next)
    assert.Equal(t, one, three.Prev)
    assert.Equal(t, 4, four.Pop())
    assert.Equal(t, nilptr, three.Next)
}

func TestDllHeadToTail(t *testing.T) {
    one := ll.NewDoublyLinkedList(1)
    two := one.Append(2)
    three := two.Append(3)
    four := three.Append(4)
    three.Print()

    assert.Equal(t, one, three.Head())
    assert.Equal(t, four, one.Tail())
}

func TestArrToDll(t *testing.T) {
    n := ll.ArrToDll([]int{1,2,3,4})

    assert.Equal(t, 1, n.Data)
    n = n.Next
    assert.Equal(t, 2, n.Data)
    n = n.Next
    assert.Equal(t, 3, n.Data)
    n = n.Next
    assert.Equal(t, 4, n.Data)
}

func TestJoining(t *testing.T) {
    n := ll.ArrToDll([]int{1,2})
    n1 := ll.ArrToDll([]int{3,4,5})
    n2 := ll.ArrToDll([]int{6,7})

    n1.PrependAll(n)
    n1.Tail().AppendAll(n2)

    fmt.Print("TestJoining: ")
    n.Print()
    for _, i := range([]int{1,2,3,4,5,6,7}) {
        assert.Equal(t, i, n.Data)
        n = n.Next
    }
}

