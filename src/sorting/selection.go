package sorting

import ct "jy.org/csgo/src/custtypes"

type SelectionSorter struct{}

func (b *SelectionSorter) GetName() string {
    return "Selection Sort"
}

func (b *SelectionSorter) Sort(arr *[]any, cmp ct.Comparator) {
    a := *arr
    for i:=0; i<len(a); i++ {
        mi := i
        for j := i; j<len(a); j++ {
            if cmp(a[j],a[mi]) < 0 {
                mi = j
            }
        }
        if mi != i {
            ival := a[i]
            a[i] = a[mi]
            a[mi] = ival
        }
    }
}

