package sorting

import ct "jy.org/csgo/src/custtypes"

type QuickSorter struct{}

func (b *QuickSorter) GetName() string {
    return "Quick Sort"
}

func (b *QuickSorter) Sort(arr *[]any, cmp ct.Comparator) {
    if len(*arr) <= 1 {
        return
    }
    pivot := (*arr)[len(*arr) / 2]

    larr := make([]any, 0)
    marr := make([]any, 0)
    garr := make([]any, 0)
    for _, e := range(*arr) {
        if e == pivot {
            marr = append(marr, e)
        } else if cmp(e, pivot) > 0 {
            garr = append(garr, e)
        } else {
            larr = append(larr, e)
        }
    }

    b.Sort(&larr, cmp)
    b.Sort(&garr, cmp)

    *arr = append(larr, marr...)
    *arr = append(*arr, garr...)
}

