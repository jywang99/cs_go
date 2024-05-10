package graph

import "errors"

func (g *DWGraph) BellmanFord(s int) ([]int, error) {
    ds := make([]*int, g.size)
    d := 0
    ds[s] = &d

    // repeat n-1 times
    for i:=0; i<g.size-1; i++ {
        g.bfRelax(&ds)
    }

    rs := make([]int, g.size)
    for i, p := range ds {
        rs[i] = *p
    }

    g.bfRelax(&ds)
    for i:=0; i<g.size; i++ {
        if *ds[i] < rs[i] {
            return rs, errors.New("Negative cycle detected!")
        }
    }

    return rs, nil
}

func (g *DWGraph) bfRelax(dsp *[]*int) {
    ds := *dsp
    // for each vertex
    for v:=0; v<g.size; v++ {
        // distance to this vertex not calculated yet
        if ds[v] == nil {
            continue
        }
        // update distance for each of its neighbors
        for vt, p := range g.adjMtx[v] {
            // not neighbor to vt
            if p == nil || v == vt {
                continue
            }
            // distance from start (s) to source (v) + edge weight
            nd := *ds[v] + *p
            // update if found the path first time or shorter than original distance
            if ds[vt] == nil || *ds[vt] > nd {
                ds[vt] = &nd
            }
        }
    }
}

