package binarysearchtree_test

import (
	"testing"

	bst "jy.org/csgo/src/binarysearchtree"
)

func TestArrToBST(t *testing.T) {
    arr := []any { 1,2,3,4,5,6,7 }
    node := bst.SortedArrToBalancedBST(arr)
    node.Print()
}

