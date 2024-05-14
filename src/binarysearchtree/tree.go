package binarysearchtree

import (
	"fmt"

	ct "jy.org/csgo/src/custtypes"
)

type Node struct {
    Data ct.Comparable
    Left *Node
    Right *Node
}

func NewNode(data ct.Comparable) *Node {
    return &Node{
        data,
        nil,
        nil,
    }
}

func (node *Node) Print() {
    printRecurse([]*Node{node})
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
        fmt.Printf("%s ", n.Data.ToString())
    }

    if !any {
        return
    }

    fmt.Println()
    printRecurse(newNodes)
}

func (node *Node) Insert(data ct.Comparable) {
    if data.CompareTo(&node.Data) < 0 {
        nextNode := node.Left
        if nextNode == nil {
            node.Left = NewNode(data)
        } else {
            nextNode.Insert(data)
        }
    } else {
        nextNode := node.Right
        if nextNode == nil {
            node.Right = NewNode(data)
        } else {
            nextNode.Insert(data)
        }
    }
}

func SortedArrToBalancedBST(arr []ct.Comparable) *Node {
    if len(arr) == 0 {
        return nil
    }

    iMid := len(arr) / 2
    root := NewNode(arr[iMid])
    root.Left = SortedArrToBalancedBST(arr[:iMid])
    root.Right = SortedArrToBalancedBST(arr[iMid+1:])

    return root
}

