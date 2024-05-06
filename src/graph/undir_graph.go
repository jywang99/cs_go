package graph

import (
	"errors"
	"fmt"
)

type UGraph struct {
    vtcs int
    adjMtx [][]bool
}

func NewUGraph(vtcs int, edges [][2]int) (*UGraph, error) {
    adjMap := make([][]bool, vtcs)
    for i := range adjMap {
        adjMap[i] = make([]bool, vtcs)
    }

    g := UGraph{
        vtcs: vtcs,
        adjMtx: adjMap,
    }

    for _, e := range edges {
        g.AddEdge(e[0], e[1])
    }

    return &g, nil
}

func NewUGraphNoEdge(vtcs int) *UGraph {
    adjMap := make([][]bool, vtcs)
    for i := range adjMap {
        adjMap[i] = make([]bool, vtcs)
    }

    return &UGraph{
        vtcs: vtcs,
        adjMtx: adjMap,
    }
}

func (g *UGraph) AddEdge(a, b int) error {
    if !g.validVtx(a) || !g.validVtx(b) {
        return errors.New(fmt.Sprintf("Cannot add edge {%v, %v}", a, b))
    }

    g.adjMtx[a][b] = true
    g.adjMtx[b][a] = true

    return nil
}

func (g *UGraph) validVtx(i int) bool {
    return 0 <= i && i < g.vtcs
}

func (g *UGraph) PrintAdjMtx() {
    fmt.Println("Adjacency matrix:")
    fmt.Print("  ")
    for i:=0; i<g.vtcs; i++ {
        fmt.Printf("%v ", i)
    }
    fmt.Println()
    for i:=0; i<g.vtcs; i++ {
        fmt.Printf("%v ", i)
        for _, v := range g.adjMtx[i] {
            var val string
            if v {
                val = "1"
            } else {
                val = " "
            }
            fmt.Printf("%v ", val)
        }
        fmt.Println()
    }
}

