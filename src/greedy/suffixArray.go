package greedy

import (
	"strings"

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

func cmpStr(cs, cb any) int {
    c := cs.(compStr)
    ct := cb.(compStr)
    hs := c.data[0]
    ht := ct.data[0]
    if hs == ht {
        hs := len(c.data)
        ht := len(ct.data)
        return hs - ht
    }
    if hs < ht {
        return -1
    }
    return 1
}

func NewSuffixArray(src string) *suffixArray {
    subs := make([]any, len(src))
    for i := range src {
        subs[i] = compStr{
            data: src[i:],
            idx: i,
        }
    }
    sort := sorting.QuickSorter{}
    sort.Sort(&subs, cmpStr)

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

