package linkedlist

type SinglyLinkedList struct {
    Data any
    Next *SinglyLinkedList
}

func NewSinglyLinkedList(d any) *SinglyLinkedList {
    return &SinglyLinkedList{
        d,
        nil,
    }
}

func (src *SinglyLinkedList) Append(d any) *SinglyLinkedList {
    n := NewSinglyLinkedList(d)
    if src.Next != nil {
        n.Next = src.Next
    }
    src.Next = n
    return n
}

type DoublyLinkedList struct {
    Data any
    prev *DoublyLinkedList
    next *DoublyLinkedList
}

func NewDoublyLinkedList(d any) *DoublyLinkedList {
    return &DoublyLinkedList {
        d,
        nil,
        nil,
    }
}

func (src *DoublyLinkedList) Next() *DoublyLinkedList {
    return src.next
}

func (src *DoublyLinkedList) Prev() *DoublyLinkedList {
    return src.prev
}

func (src *DoublyLinkedList) Prepend(d any) *DoublyLinkedList {
    n := NewDoublyLinkedList(d)

    pn := src.prev
    if pn != nil {
        n.prev = pn
        pn.next = n
    }

    src.prev = n
    n.next = src

    return n
}

func (src *DoublyLinkedList) Append(d any) *DoublyLinkedList {
    n := NewDoublyLinkedList(d)

    nn := src.next
    if nn != nil {
        n.next = nn
        nn.prev = n
    }

    src.next = n
    n.prev = src

    return n
}

