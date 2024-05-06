package stack

type Stack []any

func (s *Stack) Push(a any) {
    *s = append(*s, a)
}

func (s *Stack) Pop() any {
    st := *s
    if len(st) == 0 {
        return nil
    }
    a := (st)[len(st)-1]
    *s = st[:len(st)-1]
    return a
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) Contains(d any) bool {
    for _, v := range *s {
        if v == d {
            return true
        }
    }
    return false
}

