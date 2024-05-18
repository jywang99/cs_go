package graph

import (
	"errors"
	"fmt"
)

func (g *WGraph) BellmanFord(s int) ([]int, error) {
    ds := make([]*int, g.size)
    d := 0
    ds[s] = &d

    // relax n-1 times
    for i:=0; i<g.size-1; i++ {
        g.bfRelax(&ds)
    }

    // copy ds into rs
    // ds is relaxed one additional time
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

func (g *WGraph) bfRelax(dsp *[]*int) {
    ds := *dsp
    // for each vertex
    for v:=0; v<g.size; v++ {
        // distance to this vertex not calculated yet, skip
        if ds[v] == nil {
            continue
        }
        // update distance for each of its neighbors
        for vt, e := range g.adjMtx[v] {
            // not neighbor to vt
            if e == nil || v == vt {
                continue
            }
            // distance from start (s) to source (v) + edge weight
            nd := *ds[v] + (*e).GetWeight()
            // update if found the path first time or shorter than original distance
            if ds[vt] == nil || *ds[vt] > nd {
                ds[vt] = &nd
            }
        }
    }
}

type DijResult struct {
    Dist int
    Prev *int
}

func (g *WGraph) Dijkstra(s int) ([]*DijResult, error) {
    // make sure there's no negative edge
    for _, r := range g.adjMtx {
        for c, w := range r {
            if w != nil && (*w).GetWeight() < 0 {
                return nil, errors.New(fmt.Sprintf("Negative edge found! src=%v, dst=%v, w=%v", r, c, *w))
            }
        }
    }

    // visited vertices
    visited := make([]bool, g.size)
    // shortest distance found
    ds := make([]*DijResult, g.size)
    // initialize starting vertex
    ds[s] = &DijResult{ Dist: 0 }

    allVisited := func() bool {
        for _, v := range visited {
            if !v {
                return false
            }
        }
        return true
    }

    // get closest vertex that is not visited and has known distance
    getClosestUnvisitedV := func() (int, error) {
        var cv *int
        var md *int
        for i, d := range ds {
            // already visited or distance unknown
            if visited[i] || d == nil {
                continue
            }
            if md == nil || (*d).Dist < *md {
                ci := i
                cv = &ci
                md = &(*d).Dist
            }
        }

        if cv == nil {
            return -1, errors.New("Could not find closest unvisited vertex!")
        }

        return *cv, nil
    }

    // repeat until all vertices are visited
    for ; !allVisited(); {
        // pick the vertex with shortest known distance
        v, err := getClosestUnvisitedV()
        if err != nil {
            return nil, err
        }

        visited[v] = true
        // update distances for its unvisited neighbors
        for vt, w := range g.adjMtx[v] {
            // not neighbor
            if w == nil {
                continue
            }
            nd := (*ds[v]).Dist + (*w).GetWeight()
            if ds[vt] == nil || nd < (*ds[vt]).Dist {
                ds[vt] = &DijResult{nd, &v}
            }
        }
    }

    return ds, nil
}

