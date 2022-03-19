package main

func main() {
	m := make(map[interface{}]interface{}, 16)
	m["111"] = 1
	m["222"] = 2
	m["444"] = 4
	_ = m["444"]
	_, _ = m["444"]
	delete(m, "444")

	for range m {
	}
}
