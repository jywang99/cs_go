package strings

type KnuthMorrisPratt struct {
    txt []rune
}

func NewKnuthMorrisPratt(txt string) *KnuthMorrisPratt {
    return &KnuthMorrisPratt{
        txt: []rune(txt),
    }
}

func (b *KnuthMorrisPratt) GetName() string {
    return "Knuth Morris Pratt"
}

func (b *KnuthMorrisPratt) FindPattern(src string) []int {
    ptn := []rune(src)
    lps := b.preprocess(ptn)

    matches := make([]int, 0)
    j := 0
    for i, c := range b.txt {
        if ptn[j] == c {
            j++
            // if pattern matched entirely, mark it and start over
            if j > len(ptn)-1 {
                matches = append(matches, i - len(ptn) + 1)
                j = 0
            }
            continue
        }
        // if r[i], r[j] mismatch,
        // move j back continuously until it matches or becomes 0
        for ; ptn[j] != c && j > 0; {
            j = *lps[j]
        }
    }

    return matches
}

func (b *KnuthMorrisPratt) preprocess(s []rune) []*int {
    // index 0: padding, easier when using the array
    lps := make([]*int, len(s)+1)
    for i:=1; i<len(lps); i++ {
        l := 0
        lps[i] = &l
    }

    l := 0 // next index in pattern to match
    // index 1 (first letter) is always 0
    for i:=2; i<=len(s); i++ {
        // if matched, increment the next index to match (i.e. length matched), and record that in lps
        // iterating from i=2, but want to start from 2nd char in pattern
        if s[i-1] == s[l] {
            l ++
            lc := l
            lps[i] = &lc
        } else {
            l = 0
        }
    }
    
    return lps
}

