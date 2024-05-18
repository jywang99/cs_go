package graph

import "fmt"

type WGraph struct {
    size int
    adjMtx [][]*int
    dir bool
}

func NewWGraph(size int, dir bool) *WGraph {
    g := WGraph{
        size: size,
        adjMtx: make([][]*int, size),
        dir: dir,
    }
    for r := 0; r < len(g.adjMtx); r++  {
        g.adjMtx[r] = make([]*int, size)
    }
    return &g
}

func (g *WGraph) PrintAdjMtx() {
    fmt.Println("Adjacency matrix:")
    fmt.Print("  ")
    for i:=0; i<g.size; i++ {
        fmt.Printf("%v ", i)
    }
    fmt.Println()
    for i:=0; i<g.size; i++ {
        fmt.Printf("%v ", i)
        for _, v := range g.adjMtx[i] {
            var val any
            if v == nil {
                val = "-"
            } else {
                val = *v
            }
            fmt.Printf("%v ", val)
        }
        fmt.Println()
    }
}

func (g *WGraph) AddEdge(s, t, w int) bool {
    if !g.isValidEdgeToCreate(s, t) {
        return false
    }
    g.adjMtx[s][t] = &w
    if !g.dir {
        g.adjMtx[t][s] = &w
    }
    return true
}

func (g *WGraph) isValidEdgeToCreate(s, t int) bool {
    validVtx := func (v int) bool {
        return 0 <= v && v < g.size
    }
    return s != t && validVtx(s) && validVtx(t)
}

func (g *WGraph) HasEdge(s, t int) bool {
    return g.adjMtx[s][t] != nil
}

func (g *WGraph) HasPath(s, t int) bool {
    visited := make([]bool, g.size)
    return g.hasPathRec(s, t, &visited)
}

func (g *WGraph) hasPathRec(n, t int, visited *[]bool) bool {
    if n == t {
        return true
    }

    (*visited)[n] = true
    for nt, w := range g.adjMtx[n] {
        if (*visited)[nt] || w == nil || nt == n {
            continue
        }
        if g.hasPathRec(nt, t, visited) {
            return true
        }
    }

    return false
}

