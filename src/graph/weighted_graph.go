package graph

import "fmt"

type WEdge interface {
    GetWeight() int
}

type WGraph struct {
    size int
    adjMtx [][]*WEdge
    dir bool
}

type IntWEdge struct {
    w int
}

func NewIntEdge(w int) *IntWEdge {
    return &IntWEdge{w: w}
}

func (ie *IntWEdge) GetWeight() int {
    return ie.w
}

func NewWGraph(size int, dir bool) *WGraph {
    g := WGraph{
        size: size,
        adjMtx: make([][]*WEdge, size),
        dir: dir,
    }
    for r := 0; r < len(g.adjMtx); r++  {
        g.adjMtx[r] = make([]*WEdge, size)
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

func (g *WGraph) AddEdge(s, t int, e WEdge) bool {
    if !g.isValidEdgeToCreate(s, t) {
        return false
    }
    g.adjMtx[s][t] = &e
    if !g.dir {
        g.adjMtx[t][s] = &e
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
    has := g.adjMtx[s][t] != nil
    if g.dir {
        has = has || g.adjMtx[s][t] != nil
    }
    return has
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

