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

