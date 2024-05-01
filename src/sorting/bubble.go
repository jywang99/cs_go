package sorting

type BubbleSorter struct{}

func (b *BubbleSorter) GetName() string {
    return "Bubble Sort"
}

func (b *BubbleSorter) Sort(arr *[]int) {
    a := *arr
    for i := range a {
        for j, jval := range(a[:len(a)-i-1]) {
            if jval > a[j+1] {
                a[j] = a[j+1]
                a[j+1] = jval
            }
        }
    }
}

