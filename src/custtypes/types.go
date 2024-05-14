package custtypes

type Comparable interface {
    CompareTo(interface{}) int
    ToString() string
}

