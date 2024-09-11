package main

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/cosmos72/gomacro/fast"
	v1 "github.com/danielhookx/go_skill/exercise/gomacro/v1"
)

// func main() {
// 	interp := fast.New()
// 	interp.ImportPackage("fmt", "fmt")
// 	// interp.DeclVar("x", nil, int(2))
// 	// interp.DeclVar("num", nil, int32(0x1))

// 	// vals, _ := interp.Eval(`fmt.Sprintf("%d+%d=%d", x, x, x+x)`)
// 	vals, _ := interp.Eval(`
// 	ret := int32(0)
// 	// 逐位遍历
// 	for i := 0; i < 32; i++ {
// 		mask := int32(1) << i
// 		bit := (15 & mask) >> i
// 		ret += bit
// 	}
// 	fmt.Sprint(ret)
// 	`)
// 	// for simplicity, only use the first returned value
// 	// fmt.Println(vals[0].ReflectValue())
// 	fmt.Println(vals)
// }

func replaceExtScriptValues(extScript string, cols []*v1.ComposeColumnDataItem) string {
	// 定义要替换的字符串和替换规则
	//input := "Hello $1, how are you $2?"
	replacements := map[string]interface{}{
		// "$1": "Alice",
		// "$2": "today",
	}

	for i, col := range cols {
		replacements[fmt.Sprintf("$%d", i+1)] = col.V
		replacements[fmt.Sprintf("${%d}", i+1)] = col.V
		replacements[fmt.Sprintf("${%d.V}", i+1)] = col.V
		replacements[fmt.Sprintf("${%d.K}", i+1)] = col.K
	}

	// 定义正则表达式来匹配变量
	regex := regexp.MustCompile(`(\$[0-9]+)|\$\{\d+(\.[A-Za-z]+)?\}`)

	// 使用正则表达式替换字符串中的变量
	replaced := regex.ReplaceAllStringFunc(extScript, func(match string) string {
		if value, ok := replacements[match]; ok {
			switch v := value.(type) {
			case string:
				return v
			case *int32:
				if v == nil {
					return "nil"
				}
				return fmt.Sprintf("%d", *v)
			default:
				return fmt.Sprintf("%v", v)
			}
		}
		// 如果没有找到匹配的变量，则返回原始字符串
		return match
	})
	return replaced
}

func eval(script string) (ret string) {
	interp := fast.New()
	interp.ImportPackage("fmt", "fmt")
	vals, _ := interp.Eval(script)
	// for simplicity, only use the first returned value
	if len(vals) > 0 {
		if vals[0].ReflectValue().Kind() != reflect.String {
			return script
		}
		return vals[0].ReflectValue().String()
	}
	return script
}

// func main() {
// 	interp := fast.New()
// 	interp.ImportPackage("fmt", "fmt")
// 	// interp.DeclVar("x", nil, int(2))
// 	// interp.DeclVar("num", nil, int32(0x1))
// 	extScript := `
// import "strings"

// elems := make([]string, 0)
// if ${1.K} != 0 {
// 	elems = append(elems, "$1")
// }
// if ${2.K} != 0 {
// 	elems = append(elems, "$2")
// }
// if ${3.K} != 0 {
// 	elems = append(elems, "$3")
// }
// if ${4.K} != 0 {
// 	elems = append(elems, "$4")
// }
// if ${5.K} != 0 {
// 	elems = append(elems, "$5")
// }
// if ${6.K} != 0 {
// 	elems = append(elems, "$6")
// }
// ret := "$1"
// if len(elems) != 0 {
// 	ret = strings.Join(elems, ";")
// }
// fmt.Sprint(ret)
// 	`

// 	k1 := int32(0)
// 	k2 := int32(0)
// 	k3 := int32(0)
// 	k4 := int32(0)
// 	values := []*v1.ComposeColumnDataItem{
// 		&v1.ComposeColumnDataItem{
// 			K: &k1,
// 			V: "正常",
// 		},
// 		&v1.ComposeColumnDataItem{
// 			K: &k2,
// 			V: "异常1",
// 		},
// 		&v1.ComposeColumnDataItem{
// 			K: &k3,
// 			V: "异常2",
// 		},
// 		&v1.ComposeColumnDataItem{
// 			K: &k4,
// 			V: "异常3",
// 		},
// 	}
// 	fmt.Println(eval(replaceExtScriptValues(extScript, values)))
// }

func main() {
	interp := fast.New()
	interp.ImportPackage("fmt", "fmt")
	// interp.DeclVar("x", nil, int(2))
	// interp.DeclVar("num", nil, int32(0x1))
	extScript := `
	import "strings"

	elems := make([]string, 0)
	if 1 != 0 {
		elems = append(elems, "总压差超限")
	}
	if 0 != 0 {
		elems = append(elems, "1")
	}
	ret := "总压差超限"
	if len(elems) != 0 {
		ret = strings.Join(elems, ";")
	}
	fmt.Sprint(ret)
	`

	k1 := int32(0)
	k2 := int32(0)
	k3 := int32(0)
	k4 := int32(0)
	values := []*v1.ComposeColumnDataItem{
		&v1.ComposeColumnDataItem{
			K: &k1,
			V: "正常",
		},
		&v1.ComposeColumnDataItem{
			K: &k2,
			V: "异常1",
		},
		&v1.ComposeColumnDataItem{
			K: &k3,
			V: "异常2",
		},
		&v1.ComposeColumnDataItem{
			K: &k4,
			V: "异常3",
		},
	}
	fmt.Println(eval(replaceExtScriptValues(extScript, values)))
}
