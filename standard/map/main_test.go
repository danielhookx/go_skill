package main

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitMap(t *testing.T) {
	type People struct {
		name string
		age  int
	}

	m1 := make(map[interface{}]string)
	m1[People{
		name: "daniel",
		age:  19,
	}] = "People1"
	m1[123] = "123"
	assert.Equal(t, "People1", m1[People{
		name: "daniel",
		age:  19,
	}])
	assert.Equal(t, "123", m1[123])

	m2 := make(map[int]string)
	m2[123] = "123"
	assert.Equal(t, "123", m2[123])

	m3 := make(map[People]string)
	m3[People{
		name: "daniel",
		age:  19,
	}] = "People1"
	assert.Equal(t, "People1", m3[People{
		name: "daniel",
		age:  19,
	}])
}

func TestVisitMap(t *testing.T) {
	var m map[string]int
	assert.Equal(t, 0, m[""])

	i := 0
	for k, v := range m {
		i++
		assert.Equal(t, 0, v)
		assert.Equal(t, "", k)
	}
	assert.Equal(t, 0, i)
}

func TestAppendNilMap(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.EqualError(t, err.(error), "assignment to entry in nil map")
		}
	}()
	var m map[string]int
	m["123"] = 123
}

func TestDeleteMap(t *testing.T) {
	var m map[string]int
	delete(m, "123")
	assert.Equal(t, 0, m[""])

	m2 := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range m2 {
		if v == 2 {
			delete(m2, k)
		}
	}
	assert.EqualValues(t, map[string]int{"1": 1, "3": 3}, m2)
}

func printMemStats(title string) string {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	s := []string{
		fmt.Sprintf("%v---------------------\n", title),
		fmt.Sprintf("\t Allocated memory: %v kb\n", float64(memStats.Alloc)/1024),
		fmt.Sprintf("\t Total memory allocated (including garbage collection overhead): %v kb\n", float64(memStats.TotalAlloc)/1024),
		fmt.Sprintf("\t Memory currently in use: %v kb\n", float64(memStats.Sys-memStats.HeapReleased)/1024),
		fmt.Sprintf("\t Memory reserved for the heap: %v kb\n", float64(memStats.HeapAlloc)/1024),
		fmt.Sprintf("\t Number of bytes freed by garbage collector: %v kb\n", float64(memStats.Frees)/1024),
		fmt.Sprintf("\t Number of garbage collections: %v\n", memStats.NumGC),
		fmt.Sprintf("\t Last amount of memory freed by garbage collector: %v bytes\n", float64(memStats.LastGC-memStats.NextGC)/1024),
	}
	return strings.Join(s, "\n")
}

var intMap map[int]int
var refMap map[int]map[int]int

func TestIntMapGC(t *testing.T) {
	runtime.GC()
	t.Log(printMemStats("begin func"))
	var intMapCnt = 8192
	//init map
	intMap = make(map[int]int, intMapCnt)
	for i := 0; i < intMapCnt; i++ {
		intMap[i] = i
	}
	runtime.GC()
	t.Log(printMemStats("after map init"))

	for i := 0; i < intMapCnt; i++ {
		delete(intMap, i)
	}
	runtime.GC()
	t.Log(printMemStats("after delete"))

	intMap = nil
	runtime.GC()
	t.Log(printMemStats("after set nil and gc"))
}

func TestInterfaceMapGC(t *testing.T) {
	runtime.GC()
	t.Log(printMemStats("begin func"))
	var mapCnt = 8192
	//init map
	refMap = make(map[int]map[int]int, mapCnt)
	for i := 0; i < mapCnt; i++ {
		//如果给value的引用对象初始化，则delete的时候会释放key对应value指向的地址。
		//refMap[i] = make(map[int]int, mapCnt)
		refMap[i] = nil
	}
	runtime.GC()
	t.Log(printMemStats("after map init"))

	for i := 0; i < mapCnt; i++ {
		delete(refMap, i)
	}
	runtime.GC()
	t.Log(printMemStats("after delete"))

	refMap = nil
	runtime.GC()
	t.Log(printMemStats("after set nil and gc"))
}
