package sorting

import ct "jy.org/csgo/src/custtypes"

type BubbleSorter struct{}

func (b *BubbleSorter) GetName() string {
    return "Bubble Sort"
}

func (b *BubbleSorter) Sort(arr *[]ct.Comparable) {
    a := *arr
    for i := range a {
        for j, jval := range(a[:len(a)-i-1]) {
            if jval.CompareTo(a[j+1]) > 0 {
                a[j] = a[j+1]
                a[j+1] = jval
            }
        }
    }
}

