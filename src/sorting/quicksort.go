package sorting

import ct "jy.org/csgo/src/custtypes"

type QuickSorter struct{}

func (b *QuickSorter) GetName() string {
    return "Quick Sort"
}

func (b *QuickSorter) Sort(arr *[]ct.Comparable) {
    if len(*arr) <= 1 {
        return
    }
    pivot := (*arr)[len(*arr) / 2]

    larr := make([]ct.Comparable, 0)
    marr := make([]ct.Comparable, 0)
    garr := make([]ct.Comparable, 0)
    for _, e := range(*arr) {
        if e == pivot {
            marr = append(marr, e)
        } else if e.CompareTo(pivot) > 0 {
            garr = append(garr, e)
        } else {
            larr = append(larr, e)
        }
    }

    b.Sort(&larr)
    b.Sort(&garr)

    *arr = append(larr, marr...)
    *arr = append(*arr, garr...)
}

