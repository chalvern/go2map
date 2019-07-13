# go2map

[中文文档](./README_zh.md)

convert YAML to Go's map


## TODO

* more function of FlatMap.
* Improve English documentation (your smart guy CAN help!!!)
* convert JSON.


## example

more see [source code](./example/main.go).

```go
package main

import (
	"fmt"

	"github.com/chalvern/go2map"
)

func main() {
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
}

```