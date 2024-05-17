package sorting

import ct "jy.org/csgo/src/custtypes"

type BubbleSorter struct{}

func (b *BubbleSorter) GetName() string {
    return "Bubble Sort"
}

func (b *BubbleSorter) Sort(arr *[]any, cmp ct.Comparator) {
    a := *arr
    for i := range a {
        for j, jval := range(a[:len(a)-i-1]) {
            if cmp(jval, a[j+1]) > 0 {
                a[j] = a[j+1]
                a[j+1] = jval
            }
        }
    }
}

