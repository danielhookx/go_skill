package cache_test

import (
	"testing"

	"github.com/coocood/freecache"
)

func TestCache(t *testing.T) {
	cacheSize := 100 * 1024 * 1024
	c := freecache.NewCache(cacheSize)

	kk := "test"
	v, err := c.Get([]byte(kk))
	t.Log(err)

	t.Log(v)

}
