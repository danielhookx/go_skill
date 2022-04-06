package main

import (
	"log"
	"runtime"
)

func test() {
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println(" ===> loop end.")
}

func main() {
	log.Println("Start.")

	test()

	log.Println("force gc.")
	runtime.GC() //强制调用gc回收

	log.Println("Done.")
}
