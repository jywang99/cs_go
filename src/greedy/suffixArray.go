package greedy

type SuffixArray struct {
    str string
    sa []int
}

func NewSuffixArray(src string) *SuffixArray {

    s := &SuffixArray{
        str: src, 
        sa: make([]int, len(src)),
    }

    return s
}
