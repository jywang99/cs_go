package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gh "jy.org/csgo/src/graph"
)

func TestHuffman(t *testing.T) {
    plain := "AAAAAABBCCDDEEFFFFF"
    encoded, ht, err := gh.NewHuffmanEncode(plain)
    assert.Nil(t, err)
    assert.Equal(t, 46, len(encoded)) // since order of traversing a map is not guaranteed, specific codes can be different
    assert.NotNil(t, ht)
    assert.Equal(t, plain, ht.Decode(encoded))

    plain = "Hello World!"
    encoded, ht, err = gh.NewHuffmanEncode(plain)
    assert.Equal(t, plain, ht.Decode(encoded))
}

func getExConnGraph() *gh.WGraph {
    eg := gh.NewWGraph(7, false)
    eg.AddEdge(0, 1, gh.NewIntEdge(2))
    eg.AddEdge(0, 3, gh.NewIntEdge(3))
    eg.AddEdge(0, 2, gh.NewIntEdge(3))
    eg.AddEdge(1, 2, gh.NewIntEdge(4))
    eg.AddEdge(1, 4, gh.NewIntEdge(3))
    eg.AddEdge(2, 3, gh.NewIntEdge(5))
    eg.AddEdge(2, 4, gh.NewIntEdge(1))
    eg.AddEdge(3, 5, gh.NewIntEdge(7))
    eg.AddEdge(4, 5, gh.NewIntEdge(8))
    eg.AddEdge(5, 6, gh.NewIntEdge(9))
    return eg
}

func getExDisconnGraph() *gh.WGraph {
    eg := gh.NewWGraph(7, false)
    eg.AddEdge(0, 1, gh.NewIntEdge(2))
    eg.AddEdge(0, 3, gh.NewIntEdge(3))
    eg.AddEdge(0, 2, gh.NewIntEdge(3))
    eg.AddEdge(1, 2, gh.NewIntEdge(4))
    eg.AddEdge(1, 4, gh.NewIntEdge(3))
    eg.AddEdge(2, 3, gh.NewIntEdge(5))
    eg.AddEdge(2, 4, gh.NewIntEdge(1))
    eg.AddEdge(3, 5, gh.NewIntEdge(7))
    eg.AddEdge(4, 5, gh.NewIntEdge(8))
    return eg
}

func TestConnected(t *testing.T) {
    eg := getExConnGraph()
    assert.True(t, eg.IsConnected())
    dg := getExDisconnGraph()
    assert.False(t, dg.IsConnected())
}

func TestKruskal(t *testing.T) {
    eg := getExConnGraph()
    mst, err := eg.Kruskal()
    assert.Nil(t, err)
    assert.NotNil(t, mst)

    assert.True(t, mst.HasEdge(0, 1))
    assert.True(t, mst.HasEdge(0, 3))
    assert.True(t, mst.HasEdge(0, 2))
    assert.False(t, mst.HasEdge(1, 2))
    assert.False(t, mst.HasEdge(1, 4))
    assert.False(t, mst.HasEdge(2, 3))
    assert.True(t, mst.HasEdge(2, 4))
    assert.True(t, mst.HasEdge(3, 5))
    assert.False(t, mst.HasEdge(4, 5))
    assert.True(t, mst.HasEdge(5, 6))
}

func TestHasPath(t *testing.T) {
    eg := getExDisconnGraph()
    assert.True(t, eg.HasPath(1, 5))
    assert.True(t, eg.HasPath(0, 4))
    assert.False(t, eg.HasPath(0, 6))
    assert.False(t, eg.HasPath(3, 6))
}

func getFulkersonGraph() *gh.WGraph {
    g := gh.NewWGraph(6, true)
    g.AddEdge(0, 1, gh.NewIntEdge(10))
    g.AddEdge(0, 2, gh.NewIntEdge(10))
    g.AddEdge(1, 3, gh.NewIntEdge(25))
    g.AddEdge(2, 4, gh.NewIntEdge(15))
    g.AddEdge(3, 5, gh.NewIntEdge(10))
    g.AddEdge(4, 1, gh.NewIntEdge(6))
    g.AddEdge(4, 5, gh.NewIntEdge(10))
    return g
}

func TestFordFulkerson(t *testing.T) {
    // graph 1
    g := getFulkersonGraph()

    f, err := g.FordFulkerson(0, 5)
    assert.Nil(t, err)
    assert.Equal(t, 20, f)

    // graph 2
    g2 := gh.NewWGraph(6, true)
    g2.AddEdge(0, 1, gh.NewIntEdge(8))
    g2.AddEdge(0, 4, gh.NewIntEdge(3))
    g2.AddEdge(1, 2, gh.NewIntEdge(9))
    g2.AddEdge(2, 5, gh.NewIntEdge(2))
    g2.AddEdge(3, 5, gh.NewIntEdge(5))
    g2.AddEdge(4, 3, gh.NewIntEdge(4))
    g2.AddEdge(4, 2, gh.NewIntEdge(7))

    f2, err2 := g2.FordFulkerson(0, 5)
    assert.Nil(t, err2)
    assert.Equal(t, 6, f2)
}

func TestDfsAndDo(t *testing.T) {
    g := getFulkersonGraph()

    paths := make([][]int, 0)
    g.DoForEveryPath(0, 5, func(path []int) {
        paths = append(paths, path)
    })

    assert.Equal(t, 3, len(paths))
    assert.Contains(t, paths, []int{0,1,3,5}, []int{0,2,4,1,3,5}, []int{0,2,4,5})
}

func TestPrim(t *testing.T) {
    g := gh.NewWGraph(8, false)

    g.AddEdge(0, 1, gh.NewIntEdge(4))
    g.AddEdge(0, 3, gh.NewIntEdge(3))
    g.AddEdge(1, 2, gh.NewIntEdge(3))
    g.AddEdge(1, 3, gh.NewIntEdge(5))
    g.AddEdge(1, 4, gh.NewIntEdge(6))
    g.AddEdge(2, 4, gh.NewIntEdge(4))
    g.AddEdge(2, 7, gh.NewIntEdge(2))
    g.AddEdge(3, 4, gh.NewIntEdge(7))
    g.AddEdge(3, 5, gh.NewIntEdge(4))
    g.AddEdge(4, 5, gh.NewIntEdge(5))
    g.AddEdge(4, 6, gh.NewIntEdge(3))
    g.AddEdge(5, 6, gh.NewIntEdge(7))
    g.AddEdge(6, 7, gh.NewIntEdge(5))

    mst, err := g.Prim()

    assert.Nil(t, err)
    assert.True(t, mst.HasEdge(0, 1))
    assert.True(t, mst.HasEdge(0, 3))
    assert.True(t, mst.HasEdge(1, 2))
    assert.False(t, mst.HasEdge(1, 3))
    assert.False(t, mst.HasEdge(1, 4))
    assert.True(t, mst.HasEdge(2, 4))
    assert.True(t, mst.HasEdge(2, 7))
    assert.False(t, mst.HasEdge(3, 4))
    assert.True(t, mst.HasEdge(3, 5))
    assert.False(t, mst.HasEdge(4, 5))
    assert.True(t, mst.HasEdge(4, 6))
    assert.False(t, mst.HasEdge(5, 6))
    assert.False(t, mst.HasEdge(6, 7))
}

