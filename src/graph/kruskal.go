package graph

import (
	"errors"

	"jy.org/csgo/src/heap"
)

type edgeNode struct {
    vs [2]int
    w int
}

func cmpEdgeNode(as, at any) int {
    es := as.(*edgeNode)
    et := at.(*edgeNode)
    return es.w - et.w
}

func (g* WGraph) Kruskal() (*WGraph, error) {
    if g.dir {
        return nil, errors.New("Directional graph not supported!")
    }

    // earr := make([]*edgeNode, 0)
    // for ri, row := range g.adjMtx {
    //     for ci, w := range row {
    //         earr = append(earr, &edgeNode{
    //             vs: [2]int{ri, ci},
    //             w: *w,
    //         })
    //     }
    // }

    h := heap.NewMinHeap(len(g.adjMtx)*len(g.adjMtx[0]), cmpEdgeNode)
    g.traverseEdgesAndDo(
        func(src, dst, w int) {
            h.Insert(&edgeNode{
                vs: [2]int{src, dst},
                w: w,
            })
        },
    )

    kg := NewWGraph(g.size, false)
    for ; !kg.IsConnected(); {
        en := h.ExtractMin().(*edgeNode)
        s := en.vs[0]
        t := en.vs[1]
        if (!kg.HasPath(s, t)) {
            kg.AddEdge(s, t, en.w)
        }
    }

    return kg, nil
}

func (g* WGraph) IsConnected() bool {
    visited := make([]bool, g.size)

    visited[0] = true
    g.traverseEdgesAndDo(
        func(src, dst, _ int) {
            if visited[src] {
                visited[dst] = true
            }
        },
    )

    for _, b := range visited {
        if !b {
            return false
        }
    }
    return true
}

func (g* WGraph) traverseEdgesAndDo(f func(src, dst, w int)) {
    for ri, row := range g.adjMtx {
        for ci, w := range row {
            if w != nil {
                f(ri, ci, *w)
            }
        }
    }
}

