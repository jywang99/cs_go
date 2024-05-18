package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	g "jy.org/csgo/src/graph"
)

func TestHuffman(t *testing.T) {
    plain := "AAAAAABBCCDDEEFFFFF"
    encoded, ht, err := g.NewHuffmanEncode(plain)
    assert.Nil(t, err)
    assert.Equal(t, 46, len(encoded)) // since order of traversing a map is not guaranteed, specific codes can be different
    assert.NotNil(t, ht)
    assert.Equal(t, plain, ht.Decode(encoded))

    plain = "Hello World!"
    encoded, ht, err = g.NewHuffmanEncode(plain)
    assert.Equal(t, plain, ht.Decode(encoded))
}

func getExConnGraph() *g.WGraph {
    eg := g.NewWGraph(7, false)
    eg.AddEdge(0, 1, 2)
    eg.AddEdge(0, 3, 3)
    eg.AddEdge(0, 2, 3)
    eg.AddEdge(1, 2, 4)
    eg.AddEdge(1, 4, 3)
    eg.AddEdge(2, 3, 5)
    eg.AddEdge(2, 4, 1)
    eg.AddEdge(3, 5, 7)
    eg.AddEdge(4, 5, 8)
    eg.AddEdge(5, 6, 9)
    return eg
}

func getExDisconnGraph() *g.WGraph {
    eg := g.NewWGraph(7, false)
    eg.AddEdge(0, 1, 2)
    eg.AddEdge(0, 3, 3)
    eg.AddEdge(0, 2, 3)
    eg.AddEdge(1, 2, 4)
    eg.AddEdge(1, 4, 3)
    eg.AddEdge(2, 3, 5)
    eg.AddEdge(2, 4, 1)
    eg.AddEdge(3, 5, 7)
    eg.AddEdge(4, 5, 8)
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
}

func TestHasPath(t *testing.T) {
    eg := getExDisconnGraph()
    assert.True(t, eg.HasPath(1, 5))
    assert.True(t, eg.HasPath(0, 4))
    assert.False(t, eg.HasPath(0, 6))
    assert.False(t, eg.HasPath(3, 6))
}
