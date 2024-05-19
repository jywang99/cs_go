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
    multiSearchers := []str.StirngSearch{
        str.NewBruteForceSearch(txt),
    }
    for _, s := range(multiSearchers) {
        t.Logf("String matcher: %v", s.GetName())
        assert.Equal(t, []int{2,4}, s.FindPattern("na"))
        assert.Equal(t, []int{1,3,5}, s.FindPattern("a"))
        assert.Equal(t, []int{}, s.FindPattern("asdf"))
    }
}

