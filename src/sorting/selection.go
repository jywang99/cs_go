package sorting

type SelectionSorter struct{}

func (b *SelectionSorter) GetName() string {
    return "Selection Sort"
}

func (b *SelectionSorter) Sort(arr *[]int) {
    a := *arr
    for i:=0; i<len(a); i++ {
        mi := i
        for j := i; j<len(a); j++ {
            if a[j] < a[mi] {
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

