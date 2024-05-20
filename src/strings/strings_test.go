package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	str "jy.org/csgo/src/strings"
)

func TestStringSearches(t *testing.T) {
    // algorithms that find the first occurrence
    txt := "banana"
    searchers := []str.StirngSearch{
        str.NewSuffixArray(txt),
    }
    for _, s := range(searchers) {
        t.Logf("String matcher: %v", s.GetName())
        assert.Equal(t, []int{2}, s.FindPattern("nan"))
        assert.Equal(t, []int{0}, s.FindPattern("ban"))
        assert.Equal(t, []int{2}, s.FindPattern("nana"))
        assert.Equal(t, []int{1}, s.FindPattern("anana"))
        assert.Equal(t, []int{}, s.FindPattern("asdf"))
    }

    // algorithms that find all occurrences
    txt = "ABABDABACDABABCABAB"
    multiSearchers := []str.StirngSearch{
        str.NewBruteForceSearch(txt),
        str.NewKnuthMorrisPratt(txt),
    }
    for _, s := range(multiSearchers) {
        t.Logf("String matcher: %v", s.GetName())
        assert.Equal(t, []int{10}, s.FindPattern("ABABCABAB"))
        assert.Equal(t, []int{0,10,15}, s.FindPattern("ABAB"))
        assert.Equal(t, []int{}, s.FindPattern("CCC"))
    }
}

