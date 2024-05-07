package graph

import (
	"fmt"
)

type DGraph struct {
    vtcs int
    adjMap map[int][]int
}

func NewDGraph(vtcs int, adjList [][2]int) *DGraph {
    adjMap := make(map[int][]int)

    for _, e := range(adjList) {
        ee, exists := adjMap[e[0]]
        if !exists {
            ee = make([]int, 1)
            ee[0] = e[1]
        } else {
            ee = append(ee, e[1])
        }
        adjMap[e[0]] = ee
    }

    return &DGraph{
        vtcs: vtcs,
        adjMap: adjMap,
    }
}

func (g *DGraph) PrintEdges() {
    for from, to := range(g.adjMap) {
        fmt.Printf("%v -> ", from)
        for _, e := range to {
            fmt.Printf("%v ", e)
        }
        fmt.Println()
    }
}

type DWGraph struct {
    size int
    adjMtx [][]*int
}

func NewDWGraph(size int) *DWGraph {
    g := DWGraph{
        size: size,
        adjMtx: make([][]*int, size),
    }
    for r := 0; r < len(g.adjMtx); r++  {
        g.adjMtx[r] = make([]*int, size)
    }
    return &g
}

func (g *DWGraph) PrintAdjMtx() {
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

func (g *DWGraph) AddEdge(s, t, w int) bool {
    if !g.isValidEdge(s, t) {
        return false
    }
    g.adjMtx[s][t] = &w
    return true
}

func (g *DWGraph) isValidEdge(s, t int) bool {
    validVtx := func (v int) bool {
        return 0 <= v && v < g.size
    }
    return s != t && validVtx(s) && validVtx(t)
}
