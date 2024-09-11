package main

import (
	"fmt"
	"testing"

	"github.com/cosmos72/gomacro/fast"
)

func BenchmarkMain(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {

		for p.Next() {
			interp := fast.New()
			interp.ImportPackage("fmt", "fmt")
			interp.DeclVar("num", nil, int32(0x1))

			// vals, _ := interp.Eval(`fmt.Sprintf("%d+%d=%d", x, x, x+x)`)
			_, _ = interp.Eval(`
			ret := int32(0)
			// 逐位遍历
			for i := 0; i < 32; i++ {
				mask := int32(1) << i
				bit := (num & mask) >> i
				ret += bit
			}
			fmt.Sprint(ret)
			`)

			// for simplicity, only use the first returned value
			// fmt.Println(vals[0].ReflectValue())
			// fmt.Println(vals)
		}
	})
}

func TestScript(t *testing.T) {
	str := "ret := int32(0)\n\t// 逐位遍历\n\tfor i := 0; i < 32; i++ {\n\t\tmask := int32(1) << i\n\t\tbit := (15 & mask) >> i\n\t\tret += bit\n\t}\n\tfmt.Sprint(ret)"

	interp := fast.New()
	interp.ImportPackage("fmt", "fmt")
	// interp.DeclVar("x", nil, int(2))
	// interp.DeclVar("num", nil, int32(0x1))

	// vals, _ := interp.Eval(`fmt.Sprintf("%d+%d=%d", x, x, x+x)`)
	vals, _ := interp.Eval(str)
	// for simplicity, only use the first returned value
	// fmt.Println(vals[0].ReflectValue())
	fmt.Println(vals)
}
