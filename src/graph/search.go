package graph

import "jy.org/csgo/src/stack"

func (g *UGraph) BFS(s int) []int {
    visited := make([]bool, g.vtcs)
    rs := make([]int, g.vtcs)

    visited[0] = true
    rs[0] = s
    ri := 1
    g.bfsRecurse(s, &visited, &rs, &ri)

    return rs
}

func (g *UGraph) bfsRecurse(f int, visited *[]bool, rs *[]int, ri *int) {
    ns := g.getNeighbors(f)
    nns := make([]int, 0)
    for _, n := range ns {
        if (*visited)[n] {
            continue
        }
        (*visited)[n] = true
        (*rs)[*ri] = n
        (*ri) ++
        nns = append(nns, n)
    }
    for _, n := range nns {
        g.bfsRecurse(n, visited, rs, ri)
    }
}

func (g *UGraph) getNeighbors(f int) []int {
    rs := make([]int, 0)
    adjRow := g.adjMtx[f]
    for i:=0; i<len(adjRow); i++ {
        if adjRow[i] {
            rs = append(rs, i)
        }
    }
    return rs
}

func (g *UGraph) DFS(s int) []int {
    visited := make([]bool, g.vtcs)
    rs := make([]int, g.vtcs)
    ri := 0
    nstack := stack.Stack{}

    nstack.Push(s)
    for ok := true; ok; ok = !nstack.IsEmpty() {
        f := nstack.Pop().(int)
        visited[f] = true
        rs[ri] = f
        ri ++

        for _, i := range(g.getNeighbors(f)) {
            if !visited[i] && !nstack.Contains(i) {
                nstack.Push(i)
            }
        }
    }

    return rs
}

func (g *UGraph) DFSRecurse(s int) []int {
    visited := make([]bool, g.vtcs)
    rs := make([]int, g.vtcs)
    ri := 0
    g.dfsRecurse(s, &visited, &rs, &ri)
    return rs
}

func (g *UGraph) dfsRecurse(f int, visited *[]bool, rs *[]int, ri *int) {
    (*visited)[f] = true
    (*rs)[*ri] = f
    (*ri) ++

    for _, i := range g.getNeighbors(f) {
        if !(*visited)[i] {
            g.dfsRecurse(i, visited, rs, ri)
        }
    }
}
