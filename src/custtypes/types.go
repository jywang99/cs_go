package custtypes

import "fmt"

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

func (ci CompInt) ToString() string {
    return fmt.Sprintf("%d", ci.GetData())
}

func (ci CompInt) GetData() any {
    return ci.Data
}

