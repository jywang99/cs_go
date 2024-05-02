package heap

import (
	"errors"
)

type MinHeap struct {
    arr []*int
    size int
    cap int
}

func NewMinHeap(cap int) *MinHeap {
    return &MinHeap{
        arr: make([]*int, cap),
        size: 0,
        cap: cap,
    }
}

func (h *MinHeap) parentIdx(i int) (int, bool) {
    if i <= 0 {
        return 0, false
    }
    return i/2, true
}

func (h *MinHeap) leftIdx(i int) (int, bool) {
    li := i*2 + 1
    return li, li < h.size
}

func (h *MinHeap) rightIdx(i int) (int, bool) {
    ri := i*2 + 2
    return ri, ri < h.size
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
    li, lvalid := h.leftIdx(i)
    ri, rvalid := h.rightIdx(i)

    mi := i
    if lvalid && *h.arr[li] < *h.arr[mi] {
        mi = li
    }
    if rvalid && *h.arr[ri] < *h.arr[mi] {
        mi = ri
    }

    if mi != i {
        h.swap(mi, i)
        h.heapifyRecurse(mi)
    }
}

func (h *MinHeap) GetMin() *int {
    return h.arr[0]
}

func (h *MinHeap) ExtractMin() *int {
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

func (h *MinHeap) Insert(d int) error {
    if h.size >= h.cap {
        return errors.New("Heap is already full!")
    }
    dd := d
    h.arr[h.size] = &dd
    h.size ++
    h.swim(h.size-1)
    return nil
}

func (h *MinHeap) swim(i int) {
    for ; i != 0; {
        pi, pvalid := h.parentIdx(i)
        if pvalid && *h.arr[i] >= *h.arr[pi] {
            break
        }
        h.swap(i, pi)
        i = pi
    }
}

