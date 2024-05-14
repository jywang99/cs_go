package binarytree

import (
	"sort"

	ll "jy.org/csgo/src/linkedlist"
)

type Node struct {
    Data any
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

func (node *Node) TraverseVertical() map[int][]any {
    lvlMap := make(map[int][]any)
    node.traverseVertRecurse(0, &lvlMap)

    keys := make([]int, 0, len(lvlMap))
    for key := range lvlMap {
        keys = append(keys, key)
    }
    sort.Ints(keys)

    return lvlMap
}

func (node *Node) traverseVertRecurse(curLvl int, lvlMap *map[int][]any) {
    if node == nil {
        return
    }

    if _, ok := (*lvlMap)[curLvl]; !ok {
        (*lvlMap)[curLvl] = []any{node.Data}
    } else {
        (*lvlMap)[curLvl] = append((*lvlMap)[curLvl], node.Data)
    }

    node.Left.traverseVertRecurse(curLvl-1, lvlMap)
    node.Right.traverseVertRecurse(curLvl+1, lvlMap)
}

func (root *Node) FindLCA(d1 int, d2 int) any {
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

func (n *Node) ToDoublyLinkedList() ll.DoublyLinkedList {
    dll := n.toDllRecurse()
    head := dll.Head()
    return *head
}

func (n *Node) toDllRecurse() *ll.DoublyLinkedList {
    if n == nil {
        return nil
    }

    dll := ll.NewDoublyLinkedList(n.Data)

    lseg := n.Left.toDllRecurse()
    if lseg != nil {
        dll.Head().PrependAll(lseg)
    }

    rseg := n.Right.toDllRecurse()
    if rseg != nil {
        dll.Tail().AppendAll(rseg)
    }

    return dll
}

