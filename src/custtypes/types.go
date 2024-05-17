package custtypes

type Comparator func(any, any) int

func IntComparator(a, b any) int {
    ai := a.(int)
    bi := b.(int)
    return ai-bi
}

