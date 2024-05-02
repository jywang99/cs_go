package sorting

type QuickSorter struct{}

func (b *QuickSorter) GetName() string {
    return "Quick Sort"
}

func (b *QuickSorter) Sort(arr *[]int) {
    if len(*arr) <= 1 {
        return
    }
    pivot := (*arr)[len(*arr) / 2]

    larr := make([]int, 0)
    marr := make([]int, 0)
    garr := make([]int, 0)
    for _, e := range(*arr) {
        if e == pivot {
            marr = append(marr, e)
        } else if e > pivot {
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

