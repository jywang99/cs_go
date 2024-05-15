package greedy

import (
	"strings"

	ct "jy.org/csgo/src/custtypes"
	"jy.org/csgo/src/sorting"
)

type suffixArray struct {
    str string
    sa []int
}

type compStr struct {
    data string
    idx int
}

func (cs compStr) CompareTo(cb ct.Comparable) int {
    ct := cb.(compStr)
    hs := cs.data[0]
    ht := ct.data[0]
    if hs == ht {
        hs := len(cs.data)
        ht := len(ct.data)
        return hs - ht
    }
    if hs < ht {
        return -1
    }
    return 1
}

func NewSuffixArray(src string) *suffixArray {
    subs := make([]ct.Comparable, len(src))
    for i := range src {
        subs[i] = compStr{
            data: src[i:],
            idx: i,
        }
    }
    sort := sorting.QuickSorter{}
    sort.Sort(&subs)

    sarr := &suffixArray{
        str: src, 
        sa: make([]int, len(src)),
    }

    for i, s := range subs {
        sarr.sa[i] = s.(compStr).idx
    }

    return sarr
}

func (sa *suffixArray) FindFirstOccr(src string) int {
    for _, s := range sa.sa {
        if strings.HasPrefix(sa.str[s:], src) {
            return s
        }
    }
    return -1
}

