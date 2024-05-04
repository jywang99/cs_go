package sorting

type MergeSorter struct{}

func (b *MergeSorter) GetName() string {
    return "Merge Sort"
}

func (b *MergeSorter) Sort(arr *[]int) {
    a := *arr
    if len(a) <= 1 {
        return
    }
    if len(a) == 2 {
        if a[0] > a[1] {
            tmp := a[0]
            a[0] = a[1]
            a[1] = tmp
        }
    }

    mi := len(a) / 2
    larr := make([]int, mi)
    rarr := make([]int, len(a) - mi)
    copy(larr, a[:mi])
    copy(rarr, a[mi:])

    b.Sort(&larr)
    b.Sort(&rarr)

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
        if rv < lv {
            a[i] = rv
            ri ++
        } else {
            a[i] = lv
            li ++
        }
    }
}

