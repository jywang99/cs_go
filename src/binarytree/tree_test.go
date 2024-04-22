package binarytree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	bt "jy.org/csgo/src/binarytree"
)

func TestDepth(t *testing.T) {
    t.Log("Starting TestDepth")
    head := bt.NewNode(1)

    two := bt.NewNode(2)
    head.Left = two
    head.Right = bt.NewNode(3)

    two.Left = bt.NewNode(4)
    two.Right = bt.NewNode(5)

    assert.Equal(t, 3, head.GetDepth())
}

func TestVerticalTraverse(t *testing.T) {
    t.Log("Starting TestVerticalTraverse")

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
    
    exp := make(map[int][]int)
    exp[-2] = []int{4}
    exp[-1] = []int{2}
    exp[0] = []int{1,5,6}
    exp[1] = []int{3,8}
    exp[2] = []int{7}
    exp[3] = []int{9}

    assert.Equal(t, exp, vmap)
}

func TestFindLCA(t *testing.T) {
    t.Log("Starting TestFindLCA")

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

