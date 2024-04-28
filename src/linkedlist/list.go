package linkedlist

import "fmt"

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
    Prev *DoublyLinkedList
    Next *DoublyLinkedList
}

func NewDoublyLinkedList(d any) *DoublyLinkedList {
    return &DoublyLinkedList {
        d,
        nil,
        nil,
    }
}

func (src *DoublyLinkedList) Prepend(d any) *DoublyLinkedList {
    n := NewDoublyLinkedList(d)

    pn := src.Prev
    if pn != nil {
        n.Prev = pn
        pn.Next = n
    }

    src.Prev = n
    n.Next = src

    return n
}

func (src *DoublyLinkedList) Append(d any) *DoublyLinkedList {
    n := NewDoublyLinkedList(d)

    nn := src.Next
    if nn != nil {
        n.Next = nn
        nn.Prev = n
    }

    src.Next = n
    n.Prev = src

    return n
}

func (src *DoublyLinkedList) PrependAll(n *DoublyLinkedList) {
    nTail := n.Tail()
    nHead := n.Head()
    oPrev := src.Prev

    src.Prev = nTail
    nTail.Next = src

    nHead.Prev = oPrev
    if oPrev != nil {
        oPrev.Next = nHead
    }
}

func (src *DoublyLinkedList) AppendAll(n *DoublyLinkedList) {
    nHead := n.Head()
    nTail := n.Tail()
    oNext := src.Next

    src.Next = nHead
    nHead.Prev = src

    nTail.Next = oNext
    if oNext != nil {
        oNext.Prev = nTail
    }
}

func (src *DoublyLinkedList) Pop() any {
    pp := src.Prev
    nn := src.Next
    if pp != nil {
        pp.Next = nn
    }
    if nn != nil {
        nn.Prev = pp
    }

    src.Prev = nil
    src.Next = nil

    return src.Data
}

func (n *DoublyLinkedList) Head() *DoublyLinkedList {
    if n.Prev == nil {
        return n
    }
    return n.Prev.Head()
}

func (n *DoublyLinkedList) Tail() *DoublyLinkedList {
    if n.Next == nil {
        return n
    }
    return n.Next.Tail()
}

func (n *DoublyLinkedList) Print() {
    head := n.Head()
    fmt.Print("(nil) - ")
    head.printRecurse()
    fmt.Println("(nil)")
}

func (n *DoublyLinkedList) printRecurse() {
    if n == nil {
        return
    }
    fmt.Printf("%v - ", n.Data)
    n.Next.printRecurse()
}

func ArrToDll(arr []int) *DoublyLinkedList {
    if len(arr) == 0 {
        return nil
    }

    head := NewDoublyLinkedList(arr[0])
    n := head
    if len(arr) == 1 {
        return n
    }

    for _, i := range(arr[1:]) {
        n = n.Append(i)
    }
    return head
}

