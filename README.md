# go2map

[中文文档](./README_zh.md)

Usage of Yaml in Golang project is boring, as STRUCT must be better defined first to mapping the Yaml content. **However, Yaml content always contains many enties which makes STRUCT defination very difficult.**

Go2Map's target is to make usage of Yaml easier by **avoid defining complex structs while conveniently reading values in Yaml**.

```yaml
a: Easy!
b:
  c: 2
  d: [3, 4]
  1: test
  2:
  - 21
  - 22
```
In the above example, the easiest way is to get the value of a specific path, without defining a struct, in a way of `x.a`, `x.b.c`, `x.b.1` by once Yaml content loaded.

## TODO

* more function of FlatMap.
* convert JSON if needed (welcome your issues).

## example

more see [source code](./example/main.go).

```go
package main

import (
	"fmt"
	"path"
	"runtime"

	"github.com/chalvern/go2map"
)

func main() {
	// convert from bytes
	var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
  1: test
  2:
  - 21
  - 22
`
	m := go2map.Yaml2Map([]byte(data))
	b1, _ := m.GetStringOf("b.1")
	fmt.Println(b1)

	// convert from filePath
	_, file, _, _ := runtime.Caller(0)
	yamlFile := path.Join(path.Dir(file), "example.yaml")
	m2 := go2map.Yaml2MapFromFile(yamlFile)
	jingweiLink, _ := m2.GetStringOf("zh.blog.jingwei.link")
	fmt.Println(jingweiLink)
}

```