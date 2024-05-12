package binarytree

import ll "jy.org/csgo/src/linkedlist"

func (n *Node) TraversePreOrder() []int {
    rs := make([]int, 0)
    n.preOrderRec(&rs)
    return rs
}

func (n *Node) preOrderRec(rs *[]int) {
    if n == nil {
        return
    }
    *rs = append(*rs, n.Data)
    n.Left.preOrderRec(rs)
    n.Right.preOrderRec(rs)
}

func (n *Node) TraverseInOrder() []int {
    rs := make([]int, 0)
    n.inOrderRec(&rs)
    return rs
}

func (n *Node) inOrderRec(rs *[]int) {
    if n == nil {
        return
    }
    n.Left.inOrderRec(rs)
    *rs = append(*rs, n.Data)
    n.Right.inOrderRec(rs)
}

func (n *Node) TraversePostOrder() []int {
    rs := make([]int, 0)
    n.postOrderRec(&rs)
    return rs
}

func (n *Node) postOrderRec(rs *[]int) {
    if n == nil {
        return
    }
    n.Left.postOrderRec(rs)
    n.Right.postOrderRec(rs)
    *rs = append(*rs, n.Data)
}

func (n *Node) TraverseLevelOrder() []int {
    rs := make([]int, 0)
    nn := ll.NewSinglyLinkedList(n)
    nt := nn

    for ; nn != nil; {
        nf := nn.Data.(*Node)
        rs = append(rs, nf.Data)
        nt = nt.Tail()
        if nf.Left != nil {
            nt = nt.Append(nf.Left)
        }
        if nf.Left != nil {
            nt.Append(nf.Right)
        }
        nn = nn.Next
    }

    return rs
}

