package sorting

import ct "jy.org/csgo/src/custtypes"

type MergeSorter struct{}

func (b *MergeSorter) GetName() string {
    return "Merge Sort"
}

func (b *MergeSorter) Sort(arr *[]any, cmp ct.Comparator) {
    a := *arr
    if len(a) <= 1 {
        return
    }
    if len(a) == 2 {
        if cmp(a[0], a[1]) > 0 {
            tmp := a[0]
            a[0] = a[1]
            a[1] = tmp
        }
    }

    mi := len(a) / 2
    larr := make([]any, mi)
    rarr := make([]any, len(a) - mi)
    copy(larr, a[:mi])
    copy(rarr, a[mi:])

    b.Sort(&larr, cmp)
    b.Sort(&rarr, cmp)

    ri := 0
    li := 0
    for i := range(a) {
        if li >= len(larr) {
            a[i] = rarr[ri]
            ri ++
            continue
        }
        if ri >= len(rarr) {
            a[i] = larr[li]
            li ++
            continue
        }

        rv := rarr[ri]
        lv := larr[li]
        if cmp(rv, lv) < 0 {
            a[i] = rv
            ri ++
        } else {
            a[i] = lv
            li ++
        }
    }
}

