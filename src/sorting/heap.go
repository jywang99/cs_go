package sorting

import (
	ct "jy.org/csgo/src/custtypes"
	hp "jy.org/csgo/src/heap"
)

type HeapSorter struct{
    Sorter
}

func (b *HeapSorter) GetName() string {
    return "Heap Sort"
}

func (b *HeapSorter) Sort(arr *[]any, cmp ct.Comparator) {
    h := hp.NewMinHeap(len(*arr), cmp)
    for _, e := range *arr {
        h.Insert(e)
    }
    for i:=0; i<len(*arr); i++ {
        (*arr)[i] = h.ExtractMin()
    }
}

