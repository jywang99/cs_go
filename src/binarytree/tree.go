package binarytree

import (
	"fmt"
	"sort"
)

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

func (node *Node) GetDepth() int {
    if node == nil {
        return 0
    }
    return 1 + max(node.Left.GetDepth(), node.Right.GetDepth())
}

func (node *Node) TraverseVertical() map[int][]int {
    lvlMap := make(map[int][]int)
    node.traverseVertRecurse(0, &lvlMap)

    keys := make([]int, 0, len(lvlMap))
    for key := range lvlMap {
        keys = append(keys, key)
    }
    sort.Ints(keys)

    for _, key := range keys {
        fmt.Printf("%v\n", lvlMap[key])
    }

    return lvlMap
}

func (node *Node) traverseVertRecurse(curLvl int, lvlMap *map[int][]int) {
    if node == nil {
        return
    }

    if _, ok := (*lvlMap)[curLvl]; !ok {
        (*lvlMap)[curLvl] = []int{node.Data}
    } else {
        (*lvlMap)[curLvl] = append((*lvlMap)[curLvl], node.Data)
    }

    node.Left.traverseVertRecurse(curLvl-1, lvlMap)
    node.Right.traverseVertRecurse(curLvl+1, lvlMap)
}

func (root *Node) FindLCA(d1 int, d2 int) int {
    path1 := root.findPath(d1, []*Node{})
    path2 := root.findPath(d2, []*Node{})

    overlap := []*Node{}
    for _, n := range(path1) {
        if contains(path2, n) {
            overlap = append(overlap, n)
        }
    }

    return overlap[len(overlap)-1].Data
}

func (n *Node) findPath(d int, path []*Node) []*Node {
    if n == nil {
        return []*Node{}
    }

    if n.Data == d {
        return append(path, n)
    }

    newPath := n.Left.findPath(d, path)
    if len(newPath) == 0 {
        newPath = n.Right.findPath(d, path)
    }

    if len(newPath) == 0 {
        return []*Node{}
    }
    return append([]*Node{n}, newPath...)
}

func contains(s []*Node, t *Node) bool {
    for _, e := range(s) {
        if e == t {
            return true
        }
    }
    return false
}

