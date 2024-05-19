package graph

import (
	"errors"
)

type edgeInfo struct {
    src int
    dst int
    edge *WEdge
}

func (g *WGraph) Prim() (*WGraph, error) {
    if !g.IsConnected() {
        return nil, errors.New("Graph must be connected!")
    }

    visit, allVisited, visited := g.getVisitTracker()
    visit(0)

    mst := NewWGraph(g.size, g.dir)
    for ; !allVisited(); {
        // find 
        //   the least weight of edge outgoing to an unvisited vertex
        //   the newly visited vertex
        var medge *edgeInfo
        // for all visited vertices
        for s:=0; s<g.size; s++ {
            if !visited(s) {
                continue
            }

            // for all its unvisited outgoing edges
            g.DoForOutgoingEdges(s, func(t int, we *WEdge) {
                if visited(t) {
                    return
                }

                // update minimum weighted edge found
                einf := &edgeInfo{
                    src: s,
                    dst: t,
                    edge: we,
                }
                if medge == nil {
                    medge = einf
                } else {
                    if (*we).GetWeight() < (*medge.edge).GetWeight() {
                        medge = einf
                    }
                }
            })
        }

        // error handling
        if medge == nil {
            return nil, errors.New("Graph must be connected!")
        }

        // mark vertex visited
        visit(medge.dst)

        // update MST
        mst.AddEdge(medge.src, medge.dst, *medge.edge)
    }

    return mst, nil
}

