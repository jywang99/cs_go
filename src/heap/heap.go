package heap

import (
	"errors"

	ct "jy.org/csgo/src/custtypes"
)

type minHeap struct {
    arr []*ct.Comparable
    size int
    cap int
}

func NewMinHeap(cap int) *minHeap {
    return &minHeap{
        arr: make([]*ct.Comparable, cap),
        size: 0,
        cap: cap,
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
    if le && (*h.arr[li]).CompareTo(*h.arr[mi]) < 0 {
        mi = li
    }
    if re && (*h.arr[ri]).CompareTo(*h.arr[mi]) < 0 {
        mi = ri
    }

    if mi != i {
        h.swap(mi, i)
        h.heapifyRecurse(mi)
    }
}

func (h *minHeap) GetMin() *ct.Comparable {
    return h.arr[0]
}

func (h *minHeap) ExtractMin() *ct.Comparable {
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

func (h *minHeap) Insert(d ct.Comparable) error {
    if h.size >= h.cap {
        return errors.New("Heap is already full!")
    }
    dd := d
    h.arr[h.size] = &dd
    h.size ++
    h.swim(h.size-1)
    return nil
}

func (h *minHeap) swim(i int) {
    for ; i != 0; {
        pi, pvalid := h.parentIdx(i)
        if pvalid && (*h.arr[i]).CompareTo(*h.arr[pi]) >= 0 {
            break
        }
        h.swap(i, pi)
        i = pi
    }
}

