package misc

func BinarySearch(arr []int, t int) int {
    return bSearchRecurse(arr, 0, len(arr), t)
}

func bSearchRecurse(arr []int, s, e, t int) int {
    if s >= e {
        return -1
    }

    m := (e + s) / 2
    if arr[m] == t {
        return m
    } else {
        if e - s == 1 {
            return -1
        }
    }

    if t < arr[m] {
        return bSearchRecurse(arr, s, m, t)
    }
    return bSearchRecurse(arr, m+1, e, t)
}
