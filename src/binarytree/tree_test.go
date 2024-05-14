package binarytree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	bt "jy.org/csgo/src/binarytree"
)

func TestDepth(t *testing.T) {
    head := bt.NewNode(1)

    two := bt.NewNode(2)
    head.Left = two
    head.Right = bt.NewNode(3)

    two.Left = bt.NewNode(4)
    two.Right = bt.NewNode(5)

    assert.Equal(t, 3, head.GetDepth())
}

func TestVerticalTraverse(t *testing.T) {
    head := bt.NewNode(1)

    two := bt.NewNode(2)
    three := bt.NewNode(3)
    head.Left = two
    head.Right = three

    two.Left = bt.NewNode(4)
    two.Right = bt.NewNode(5)
    three.Left = bt.NewNode(6)
    seven := bt.NewNode(7)
    three.Right = seven

    seven.Left = bt.NewNode(8)
    seven.Right = bt.NewNode(9)

    vmap := head.TraverseVertical()
    
    exp := make(map[int][]any)
    exp[-2] = []any{4}
    exp[-1] = []any{2}
    exp[0] = []any{1,5,6}
    exp[1] = []any{3,8}
    exp[2] = []any{7}
    exp[3] = []any{9}

    assert.Equal(t, exp, vmap)
}

func TestFindLCA(t *testing.T) {
    head := bt.NewNode(1)

    two := bt.NewNode(2)
    three := bt.NewNode(3)
    head.Left = two
    head.Right = three

    two.Left = bt.NewNode(4)
    two.Right = bt.NewNode(5)
    three.Left = bt.NewNode(6)
    three.Right = bt.NewNode(7)

    assert.Equal(t, 3, head.FindLCA(6,7))
    assert.Equal(t, 2, head.FindLCA(4,5))
    assert.Equal(t, 1, head.FindLCA(5,6))
}

func TestToDoublyLinkedList(t *testing.T) {
    ten := bt.NewNode(10)

    twelve := bt.NewNode(12)
    fifteen := bt.NewNode(15)
    ten.Left = twelve
    ten.Right = fifteen

    twelve.Left = bt.NewNode(25)
    twelve.Right = bt.NewNode(30)

    fifteen.Left = bt.NewNode(36)

    dll := ten.ToDoublyLinkedList()
    n := &dll
    // n.Print()

    assert.Equal(t, 25, n.Data)
    n = n.Next
    assert.Equal(t, 12, n.Data)
    n = n.Next
    assert.Equal(t, 30, n.Data)
    n = n.Next
    assert.Equal(t, 10, n.Data)
    n = n.Next
    assert.Equal(t, 36, n.Data)
    n = n.Next
    assert.Equal(t, 15, n.Data)
}

func TestFullAndCompleteTree(t *testing.T) {
    head := bt.NewNode(1)

    b := bt.NewNode(2)
    c := bt.NewNode(3)
    head.Left = b
    head.Right = c

    d := bt.NewNode(4)
    e := bt.NewNode(5)
    f := bt.NewNode(6)
    b.Left = d
    b.Right = e
    c.Left = f

    assert.False(t, head.IsFull())
    assert.True(t, head.IsComplete())

    g := bt.NewNode(7)
    c.Right = g

    assert.True(t, head.IsFull())
    assert.True(t, head.IsComplete())

    h := bt.NewNode(8)
    g.Right = h

    assert.False(t, head.IsFull())
    assert.False(t, head.IsComplete())
}

func TestBalanced(t *testing.T) {
    head := bt.NewNode(0)

    one := bt.NewNode(1)
    two := bt.NewNode(2)
    head.Left = one
    head.Right = two

    three := bt.NewNode(3)
    four := bt.NewNode(4)
    one.Left = three
    one.Right = four

    assert.True(t, head.IsBalanced())

    five := bt.NewNode(5)
    three.Left = five

    assert.False(t, head.IsBalanced())
}

func TestTraversals(t *testing.T) {
    one := bt.NewNode(1)
    two := bt.NewNode(2)
    three := bt.NewNode(3)
    four := bt.NewNode(4)
    five := bt.NewNode(5)
    six := bt.NewNode(6)
    seven := bt.NewNode(7)

    one.Left = two
    one.Right = three
    two.Left = four
    two.Right = five
    three.Left = six
    three.Right = seven

    assert.Equal(t, []any{4,2,5,1,6,3,7}, one.TraverseInOrder())
    assert.Equal(t, []any{1,2,4,5,3,6,7}, one.TraversePreOrder())
    assert.Equal(t, []any{4,5,2,6,7,3,1}, one.TraversePostOrder())
    assert.Equal(t, []any{1,2,3,4,5,6,7}, one.TraverseLevelOrder())
}
