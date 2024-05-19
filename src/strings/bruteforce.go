package strings

type BruteForceSearch struct {
    txt []rune
}

func NewBruteForceSearch(txt string) *BruteForceSearch {
    return &BruteForceSearch{
        txt: []rune(txt),
    }
}

func (b *BruteForceSearch) GetName() string {
    return "Brute Force Search"
}

func (b *BruteForceSearch) FindPattern(s string) []int {
    found := make([]int, 0)
    s = string([]rune(s))

    for i := range b.txt {
        // pattern length exceeds remaining chars
        if len(b.txt) - i < len(s) {
            break
        }

        matched := true
        for j, t := range s {
            // character unmatch, progress to next char in outer loop
            if b.txt[i+j] != t {
                matched = false
                break
            }
        }
        // all chars in s matched, mark idx
        if matched {
            found = append(found, i)
        }
    }

    return found
}

