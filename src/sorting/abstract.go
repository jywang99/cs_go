package sorting

type Sorter interface {
    GetName() string
    Sort(arr *[]int)
}

