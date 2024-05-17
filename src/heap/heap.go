package heap

import (
	"errors"

	ct "jy.org/csgo/src/custtypes"
)

type minHeap struct {
    arr []any
    size int
    cap int
    cmp ct.Comparator
}

func NewMinHeap(cap int, cmp ct.Comparator) *minHeap {
    return &minHeap{
        arr: make([]any, cap),
        size: 0,
        cap: cap,
        cmp: cmp,
    }
}

func (h *minHeap) parentIdx(i int) (int, bool) {
    if i <= 0 {
        return 0, false
    }
    return i/2, true
}

func (h *minHeap) leftIdx(i int) (int, bool) {
    li := i*2 + 1
    return li, li < h.size
}

func (h *minHeap) rightIdx(i int) (int, bool) {
    ri := i*2 + 2
    return ri, ri < h.size
}

func (h *minHeap) swap(i, j int) {
    tmp := h.arr[i]
    h.arr[i] = h.arr[j]
    h.arr[j] = tmp
}

func (h *minHeap) heapify() {
    h.heapifyRecurse(0)
}

func (h *minHeap) heapifyRecurse(i int) {
    li, le := h.leftIdx(i)
    ri, re := h.rightIdx(i)

    mi := i
    if le && h.cmp(h.arr[li], h.arr[mi]) < 0 {
        mi = li
    }
    if re && h.cmp(h.arr[ri], h.arr[mi]) < 0 {
        mi = ri
    }

    if mi != i {
        h.swap(mi, i)
        h.heapifyRecurse(mi)
    }
}

func (h *minHeap) GetMin() any {
    return h.arr[0]
}

func (h *minHeap) ExtractMin() any {
    if h.size <= 0 {
        return nil
    }
    r := h.arr[0]

    if h.size == 1 {
        h.arr[0] = nil
        h.size --
        return r
    }

    h.arr[0] = h.arr[h.size-1]
    h.size --
    h.heapify()
    return r
}

func (h *minHeap) Insert(d any) error {
    if h.size >= h.cap {
        return errors.New("Heap is already full!")
    }
    h.arr[h.size] = d
    h.size ++
    h.swim(h.size-1)
    return nil
}

func (h *minHeap) swim(i int) {
    for ; i != 0; {
        pi, pvalid := h.parentIdx(i)
        if pvalid && h.cmp(h.arr[i], h.arr[pi]) >= 0 {
            break
        }
        h.swap(i, pi)
        i = pi
    }
}

