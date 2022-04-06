package subset

import (
	"testing"
)

func TestSubset(t *testing.T) {
	t.Log(subsets([]int{1, 2, 3}))
}

func Test_subsets2(t *testing.T) {
	t.Log(subsets2([]int{1, 2, 3}))
}
