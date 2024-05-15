package sorting

import "jy.org/csgo/src/custtypes"

type Sorter interface {
    GetName() string
    Sort(arr *[]custtypes.Comparable)
}

