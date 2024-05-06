package binarytree

import "math"

func (node *Node) GetDepth() int {
    if node == nil {
        return 0
    }
    return 1 + max(node.Left.GetDepth(), node.Right.GetDepth())
}

func (n *Node) IsFull() bool {
    ln := n.Left == nil
    rn := n.Right == nil
    if ln && rn {
        return true
    }
    if ln || rn {
        return false
    }
    return n.Left.IsFull() && n.Right.IsFull()
}

func (n *Node) IsComplete() bool {
    ln := n.Left == nil
    rn := n.Right == nil
    if ln && rn {
        return true
    }
    if ln && !rn {
        return false
    }
    rcomp := true
    if !rn {
        rcomp = n.Right.IsComplete()
    }
    return n.Left.IsComplete() && rcomp
}

func (n *Node) IsBalanced() bool {
    if n.Left == nil && n.Right == nil {
        return true
    }

    ldepth := n.Left.GetDepth()
    rdepth := n.Right.GetDepth()
    if math.Abs(float64(ldepth - rdepth)) > 1 {
        return false
    }

    return n.Left.IsBalanced() && n.Right.IsBalanced()
}

