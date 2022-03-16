package pointer

import (
	"testing"
	"unsafe"
)

func Test_Pointer(t *testing.T) {
	type inner struct {
		name string
		age  int
	}

	var user1 inner
	user1.name = "zhangsan"

	up1 := unsafe.Pointer(&user1)
	t.Logf("unsafe.Pointer(user1) address=%v, name=%v", up1, (*inner)(up1).name)

	uptr1 := uintptr(up1)
	t.Logf("uintptr(up1) address=%x", uptr1)

	t.Logf("(*inner)(up1) name=%s, age=%d", (*inner)(up1).name, (*inner)(up1).age)
	// *inner
	p := (*int)(up1)
	t.Logf("v address=%x,val=%d", v, *v)

}
