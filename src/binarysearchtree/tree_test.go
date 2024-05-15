package binarysearchtree_test

import (
	"testing"

	bst "jy.org/csgo/src/binarysearchtree"
	ct "jy.org/csgo/src/custtypes"
)

func TestArrToBST(t *testing.T) {
    arr := []ct.Comparable {
        &ct.CompInt{Data: 1},
        &ct.CompInt{Data: 2},
        &ct.CompInt{Data: 3},
        &ct.CompInt{Data: 4},
        &ct.CompInt{Data: 5},
        &ct.CompInt{Data: 6},
    }
    node := bst.SortedArrToBalancedBST(arr)
    node.Print()
}

