

以下代码能否编译？

```
package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}

```
答案解析：
结果

编译失败，cannot assign to struct field list["student"].Name in map

分析

map[string]Student 的 value 是一个 Student 结构值，所以当list["student"] = student,是一个值拷贝过程。而list["student"]则是一个值引用。那么值引用的特点是只读。所以对list["student"].Name = "LDB"的修改是不允许的。

如何成功修改map的值？

方法一：
```
package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	//list["student"].Name = "LDB"

    /*
        方法1:
    */
    tmpStudent := list["student"]
    tmpStudent.Name = "LDB"
    list["student"] = tmpStudent

	fmt.Println(list["student"])
}
```

其中

```
/**
方法1:
*/
tmpStudent := list["student"]
tmpStudent.Name = "LDB"
list["student"] = tmpStudent

是先做一次值拷贝，做出一个tmpStudent副本,然后修改该副本，然后再次发生一次值拷贝复制回去，list["student"] = tmpStudent,但是这种会在整体过程中发生 2 次结构体值拷贝，性能很差。
```

方法二：
```
package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]*Student

func main() {

	list = make(map[string]*Student)

	student := Student{"Aceld"}

	list["student"] = &student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}
```
我们将 map 的类型的 value 由 Student 值，改成 Student 指针。

var list map[string]*Student
这样，我们实际上每次修改的都是指针所指向的 Student 空间，指针本身是常指针，不能修改，只读属性，但是指向的 Student 是可以随便修改的，而且这里并不需要值拷贝。只是一个指针的赋值。

关于所谓值引用，解释：

在 Go (Golang) 中，理解值类型和值引用是非常重要的概念。虽然 Go 本身并没有明确的“值引用”这个术语，但可以通过讨论值类型和引用类型来解释相关的概念。

### 值类型
值类型是指变量直接存储其值的数据类型。当你将一个值类型的变量赋值给另一个变量时，会创建该值的副本，因此两个变量是相互独立的。

Go 中的值类型包括：
- 基本数据类型：`int`、`float`、`bool`、`string`等
- 结构体类型 (`struct`)
- 数组 (`array`)

例如：
```go
a := 10
b := a
b = 20
fmt.Println(a) // 输出 10
```
在上面的例子中，`a` 和 `b` 是两个不同的变量，即使 `b` 被赋值为 `a`，修改 `b` 并不会影响 `a` 的值，因为 `a` 和 `b` 是独立的副本。

### 引用类型
引用类型是指变量存储的是数据的地址（内存地址），而不是数据本身。当你将一个引用类型的变量赋值给另一个变量时，它们指向同一个内存地址，因此修改一个变量的值会影响另一个变量。

Go 中的引用类型包括：
- 指针 (`pointer`)
- 切片 (`slice`)
- 映射 (`map`)
- 通道 (`channel`)
- 函数 (`function`)

例如：
```go
a := []int{1, 2, 3}
b := a
b[0] = 10
fmt.Println(a) // 输出 [10 2 3]
```
在上面的例子中，`a` 和 `b` 都指向相同的底层数组，因此修改 `b` 中的元素也会影响 `a` 中的元素。

### 总结
在 Go 中，**值类型**直接包含数据，赋值时会复制数据本身；而**引用类型**包含数据的地址，赋值时复制的是地址，因此多个变量可以共享同一块数据。你可以通过使用指针来显式地进行引用传递，这样函数或方法中对数据的修改可以反映到外部。


map的元素是存的值还是引用?

在 Go 中，`map` 是一种引用类型，但它的元素存储的是**值**，而不是引用。

### 具体解释

- **map 本身是引用类型**：当你将一个 `map` 赋值给另一个变量时，这两个变量实际上指向的是同一个底层数据结构（内存地址）。因此，对其中一个 `map` 的修改会反映在另一个 `map` 上。

- **map 的元素是存储的值**：尽管 `map` 是引用类型，但它的每个键对应的值（`map[key] = value` 中的 `value`）是以值的形式存储的。这意味着当你向 `map` 中插入一个值时，Go 会在 `map` 中保存该值的副本，而不是保存该值的引用。

### 例子
```go
package main

import "fmt"

type MyStruct struct {
    Field int
}

func main() {
    m := make(map[string]MyStruct)
    m["key1"] = MyStruct{Field: 1}
    
    // 尝试修改 map 中的值
    value := m["key1"]
    value.Field = 2
    
    // 检查 map 中的值
    fmt.Println(m["key1"].Field) // 输出 1, 而不是 2
}
```

在上面的例子中，`map["key1"]` 存储的是 `MyStruct` 类型的一个副本。当我们通过 `value := m["key1"]` 取出 `map` 中的元素时，得到的也是该值的一个副本，因此修改 `value.Field` 不会影响 `map` 中的实际数据。

如果希望在 `map` 中存储的元素能够被修改，可以使用指针类型来存储引用，例如：

```go
package main

import "fmt"

type MyStruct struct {
    Field int
}

func main() {
    m := make(map[string]*MyStruct)
    m["key1"] = &MyStruct{Field: 1}
    
    // 修改 map 中的值
    m["key1"].Field = 2
    
    // 检查 map 中的值
    fmt.Println(m["key1"].Field) // 输出 2
}
```

在这个例子中，由于 `map` 存储的是 `MyStruct` 的指针，因此修改 `m["key1"].Field` 会直接影响 `map` 中存储的值。

### 总结

- **map 是引用类型**：`map` 的变量本身是引用类型。
- **map 存储的是值**：`map` 中存储的是值的副本，而不是引用。
- **使用指针存储引用**：如果希望 `map` 中的元素能够被修改，可以将值的指针存储在 `map` 中。


为什么上述第一段代码中如果将 m["key1"].Field = 2 无法通过编译？

在 Go 语言中，`map` 的值是通过值拷贝的方式存储的。这意味着当你访问 `map` 中的元素时，实际上获得的是该值的一个拷贝。如果这个值是一个结构体，那么你获得的是该结构体的一个副本，而不是对原始数据的引用。

### 代码无法编译的原因
在第一段代码中：
```go
package main

import "fmt"

type MyStruct struct {
    Field int
}

func main() {
    m := make(map[string]MyStruct)
    m["key1"] = MyStruct{Field: 1}
    
    // 尝试直接修改 map 中的值
    m["key1"].Field = 2 // 这行代码会导致编译错误
    
    // 检查 map 中的值
    fmt.Println(m["key1"].Field)
}
```
这里的问题在于，`m["key1"]` 返回的是 `MyStruct` 类型的值，而不是指向 `MyStruct` 的引用或指针。因此，`m["key1"].Field` 这行代码试图直接修改这个副本的字段，这在 Go 中是不允许的，因为 `map` 的索引操作返回的不是可取地址的值（即非左值）。

具体来说，`map` 索引操作（`m["key1"]`）的结果是一个临时值，而不是一个可以直接修改的左值。因此，你不能直接通过 `m["key1"].Field = 2` 来修改结构体的字段，因为这相当于试图修改一个临时值的字段。

### 解决方法
如果你想要修改 `map` 中存储的结构体的某个字段，有两种常见的解决方法：

1. **通过变量修改后再更新回 `map` 中**：
    ```go
    value := m["key1"]  // 获取副本
    value.Field = 2      // 修改副本
    m["key1"] = value    // 将修改后的副本放回 map 中
    ```

2. **使用指针存储结构体**：
    如果你希望能够直接修改结构体的字段，可以在 `map` 中存储结构体的指针：
    ```go
    package main

    import "fmt"

    type MyStruct struct {
        Field int
    }

    func main() {
        m := make(map[string]*MyStruct)
        m["key1"] = &MyStruct{Field: 1}
        
        // 直接修改 map 中的值
        m["key1"].Field = 2
        
        // 检查 map 中的值
        fmt.Println(m["key1"].Field) // 输出 2
    }
    ```

通过这种方式，`m["key1"]` 返回的是指针，所以你可以直接通过指针来修改结构体的字段。