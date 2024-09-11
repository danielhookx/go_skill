package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func bigRespHandler(w http.ResponseWriter, r *http.Request) {
	// 创建一个大量数据的字符串（例如10MB）
	largeData := strings.Repeat("A", 10*1024*1024)

	// 将数据转换为Reader
	dataReader := strings.NewReader(largeData)

	fmt.Fprintln(w, dataReader)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bigresp", bigRespHandler)
	http.ListenAndServe(":8080", nil)
}
