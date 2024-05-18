package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gh "jy.org/csgo/src/graph"
)

func TestCreateDGraph(t *testing.T) {
    edges := [][2]int {
        {0, 1},
        {1, 2},
        {2, 0},
    }
    g := gh.NewDGraph(3, edges)
    // fmt.Println("Graph 1:")
    // g.PrintEdges()

    edges = [][2]int {
        {0, 1},
        {1, 2},
        {1, 3},
        {2, 3},
        {3, 0},
    }
    g = gh.NewDGraph(3, edges)
    // fmt.Println("Graph 2:")
    // g.PrintEdges()
    assert.NotNil(t, g)
}

func TestCreateUGraph(t *testing.T) {
    edges := [][2]int {
        {0, 1},
        {0, 2},
        {1, 2},
        {1, 3},
        {3, 4},
        {2, 4},
    }
    g, _ := gh.NewUGraph(5, edges)
    assert.NotNil(t, g)
    // g.PrintAdjMtx()
}

func TestBfs(t *testing.T) {
    g := gh.NewUGraphNoEdge(5)
    g.AddEdge(0, 1);
    g.AddEdge(0, 2);
    g.AddEdge(1, 3);
    g.AddEdge(1, 4);
    g.AddEdge(2, 4);
    // g.PrintAdjMtx()

    expArr := []int{ 0, 1, 2, 3, 4 }
    assert.Equal(t, expArr, g.BFS(0))
}

func TestDfs(t *testing.T) {
    g := gh.NewWGraph(5, false)
    g.AddEdge(0, 1, gh.NewIntEdge(0))
    g.AddEdge(0, 2, gh.NewIntEdge(0))
    g.AddEdge(0, 3, gh.NewIntEdge(0))
    g.AddEdge(2, 3, gh.NewIntEdge(0))
    g.AddEdge(2, 4, gh.NewIntEdge(0))

    exps := [][]int {
        {0, 3, 2, 4, 1},
        {0, 1, 2, 4, 3},
        {0, 1, 2, 4, 3},
        {0, 1, 2, 3, 4},
    }

    dfsRs := g.DFS(0)
    assert.Contains(t, exps, dfsRs)
    dfsRs = g.DFSRecurse(0)
    assert.Contains(t, exps, dfsRs)
}

func TestCreateWeightedGraph(t *testing.T) {
    g := gh.NewWGraph(6, false)
    // g.PrintAdjMtx()

    g.AddEdge(0, 1, gh.NewIntEdge(5))
    g.AddEdge(1, 2, gh.NewIntEdge(1))
    g.AddEdge(1, 3, gh.NewIntEdge(2))
    g.AddEdge(2, 4, gh.NewIntEdge(1))
    g.AddEdge(4, 3, gh.NewIntEdge(-1))
    g.AddEdge(3, 5, gh.NewIntEdge(2))
    g.AddEdge(5, 4, gh.NewIntEdge(-3))
    // g.PrintAdjMtx()

    assert.True(t, g.HasEdge(0, 1))
    assert.True(t, g.HasEdge(1, 2))
    assert.False(t, g.HasEdge(0, 5))
}

func TestBellmanFord(t *testing.T) {
    g := gh.NewWGraph(6, true)

    g.AddEdge(0, 1, gh.NewIntEdge(10))
    g.AddEdge(0, 5, gh.NewIntEdge(8))
    g.AddEdge(1, 3, gh.NewIntEdge(2))
    g.AddEdge(2, 1, gh.NewIntEdge(1))
    g.AddEdge(3, 2, gh.NewIntEdge(-2))
    g.AddEdge(4, 3, gh.NewIntEdge(-1))
    g.AddEdge(4, 1, gh.NewIntEdge(-4))
    g.AddEdge(5, 4, gh.NewIntEdge(1))

    ds, err := g.BellmanFord(0)
    assert.Nil(t, err)
    assert.Equal(t, []int{0,5,5,7,9,8}, ds)

    g.AddEdge(3, 4, gh.NewIntEdge(1)) // create negative cycle
    ds, err = g.BellmanFord(0)
    assert.NotNil(t, err)
}

func TestDijkstra(t *testing.T) {
    g := gh.NewWGraph(5, false)
    g.AddEdge(0, 1, gh.NewIntEdge(6))
    g.AddEdge(0, 3, gh.NewIntEdge(1))
    g.AddEdge(1, 3, gh.NewIntEdge(2))
    g.AddEdge(1, 4, gh.NewIntEdge(2))
    g.AddEdge(1, 2, gh.NewIntEdge(5))
    g.AddEdge(2, 4, gh.NewIntEdge(5))
    g.AddEdge(4, 3, gh.NewIntEdge(1))

    ds, err := g.Dijkstra(0)
    assert.Nil(t, err)

    d := *ds[0]
    assert.Equal(t, 0, d.Dist)
    var nili *int
    assert.Equal(t, nili, d.Prev)
    d = *ds[1]
    assert.Equal(t, 3, d.Dist)
    assert.Equal(t, 3, *d.Prev)
    d = *ds[2]
    assert.Equal(t, 7, d.Dist)
    assert.Equal(t, 4, *d.Prev)
    d = *ds[3]
    assert.Equal(t, 1, d.Dist)
    assert.Equal(t, 0, *d.Prev)
    d = *ds[4]
    assert.Equal(t, 2, d.Dist)
    assert.Equal(t, 3, *d.Prev)
}
