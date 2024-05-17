package greedy

import (
	"jy.org/csgo/src/heap"
)

type huffmanRune struct {
    c rune
    freq int
}

func cmpHuffmanRune(c, ct any) int {
    h := c.(*huffmanRune)
    ht := ct.(*huffmanRune)
    if h.freq == ht.freq {
        return 0
    }
    if h.freq > ht.freq {
        return 1
    }
    return -1
}

func HuffmanEncode(src string) (string, map[rune]string) {
    // make frequency table
    freq := make(map[rune]int)
    for _, c := range src {
        f, e := freq[c]
        if e {
            freq[c] = 1
        } else {
            freq[c] = f + 1
        }
    }

    // create a priority queue
    h := heap.NewMinHeap(len(src), cmpHuffmanRune)
    for k, v := range(freq) {
        h.Insert(&huffmanRune{c: k, freq:v})
    }

    // get hashcode from queue
    hashes := make(map[rune]string)

    // encode
    encoded := ""

    return encoded, hashes
}
