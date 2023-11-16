package sort

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~int64
}
type Float interface {
	~float32 | ~float64
}

type Orderliness interface {
	Int | Uint | Float | ~string
}
