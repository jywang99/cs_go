package graph

import (
	"errors"

	"jy.org/csgo/src/misc"
)

type flowEdge struct {
    cap int
    flow int
}

func (f flowEdge) GetWeight() int {
    return f.cap
}

func (g *WGraph) FordFulkerson(s, t int) (int, error) {
    if !g.dir {
        return 0, errors.New("Graph has to be directional!")
    }

    // residual graph
    fg := NewWGraph(g.size, true)
    g.traverseEdgesAndDo(func(src, dst int, we *WEdge) {
        // forward edge
        fg.AddEdge(src, dst, &flowEdge{
            cap: (*we).GetWeight(),
            // flow: 0
        })
        // residual edge
        fg.AddEdge(dst, src, &flowEdge{
            flow: -(*we).GetWeight(),
            // cap: 0
        })
    })

    sflow := 0
    fg.DoForEveryPath(s, t, func(path []int) {
        // get the edges and bottleneck flow value
        edges := make([]*flowEdge, len(path)-1)
        var bflow *int
        for i:=0; i<len(path)-1; i++ {
            // add edge to list
            e := (*fg.adjMtx[path[i]][path[i+1]]).(*flowEdge)
            edges[i] = e

            // get flow for current edge
            ef := e.cap - e.flow
            if ef <= 0 {
                // no capacity left for one of the edges = invalid path
                return
            }

            // update bottleneck flow
            if bflow == nil {
                bflow = &ef
            } else {
                nbf := min(*bflow, ef)
                bflow = &nbf
            }
        }

        // udpate flow value for each edge traversed
        for _, e := range edges {
            nf := e.flow + *bflow
            if nf > e.cap {
                // new flow value exceeds cap for one of the edges = invalid path
                return
            }
            e.flow = nf
        }

        // sum bottleneck value
        sflow += *bflow
    })

    return sflow, nil
}

func (g *WGraph) DoForEveryPath(s, t int, do func (path []int)) {
    path := make([]int, 0)
    g.dfsAndDoRec(s, t, path, do)
}

func (g *WGraph) dfsAndDoRec(n, t int, path []int, do func (path []int)) {
    path = append(path, n)
    if n == t {
        do(path)
        return
    }
    if len(path) == g.size {
        return
    }

    for nn, e := range g.adjMtx[n] {
        if misc.BinarySearch(path, nn) >= 0 || e == nil {
            continue
        }
        g.dfsAndDoRec(nn, t, path, do)
    }
}
