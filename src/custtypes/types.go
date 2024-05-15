package custtypes

type Comparable interface {
    CompareTo(Comparable) int
}

type CompInt struct {
    Data int
}

func (ci CompInt) CompareTo(ct Comparable) int {
    cti := ct.(CompInt).Data
    if ci.Data == cti {
        return 0
    }
    if ci.Data < cti {
        return -1
    }
    return 1
}

