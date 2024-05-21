package misc

func PrimeGen(size int) []int {
    if size <= 0 {
        return []int{}
    }

    // array of primes, next idx to insert
    primes := make([]int, size)
    pi := 0

    // for every positive number starting from 2
    for n:=2; ; n++ {
        // found desired # of primes, break
        if pi >= size {
            break
        }

        // check if n is prime
        pri := true
        for i:=0; i<pi; i++ {
            p := primes[i]
            if p > n/2 {
                break
            }
            if n % p == 0 {
                pri = false
            }
        }

        // found prime, add to arr
        if pri {
            primes[pi] = n
            pi ++
        }
    }

    return primes
}

