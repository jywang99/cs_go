package binarysearchtree

import "fmt"

type Node struct {
    Data int
    Left *Node
    Right *Node
}

func NewNode(data int) *Node {
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
        fmt.Printf("%d ", n.Data)
    }

    if !any {
        return
    }

    fmt.Println()
    printRecurse(newNodes)
}

func (node *Node) Insert(data int) {
    if data <= node.Data {
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

func SortedArrToBalancedBST(arr []int) *Node {
    if len(arr) == 0 {
        return nil
    }

    iMid := len(arr) / 2
    root := NewNode(arr[iMid])
    root.Left = SortedArrToBalancedBST(arr[:iMid])
    root.Right = SortedArrToBalancedBST(arr[iMid+1:])

    return root
}

