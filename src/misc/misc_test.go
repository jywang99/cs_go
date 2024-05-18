package misc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jy.org/csgo/src/misc"
)

func TestBinarySearch(t *testing.T) {
    assert.Equal(t, 3, misc.BinarySearch([]int{1,2,3,4,5,6,7,8}, 4))
    assert.Equal(t, 5, misc.BinarySearch([]int{1,2,3,4,5,6,7,8}, 6))
    assert.Equal(t, 7, misc.BinarySearch([]int{1,2,3,4,5,6,7,8}, 8))
    assert.Equal(t, 1, misc.BinarySearch([]int{0,10,20,30,40,50,60}, 10))
    assert.Equal(t, -1, misc.BinarySearch([]int{0,10,20,30,40,50,60}, 100))
}
