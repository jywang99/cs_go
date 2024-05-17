package sorting

import ct "jy.org/csgo/src/custtypes"

type Sorter interface {
    GetName() string
    Sort(arr *[]any, cmp ct.Comparator)
}

