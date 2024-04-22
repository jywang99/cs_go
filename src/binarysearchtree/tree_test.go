package binarysearchtree_test

import (
	"testing"

	bst "jy.org/csgo/src/binarysearchtree"
)

func TestArrToBST(t *testing.T) {
    arr := []int{1,2,3,4,5,6}
    node := bst.SortedArrToBalancedBST(arr)
    node.Print()
}

