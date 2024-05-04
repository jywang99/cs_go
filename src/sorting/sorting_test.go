package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	s "jy.org/csgo/src/sorting"
)

func TestBubbleSort(t *testing.T) {
    sorters := []s.Sorter{
        &s.BubbleSorter{},
        &s.InsertionSorter{},
        &s.SelectionSorter{},
        &s.HeapSorter{},
        &s.QuickSorter{},
        &s.MergeSorter{},
    }

    for _, sorter := range(sorters) {
        name := sorter.GetName()
        t.Logf("Testing %s", name)
        arr := []int{4,3,1,5,7,6,2}
        sorter.Sort(&arr)

        expArr := []int{1,2,3,4,5,6,7}
        assert.Equal(t, expArr, arr)
        t.Logf("Finished testing %s", name)
    }
}

