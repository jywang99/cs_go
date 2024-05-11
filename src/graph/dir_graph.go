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

