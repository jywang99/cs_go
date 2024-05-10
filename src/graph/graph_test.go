package graph_test

import (
	// "fmt"

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
    g := gh.NewUGraphNoEdge(5)
    g.AddEdge(0, 1)
    g.AddEdge(0, 2)
    g.AddEdge(0, 3)
    g.AddEdge(2, 3)
    g.AddEdge(2, 4)

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
    g := gh.NewDWGraph(6)

    g.AddEdge(0, 1, 5)
    g.AddEdge(1, 2, 1)
    g.AddEdge(1, 3, 2)
    g.AddEdge(2, 4, 1)
    g.AddEdge(4, 3, -1)
    g.AddEdge(3, 5, 2)
    g.AddEdge(5, 4, -3)
    // g.PrintAdjMtx()
}

func TestBellmanFord(t *testing.T) {
    g := gh.NewDWGraph(6)

    g.AddEdge(0, 1, 10)
    g.AddEdge(0, 5, 8)
    g.AddEdge(1, 3, 2)
    g.AddEdge(2, 1, 1)
    g.AddEdge(3, 2, -2)
    g.AddEdge(4, 3, -1)
    g.AddEdge(4, 1, -4)
    g.AddEdge(5, 4, 1)

    ds, err := g.BellmanFord(0)
    assert.Nil(t, err)
    assert.Equal(t, []int{0,5,5,7,9,8}, ds)

    g.AddEdge(3, 4, 1) // create negative cycle
    ds, err = g.BellmanFord(0)
    assert.NotNil(t, err)
}
