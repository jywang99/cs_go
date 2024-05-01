package sorting

type InsertionSorter struct{}

func (b *InsertionSorter) GetName() string {
    return "Selection Sort"
}

func (b *InsertionSorter) Sort(arr *[]int) {
    a := *arr
    if len(a) <= 1 {
        return
    }
    for i:=1; i<len(a); i++ {
        ival := a[i]
        j := i-1
        for ; j>=0 && a[j]>ival; j-- {
            a[j+1] = a[j]
        }
        a[j+1] = ival
    }
}

