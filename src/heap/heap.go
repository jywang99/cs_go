package heap

import (
	"errors"

	ct "jy.org/csgo/src/custtypes"
)

type MinHeap struct {
    arr []any
    size int
    cap int
    cmp ct.Comparator
}

func NewMinHeap(cap int, cmp ct.Comparator) *MinHeap {
    return &MinHeap{
        arr: make([]any, cap),
        size: 0,
        cap: cap,
        cmp: cmp,
    }
}

func (h *MinHeap) Size() int {
    return h.size
}

func (h *MinHeap) parentIdx(i int) (int, bool) {
    if i <= 0 {
        return 0, false
    }
    return i/2, true
}

func (h *MinHeap) LeftIdx(i int) (int, bool) {
    li := i*2 + 1
    return li, li < h.size
}

func (h *MinHeap) RightIdx(i int) (int, bool) {
    ri := i*2 + 2
    return ri, ri < h.size
}

func (h *MinHeap) GetValue(i int) any {
    return h.arr[i]
}

func (h *MinHeap) swap(i, j int) {
    tmp := h.arr[i]
    h.arr[i] = h.arr[j]
    h.arr[j] = tmp
}

func (h *MinHeap) heapify() {
    h.heapifyRecurse(0)
}

func (h *MinHeap) heapifyRecurse(i int) {
    li, le := h.LeftIdx(i)
    ri, re := h.RightIdx(i)

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

func (h *MinHeap) GetMin() any {
    return h.arr[0]
}

func (h *MinHeap) ExtractMin() any {
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

func (h *MinHeap) Insert(d any) error {
    if h.size >= h.cap {
        return errors.New("Heap is already full!")
    }
    h.arr[h.size] = d
    h.size ++
    h.swim(h.size-1)
    return nil
}

func (h *MinHeap) swim(i int) {
    for ; i != 0; {
        pi, pvalid := h.parentIdx(i)
        if pvalid && h.cmp(h.arr[i], h.arr[pi]) >= 0 {
            break
        }
        h.swap(i, pi)
        i = pi
    }
}

