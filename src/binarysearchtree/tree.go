package binarysearchtree

import (
	"fmt"

	ct "jy.org/csgo/src/custtypes"
)

type Node struct {
    Data any
    Left *Node
    Right *Node
}

type BST struct {
    Root *Node
    Cmp ct.Comparator
}

func NewNode(data any) *Node {
    return &Node{
        data,
        nil,
        nil,
    }
}

func (bst *BST) Print() {
    printRecurse([]*Node{bst.Root})
}

func printRecurse(nodes []*Node) {
    newNodes := make([]*Node, len(nodes)*2)
    any := false

    for _, n := range(nodes) {
        if n == nil {
            continue
        }
        newNodes = append(newNodes, n.Left)
        newNodes = append(newNodes, n.Right)
        any = true
        fmt.Printf("%v ", n.Data)
    }

    if !any {
        return
    }

    fmt.Println()
    printRecurse(newNodes)
}

func (bst *BST) Insert(data any) {
    bst.insertRec(bst.Root, data)
}

func (bst *BST) insertRec(node *Node, data any) {
    if bst.Cmp(data, node.Data) < 0 {
        nextNode := node.Left
        if nextNode == nil {
            node.Left = NewNode(data)
        } else {
            bst.insertRec(nextNode, data)
        }
    } else {
        nextNode := node.Right
        if nextNode == nil {
            node.Right = NewNode(data)
        } else {
            bst.insertRec(nextNode, data)
        }
    }
}

func SortedArrToBalancedBST(arr []any) *BST {
    bst := &BST{
        nil,
        nil,
    }

    bst.Root = bst.fromSortedArrRec(arr)
    return bst
}

func (bst *BST) fromSortedArrRec(arr []any) *Node {
    if len(arr) == 0 {
        return nil
    }

    iMid := len(arr) / 2
    root := NewNode(arr[iMid])
    root.Left = bst.fromSortedArrRec(arr[:iMid])
    root.Right = bst.fromSortedArrRec(arr[iMid+1:])

    return root
}

