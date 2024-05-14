package binarysearchtree_test

import (
	"fmt"
	"testing"

	bst "jy.org/csgo/src/binarysearchtree"
	ct "jy.org/csgo/src/custtypes"
)

type compInt struct {
    data int
}

func (ci *compInt) CompareTo(ct interface{}) int {
    cti := ct.(compInt).data
    if ci.data == cti {
        return 0
    }
    if ci.data < cti {
        return -1
    }
    return 1
}

func (ci *compInt) ToString() string {
    return fmt.Sprintf("%d", ci.data)
}

func TestArrToBST(t *testing.T) {
    arr := []ct.Comparable{
        &compInt{1},
        &compInt{2},
        &compInt{3},
        &compInt{4},
        &compInt{5},
        &compInt{6},
    }
    node := bst.SortedArrToBalancedBST(arr)
    node.Print()
}

