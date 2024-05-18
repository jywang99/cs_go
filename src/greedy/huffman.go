package greedy

import (
	"errors"
	"fmt"

	"jy.org/csgo/src/heap"
)

type HuffmanTree struct {
    head *huffNode
}

type huffNode struct {
    ch *rune
    freq int
    left *huffNode
    right *huffNode
}

func cmpHuffmanNode(c, ct any) int {
    h := c.(*huffNode)
    ht := ct.(*huffNode)
    if h.freq == ht.freq {
        return 0
    }
    if h.freq > ht.freq {
        return 1
    }
    return -1
}

func NewHuffmanEncode(src string) (string, *HuffmanTree, error) {
    // make frequency table
    freq := make(map[rune]int)
    for _, c := range src {
        f, e := freq[c]
        if e {
            freq[c] = f + 1
        } else {
            freq[c] = 1
        }
    }

    // create a min heap
    l := len(freq)
    h := heap.NewMinHeap(l, cmpHuffmanNode)
    for k, v := range(freq) {
        kc := k
        h.Insert(&huffNode{ch: &kc, freq:v})
    }

    // create huffman tree
    for ; h.Size()>1; {
        ns := h.ExtractMin().(*huffNode)
        nl := h.ExtractMin().(*huffNode)
        np := &huffNode{
            freq: ns.freq + nl.freq,
            left: ns,
            right: nl,
        }
        h.Insert(np)
    }

    head := h.GetMin().(*huffNode)
    ht := HuffmanTree{head: head}

    encoded, err := ht.Encode(src)
    if err != nil {
        return "", nil, err
    }
    return encoded, &ht, nil
}

func (ht *HuffmanTree) Encode(src string) (string, error) {
    // get hashcodes from heap
    hashes := make(map[rune]string)
    getHashcodes(ht.head, "", &hashes)

    // encode
    encoded := ""
    for _, c := range src {
        code, e := hashes[c]
        if !e {
            return "", errors.New(fmt.Sprintf("Character %v could not be encoded!", string(c)))
        }
        encoded += code
    }

    return encoded, nil
}

func getHashcodes(h *huffNode, path string, cmap *map[rune]string) {
    if h == nil {
        return
    }
    if h.ch != nil {
        (*cmap)[*h.ch] = path
    }
    getHashcodes(h.left, path + "0", cmap)
    getHashcodes(h.right, path + "1", cmap)
}

func (ht *HuffmanTree) Decode(src string) string {
    spos := 0
    decoded := ""
    for ; spos < len(src); {
        n := ht.head
        decoded += string(ht.decodeFirstChar(&src, n, &spos))
    }
    return decoded
}

func (ht *HuffmanTree) decodeFirstChar(src *string, n *huffNode, spos *int) rune {
    if n.left == nil && n.right == nil {
        return *n.ch
    }
    goleft := (*src)[*spos] == '0'
    (*spos) ++
    if goleft {
        return ht.decodeFirstChar(src, n.left, spos)
    }
    return ht.decodeFirstChar(src, n.right, spos)
}

